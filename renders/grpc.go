package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderGrpc() *File {
	file := NewFile("grpc")

	return file
}
