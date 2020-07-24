package pkg

import (
	"os"

	"os/exec"

	"path/filepath"

	"strings"

	"log"
)

func ListFiles(dirPath string) []string {

	filePaths := []string{}

	err := filepath.Walk(dirPath,
    func(filePath string, info os.FileInfo, err error) error {
	
		if err != nil {
        	return err
		}

		if filepath.Ext(filePath) == ".exe" || filepath.Ext(filePath) == ".msi" {	
			filePaths = append(filePaths, filePath)
		}

    	return nil
	})
	
	if err != nil {
		log.Println(err)
	}

    return filePaths
}

func InstallProgram(programPath string) bool {

	filePathReplaced := strings.Replace(programPath, " ", "^ ", -1)
	
	c := exec.Command("cmd.exe", "/C",filePathReplaced) 

	if err := c.Run(); err != nil { 
		log.Println(err) 
		return false
	}else{
		return true
	}

}