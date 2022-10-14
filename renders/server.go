package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderServer() *File {

	/*
		package cmd

		import (
			"fmt"
			hellov1 "github.com/company/blanksvc/gen/proto/go/hello/v1"
			"github.com/company/blanksvc/pkg/endpoints"
			"github.com/company/blanksvc/pkg/service"
			grpctransport "github.com/company/blanksvc/pkg/transport/grpc"
			httptransport "github.com/company/blanksvc/pkg/transport/http"
			"github.com/company/blanksvc/pkg/utils/healthcheck"
			"github.com/go-kit/log"
			"github.com/gorilla/mux"
			"google.golang.org/grpc"
			"net"
			"net/http"
			"os"
		)

	*/

	file := NewFile("cmd")

	hellov1Package := path.Join(svc.ModulePath, "gen/proto/go/hello/v1")
	file.ImportAlias(hellov1Package, "hellov1")
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

	/*
		func RunServer() error {

			doneC := make(chan error)

			// Init config
			config := NewConfig()

			// Create a single logger, which we'll use and give to other components.
			var logger log.Logger
			logger = log.NewLogfmtLogger(os.Stderr)
			logger = log.With(logger, "ts", log.DefaultTimestampUTC)
			logger = log.With(logger, "caller", log.DefaultCaller)

			// Build the layers of the service "onion" from the inside out
			service := service.New(logger)
			endpoints := endpoints.New(service, logger)
			httpHandler := httptransport.NewHTTPHandler(endpoints, logger)
			grpcServer := grpctransport.NewGRPCServer(endpoints, logger)

			// Configure the HTTP server
			rootMux := mux.NewRouter()
			// Configure health checks
			healthchecker := healthcheck.New()
			healthchecker.AddReadinessChecks(readinessCheck)
			rootMux.Handle(config.HTTPServer.ReadinessEndpoint, healthchecker.ReadinessHandler())
			rootMux.Handle(config.HTTPServer.LivenessEndpoint, healthchecker.LivenessHandler())
			// Configure REST API
			subrouter := rootMux.PathPrefix("/api/v1").Subrouter()
			subrouter.Handle("/hello", httpHandler)
			// Start the HTTP server
			httpServerAddr := fmt.Sprintf("0.0.0.0:%d", config.HTTPServer.Port)
			logger.Log("transport", "HTTP", "addr", httpServerAddr)
			go func() {
				doneC <- http.ListenAndServe(httpServerAddr, rootMux)
			}()

			// Configure the GRPC server
			grpcServerAddr := fmt.Sprintf("0.0.0.0:%d", config.GRPCServer.Port)
			grpcListener, err := net.Listen("tcp", grpcServerAddr)
			if err != nil {
				logger.Log("transport", "gRPC", "during", "Listen", "err", err)
				return err
			}
			baseServer := grpc.NewServer()
			hellov1.RegisterHelloServiceServer(baseServer, grpcServer)
			// Start the GRPC server
			logger.Log("transport", "GRPC", "addr", grpcServerAddr)
			go func() {
				doneC <- baseServer.Serve(grpcListener)
			}()

			// waiting for the errors from servers
			if err := <-doneC; err != nil {
				logger.Log("err", err)
				return err
			}

			return nil
		}
	*/

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
		Qual(hellov1Package, "RegisterHelloServiceServer").Call(Id("baseServer"), Id("grpcServer")),
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

	/*
		func readinessCheck() error {
			return nil
		}
	*/

	file.Func().Id("readinessCheck").Params().Error().Block(
		Return(Nil()),
	)

	return file
}
