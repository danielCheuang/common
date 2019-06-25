package utils

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

//获取今天0点
func GetToday() time.Time {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//fmt.Printf("%v\n", t)
	return tm1
}

//获取前一天0点
func GetLastDay() time.Time {
	t := time.Now().AddDate(0,0, -1)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//fmt.Println(tm1)
	return tm1
}


//获取明天0点
func GetNextDay() time.Time {
	t := time.Now().AddDate(0,0, 1)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//fmt.Println(tm1)
	return tm1
}


/**
	获取周一 00:00:00的时间time
 */
func GetWeekStartTime() time.Time {
	var weekDate int
	weekDate = int( time.Now().Weekday() + 0 )
	return GetToday().AddDate(0, 0, - weekDate + 1 )
}

/**
	获取本月1号 00:00:00的时间time
 */
func GetMonthStartTime() time.Time {
	var mouthDate int
	mouthDate = int( time.Now().Day() )
	return GetToday().AddDate(0, 0, - mouthDate + 1 )
}


/**
	获取days天前 00:00:00的时间time
 */
func GetBeforedDayTime( days int ) time.Time {
	return GetToday().AddDate(0, 0, - days)
}



/**
	计算date1, date2 的时间差天数
*/
func ComDateDiff(date1, date2 string) int {
	fmt.Println("计算时间差")

	s1 := strings.Split(date1, "-")
	y1,_ := strconv.Atoi(s1[0])
	m1,_ := strconv.Atoi(s1[1])
	d1,_ := strconv.Atoi(s1[2])
	t1 := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.Local)

	s2 := strings.Split(date2, "-")
	y2,_ := strconv.Atoi(s2[0])
	m2,_ := strconv.Atoi(s2[1])
	d2,_ := strconv.Atoi(s2[2])
	t2 := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.Local)

	days := int(t2.Sub(t1) /(1000000000 * 60 * 60 * 24) )
	return days
}


/**
	获取日期
*/
func GetDate(t time.Time) string {
	dateStr := t.Format("2006-01-02")
	return dateStr
}

func DateAddDays(dateStr string, days int) string {
	strs := strings.Split(dateStr, "-")
	y,_ := strconv.Atoi(strs[0])
	m,_ := strconv.Atoi(strs[1])
	d,_ := strconv.Atoi(strs[2])
	t1 := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	t2 := t1.AddDate(0,0, days)
	return GetDate(t2)
}

func AddZone(t1 time.Time, hour int64) time.Time {
	return t1.Add(time.Hour * time.Duration(hour))
}


/**
	定时任务器
*/
func StartTimer(f func(n string), hour ,min, sec int, param string ) {
	go func() {
		timerJobExecCnt := 0
		for {
			if timerJobExecCnt > 0 {
				f(param)
			}
			timerJobExecCnt ++
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), hour, min, sec, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}




/**
	从请求体读取数据
*/
func ReadDataFromRequestBody(body io.ReadCloser)  []byte {
	var cnt bytes.Buffer

	if _, err := io.Copy(&cnt, body); err != nil {
		fmt.Println("+++bet_handle.SaveBetResolve+++从请求body读数据err:", err)
		panic(err)
	}
	return cnt.Bytes()
}

func IsDateStr(dateStr string) bool {
	fmt.Printf("---IsDateStr()---dateStr:%v\n", dateStr)
	dateStr = dateStr + " 00:00:00"
	the_time, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return false
	}
	unix_time := the_time.Unix()
	fmt.Println(unix_time)

	return true
}

func IsTimeStr(timeStr string) bool {
	fmt.Printf("---IsTimeStr()---timeStr:%v\n", timeStr)
	timeStr = "2019-01-01 " + timeStr
	the_time, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return false
	}
	unix_time := the_time.Unix()
	fmt.Println(unix_time)
	return true
}


