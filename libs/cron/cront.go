package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {

	// Seconds field, required
	// cron.New(cron.WithSeconds())

	// Seconds field, optional
	// cron.New(cron.WithParser(cron.NewParser(
	// 	cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	// )))

	fmt.Println("go")

	c := cron.New()
	c.AddFunc("1,11,21,31,41,51 * * * *", func() {
		// p.CachePaymentList()
	})
	c.AddFunc("23 10 */2 * *", func() {

	})
}

// 注意： 和linux cron表达式有不同

// "github.com/robfig/cron/v3"

//  c := cron.New()
// // 也可以秒级任务
// //c := cron.New(cron.WithSeconds()) // 创建定时任务 秒
// // spec := "0 */1 * * * *" // 每一分钟 如果用到分钟, 秒要设置为0
// // spec := "* */1 * * * *"
// fmt.Println(c)
// spec := "*/1 * * * *"    // 每一分钟，
