package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderServer() *File {
	file := NewFile("cmd")

	endpointsPackage := path.Join(svc.ModulePath, "pkg/endpoints")
	file.ImportName(endpointsPackage, "endpoints")
	servicePackage := path.Join(svc.ModulePath, "pkg/service")
	file.ImportName(servicePackage, "service")
	grpctransportPackage := path.Join(svc.ModulePath, "pkg/transport/grpc")
	file.ImportAlias(grpctransportPackage, "grpctransport")
	httptransportPackage := path.Join(svc.ModulePath, "pkg/transport/http")
	file.ImportAlias(httptransportPackage, "httptransport")
	healthcheckPackage := path.Join(svc.ModulePath, "pkg/utils/healthcheck")
	file.ImportName(healthcheckPackage, "healthcheck")

	logPackage := "github.com/go-kit/log"
	file.ImportName(logPackage, "log")
	muxPackage := "github.com/gorilla/mux"
	file.ImportName(muxPackage, "mux")
	grpcPackage := "google.golang.org/grpc"
	file.ImportName(grpcPackage, "grpc")
	httpPackage := "net/http"
	file.ImportName(httpPackage, "http")

	file.Func().Id("RunServer").Params().Error().Block(
		Line(),
		Id("doneC").Op(":=").Make(Chan().Error()),
		Line(),
		Comment("Init config"),
		Id("config").Op(":=").Id("NewConfig").Call(),
		Line(),
		Comment("Create a single logger, which we'll use and give to other components."),
		Var().Id("logger").Qual(logPackage, "Logger"),
		Id("logger").Op("=").Qual(logPackage, "NewLogfmtLogger").Call(Qual("os", "Stderr")),
		Id("logger").Op("=").Qual(logPackage, "With").Call(Id("logger"), Lit("ts"), Qual(logPackage, "DefaultTimestampUTC")),
		Id("logger").Op("=").Qual(logPackage, "With").Call(Id("logger"), Lit("caller"), Qual(logPackage, "DefaultCaller")),
		Line(),
		Comment(`Build the layers of the service "onion" from the inside out`),
		Id("service").Op(":=").Qual(servicePackage, "New").Call(Id("logger")),
		Id("endpoints").Op(":=").Qual(endpointsPackage, "New").Call(Id("service"), Id("logger")),
		Id("httpHandler").Op(":=").Qual(httptransportPackage, "NewHTTPHandler").Call(Id("endpoints"), Id("logger")),
		Id("grpcServer").Op(":=").Qual(grpctransportPackage, "NewGRPCServer").Call(Id("endpoints"), Id("logger")),
		Line(),
		Comment("Configure the HTTP server"),
		Id("rootMux").Op(":=").Qual(muxPackage, "NewRouter").Call(),
		Comment("Configure health checks"),
		Id("healthchecker").Op(":=").Qual(healthcheckPackage, "New").Call(),
		Id("healthchecker").Dot("AddReadinessChecks").Call(Id("readinessCheck")),
		Id("rootMux").Dot("Handle").Call(Id("config").Dot("HTTPServer").Dot("ReadinessEndpoint"), Id("healthchecker").Dot("ReadinessHandler").Call()),
		Id("rootMux").Dot("Handle").Call(Id("config").Dot("HTTPServer").Dot("LivenessEndpoint"), Id("healthchecker").Dot("LivenessHandler").Call()),
		Comment("Configure REST API"),
		Id("subrouter").Op(":=").Id("rootMux").Dot("PathPrefix").Call(Lit("/api/v1")).Dot("Subrouter").Call(),
		Id("subrouter").Dot("Handle").Call(Lit("/hello"), Id("httpHandler")),
		Comment("Start the HTTP server"),
		Id("httpServerAddr").Op(":=").Qual("fmt", "Sprintf").Call(Lit("0.0.0.0:%d"), Id("config").Dot("HTTPServer").Dot("Port")),
		Id("logger").Dot("Log").Call(Lit("transport"), Lit("HTTP"), Lit("addr"), Id("httpServerAddr")),
		Go().Func().Params().Block(
			Id("doneC").Op("<-").Qual(httpPackage, "ListenAndServe").Call(Id("httpServerAddr"), Id("rootMux")),
		).Call(),
		Line(),
		Comment("Configure the GRPC server"),
		Id("grpcServerAddr").Op(":=").Qual("fmt", "Sprintf").Call(Lit("0.0.0.0:%d"), Id("config").Dot("GRPCServer").Dot("Port")),
		List(Id("grpcListener"), Err()).Op(":=").Qual("net", "Listen").Call(Lit("tcp"), Id("grpcServerAddr")),
		If(
			Err().Op("!=").Nil(),
		).Block(
			Id("logger").Dot("Log").Call(Lit("transport"), Lit("gRPC"), Lit("during"), Lit("Listen"), Lit("err"), Err()),
			Return(Err()),
		),
		Id("baseServer").Op(":=").Qual(grpcPackage, "NewServer").Call(),
		Qual(grpctransportPackage, "RegisterHelloServer").Call(Id("baseServer"), Id("grpcServer")),
		Comment("Start the GRPC server"),
		Id("logger").Dot("Log").Call(Lit("transport"), Lit("GRPC"), Lit("addr"), Id("grpcServerAddr")),
		Go().Func().Params().Block(
			Id("doneC").Op("<-").Id("baseServer").Dot("Serve").Call(Id("grpcListener")),
		).Call(),
		Line(),
		Comment("waiting for the errors from servers"),
		If(
			Err().Op(":=").Op("<-").Id("doneC"),
			Err().Op("!=").Nil(),
		).Block(
			Id("logger").Dot("Log").Call(Lit("err"), Err()),
			Return(Err()),
		),
		Line(),
		Return(Nil()),
	)

	file.Func().Id("readinessCheck").Params().Error().Block(
		Return(Nil()),
	)

	return file
}
