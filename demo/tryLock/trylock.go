package trylock

import "time"

type TryLock chan struct{}

func NewTryLock() TryLock {
	ch := make(chan struct{}, 1)
	return ch
}

func (tl *TryLock) Lock() {
	(*tl) <- struct{}{}
}

func (tl *TryLock) UnLock() {
	<-(*tl)
}

func (tl *TryLock) TryLock() bool {
	select {
	case *tl <- struct{}{}:
		return true
	default:
		return false
	}
}

func (tl *TryLock) TryLockWithTimeOut(d time.Duration) bool {
	t := time.NewTimer(d)
	select {
	case <-t.C:
		return false
	case *tl <- struct{}{}:
		t.Stop()
		return true
	}
}
