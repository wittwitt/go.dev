package ctx1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Wait() {
	resultCh := make(chan CC, 100)
	ctx, cancel := context.WithCancel(context.Background())
	tk := time.NewTicker(13 * time.Millisecond)
	go func() {
		lastNumber := -1
		for rs := range resultCh {
			fmt.Println(rs)
			if rs.Number == 7 && lastNumber == 1 {
				fmt.Println("cancel")
				tk.Stop()
				cancel()
				break
			}
			lastNumber = rs.Number
		}
		fmt.Println("oooover")
	}()

	j := 0

	// ctx2, _ := context.WithTimeout(ctx, 12*time.Millisecond)
FOR:
	for {
		select {
		case <-tk.C:
			j++
			go Get1(ctx, resultCh, j)
		case <-ctx.Done():
			fmt.Println("tkstop")
			break FOR
		}
	}
	log.Println("ok")
}

type CC struct {
	J      int
	Number int `json:"number"`
	C      int `json:"c"`
}

// Get1 Get1
func Get1(ctx context.Context, rsCh chan CC, j int) {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8000", nil)

	req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("do: %v", err)
		return
	}

	bd, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ReadAll: %v", err)
		return
	}

	o := &CC{}
	if err := json.Unmarshal(bd, o); err != nil {
		log.Printf("Unmarshal: %v", err)
		return
	}

	// fmt.Println("xxxx", j)
	o.J = j

	select {
	case <-ctx.Done():
	default:
		rsCh <- *o
	}
}

func Do(ctx context.Context, outCh chan int) {
	tk := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-tk.C:
			outCh <- 1
		case <-ctx.Done():
			fmt.Println("do Done")
			return
		}
	}
}

func Do2(ctx context.Context, outCh chan int) {
	tk := time.NewTimer(3 * time.Second)
	for {
		select {
		case <-tk.C:
			outCh <- 3
		case <-ctx.Done():
			fmt.Println("do2 Done")
			return
		}
	}
}

func Wait2() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	outCh := make(chan int, 10)
	outCh2 := make(chan int, 10)

	go Do(ctx, outCh)
	go Do2(ctx, outCh2)

	for {
		select {
		case <-outCh:
			fmt.Println("ooo1")
			cancel()
			return
		case <-outCh2:
			fmt.Println("ooo2")
			cancel()
			return
		}
	}
}
