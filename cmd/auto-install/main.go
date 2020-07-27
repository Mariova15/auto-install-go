package main

import (
	"fmt"

	"github.com/mariova15/auto-install-go/internal"

	"github.com/mariova15/auto-install-go/pkg"
)

func main() {
	fmt.Println(internal.Message)
	programs := pkg.ListFiles("C:\\INSTALL")
	fmt.Println(programs)
	pkg.InstallProgram(programs[1])
}
