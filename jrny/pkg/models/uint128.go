package models

import "encoding/binary"

type UInt128 struct {
	high uint64
	low  uint64
}

func (u *UInt128) Set(value [16]byte) {
	u.high = binary.BigEndian.Uint64(value[0:8])
	u.low = binary.BigEndian.Uint64(value[8:])
}

func (u *UInt128) SmallerEqual(value *UInt128) bool {
	return u.Smaller(value) || u.Equal(value)
}

func (u *UInt128) Equal(value *UInt128) bool {
	if value.high == u.high && value.low == u.low {
		return true
	}
	return false
}

func (u *UInt128) Smaller(value *UInt128) bool {
	if value.high == u.high && u.low < value.low {
		return true
	}
	if u.high < value.high {
		return true
	}
	return false
}

func (u *UInt128) BiggerEqual(value *UInt128) bool {
	return u.Bigger(value) || u.Equal(value)

}
func (u *UInt128) Bigger(value *UInt128) bool {

	if u.high == value.high && u.low > value.low {
		return true
	}
	if u.high > value.high {
		return true
	}
	return false

}
