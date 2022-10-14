package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderService() *File {

	/*
		package service

		import (
			"fmt"
			"github.com/go-kit/log"
		)
	*/

	file := NewFile("service")

	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")

	/*
		type Service interface {
			Hello(name string) (string, error)
		}
		type service struct {
			logger log.Logger
		}
	*/

	file.Type().Id("Service").Interface(
		Id("Hello").Params(Id("name").String()).Params(String(), Error()),
	)

	file.Type().Id("service").Struct(
		Id("logger").Qual(logPackage, "Logger"),
	)

	/*
		func New(logger log.Logger) Service {
			return service{logger: logger}
		}
	*/

	file.Func().Id("New").Params(Id("logger").Qual(logPackage, "Logger")).Id("Service").Block(
		Return(
			Id("service").Values(Dict{
				Id("logger"): Id("logger"),
			}),
		),
	)

	/*
		func (s service) Hello(name string) (string, error) {
			if name == "" {
				return "", ErrEmptyString
			}

			return fmt.Sprintf("Hello %s", name), nil
		}
	*/

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
