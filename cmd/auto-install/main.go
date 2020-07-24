package main

import (
	"fmt"

	"github.com/mariova15/auto-install-go/internal"

	"github.com/mariova15/auto-install-go/pkg"
)

func main() {
	fmt.Println(internal.Message)
	fmt.Println(pkg.ListFiles("C:\\INSTALL"))
	pkg.InstallProgram("C:\\INSTALL\\TSC 1.1.exe")
}
