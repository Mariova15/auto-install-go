package pkg

import (
	"os"

	"path/filepath"
)

func ListFiles(dirPath string) []string {

	filePaths := []string{}

	err := filepath.Walk(dirPath,
    func(path string, info os.FileInfo, err error) error {
	
		if err != nil {
        	return err
		}

		filePaths = append(filePaths, path)

    	return nil
	})
	
	if err != nil {
		// log.Println(err)
	}

    return filePaths
}