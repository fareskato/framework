package main

import (
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath, targetFile string) error {
	// check if the target file doesn't already exist
	if isFileExists(targetFile) {
		return errors.New(fmt.Sprintf("%s already exists", targetFile))
	}
	// read files
	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}
	// write data to target file
	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}
	return nil
}

func copyDataToFile(data []byte, destination string) error {
	err := ioutil.WriteFile(destination, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func isFileExists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}
