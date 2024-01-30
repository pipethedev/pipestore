package enums

type ResponseCode int

const (
	Successful ResponseCode = 211

	BadRequest ResponseCode = 411

	InternalServerError ResponseCode = 511
)
