package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderEndpoints() *File {
	file := NewFile("endpoints")

	return file
}
