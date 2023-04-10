package use

import (
	"testing"
	"time"
)

// TestTimeNow 获取当前时间
func TestTimeNow(t *testing.T) {
	t.Log(time.Now())
}

// TestTimeNow 时间的计算
func TestTimeCalculate(t *testing.T) {
	// 加指定时间

	// 当前时间增加10秒
	t.Log(time.Now().Add(time.Second * 10))

	// 减指定时间

	// 当前时间减少10秒
	t.Log(time.Now().Add(time.Second * -10))

	// 按照年月日来指定增减或者减少时间

	// 增加一天
	t.Log(time.Now().AddDate(0, 0, 1))
	// 减少一年
	t.Log(time.Now().AddDate(-1, 0, 0))

	// 获取指定的两个时间点之间的时间差

	start := time.Now()

	end := start.AddDate(0, 0, 1)
	// 结束时间比开始时间大86400秒
	t.Log(end.Sub(start).Seconds())
}

// 时间点的比较🆚
func TestTimeCompare(t *testing.T) {
	timeOne := time.Now()

	timeTwo := time.Now().Add(time.Second * 1000)

	// 比较两个时间是否相等
	t.Log(timeOne.Equal(timeTwo))

	// 判断时间One 是否在时间Two之后
	t.Log(timeOne.After(timeTwo))

	// 判断时间One是否在时间Two之前
	t.Log(timeOne.Before(timeTwo))
}

// 获取时间戳

func TestTimeUnix(t *testing.T) {
	// 秒级时间戳
	t.Log(time.Now().Unix())

	// 毫秒级时间戳
	t.Log(time.Now().UnixMilli())

	// 微妙级时间戳
	t.Log(time.Now().UnixMicro())

	// 纳秒时间戳
	t.Log(time.Now().UnixNano())
}

// 时间的格式化和解析

func TestFormat(t *testing.T) {

	// 按如下格式输出当前时间：

	// 12小时制 2006-01-02 03:04:05 PM
	t.Log(time.Now().Format("2006-01-02 03:04:05 PM"))

	// 24小时制
	t.Log(time.Now().Format("2006-01-02 15:04:05"))

	// 月份日期中不带前导0
	t.Log(time.Now().Format("2006-1-2 15:04:05"))

	// 时分秒中不带前导0
	t.Log(time.Now().Format("2006-01-02 3:4:5"))

	// 解析 2023-04-10 21:32:38 解析字符串
	parse, err := time.Parse("2006-01-02 15:04:05", "2023-04-10 21:32:38")
	if err != nil {
		t.Error(err)
	}
	t.Log(parse)
}

// 时区的处理
func TestLoc(t *testing.T) {
	// 创建时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	// 指定时区
	f := time.Now().In(location).Format("2006-01-02 15:04:05")
	t.Log(f)
}
