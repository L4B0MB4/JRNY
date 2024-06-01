package models

import (
	"testing"

	"github.com/google/uuid"
)

func TestEqual(t *testing.T) {
	uu := uuid.New()
	u := UInt128{}
	v := UInt128{}
	u.Set(uu)
	v.Set(uu)
	if !u.Equal(&v) {
		t.Error("u should be the same as v")
	}
}

func TestSmaller(t *testing.T) {

	u := UInt128{}
	u.Set(uuid.Max)
	v := UInt128{}
	uu, _ := uuid.FromBytes(make([]byte, 16))
	v.Set(uu)
	if u.Smaller(&v) {
		t.Error("u should be bigger than v")
	}
	if !v.Smaller(&u) {
		t.Error("v should be smaller than u")
	}
}

func TestBiggerEqual(t *testing.T) {
	u := UInt128{}
	uu := uuid.New()
	uu[15] = 1
	u.Set(uu)
	uu[15] = 0
	v := UInt128{}
	uu, _ = uuid.FromBytes(make([]byte, 16))
	v.Set(uu)
	if !u.Bigger(&v) {
		t.Error("u should be bigger than v")
	}
	if v.Bigger(&u) {
		t.Error("v should be smaller than u")
	}
	u.Set(uu)
	if !v.BiggerEqual(&u) {
		t.Error("v should be smaller than u")
	}
}
func TestBigger(t *testing.T) {

	u := UInt128{}
	u.Set(uuid.Max)
	v := UInt128{}
	uu, _ := uuid.FromBytes(make([]byte, 16))
	v.Set(uu)
	if !u.Bigger(&v) {
		t.Error("u should be bigger than v")
	}
	if v.Bigger(&u) {
		t.Error("v should be smaller than u")
	}
}
