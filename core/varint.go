package core

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

// getLSB returns the least significant `n` bits from
// the byte value `x`.
func getLSB(x byte, n uint8) byte {
	if n > 8 {
		panic("can extract at max 8 bits from the number")
	}
	return byte(x) & BITMASK[n-1]
}

// TOOD: not thread safe
var buf [11]byte
var bitShifts []uint8 = []uint8{7, 7, 7, 7, 7, 7, 7, 7, 7, 1}

// EncodeInt64 encodes the unsigned 64 bit integer value into a varint
// and returns an array of bytes (little endian encoded)
func EncodeUInt64(x uint64) []byte {
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

// DecodeUInt64 decodes the array of bytes and returns an unsigned 64 bit integer
func DecodeUInt64(vint []byte) uint64 {
	var i int = 0
	var v uint64 = 0
	for i = 0; i < len(vint); i++ {
		b := getLSB(vint[i], 7)
		v = v | uint64(b)<<(7*i)
	}
	return v
}
