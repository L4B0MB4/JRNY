package models

type IdentifierReference struct {
	Self   [16]byte
	Linked [][16]byte
}
