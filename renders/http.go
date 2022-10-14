package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHttp() *File {

	/*
		package http

		import (
			"context"
			"encoding/json"
			"github.com/company/blanksvc/pkg/endpoints"
			"github.com/go-kit/kit/endpoint"
			"github.com/go-kit/kit/transport"
			httptransport "github.com/go-kit/kit/transport/http"
			"github.com/go-kit/log"
			"net/http"
		)
	*/

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

	/*
		func NewHTTPHandler(endpoints endpoints.Endpoints, logger log.Logger) http.Handler {
			options := []httptransport.ServerOption{httptransport.ServerErrorEncoder(errorEncoder), httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger))}
			return httptransport.NewServer(endpoints.HelloEndpoint, decodeHTTPHelloRequest, encodeHTTPGenericResponse, options...)
		}
	*/

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

	/*
		func decodeHTTPHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
			name := r.URL.Query().Get("name")
			return endpoints.HelloRequest{Name: name}, nil
		}
	*/

	file.Func().Id("decodeHTTPHelloRequest").Params(Id("_").Qual("context", "Context"), Id("r").Op("*").Qual(httpPackage, "Request")).Params(
		Interface(), Error(),
	).Block(
		Id("name").Op(":=").Id("r").Dot("URL").Dot("Query").Call().Dot("Get").Call(Lit("name")),
		Return(Qual(endpointsPackage, "HelloRequest").Values(Dict{
			Id("Name"): Id("name"),
		}), Nil()),
	)

	/*
		func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
				errorEncoder(ctx, f.Failed(), w)
				return nil
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			return json.NewEncoder(w).Encode(response)
		}
	*/

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
