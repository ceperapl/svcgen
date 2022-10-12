package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderService() *File {
	file := NewFile("service")

	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")

	file.Type().Id("Service").Interface(
		Id("Hello").Params(Id("name").String()).Params(String(), Error()),
	)

	file.Type().Id("service").Struct(
		Id("logger").Qual(logPackage, "Logger"),
	)

	file.Func().Id("New").Params(Id("logger").Qual(logPackage, "Logger")).Id("Service").Block(
		Return(
			Id("service").Values(Dict{
				Id("logger"): Id("logger"),
			}),
		),
	)

	file.Func().Params(Id("s").Id("service")).Id("Hello").Params(Id("name").String()).Params(String(), Error()).Block(
		If(
			Id("name").Op("==").Lit(""),
		).Block(
			Return(Lit(""), Id("ErrEmptyString")),
		),
		Line(),
		Return(Qual("fmt", "Sprintf").Call(Lit("Hello %s"), Id("name")), Nil()),
	)

	return file
}
