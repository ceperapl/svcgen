package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHttp() *File {
	file := NewFile("http")

	endpointsPackage := path.Join(svc.ModulePath, "pkg/endpoints")
	file.ImportName(endpointsPackage, "endpoints")
	httptransportPackage := "github.com/go-kit/kit/transport/http"
	file.ImportAlias(httptransportPackage, "httptransport")

	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")
	endpointPackage := "github.com/go-kit/kit/endpoint"
	file.ImportName(endpointPackage, "endpoint")
	transportPackage := "github.com/go-kit/kit/transport"
	file.ImportName(transportPackage, "transport")
	jsonPackage := "encoding/json"
	file.ImportName(jsonPackage, "json")
	httpPackage := "net/http"
	file.ImportName(httpPackage, "http")

	file.Func().Id("NewHTTPHandler").Params(
		Id("endpoints").Qual(endpointsPackage, "Endpoints"),
		Id("logger").Qual(logPackage, "Logger"),
	).Qual(httpPackage, "Handler").Block(
		Id("options").Op(":=").Index().Qual(httptransportPackage, "ServerOption").Values(
			Qual(httptransportPackage, "ServerErrorEncoder").Call(Id("errorEncoder")),
			Qual(httptransportPackage, "ServerErrorHandler").Call(Qual(transportPackage, "NewLogErrorHandler").Call(Id("logger"))),
		),
		Return(
			Qual(httptransportPackage, "NewServer").Call(
				Qual(endpointsPackage, "HelloEndpoint"),
				Id("decodeHTTPHelloRequest"),
				Id("encodeHTTPGenericResponse"),
				Id("options").Op("..."),
			),
		),
	)

	file.Func().Id("decodeHTTPHelloRequest").Params(Id("_").Qual("context", "Context"), Id("r").Op("*").Qual(httpPackage, "Request")).Params(
		Interface(), Error(),
	).Block(
		Id("name").Op(":=").Id("r").Dot("URL").Dot("Query").Call().Dot("Get").Call(Lit("name")),
		Return(Qual(endpointsPackage, "HelloRequest").Values(Dict{
			Id("Name"): Id("name"),
		}), Nil()),
	)

	file.Func().Id("encodeHTTPGenericResponse").Params(
		Id("ctx").Qual("context", "Context"),
		Id("w").Qual(httpPackage, "ResponseWriter"),
		Id("response").Interface(),
	).Error().Block(
		If(
			List(Id("f"), Id("ok")).Op(":=").Id("response").Assert(Qual(endpointPackage, "Failer")),
			Id("ok").Op("&&").Id("f").Dot("Failed").Call().Op("!=").Nil(),
		).Block(
			Id("errorEncoder").Call(Id("ctx"), Id("f").Dot("Failed").Call(), Id("w")),
			Return(Nil()),
		),
		Id("w").Dot("Header").Call().Dot("Set").Call(Lit("Content-Type"), Lit("application/json; charset=utf-8")),
		Return(Qual(jsonPackage, "NewEncoder").Call(Id("w")).Dot("Encode").Call(Id("response"))),
	)

	return file
}
