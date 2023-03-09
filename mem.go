package upress

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type MemPresser struct {
	sizeMB int
	data   []byte
}

func NewMemPresser(sizeMB int) *MemPresser {
	return &MemPresser{
		sizeMB: sizeMB,
	}
}

//go:noinline
func (mp *MemPresser) Do() {
	fmt.Printf("pressing memory, size=%dMB\n", mp.sizeMB)
	mp.data = make([]byte, mp.sizeMB*1024*1024)
	for i := range mp.data {
		mp.data[i] = 255
	}
}

func (mp *MemPresser) Stop() {
	mp.data = nil
	pre := debug.SetGCPercent(-1)
	runtime.GC()
	debug.SetGCPercent(pre)
}
