package enums

type Operation int

const (
	CreateOperation Operation = iota

	ReadOperation

	UpdateOperation

	DeleteOperation
)
