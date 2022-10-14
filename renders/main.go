package renders

import (
	"path"

	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderMain() *File {

	/*
		package main

		import (
			"github.com/company/blanksvc/cmd"
			"os"
		)
	*/

	file := NewFile("main")
	cmdPackage := path.Join(svc.ModulePath, "cmd")
	file.ImportName(cmdPackage, "cmd")

	/*
		func main() {
			if err := cmd.RunServer(); err != nil {
				os.Exit(1)
			}
		}
	*/

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
