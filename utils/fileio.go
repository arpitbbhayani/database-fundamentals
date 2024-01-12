package utils

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/arpitbbhayani/database-fundamentals/core"
)

func PersistSliceUint64(arr []uint64, filename string, shouldVarint bool) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("error creating file:", err)
	}
	defer file.Close()

	for i := range arr {
		var buf []byte
		if shouldVarint {
			buf = core.EncodeUInt64(arr[i])
		} else {
			buf = make([]byte, 8)
			binary.LittleEndian.PutUint64(buf, arr[i])
		}
		file.Write(buf)
	}
}
