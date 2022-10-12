package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHealthcheck() *File {
	file := NewFile("healthcheck")

	httpPackage := "net/http"
	file.ImportName(httpPackage, "http")

	file.Type().Id("Check").Func().Params().Error()

	file.Type().Id("HealthChecker").Interface(
		Id("AddLivenessChecks").Params(Id("check").Id("Check")),
		Id("AddReadinessChecks").Params(Id("check").Id("Check")),
		Id("LivenessHandler").Params().Qual(httpPackage, "HandlerFunc"),
		Id("ReadinessHandler").Params().Qual(httpPackage, "HandlerFunc"),
	)

	file.Func().Id("New").Params().Id("HealthChecker").Block(
		Return(Op("&").Id("healthcheck")).Values(),
	)

	file.Type().Id("healthcheck").Struct(
		Id("lock").Qual("sync", "RWMutex"),
		Line(),
		Id("livenessChecks").Index().Id("Check"),
		Id("readinessChecks").Index().Id("Check"),
	)

	file.Func().Params(Id("h").Op("*").Id("healthcheck")).Id("LivenessHandler").Params().Qual(httpPackage, "HandlerFunc").Block(
		Return(
			Func().Params(
				Id("rw").Qual(httpPackage, "ResponseWriter"),
				Id("r").Op("*").Qual(httpPackage, "Request"),
			).Block(
				If(
					Err().Op(":=").Id("h").Dot("checkLiveness").Call(),
					Err().Op("!=").Nil(),
				).Block(
					Id("rw").Dot("WriteHeader").Call(Qual(httpPackage, "StatusServiceUnavailable")),
					Return(),
				),
				Id("rw").Dot("WriteHeader").Call(Qual(httpPackage, "StatusOK")),
			),
		),
	)

	file.Func().Params(Id("h").Op("*").Id("healthcheck")).Id("ReadinessHandler").Params().Qual(httpPackage, "HandlerFunc").Block(
		Return(
			Func().Params(
				Id("rw").Qual(httpPackage, "ResponseWriter"),
				Id("r").Op("*").Qual(httpPackage, "Request"),
			).Block(
				If(
					Err().Op(":=").Id("h").Dot("checkReadiness").Call(),
					Err().Op("!=").Nil(),
				).Block(
					Id("rw").Dot("WriteHeader").Call(Qual(httpPackage, "StatusServiceUnavailable")),
					Return(),
				),
				Id("rw").Dot("WriteHeader").Call(Qual(httpPackage, "StatusOK")),
			),
		),
	)

	file.Func().Params(Id("h").Op("*").Id("healthcheck")).Id("AddLivenessChecks").Params(Id("check").Id("Check")).Block(
		Id("h").Dot("lock").Dot("Lock").Call(),
		Defer().Id("h").Dot("lock").Dot("Unlock").Call(),
		Line(),
		Id("h").Dot("livenessChecks").Op("=").Append(Id("h").Dot("livenessChecks"), Id("check")),
	)

	file.Func().Params(Id("h").Op("*").Id("healthcheck")).Id("AddReadinessChecks").Params(Id("check").Id("Check")).Block(
		Id("h").Dot("lock").Dot("Lock").Call(),
		Defer().Id("h").Dot("lock").Dot("Unlock").Call(),
		Line(),
		Id("h").Dot("readinessChecks").Op("=").Append(Id("h").Dot("readinessChecks"), Id("check")),
	)

	file.Func().Params(Id("h").Op("*").Id("healthcheck")).Id("checkReadiness").Params().Error().Block(
		Id("h").Dot("lock").Dot("RLock").Call(),
		Defer().Id("h").Dot("lock").Dot("RUnlock").Call(),
		Line(),
		For(
			List(Id("_"), Id("check")).Op(":=").Range().Id("h").Dot("readinessChecks").Block(
				If(
					Err().Op(":=").Id("check").Call(),
					Err().Op("!=").Nil(),
				).Block(
					Return(Err()),
				),
			),
		),
		Return(Nil()),
	)

	file.Func().Params(Id("h").Op("*").Id("healthcheck")).Id("checkLiveness").Params().Error().Block(
		Id("h").Dot("lock").Dot("RLock").Call(),
		Defer().Id("h").Dot("lock").Dot("RUnlock").Call(),
		Line(),
		For(
			List(Id("_"), Id("check")).Op(":=").Range().Id("h").Dot("livenessChecks").Block(
				If(
					Err().Op(":=").Id("check").Call(),
					Err().Op("!=").Nil(),
				).Block(
					Return(Err()),
				),
			),
		),
		Return(Nil()),
	)

	return file
}
