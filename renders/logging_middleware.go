package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderLoggingMiddleware() *File {
	file := NewFile("middleware")

	endpointPackage := "github.com/go-kit/kit/endpoint"
	file.ImportName(endpointPackage, "endpoint")
	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")

	file.Comment("LoggingMiddleware returns an endpoint middleware that logs the")
	file.Comment("duration of each invocation, and the resulting error, if any.")
	file.Func().Id("LoggingMiddleware").Params(
		Id("logger").Qual(logPackage, "Logger"),
	).Qual(endpointPackage, "Middleware").Block(
		Return(
			Func().Params(Id("next").Qual(endpointPackage, "Endpoint")).Qual(endpointPackage, "Endpoint").Block(
				Return(
					Func().Params(
						Id("ctx").Qual("context", "Context"),
						Id("request").Interface(),
					).Params(Id("response").Interface(), Err().Error()).Block(
						Defer().Func().Params(Id("begin").Qual("time", "Time")).Block(
							Id("logger").Dot("Log").Call(Lit("transport_error"), Err(), Lit("took"), Qual("time", "Since").Call(Id("begin"))),
						).Call(Qual("time", "Now").Call()),
						Return(Id("next").Call(Id("ctx"), Id("request"))),
					),
				),
			),
		),
	)

	return file
}
