package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderMain() *File {
	file := NewFile("main")
	cmdPackage := path.Join(svc.ModulePath, "cmd")
	file.ImportName(cmdPackage, "cmd")

	file.Func().Id("main").Params().Block(
		If(
			Err().Op(":=").Qual(cmdPackage, "RunServer").Call(),
			Err().Op("!=").Nil(),
		).Block(
			Qual("os", "Exit").Call(Lit(1)),
		),
	)
	return file
}
