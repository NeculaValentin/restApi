package common

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	BadRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "BAD_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Bad Request", "Unauthorized"}[r-1]
}
