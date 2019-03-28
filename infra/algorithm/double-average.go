package algorithm

import (
	"math/rand"
	"time"
)

func DoubleAverage(count, amount int64) int64 {
	if count == 0 {
		return amount
	}
	//计算最大可用金额
	max := amount - min*count
	//计算最大可用平均值 max<count 得到值0
	avg := max / count
	//二倍均值
	doubleAvg := avg*2 + min
	//随机红包金额序列 把二倍均值作为随机数的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(doubleAvg) + min
	return x
}
