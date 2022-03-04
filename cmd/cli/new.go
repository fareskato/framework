package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

///////////////////////////
// Generate New Application
///////////////////////////

var appURL string

// doNew generate new application, takes the application name(arg2)
func doNew(appName string) {
	// to lower case
	appName = strings.ToLower(appName)
	appURL = appName
	// sanitize the application name: http://github.com/user/therepo
	// we need the name after the last / which is therepo
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[(len(exploded) - 1)]
	}
	// clone the skeleton application from GitHub
	color.Green("\t cloning repository ...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/fareskato/kabarda-app.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		exitGracefully(err)
	}
	// remove .git directory
	err = os.RemoveAll(fmt.Sprintf("./%s'.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	// create .env file with needed data
	color.Yellow("\tCreating .env file... ")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}
	envData := string(data)
	envData = strings.ReplaceAll(envData, "${APP_NAME}", appName)
	envData = strings.ReplaceAll(envData, "${KEY}", kbr.RandomString(32))
	err = copyDataToFile([]byte(envData), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		exitGracefully(err)
	}
	// create a makefile
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.windows", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()
		destination, err := os.Create(fmt.Sprintf("./%s/MakeFile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.unix", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()
		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	}
	_ = os.Remove("./" + appName + "/Makefile.unix")
	_ = os.Remove("./" + appName + "/Makefile.windows")

	// update the go.mod file
	color.Yellow("\tCreating go.mod file... ")
	_ = os.Remove("./" + appName + "/go.mod")
	data, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		exitGracefully(err)
	}
	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)
	err = copyDataToFile([]byte(mod), fmt.Sprintf("./%s/go.mod", appName))
	if err != nil {
		exitGracefully(err)
	}

	// update existing .go files with correct imports and data
	os.Chdir("./" + appName)
	time.Sleep(3 * time.Second)
	color.Yellow("\tUpdating source files... ")
	updateSource()
	// run go mod tidy in the project directory
	time.Sleep(1 * time.Second)
	color.Yellow("\tRunning go mod tidy... ")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Start()
	if err != nil {
		exitGracefully(err)
	}
	time.Sleep(1 * time.Second)
	color.Green("")
	color.Green("Done building " + appURL)
	color.Green("")
	color.Blue("Please set application database, session, email ...etc in .env file")
	color.Green("")
	color.Green("First of all navigate to your application root directory and type: make restart ")
	color.Green("this command will start web server...")
	color.Green("")
	color.Green("Second run: copy the binary file to your application root so you can use it")
	color.Green("to run commands")
	color.Green("")
	color.Green("Third run: binFile make auth, this command will generate auth built in package")
	color.Green("with all needed functionality")
	color.Green("")
	color.Red("to see all commands run: ./binfile help")
	color.Green("")
	color.Blue("Happy coding )")
	color.Green("")
}
