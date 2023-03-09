package upress

import (
	"fmt"
	"sync"
)

type CpuPresser struct {
	concurrency int
	wg          sync.WaitGroup
	stopChan    chan struct{}
}

func NewCpuPresser(n int) *CpuPresser {
	return &CpuPresser{
		concurrency: n,
		stopChan:    make(chan struct{}),
	}
}

func (cp *CpuPresser) Do() {
	fmt.Printf("pressing %d cpu core\n", cp.concurrency)
	for i := 0; i < cp.concurrency; i++ {
		cp.wg.Add(1)
		go func() {
			defer func() {
				cp.wg.Done()
			}()
			for {
				select {
				case <-cp.stopChan:
					return
				default:
				}
			}
		}()
	}
}

func (cp *CpuPresser) Stop() {
	close(cp.stopChan)
	cp.wg.Wait()
}
