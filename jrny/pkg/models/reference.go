package models

type IdentifierReference struct {
	Self   [16]byte
	Linked []*IdentifierReference
}
