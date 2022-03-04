package main

import (
	"errors"
	"os"

	"github.com/fareskato/kabarda"

	"github.com/fatih/color"
)

const version = "1.0.0"

var kbr kabarda.Kabarda

func main() {
	var message string
	arg1, arg2, arg3, arg4, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	// add database type to kbr
	setup(arg1, arg2)
	switch arg1 {
	case "help":
		showHelp()
	case "new":
		if arg2 == "" {
			exitGracefully(errors.New("new requires an application name "))
		}
		doNew(arg2)
	case "version":
		color.Yellow("Kabarda version: " + version)

	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
		message = "Migrations completed!"
	case "make":
		if arg2 == "" {
			exitGracefully(errors.New("make requires sub command: migration|model|handler ...etc"))
		}
		err = doMake(arg2, arg3, arg4)
		if err != nil {
			exitGracefully(err)
		}
	case "key":
		rnd := kbr.RandomString(32)
		color.Blue("32 character encryption key: %s", rnd)
	default:
		showHelp()
	}
	exitGracefully(nil, message)
}

func validateInput() (string, string, string, string, error) {
	var str1, str2, str3, str4 string
	// get command line arguments
	if len(os.Args) > 1 {
		str1 = os.Args[1]
		if len(os.Args) >= 3 {
			str2 = os.Args[2]
		}
		if len(os.Args) >= 4 {
			str3 = os.Args[3]
		}
		if len(os.Args) >= 5 {
			str4 = os.Args[4]
		}
	} else {
		showHelp()
		return "", "", "", "", errors.New("command required")
	}
	return str1, str2, str3, str4, nil
}
