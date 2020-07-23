package pkg

import (
	"os"

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

			filePathReplaced := filepath.Join(filepath.Dir(filePath),strings.Replace(filepath.Base(filePath), " ", "_", -1))
			// os.Rename(filePath, filePathReplaced)

			filePaths = append(filePaths, filePathReplaced)
		}

    	return nil
	})
	
	if err != nil {
		log.Println(err)
	}

    return filePaths
}