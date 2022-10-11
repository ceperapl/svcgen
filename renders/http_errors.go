package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHTTPErrors() *File {
	file := NewFilePathName("errors", "http")

	return file
}
