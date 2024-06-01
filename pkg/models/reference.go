package models

type IdentifierReference struct {
	Self   [16]byte
	Linked map[[16]byte]*IdentifierReference
}
