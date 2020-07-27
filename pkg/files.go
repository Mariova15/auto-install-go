package pkg

import (
	"os"
	"strconv"

	"os/exec"

	"path/filepath"

	"strings"

	"debug/pe"

	"log"
)

type Program struct {
	productName string
	path        string
	params      string
}

func ListFiles(dirPath string) []Program {

	programs := []Program{}

	err := filepath.Walk(dirPath,
		func(filePath string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if filepath.Ext(filePath) == ".exe" || filepath.Ext(filePath) == ".msi" {

				program := Program{
					productName: filepath.Base(filePath),
					path:        strings.Replace(filePath, " ", "^ ", -1),
					params:      GetFileinfo(filePath)}

				programs = append(programs, program)

			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return programs
}

func InstallProgram(program Program) bool {

	c := exec.Command("cmd.exe", "/C", program.path+" "+program.params)

	if err := c.Run(); err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}

}

func UninstallProgram(productName string) bool {

	c := exec.Command("cmd.exe", "/C", "wmic product where name= \""+productName+"\" call uninstall /nointeractive")

	if err := c.Run(); err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}

}

func GetFileinfo(filePath string) string {

	if filepath.Ext(filePath) != ".msi" {
		file, err := pe.Open(filePath)

		if err != nil {
			log.Println(err)
		}

		if file != nil {
			log.Println(file.FileHeader.Characteristics)

			installerType := int(file.FileHeader.Characteristics)

			file.Close()

			return GetParams(installerType)

		}
	}

	// MSI
	return "/qn /norestart"

}

func GetParams(installerType int) string {

	if installerType == 33167 {
		// return "INNO"
		return "/VERYSILENT /SUPPRESSMSGBOXES /NORESTART"

	} else if installerType == 271 || installerType == 258 {
		// return "NSIS"
		return "/S"

	} else if installerType == 259 {
		// return "InstallShield"
		return "/s"

	} else {
		return strconv.FormatInt(int64(installerType), 10)
	}

}
