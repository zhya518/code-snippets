package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func task() {
	c := cron.New()
	c.AddFunc("30 * * * *", func() {
		fmt.Println("在每个小时的30分钟的时候实行")
	})
	c.AddFunc("30 3-6,20-23 * * *", func() {
		fmt.Println("在每天早上3-6点, 晚上8-11点的30分钟执行")
	})
	c.AddFunc("CRON_TZ=Asia/Shanghai 30 04 * * *", func() {
		fmt.Println("在每天的北京时间04:30执行")
	}) // 如果不指定，默认使用机器的时区
	c.AddFunc("@hourly", func() { fmt.Println("每一小时执行。从1小时后开始") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("每一小时30分执行，1小时30分开始") })
	c.Start()
	//...
	// 函数在它们的goroutine中异步的执行
	//...
	// 期间可以安全的增加任务
	c.AddFunc("@daily", func() { fmt.Println("每天执行") })
	//...
	// 可以检查任务的状态
	inspect(c.Entries())
	//...
	c.Remove(entryID) // 移除某个任务
	//...
	c.Stop() // 不再执行后续的任务
}

// cron的格式
// 字段名        | 强制设置?   | 允许值           | 允许的特殊字符
// ----------   | ---------- | --------------  | -----------
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
//
// 维基百科也介绍了cron的格式:
// # ┌───────────── minute (0 - 59)
// # │ ┌───────────── hour (0 - 23)
// # │ │ ┌───────────── day of the month (1 - 31)
// # │ │ │ ┌───────────── month (1 - 12)
// # │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
// # │ │ │ │ │                                   7 is also Sunday on some systems)
// # │ │ │ │ │
// # │ │ │ │ │
// # * * * * * <command to execute>
//
// 特殊字符代表的意义：
// *: 代表满足这个字段的每一个值。
// /: 代表对于时间范围的步数，比如第一个字段*/5代表每5分钟,也就是5,10,15,20,25,30,35,40,45,50,55,00分钟。
// ,: 代表一组列表，比如第五个字段MON,WED,FRI代表每星期一、星期三、星期五会触发。
// -: 代表一个范围，比如第二个字段10-15代表会在每天10点、11点、12点、13点、14点、15点被触发。
// ?: 有时候可以在第三个和第五个字段代替*
// 同时这个库还提供预定义的几种形式：
//
//
// 预定义类型             | 描述                               | 等价格式
// -----                  | -----                              | ------
// @yearly (or @annually) | 在每年的一月一日零点               | 0 0 1 1 *
// @monthly               | 在每月的第一天零点                 | 0 0 1 * *
// @weekly                | 在每周的第一天零点，也就是周日零点 | 0 0 * * 0
// @daily (or @midnight)  | 每天零点运行                       | 0 0 * * *
// @hourly                | 每小时开始的时候运行               | 0 * * * *
