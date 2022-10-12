package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderGrpc() *File {
	file := NewFile("grpc")

	endpointsPackage := path.Join(svc.ModulePath, "pkg/endpoints")
	file.ImportName(endpointsPackage, "endpoints")

	grpctransportPackage := "github.com/go-kit/kit/transport/grpc"
	file.ImportAlias(grpctransportPackage, "grpctransport")
	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")
	transportPackage := "github.com/go-kit/kit/transport"
	file.ImportName(transportPackage, "transport")

	file.Type().Id("grpcServer").Struct(
		Id("hello").Qual(grpctransportPackage, "Handler"),
	)

	file.Func().Id("NewGRPCServer").Params(
		Id("endpoints").Qual(endpointsPackage, "Endpoints"),
		Id("logger").Qual(logPackage, "Logger"),
	).Id("HelloServer").Block(
		Id("options").Op(":=").Index().Qual(grpctransportPackage, "ServerOption").Values(
			Qual(grpctransportPackage, "ServerErrorHandler").Call(Qual(transportPackage, "NewLogErrorHandler").Call(Id("logger"))),
		),
		Line(),
		Return(
			Op("&").Id("grpcServer").Values(Dict{
				Id("hello"): Qual(grpctransportPackage, "NewServer").Call(
					Qual(endpointsPackage, "HelloEndpoint"),
					Id("decodeGRPCHelloRequest"),
					Id("encodeGRPCHelloResponse"),
					Id("options").Op("..."),
				),
			}),
		),
	)

	file.Func().Params(Id("g").Op("*").Id("grpcServer")).Id("Hello").Params(
		Id("ctx").Qual("context", "Context"),
		Id("req").Op("*").Id("HelloRequest"),
	).Params(Op("*").Id("HelloResponse"), Error()).Block(
		List(Id("_"), Id("resp"), Err()).Op(":=").Id("g").Dot("hello").Dot("ServeGRPC").Call(Id("ctx"), Id("req")),
		If(
			Err().Op("!=").Nil(),
		).Block(
			Return(Nil(), Err()),
		),
		Return(Id("resp").Assert(Op("*").Id("HelloResponse")), Nil()),
	)

	file.Func().Id("decodeGRPCHelloRequest").Params(Id("_").Qual("context", "Context"), Id("grpcReq").Interface()).Params(
		Interface(), Error(),
	).Block(
		Id("req").Op(":=").Id("grpcReq").Assert(Op("*").Id("HelloRequest")),
		Return(Qual(endpointsPackage, "HelloRequest").Values(Dict{
			Id("Name"): Id("req").Dot("Name"),
		}), Nil()),
	)

	file.Func().Id("encodeGRPCHelloResponse").Params(Id("_").Qual("context", "Context"), Id("response").Interface()).Params(
		Interface(), Error(),
	).Block(
		Id("resp").Op(":=").Id("response").Assert(Qual(endpointsPackage, "HelloResponse")),
		Return(Op("&").Id("HelloResponse").Values(Dict{
			Id("Greeting"): Id("resp").Dot("Greeting"),
		}), Nil()),
	)

	return file
}
