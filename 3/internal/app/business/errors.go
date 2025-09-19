package business

// Allowed values for Error
const (
	ErrEmptyOrder     Error = "EMPTY_ORDER"
	ErrInvalidProduct Error = "INVALID_PRODUCT"
	ErrInvalidStratum Error = "INVALID_STRATUM"
)

// Error describes a business error
type Error string

func (e Error) Error() string {
	const errorPrefix = "ERROR_"
	return errorPrefix + string(e)
}
