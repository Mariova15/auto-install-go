// go: generate auto-install -icon = Logo.ico
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mariova15/auto-install-go/pkg"
)

var exit bool

func main() {

	// TEST
	// fmt.Println(internal.Message)
	// programs := pkg.ListFiles("C:\\INSTALL")
	// fmt.Println(programs)
	// pkg.InstallProgram(programs[1])

	programs := pkg.InitPrograms()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------")
	fmt.Println("-----Auto-Install------")
	fmt.Println("-----------------------")

	for !exit {
		fmt.Println("1-Scan directory")
		fmt.Println("2-Install programs")
		fmt.Println("3-Exit")

		option, _ := reader.ReadString('\n')
		option = option[0 : len(option)-2]

		switch option {
		case "1":
			fmt.Println("Scan directory")
			fmt.Println("Write path:")
			pathScan, _ := reader.ReadString('\n')
			pathScan = pathScan[0 : len(pathScan)-2]

			if pkg.CheckDir(pathScan) {
				programs = pkg.ListFiles(pathScan)

				for _, program := range programs {
					fmt.Println(program)
				}
			}

		case "2":
			fmt.Println("Install programs")
			pkg.InstallPrograms(programs)

		case "3":
			fmt.Println("BYE")
			exit = true

		default:
			fmt.Println("Choose an option")
		}

	}

}
