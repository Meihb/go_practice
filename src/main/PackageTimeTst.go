package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	now = time.Now()             //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)

	timeObj := time.Unix(timestamp1, 0)
	fmt.Println(timeObj)
	fmt.Println("许大茂",time.Now().Format("2006-01-21 13:18:20"))

	year = timeObj.Year()     //年
	month = timeObj.Month()   //月
	day = timeObj.Day()       //日
	hour = timeObj.Hour()     //小时
	minute = timeObj.Minute() //分钟
	second = timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	now = time.Now()
	fmt.Println("番号",now)
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mwon Jan"))//format是这个意思啊，string里面是Patern!
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
