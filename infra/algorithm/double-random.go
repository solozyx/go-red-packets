package algorithm

import (
	"math/rand"
	"time"
)

func DoubleRand(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	//计算最大可调度金额
	max := amount - min*count
	//一次随机 计算出1个种子金额作为基数
	rand.Seed(time.Now().UnixNano())
	seed := rand.Int63n(count*2) + 1
	n := max/seed + min
	//二次随机 计算出红包金额序列
	x := rand.Int63n(n)
	return x + min
}
