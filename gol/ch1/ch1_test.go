package ch1

import "testing"

func TestCh1CloseRead(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	<-ch // read close ch, no painc, and return default value
	t.Log("ok")
}
