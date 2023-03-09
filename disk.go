package upress

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var zeroFile1MB = make([]byte, 1024*1024)

type DiskPresser struct {
	path   string
	size   int
	num    int
	prefix string
}

func NewDiskPresser(path string, size, num int) *DiskPresser {
	return &DiskPresser{
		path:   path,
		size:   size,
		num:    num,
		prefix: "tmpfile_",
	}
}

func (dp *DiskPresser) Do() {
	fmt.Printf("pressing disk, path=%v, size=%dMB, num=%d\n", dp.path, dp.size, dp.num)
	if dir, err := os.Stat(dp.path); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dp.path, os.ModePerm)
			if err != nil {
				fmt.Println("mkdir error:", err)
				return
			}
		} else {
			fmt.Println("open dir error:", err)
			return
		}
	} else if !dir.IsDir() {
		fmt.Println("tmpfile path is not a directory")
		return
	} else {
		dp.deleteTmpfile()
	}

	for i := 0; i < dp.num; i++ {
		filename := dp.prefix + strconv.Itoa(int(rand.Int31()))
		f, err := os.Create(filepath.Join(dp.path, filename))
		if err != nil {
			continue
		}
		for j := 0; j < dp.size; j++ {
			f.Write(zeroFile1MB)
		}
		f.Close()
	}
}

func (dp *DiskPresser) Stop() {
	dp.deleteTmpfile()
}

func (dp *DiskPresser) deleteTmpfile() {
	filepath.Walk(dp.path, func(path string, info fs.FileInfo, err error) error {
		if dp.path != path && info.IsDir() {
			return filepath.SkipDir
		}
		if strings.HasPrefix(info.Name(), dp.prefix) {
			os.Remove(path)
		}
		return nil
	})
}
