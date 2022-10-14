package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHTTPErrors() *File {

	/*
		package http

		import (
			"context"
			"encoding/json"
			"github.com/company/blanksvc/pkg/service"
			"net/http"
		)
	*/

	file := NewFile("http")

	servicePackage := path.Join(svc.ModulePath, "pkg/service")
	file.ImportName(servicePackage, "service")
	jsonPackage := "encoding/json"
	file.ImportName(jsonPackage, "json")
	httpPackage := "net/http"
	file.ImportName(httpPackage, "http")

	/*
		type errorWrapper struct {
			Error string `json:"error"`
		}
	*/

	file.Type().Id("errorWrapper").Struct(
		Id("Error").String().Tag(map[string]string{"json": "error"}),
	)

	/*
		func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
			w.WriteHeader(errorToCode(err))
			json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
		}
	*/

	file.Func().Id("errorEncoder").Params(Id("_").Qual("context", "Context"), Err().Error(), Id("w").Qual(httpPackage, "ResponseWriter")).Block(
		Id("w").Dot("WriteHeader").Call(Id("errorToCode").Call(Err())),
		Qual(jsonPackage, "NewEncoder").Call(Id("w")).Dot("Encode").Call(Id("errorWrapper").Values(Dict{
			Id("Error"): Err().Dot("Error").Call(),
		})),
	)

	/*
		func errorToCode(err error) int {
			switch err {
			case service.ErrEmptyString:
				return http.StatusBadRequest
			}
			return http.StatusInternalServerError
		}
	*/

	file.Func().Id("errorToCode").Params(Err().Error()).Int().Block(
		Switch(Err()).Block(
			Case(Qual(servicePackage, "ErrEmptyString")).Block(
				Return(Qual(httpPackage, "StatusBadRequest")),
			),
		),
		Return(Qual(httpPackage, "StatusInternalServerError")),
	)

	return file
}
