package algorithm

import (
	"testing"
	// 静态导入
	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleRand(t *testing.T) {
	// 10个红包 100元
	count, amount := int64(10), int64(100*100)
	// 剩余金额
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain -= x
		sum += x
	}
	// msg 输出到控制台 webui 提示性语句
	// t *testing.T
	// logic 需要验证逻辑 func
	// convey.Convey()
	Convey("简单随机算法", t, func() {
		// convey框架断言用So()函数
		// actual 实际值
		// 断言函数 Should
		// expected 预期值
		So(sum, ShouldEqual, amount)
	})
}

func TestDoubleRand(t *testing.T) {
	ForTest("二次随机算法", t, DoubleRand)
}

func TestBeforeShuffle(t *testing.T) {
	ForTest("先洗牌算法", t, BeforeShuffle)
}

func TestDoubleAverage(t *testing.T) {
	ForTest("而被均值算法", t, DoubleAverage)
}

func ForTest(message string, t *testing.T, fn func(count, amount int64) int64) {
	// 10个红包 100元
	count, amount := int64(10), int64(100*100)
	// 剩余金额
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := fn(count-i, remain)
		remain -= x
		sum += x
	}
	// msg 输出到控制台 webui 提示性语句
	// t *testing.T
	// logic 需要验证逻辑 func
	// convey.Convey()
	Convey(message, t, func() {
		// convey框架断言用So()函数
		// actual 实际值
		// 断言函数 Should
		// expected 预期值
		So(sum, ShouldEqual, amount)
	})
}
