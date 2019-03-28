package algorithm

import "math/rand"

//后洗牌算法 - 只能在发红包时生成 只能生成1个序列
func AfterShuffle(count, amount int64) []int64 {
	inds := make([]int64, 0)
	//计算最大可调度金额
	max := amount - min*count
	//剩余金额
	remain := max
	//随机生成初级红包序列
	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain = remain - x
		inds = append(inds, x)
	}
	//后洗牌初级红包序列 需要洗牌数据长度
	//swap函数 随机生成2个索引 i j 数值进行交换
	rand.Shuffle(len(inds), func(i, j int) {
		inds[i], inds[j] = inds[j], inds[i]
	})
	return inds
}
