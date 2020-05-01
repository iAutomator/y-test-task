package visits

import (
	"sync"
	"testing"
	"time"
)

func TestVisitsCounterIsResetOnInit(t *testing.T) {
	c := visitsCounter{}
	if c.visitsCount() != 0 {
		t.Fail()
	}
}

func TestVisitsCounterOnParallelCalls(t *testing.T) {
	var c visitsCounter
	var wg sync.WaitGroup

	const nVisits = 100
	for i := 0; i < nVisits; i++ {
		wg.Add(1)
		i := i
		go func() {
			time.Sleep(time.Second * time.Duration(i%5))
			c.newVisit()
			wg.Done()
		}()
	}
	wg.Wait()
	if c.visitsCount() != nVisits {
		t.Errorf("actual = %d", c.visitsCount())
	}
}
