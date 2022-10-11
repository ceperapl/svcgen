//go:generate go-bindata -ignore .*template-bindata\.go -pkg templates -o templates/template-bindata.go templates/...
package main

import "github.com/company/svcgen/cmd"

func main() {
	cmd.Execute()
}
