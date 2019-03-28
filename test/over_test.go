package test

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"log"
	"testing"
)

//基准测试 Benchmark 开头 首字母B大写
func BenchmarkUpdateForLock(b *testing.B) {
	g := GoodsSigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(10000000)
	_, err := db.Insert(g)
	if err != nil {
		log.Println(err)
		return
	}
	//b.N 测试函数执行次数
	//通过执行次数 和 Benchmark函数执行总时间 计算出平均耗时
	for i := 0; i < b.N; i++ {
		UpdateForLock(g.Goods)
	}
}

func BenchmarkUpdateForUnsigned(b *testing.B) {
	g := GoodsUnsigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(10000000)
	_, err := db.Insert(g)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		UpdateForUnsigned(g.Goods)
	}
}

func BenchmarkUpdateForOptimistic(b *testing.B) {
	g := GoodsSigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(10000000)
	_, err := db.Insert(g)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		UpdateForOptimistic(g.Goods)
	}
}

func BenchmarkUpdateForOptimisticUngigned(b *testing.B) {
	g := GoodsUnsigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(10000000)
	_, err := db.Insert(g)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		UpdateForOptimisticUngigned(g.Goods)
	}
}
