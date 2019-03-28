package test

import (
	// 导入1个包没有使用到Go会把导入的包删除，加上下划线表示
	// 会使用这个包但是该源码文件不使用包中的具体内容，所以通过下划线方式导入驱动
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"log"
)

var (
	// 指针
	db *dbx.Database
)

// 包初始化代码函数初始化数据库
func init() {
	settings := dbx.Settings{
		//驱动名
		DriverName: "mysql",
		Host:       "127.0.0.1:3306",
		User:       "root",
		Password:   "root",
		Database:   "redenvelope",
		Options: map[string]string{
			//时间戳类型解析为time类型
			"parseTime": "true",
		},
	}
	var err error
	db, err = dbx.Open(settings)
	if err != nil {
		log.Println(err)
	}
	// defer db.Close()

	// 持久化对象 表名
	db.RegisterTable(&GoodsSigned{}, "goods")
	db.RegisterTable(&GoodsUnsigned{}, "goods_unsigned")
}

//事务锁方案
func UpdateForLock(g Goods) {
	//通过 db.tx 函数构建事务锁代码块
	err := db.Tx(func(runner *dbx.TxRunner) error {
		//1.锁定需要更新的数据表记录
		//事务锁查询语句,使用 for update 语句锁定记录行,查询数据行的同时锁定该行
		//锁定该行后当前数据库连接才能操作该行数据,其他所有数据库连接都不能操作
		//需要等待锁的释放
		//TODO 生产代码不使用 * 查询什么返回什么
		query := `SELECT * FROM goods WHERE envelope_no=? FOR UPDATE`
		out := &GoodsSigned{}
		ok, err := runner.Get(out, query, g.EnvelopeNo)
		if !ok || err != nil {
			//返回错误 事务回滚
			return err
		}
		//2.计算剩余金额 剩余数量
		subAmount := decimal.NewFromFloat(0.01)
		remainAmount := out.RemainAmount.Sub(subAmount)
		remainQuantity := out.RemainQuantity - 1
		//3.执行更新
		update := `UPDATE goods SET remain_amount=?,remain_quantity=? WHERE envelope_no=?`
		//返回插入Id 这里是更新 返回0 使用_下划线忽略
		//受影响的行数 更新成功返回1 更新失败返回0 因此判断是否更新成功
		//执行更新过程中是否发生错误 err
		_, rows, err := db.Execute(update, remainAmount, remainQuantity, g.EnvelopeNo)
		if err != nil {
			//返回错误 事务回滚
			return err
		}
		if rows < 1 {
			return errors.New("库存扣减失败")
		}
		//该事务函数func如果返回非空err 本次事务就会回滚 返回err是nil本次事务会被提交
		//返回err为nil 提交事务
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

//数据库无符号字段 直接更新
func UpdateForUnsigned(g Goods) {
	//事务锁在程序逻辑计算
	//无符号字段直接更新在SQL中计算
	update := `UPDATE goods_unsigned SET 
				remain_amount=remain_amount-?,
				remain_quantity=remain_quantity-?
				WHERE envelope_no=?`
	_, rows, err := db.Execute(update, 0.01, 1, g.EnvelopeNo)
	if err != nil {
		log.Println(err)
	}
	if rows < 1 {
		log.Println("库存扣减失败")
	}
}

//乐观锁 + WHERE条件
func UpdateForOptimistic(g Goods) {
	//乐观锁需要在WHERE条件确定需要扣减的资源是否足够
	update := `UPDATE goods SET 
				remain_amount=remain_amount-?,
				remain_quantity=remain_quantity-?
				WHERE envelope_no=? 
				AND remain_amount>=? AND remain_quantity>=?`
	_, rows, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		log.Println(err)
	}
	if rows < 1 {
		log.Println("库存扣减失败")
	}
}

//乐观锁 + WHERE条件 + 无符号字段 双保险
func UpdateForOptimisticUngigned(g Goods) {
	//乐观锁需要在WHERE条件确定需要扣减的资源是否足够
	update := `UPDATE goods_unsigned SET 
				remain_amount=remain_amount-?,
				remain_quantity=remain_quantity-?
				WHERE envelope_no=? 
				AND remain_amount>=? AND remain_quantity>=?`
	_, rows, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		log.Println(err)
	}
	if rows < 1 {
		log.Println("库存扣减失败")
	}
}
