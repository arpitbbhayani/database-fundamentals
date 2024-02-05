package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const FILENAME = "data.txt"
const FILE_SIZE = 1 * 1024 * 1024 * 1024
const MB1 = 1 * 1024 * 1024

func setupFile() {

	if _, err := os.Stat(FILENAME); err == nil {
		return
	}

	file, err := os.OpenFile(FILENAME, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < FILE_SIZE; i += MB1 {
		buffer := make([]byte, MB1)
		_, err := rand.Read(buffer)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(buffer)
		if err != nil {
			panic(err)
		}
	}
}

type stat struct {
	random     int64
	sequential int64
}

type paramters struct {
	name     string
	pageSize int
	stats    *stat
}

var params = []paramters{
	{"1KB", 1024, &stat{0, 0}},
	{"2KB", 2 * 1024, &stat{0, 0}},
	{"4KB", 4 * 1024, &stat{0, 0}},
	{"8KB", 8 * 1024, &stat{0, 0}},
	{"16KB", 16 * 1024, &stat{0, 0}},
	{"32KB", 32 * 1024, &stat{0, 0}},
	{"64KB", 64 * 1024, &stat{0, 0}},
	{"128KB", 128 * 1024, &stat{0, 0}},
	{"256KB", 256 * 1024, &stat{0, 0}},
	{"512KB", 512 * 1024, &stat{0, 0}},
	{"1024KB", 1024 * 1024, &stat{0, 0}},
}

func benchmark() {
	setupFile()
	fp, err := os.OpenFile(FILENAME, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	for p := range params {
		var buf []byte = make([]byte, params[p].pageSize)

		st := time.Now()
		for i := 0; i < FILE_SIZE/params[p].pageSize; i++ {
			r := rand.Int() % FILE_SIZE
			offset := r - r%params[p].pageSize
			fp.Seek(int64(offset), 0)
			fp.Read(buf)
		}
		params[p].stats.random = time.Since(st).Milliseconds()
	}

	for p := range params {
		fp.Seek(0, 0)
		var buf []byte = make([]byte, params[p].pageSize)

		st := time.Now()
		for i := 0; i < int(FILE_SIZE/params[p].pageSize); i++ {
			fp.Read(buf)
		}
		params[p].stats.sequential = time.Since(st).Milliseconds()
	}
}

func main() {
	benchmark()
	fmt.Printf("buf_size,random (ms), sequential (ms)\n")
	for p := range params {
		fmt.Printf("%s,%d,%d\n", params[p].name, params[p].stats.random, params[p].stats.sequential)
	}
}
