package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderService() *File {
	file := NewFile("service")

	return file
}
