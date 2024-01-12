package delta

import (
	"math/rand"
	"testing"

	"github.com/arpitbbhayani/database-fundamentals/utils"
)

const TOTAL_NUMBERS int = 1000000

func TestMain(t *testing.T) {
	var numbers []uint64 = make([]uint64, TOTAL_NUMBERS)
	numbers[0] = 10000
	for i := 1; i < TOTAL_NUMBERS; i++ {
		numbers[i] = numbers[i-1] + rand.Uint64()%5
	}
	utils.PersistSliceUint64(numbers, "numbers_raw.dat", false)
	utils.PersistSliceUint64(numbers, "numbers_varint.dat", true)

	var numbersDelta []uint64 = make([]uint64, len(numbers))
	numbersDelta[0] = numbers[0]
	for i := 1; i < TOTAL_NUMBERS; i++ {
		numbersDelta[i] = numbers[i] - numbers[i-1]
	}
	utils.PersistSliceUint64(numbersDelta, "numbers_delta.dat", true)
}
