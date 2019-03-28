package algorithm

import (
	"math/rand"
	"time"
)

const (
	//单个红包最小金额
	min = int64(1)
)

//简单随机算法
//param count 红包数量
//param amount 红包金额
//TODO - Golang 中 float32 / float64 计算会丢失精度 涉及金额的计算不能使用 float32/64类型
//把金额数扩大 用最小金额单位 1元=100分 计算完成 分->元 金额使用 int64 计算
func SimpleRand(count, amount int64) int64 {
	//红包数量剩余1个 直接返回剩余金额
	if count == 1 {
		return amount
	}
	//计算最大可调度金额 = 红包总金额 - 单个红包最小金额 * 红包数量
	max := amount - min*count
	//随机金额 可能是0 + min 避免0金额产生
	//TODO - 随机数种子
	//随机数生成没有设置种子，种子数为1，导致每次运行得到的随机数相同
	//只有当程序不停止时得到的随机数才是不同的
	//所以每次重新运行程序都会得到相同的随机数
	//服务上线重启得到相同的红包随机数 不合理
	//给定种子 使用系统时间戳 Unix返回秒 UnixNano纳秒
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(max) + min
	return x
}
