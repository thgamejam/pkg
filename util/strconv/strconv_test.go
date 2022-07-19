package strconv

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestStrconvItoa(t *testing.T) {
	var i int = int((^uint(0)) >> 1)
	si := Itoa(i)
	assert.Equal(t, strconv.FormatInt(int64(i), 10), si)

	var i8 int8 = 127
	si8 := Itoa(i8)
	assert.Equal(t, "127", si8)

	var i16 int16 = 32767
	si16 := Itoa(i16)
	assert.Equal(t, "32767", si16)

	var i32 int32 = 2147483647
	si32 := Itoa(i32)
	assert.Equal(t, "2147483647", si32)

	var i64 int64 = 9223372036854775807
	si64 := Itoa(i64)
	assert.Equal(t, "9223372036854775807", si64)
}

func TestStrconvParse(t *testing.T) {
	var i int = int((^uint(0)) >> 1)
	var vi int
	si := strconv.FormatInt(int64(i), 10)
	err := Parse(si, &vi)
	assert.NoError(t, err)
	assert.Equal(t, i, vi)

	var i8 int8 = 127
	var vi8 int8
	si8 := strconv.FormatInt(int64(i8), 10)
	err = Parse(si8, &vi8)
	assert.NoError(t, err)
	assert.Equal(t, i8, vi8)

	var i16 int16 = 32767
	var vi16 int16
	si16 := strconv.FormatInt(int64(i16), 10)
	err = Parse(si16, &vi16)
	assert.NoError(t, err)
	assert.Equal(t, i16, vi16)

	var i32 int32 = 2147483647
	var vi32 int32
	si32 := strconv.FormatInt(int64(i32), 10)
	err = Parse(si32, &vi32)
	assert.NoError(t, err)
	assert.Equal(t, i32, vi32)

	var i64 int64 = 9223372036854775807
	var vi64 int64
	si64 := strconv.FormatInt(i64, 10)
	err = Parse(si64, &vi64)
	assert.NoError(t, err)
	assert.Equal(t, i64, vi64)
}

func TestStrconvUItoa(t *testing.T) {
	var i uint = ^uint(0)
	si := UItoa(i)
	assert.Equal(t, strconv.FormatUint(uint64(i), 10), si)

	var i8 uint8 = 255
	si8 := UItoa(i8)
	assert.Equal(t, "255", si8)

	var i16 uint16 = 65535
	si16 := UItoa(i16)
	assert.Equal(t, "65535", si16)

	var i32 uint32 = 4294967295
	si32 := UItoa(i32)
	assert.Equal(t, "4294967295", si32)

	var i64 uint64 = 18446744073709551615
	si64 := UItoa(i64)
	assert.Equal(t, "18446744073709551615", si64)
}

func TestStrconvUParse(t *testing.T) {
	var i uint = ^uint(0)
	var vi uint
	si := strconv.FormatUint(uint64(i), 10)
	err := UParse(si, &vi)
	assert.NoError(t, err)
	assert.Equal(t, i, vi)

	var i8 uint8 = 255
	var vi8 uint8
	si8 := strconv.FormatUint(uint64(i8), 10)
	err = UParse(si8, &vi8)
	assert.NoError(t, err)
	assert.Equal(t, i8, vi8)

	var i16 uint16 = 65535
	var vi16 uint16
	si16 := strconv.FormatUint(uint64(i16), 10)
	err = UParse(si16, &vi16)
	assert.NoError(t, err)
	assert.Equal(t, i16, vi16)

	var i32 uint32 = 4294967295
	var vi32 uint32
	si32 := strconv.FormatUint(uint64(i32), 10)
	err = UParse(si32, &vi32)
	assert.NoError(t, err)
	assert.Equal(t, i32, vi32)

	var i64 uint64 = 18446744073709551615
	var vi64 uint64
	si64 := strconv.FormatUint(i64, 10)
	err = UParse(si64, &vi64)
	assert.NoError(t, err)
	assert.Equal(t, i64, vi64)
}
