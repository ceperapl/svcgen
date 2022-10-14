package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHealthcheck() *File {

	/*
		package healthcheck

		import (
			"net/http"
			"sync"
		)
	*/

	file := NewFile("healthcheck")

	httpPackage := "net/http"
	file.ImportName(httpPackage, "http")

	/*
		type Check func() error
		type HealthChecker interface {
			AddLivenessChecks(check Check)
			AddReadinessChecks(check Check)
			LivenessHandler() http.HandlerFunc
			ReadinessHandler() http.HandlerFunc
		}
	*/

	file.Type().Id("Check").Func().Params().Error()

	file.Type().Id("HealthChecker").Interface(
		Id("AddLivenessChecks").Params(Id("check").Id("Check")),
		Id("AddReadinessChecks").Params(Id("check").Id("Check")),
		Id("LivenessHandler").Params().Qual(httpPackage, "HandlerFunc"),
		Id("ReadinessHandler").Params().Qual(httpPackage, "HandlerFunc"),
	)

	/*
		func New() HealthChecker {
			return &healthcheck{}
		}
	*/

	file.Func().Id("New").Params().Id("HealthChecker").Block(
		Return(Op("&").Id("healthcheck")).Values(),
	)

	/*
		type healthcheck struct {
			lock sync.RWMutex

			livenessChecks  []Check
			readinessChecks []Check
		}
	*/

	file.Type().Id("healthcheck").Struct(
		Id("lock").Qual("sync", "RWMutex"),
		Line(),
		Id("livenessChecks").Index().Id("Check"),
		Id("readinessChecks").Index().Id("Check"),
	)

	/*
		func (h *healthcheck) LivenessHandler() http.HandlerFunc {
			return func(rw http.ResponseWriter, r *http.Request) {
				if err := h.checkLiveness(); err != nil {
					rw.WriteHeader(http.StatusServiceUnavailable)
					return
				}
				rw.WriteHeader(http.StatusOK)
			}
		}
	*/

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

	/*
		func (h *healthcheck) ReadinessHandler() http.HandlerFunc {
			return func(rw http.ResponseWriter, r *http.Request) {
				if err := h.checkReadiness(); err != nil {
					rw.WriteHeader(http.StatusServiceUnavailable)
					return
				}
				rw.WriteHeader(http.StatusOK)
			}
		}
	*/

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

	/*
		func (h *healthcheck) AddLivenessChecks(check Check) {
			h.lock.Lock()
			defer h.lock.Unlock()

			h.livenessChecks = append(h.livenessChecks, check)
		}
		func (h *healthcheck) AddReadinessChecks(check Check) {
			h.lock.Lock()
			defer h.lock.Unlock()

			h.readinessChecks = append(h.readinessChecks, check)
		}
	*/

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

	/*
		func (h *healthcheck) checkReadiness() error {
			h.lock.RLock()
			defer h.lock.RUnlock()

			for _, check := range h.readinessChecks {
				if err := check(); err != nil {
					return err
				}
			}
			return nil
		}
	*/

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

	/*
		func (h *healthcheck) checkLiveness() error {
			h.lock.RLock()
			defer h.lock.RUnlock()

			for _, check := range h.livenessChecks {
				if err := check(); err != nil {
					return err
				}
			}
			return nil
		}
	*/

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
