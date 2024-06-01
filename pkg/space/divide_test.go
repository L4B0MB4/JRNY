package space

import (
	"math/big"
	"testing"
)

func TestEqualDivideSpace(t *testing.T) {
	MaxSpaceSize = big.NewInt(0xFFFFFFFF + 1)
	firstSpace, remainingSpace := equalDivideSpace(1)
	if firstSpace.Int64() != 0x100000000 || remainingSpace.Int64() != firstSpace.Int64() {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = equalDivideSpace(2)
	if firstSpace.Int64() != 0x80000000 || remainingSpace.Int64() != firstSpace.Int64() {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = equalDivideSpace(3)
	if firstSpace.Int64() != 0x55555556 || remainingSpace.Int64() != (firstSpace.Int64()-1) {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = equalDivideSpace(4)
	if firstSpace.Int64() != 0x40000000 || remainingSpace.Int64() != firstSpace.Int64() {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = equalDivideSpace(5)
	if firstSpace.Int64() != 0x33333334 || remainingSpace.Int64() != (firstSpace.Int64()-1) {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
}
