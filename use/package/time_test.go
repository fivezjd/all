package _package

import (
	"sync"
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

	// 判断时间之间的差值
	t.Log(timeTwo.Sub(timeOne))

	// 判断当前时间和指定时间的差值 可以间接用来计算程序的执行时间

	t.Log(time.Since(timeOne))
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

	// 设置默认时区
	time.Local = location
}

// 时间的常量定义

func TestConst(t *testing.T) {
	// 时
	t.Log(time.Hour)

	// 分
	t.Log(time.Minute)

	// 秒
	t.Log(time.Second)
}

// 如果我不设置时区，默认的时区是哪里？

func TestDefaultLoc(t *testing.T) {
	// 设置本地时区，并且返回本地时间
	t.Log(time.Now().Local())

	// 创建时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}
	// 设置默认时区
	time.Local = location

	// 获取设置的时区
	t.Log(time.Now().Location())
	// out: Asia/Shanghai
}

const Format = "2006-01-02 15:04:05"

/**
字符串和时间的转换
*/

// 字符串转换为时间
func TestStringToTime(t *testing.T) {
	timeString := "2023-04-12 12:31:54"
	parse, err := time.Parse(Format, timeString)
	if err != nil {
		return
	}
	t.Log(parse)
}

// 将字符串按指定的时区转换为时间

func TestStringToTimeInLoc(t *testing.T) {
	timeString := "2023-04-12 12:31:54"
	parse, err := time.ParseInLocation(Format, timeString, time.Local)
	if err != nil {
		return
	}
	t.Log(parse)
}

// 定时器和睡眠

// 指定时间后执行任务

func TestTimerFuc(t *testing.T) {
	timer := time.NewTimer(time.Second * 5)
	select {
	case v := <-timer.C:
		t.Log(v)
	}
}

// 指定时间后执行函数

func TestTimeFunc1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	timer := time.AfterFunc(time.Second*3, func() {
		wg.Done()
		t.Log("3秒后执行了此函数")
	})
	wg.Wait()
	defer timer.Stop()
}

// 定时任务

func TestTicker(t *testing.T) {
	//
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case v := <-ticker.C:
			t.Log(v)
		}
	}
}

// 重置定时器

func TestRest(t *testing.T) {
	ticker := time.NewTicker(time.Second * 3)
	ticker.Reset(time.Second * 2)
}

// 停止定时器

func TestStop(t *testing.T) {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
}

// 睡眠指定时间

func TestSleep(t *testing.T) {
	time.Sleep(time.Second)
}
