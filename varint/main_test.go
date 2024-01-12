package varint

import (
	"testing"

	"github.com/arpitbbhayani/database-fundamentals/core"
)

func TestMain(t *testing.T) {
	t.Log(core.EncodeUInt64(123))
	t.Log(core.EncodeUInt64(292))
}
