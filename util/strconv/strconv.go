package strconv

import "strconv"

type Int interface {
	int | int8 | int16 | int32 | int64
}

// Itoa is equivalent to strconv.FormatInt(int64(i), 10).
func Itoa[T Int](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

// Parse is equivalent to strconv.ParseInt(s, 10, 32).
func Parse[T Int](s string, i *T) error {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*i = T(n)
	return nil
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// UItoa is equivalent to strconv.FormatUint(uint64(i), 10).
func UItoa[T UInt](i T) string {
	return strconv.FormatUint(uint64(i), 10)
}

// ParseUint32
// Deprecated: Use strconv.UParse instead.
func ParseUint32(s string) (uint32, error) {
	var n uint32
	err := UParse(s, &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// UParse is equivalent to strconv.ParseUint(s, 10, 32).
func UParse[T UInt](s string, i *T) error {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}

	*i = T(n)
	return nil
}
