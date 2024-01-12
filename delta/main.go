package main

import (
	"encoding/binary"
	"log"
	"math/rand"
	"os"
)

const TOTAL_NUMBERS int = 1000000

var BITMASK = []byte{
	0b00000001,
	0b00000011,
	0b00000111,
	0b00001111,
	0b00011111,
	0b00111111,
	0b01111111,
	0b11111111,
}

var buf [11]byte
var bitShifts []uint8 = []uint8{7, 7, 7, 7, 7, 7, 7, 7, 7, 1}

func getLSB(x byte, n uint8) byte {
	if n > 8 {
		panic("can extract at max 8 bits from the number")
	}
	return byte(x) & BITMASK[n-1]
}

func EncodeUInt(x uint64) []byte {
	var i int = 0
	for i = 0; i < len(bitShifts); i++ {
		buf[i] = getLSB(byte(x), bitShifts[i]) | 0b10000000 // marking the continuation bit
		x = x >> bitShifts[i]
		if x == 0 {
			break
		}
	}
	buf[i] = buf[i] & 0b01111111 // marking the termination bit
	return append(make([]byte, 0, i+1), buf[:i+1]...)
}

func persist(arr []uint64, filename string, shouldVarint bool) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("error creating file:", err)
	}
	defer file.Close()

	for i := range arr {
		var buf []byte
		if shouldVarint {
			buf = EncodeUInt(arr[i])
		} else {
			buf = make([]byte, 8)
			binary.LittleEndian.PutUint64(buf, arr[i])
		}
		file.Write(buf)
	}
}

func main() {
	var numbers []uint64 = make([]uint64, TOTAL_NUMBERS)
	numbers[0] = 10000
	for i := 1; i < TOTAL_NUMBERS; i++ {
		numbers[i] = numbers[i-1] + rand.Uint64()%5
	}
	persist(numbers, "numbers_raw.dat", false)
	persist(numbers, "numbers_varint.dat", true)

	var numbersDelta []uint64 = make([]uint64, len(numbers))
	numbersDelta[0] = numbers[0]
	for i := 1; i < TOTAL_NUMBERS; i++ {
		numbersDelta[i] = numbers[i] - numbers[i-1]
	}
	persist(numbersDelta, "numbers_delta.dat", true)
}
