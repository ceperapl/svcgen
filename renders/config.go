package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderConfig() *File {
	file := NewFilePathName("config", "cmd")

	viperPackage := "github.com/spf13/viper"
	file.ImportName(viperPackage, "viper")
	pflagPackage := "github.com/spf13/pflag"
	file.ImportName(pflagPackage, "pflag")

	file.Const().Defs(
		Id("envPrefix").Op("=").Lit("BLANK"),
		Line(),
		Id("httpServerPortEnv").Op("=").Lit("HTTP_SERVER_PORT"),
		Id("httpServerPortDefault").Op("=").Lit(8080),
		Line(),
		Id("httpReadinessEndpointEnv").Op("=").Lit("HTTP_READINESS_ENDPOINT"),
		Id("httpReadinessEndpointDefault").Op("=").Lit("/ready"),
		Line(),
		Id("httpLivenessEndpointEnv").Op("=").Lit("HTTP_LIVENESS_ENDPOINT"),
		Id("httpLivenessEndpointDefault").Op("=").Lit("/health"),
		Line(),
		Id("grpcServerPortEnv").Op("=").Lit("GRPC_SERVER_PORT"),
		Id("grpcServerPortDefault").Op("=").Lit("9000"),
	)

	file.Type().Id("Config").Struct(
		Id("HTTPServer").Struct(
			Id("Port").Uint16(),
			Id("ReadinessEndpoint").String(),
			Id("LivenessEndpoint").String(),
		),
		Id("GRPCServer").Struct(
			Id("Port").Uint16(),
		),
	)

	file.Func().Id("NewConfig").Params().Id("Config").Block(
		Id("config").Op(":=").Id("Config").Values(),
		Line(),
		Comment("Init config via env variables"),
		Qual(viperPackage, "SetEnvPrefix").Call(Id("envPrefix")),
		Line(),
		Qual(viperPackage, "BindEnv").Call(Id("httpServerPortEnv")),
		Qual(viperPackage, "SetDefault").Call(Id("httpServerPortEnv"), Id("httpServerPortDefault")),
		Id("config").Dot("HTTPServer").Dot("Port").Op("=").Qual(viperPackage, "GetUint16").Call(Id("httpServerPortEnv")),
		Line(),
		Qual(viperPackage, "BindEnv").Call(Id("httpReadinessEndpointEnv")),
		Qual(viperPackage, "SetDefault").Call(Id("httpReadinessEndpointEnv"), Id("httpReadinessEndpointDefault")),
		Id("config").Dot("HTTPServer").Dot("ReadinessEndpoint").Op("=").Qual(viperPackage, "GetString").Call(Id("httpReadinessEndpointEnv")),
		Line(),
		Qual(viperPackage, "BindEnv").Call(Id("httpLivenessEndpointEnv")),
		Qual(viperPackage, "SetDefault").Call(Id("httpLivenessEndpointEnv"), Id("httpLivenessEndpointDefault")),
		Id("config").Dot("HTTPServer").Dot("LivenessEndpoint").Op("=").Qual(viperPackage, "GetString").Call(Id("httpLivenessEndpointEnv")),
		Line(),
		Qual(viperPackage, "BindEnv").Call(Id("grpcServerPortEnv")),
		Qual(viperPackage, "SetDefault").Call(Id("grpcServerPortEnv"), Id("grpcServerPortDefault")),
		Id("config").Dot("GRPCServer").Dot("Port").Op("=").Qual(viperPackage, "GetUint16").Call(Id("grpcServerPortEnv")),
		Line(),
		Comment("Init config via flags"),
		Qual(pflagPackage, "Uint16Var").Call(
			Op("&").Id("config").Dot("HTTPServer").Dot("Port"),
			Lit("httpserver.port"),
			Id("config").Dot("HTTPServer").Dot("Port"),
			Qual("fmt", "Sprintf").Call(Lit("HTTP Server port; env: %s"), Id("httpServerPortEnv")),
		),
		Qual(pflagPackage, "StringVar").Call(
			Op("&").Id("config").Dot("HTTPServer").Dot("ReadinessEndpoint"),
			Lit("httpserver.readiness"),
			Id("config").Dot("HTTPServer").Dot("ReadinessEndpoint"),
			Qual("fmt", "Sprintf").Call(Lit("HTTP Server readiness endpoint name; env: %s"), Id("httpReadinessEndpointEnv")),
		),
		Qual(pflagPackage, "StringVar").Call(
			Op("&").Id("config").Dot("HTTPServer").Dot("LivenessEndpoint"),
			Lit("httpserver.liveness"),
			Id("config").Dot("HTTPServer").Dot("LivenessEndpoint"),
			Qual("fmt", "Sprintf").Call(Lit("HTTP Server liveness endpoint name; env: %s"), Id("httpLivenessEndpointEnv")),
		),
		Qual(pflagPackage, "Uint16Var").Call(
			Op("&").Id("config").Dot("GRPCServer").Dot("Port"),
			Lit("grpcserver.port"),
			Id("config").Dot("GRPCServer").Dot("Port"),
			Qual("fmt", "Sprintf").Call(Lit("gRPC Server port; env: %s"), Id("grpcServerPortEnv")),
		),
		Line(),
		Qual(pflagPackage, "Parse").Call(),
		Line(),
		Return(Id("config")),
	)

	return file
}
