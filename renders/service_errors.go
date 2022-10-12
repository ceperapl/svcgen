package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderServiceErrors() *File {
	file := NewFile("service")

	file.Var().Defs(
		Id("ErrEmptyString").Op("=").Qual("errors", "New").Call(Lit("empty string")),
	)

	return file
}
