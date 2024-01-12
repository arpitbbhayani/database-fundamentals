package varint

import (
	"testing"

	"github.com/arpitbbhayani/database-fundamentals/utils"
)

func TestMain(t *testing.T) {
	t.Log(utils.EncodeUInt64(123))
	t.Log(utils.EncodeUInt64(292))
}
