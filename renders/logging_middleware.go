package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderLoggingMiddleware() *File {
	file := NewFilePathName("logging", "middleware")

	return file
}
