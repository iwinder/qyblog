package errors

import "net/http"

var (
	unknownCoder defaultCoder = defaultCoder{1, http.StatusInternalServerError, "An internal server error occurred", "http://github.com/marmotedu/errors/README.md"}
)

type Coder interface {
	// HTTP status that should be used for the associated error code.
	HTTPStatus() int

	// External (user) facing error text.
	String() string

	// Reference returns the detail documents for user.
	Reference() string

	// Code returns the code of the coder
	Code() int
}

type defaultCoder struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

func (coder defaultCoder) Code() int {
	return coder.C

}

func (coder defaultCoder) String() string {
	return coder.Ext
}
func (coder defaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}

	return coder.HTTP
}

func (coder defaultCoder) Reference() string {
	return coder.Ref
}

var codes = map[int]Coder{}

func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}

	return unknownCoder
}
