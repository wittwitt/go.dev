package ch1

import "testing"

func TestCh1CloseRead(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	<-ch // read close ch, no painc, and return default value
	t.Log("ok")
}

func Test_ch_func(t *testing.T) {

	type sfn func()

	ch := make(chan sfn)

	go func() {
		i := 10
		ch <- func() {
			t.Log("this i is: ", i)
		}
	}()

	fnn := <-ch

	fnn()
}
