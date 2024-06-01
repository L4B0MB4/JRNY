package spread

import "testing"

func TestEqualDivideSpace(t *testing.T) {
	firstSpace, remainingSpace := EqualDivideSpace(1)
	if firstSpace != 0x100000000 || remainingSpace != firstSpace {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = EqualDivideSpace(2)
	if firstSpace != 0x80000000 || remainingSpace != firstSpace {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = EqualDivideSpace(3)
	if firstSpace != 0x55555556 || remainingSpace != (firstSpace-1) {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = EqualDivideSpace(4)
	if firstSpace != 0x40000000 || remainingSpace != firstSpace {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
	firstSpace, remainingSpace = EqualDivideSpace(5)
	if firstSpace != 0x33333334 || remainingSpace != (firstSpace-1) {
		t.Error("Dividing up spaces does not create proper space sizes")
	}
}
