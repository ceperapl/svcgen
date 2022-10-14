package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderGrpc() *File {

	/*
		package grpc

		import (
			"context"
			hellov1 "github.com/company/blanksvc/gen/proto/go/hello/v1"
			"github.com/company/blanksvc/pkg/endpoints"
			"github.com/go-kit/kit/transport"
			grpctransport "github.com/go-kit/kit/transport/grpc"
			"github.com/go-kit/log"
		)
	*/

	file := NewFile("grpc")

	hellov1Package := path.Join(svc.ModulePath, "gen/proto/go/hello/v1")
	file.ImportAlias(hellov1Package, "hellov1")
	endpointsPackage := path.Join(svc.ModulePath, "pkg/endpoints")
	file.ImportName(endpointsPackage, "endpoints")

	grpctransportPackage := "github.com/go-kit/kit/transport/grpc"
	file.ImportAlias(grpctransportPackage, "grpctransport")
	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")
	transportPackage := "github.com/go-kit/kit/transport"
	file.ImportName(transportPackage, "transport")

	/*
		type grpcServer struct {
			hello grpctransport.Handler
		}
	*/

	file.Type().Id("grpcServer").Struct(
		Id("hello").Qual(grpctransportPackage, "Handler"),
	)

	/*
		func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) hellov1.HelloServiceServer {
			options := []grpctransport.ServerOption{grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger))}

			return &grpcServer{hello: grpctransport.NewServer(endpoints.HelloEndpoint, decodeGRPCHelloRequest, encodeGRPCHelloResponse, options...)}
		}
	*/

	file.Func().Id("NewGRPCServer").Params(
		Id("endpoints").Qual(endpointsPackage, "Endpoints"),
		Id("logger").Qual(logPackage, "Logger"),
	).Qual(hellov1Package, "HelloServiceServer").Block(
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

	/*
		func (g *grpcServer) Hello(ctx context.Context, req *hellov1.HelloRequest) (*hellov1.HelloResponse, error) {
			_, resp, err := g.hello.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.(*hellov1.HelloResponse), nil
		}
	*/

	file.Func().Params(Id("g").Op("*").Id("grpcServer")).Id("Hello").Params(
		Id("ctx").Qual("context", "Context"),
		Id("req").Op("*").Qual(hellov1Package, "HelloRequest"),
	).Params(Op("*").Qual(hellov1Package, "HelloResponse"), Error()).Block(
		List(Id("_"), Id("resp"), Err()).Op(":=").Id("g").Dot("hello").Dot("ServeGRPC").Call(Id("ctx"), Id("req")),
		If(
			Err().Op("!=").Nil(),
		).Block(
			Return(Nil(), Err()),
		),
		Return(Id("resp").Assert(Op("*").Qual(hellov1Package, "HelloResponse")), Nil()),
	)

	/*
		func decodeGRPCHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
			req := grpcReq.(*hellov1.HelloRequest)
			return endpoints.HelloRequest{Name: req.Name}, nil
		}
	*/

	file.Func().Id("decodeGRPCHelloRequest").Params(Id("_").Qual("context", "Context"), Id("grpcReq").Interface()).Params(
		Interface(), Error(),
	).Block(
		Id("req").Op(":=").Id("grpcReq").Assert(Op("*").Qual(hellov1Package, "HelloRequest")),
		Return(Qual(endpointsPackage, "HelloRequest").Values(Dict{
			Id("Name"): Id("req").Dot("Name"),
		}), Nil()),
	)

	/*
		func encodeGRPCHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
			resp := response.(endpoints.HelloResponse)
			return &hellov1.HelloResponse{Greeting: resp.Greeting}, nil
		}
	*/

	file.Func().Id("encodeGRPCHelloResponse").Params(Id("_").Qual("context", "Context"), Id("response").Interface()).Params(
		Interface(), Error(),
	).Block(
		Id("resp").Op(":=").Id("response").Assert(Qual(endpointsPackage, "HelloResponse")),
		Return(Op("&").Qual(hellov1Package, "HelloResponse").Values(Dict{
			Id("Greeting"): Id("resp").Dot("Greeting"),
		}), Nil()),
	)

	return file
}
