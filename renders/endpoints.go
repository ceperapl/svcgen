package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderEndpoints() *File {
	file := NewFile("endpoints")

	middlewarePackage := path.Join(svc.ModulePath, "pkg/endpoints/middleware")
	file.ImportName(middlewarePackage, "middleware")
	servicePackage := path.Join(svc.ModulePath, "pkg/service")
	file.ImportName(servicePackage, "service")

	endpointPackage := "github.com/go-kit/kit/endpoint"
	file.ImportName(endpointPackage, "endpoint")
	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")

	file.Type().Id("HelloRequest").Struct(
		Id("Name").String().Tag(map[string]string{"json": "name"}),
	)

	file.Type().Id("HelloResponse").Struct(
		Id("Greeting").String().Tag(map[string]string{"json": "greeting"}),
	)

	file.Type().Id("Endpoints").Struct(
		Id("HelloEndpoint").Qual(endpointPackage, "Endpoint"),
	)

	file.Func().Id("New").Params(Id("svc").Qual(servicePackage, "Service"), Id("logger").Qual(logPackage, "Logger")).Id("Endpoints").Block(
		Return(
			Id("Endpoints").Values(Dict{
				Id("HelloEndpoint"): Qual(middlewarePackage, "LoggingMiddleware").Call(
					Id("log").Dot("With").Call(
						Id("logger"), Lit("method"), Lit("Hello"),
					),
				).Call(Id("MakeHelloEndpoint").Call(Id("svc"))),
			}),
		),
	)

	// HelloEndpoint: middleware.LoggingMiddleware(log.With(logger, "method", "Hello"))(MakeHelloEndpoint(svc)),

	file.Func().Id("MakeHelloEndpoint").Params(Id("svc").Qual(servicePackage, "Service")).Qual(endpointPackage, "Endpoint").Block(
		Return(
			Func().Params(Id("ctx").Qual("context", "Context"), Id("request").Interface()).Params(
				Id("response").Interface(),
				Err().Error(),
			).Block(
				Id("req").Op(":=").Id("request").Assert(Id("HelloRequest")),
				List(Id("greeting"), Err()).Op(":=").Id("svc").Dot("Hello").Call(Id("req").Dot("Name")),
				If(
					Err().Op("!=").Nil(),
				).Block(
					Return(Nil(), Err()),
				),
				Return(
					Id("HelloResponse").Values(Dict{Id("Greeting"): Id("greeting")}),
					Nil(),
				),
			),
		),
	)

	return file
}

// func MakeHelloEndpoint(svc service.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		req := request.(HelloRequest)
// 		greeting, err := svc.Hello(req.Name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return HelloResponse{Greeting: greeting}, nil
// 	}
// }
