package metrics

import (
	"log"
	"os"
	"testing"
	"time"
	"github.com/rcrowley/go-metrics"
)

//gauge 是一个最简单的度量标准，为瞬时状态，返回存储的值
func Test_metricsGauge(t *testing.T){
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)


	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		g.Update(j)
		j++
	}
}

//counter 计数器，对gauge的封装
func Test_metricsCounter(t *testing.T){
	c := metrics.NewCounter()
	metrics.Register("foo", c)

	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	for true {
		time.Sleep(time.Second * 1)
		c.Inc(1)
	}
}

//Histogram统计数据的分布情况。比如最小值，最大值，中间值，还有中位数，75百分位, 90百分位, 95百分位, 98百分位, 99百分位, 和 99.9百分位的值(percentiles)
// mean 为算术平均数
// stddev 为标准差
//如果将一组数据从小到大排序，并计算相应的累计百分位，则某一百分位所对应数据的值就称为这一百分位的百分位数。可表示为：一组n个观测值按数值大小排列。如，处于p%位置的值称第p百分位数。
func Test_metricsHistogram(t *testing.T){
	// 第一参数是内部存储数据的个数
	//第二个参数是指数后乘以的数值 具体位置 ?? 什么意思？？
	s := metrics.NewExpDecaySample(1024, 0.015) // or metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	metrics.Register("baz", h)
	h.Update(1)

	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		h.Update(j)
	}
}

// Meter度量一系列事件发生的速率(rate)，例如TPS。Meters会统计最近1分钟，5分钟，15分钟，还有全部时间的速率。
func Test_metricMeter(t *testing.T){
	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(1)

	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		m.Mark(j)
	}
}
