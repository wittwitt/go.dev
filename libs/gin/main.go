package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/list", list)
	r.GET("/p1", p1)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func list(c *gin.Context) {
	var sleepI int32 = 0

	h := c.Request.Header.Get("X-Cloud-Trace-Context")

	header_fpath := c.Request.Header.Get("FilePath")

	// /scfs/disk03/{}

	i, _ := strconv.ParseInt(header_fpath[13:], 10, 64)

	sleepI = atomic.AddInt32(&sleepI, 1)
	if sleepI > 5 {
		sleepI = 0
	}

	fmt.Printf("req: %d, %v , %v , %v, \n", i, h, c.Request.RemoteAddr, sleepI)

	t1 := time.Now()

	time.Sleep(time.Duration(5) * time.Second)
	t2 := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("i: %d,  t1: %s, t2: %s", i,
			t1.Format("15:04:05.000"),
			t2.Format("15:04:05.000")),
	})
	t3 := time.Now()

	fmt.Printf("i: %d,  t1(%v): %s, t2(%v): %s , t3: %s \n",
		i,
		t3.Sub(t1), t1.Format("15:04:05.000"),
		t3.Sub(t2), t2.Format("15:04:05.000"),
		t3.Format("15:04:05.000"))

}

func p1(c *gin.Context) {
	fmt.Println("start:", time.Now(), c.Request.RemoteAddr)

LOOP:
	for i := 0; i < 10; i++ {
		select {
		case <-c.Done():
			fmt.Println("cccc:", time.Now())
			break LOOP
		case <-c.Request.Context().Done():
			fmt.Println("ctx done:", time.Now())
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("default:", i, time.Now())
		}
	}
	fmt.Println("end:", time.Now())

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok, abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde",
	})

	fmt.Println("end2:", time.Now())
}
