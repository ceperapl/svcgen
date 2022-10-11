package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHealthcheck() *File {
	file := NewFile("healthcheck")

	return file
}
