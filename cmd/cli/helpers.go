package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

// checkForDB checks if the database is set up in .env file and ensure that
// database.yml file is there(needed for fizz)
func checkForDB() {
	dbType := kbr.DB.DataBaseType
	if dbType == "" {
		exitGracefully(errors.New("no database connection provided in .env file"))
	}
	if !isFileExists(kbr.RootPath + "/config/database.yml") {
		exitGracefully(errors.New("config/database.yml file does not exist"))
	}
}

func showHelp() {
	color.Yellow("")
	color.Yellow(`Available commands:
	help                                - show help commands
	version                             - print application version
	migrate                             - runs  all migrations that have not run previously
	migrate down                        - roll back the most recent migration
	migrate reset                       - runs all down migrations in revers order then run all migrations up
	make migration <name> <format>      - creates migration up and down(fizz or sql, the default is fizz) migration files in migrations folder
	make auth                           - creates and run auth migrations files, models and middlewares 
	make handler <name>                 - creates handler in handlers directory
	make model <name>                   - creates model in data directory
	make session                        - creates table in the database as session store
	make key                            - creates 32 character encryption key
	make mail <name>                    - creates starter mail templates(html and plain text) in the mail directory
`)
	color.Yellow("")
}

// setup will set some needed fields in kbr(of type Kabarda)
func setup(arg1, arg2 string) {
	if arg1 != "new" && arg1 != "help" && arg1 != "version" {
		// load .env file
		err := godotenv.Load()
		if err != nil {
			exitGracefully(err)
		}
		// get root dir
		dir, err := os.Getwd()
		if err != nil {
			exitGracefully(err)
		}
		// set needed kbr fields
		kbr.RootPath = dir
		kbr.DB.DataBaseType = os.Getenv("DATABASE_TYPE")
	}

}

// exitGracefully exit the program without panic and display colored message
// using color package
func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	if err != nil {
		color.Red("Error: %v\n", err)
	}
	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
		color.Green("")
	}
	os.Exit(0)
}

// getDSN generate dsn string to be used with migrations commands
func getDSN() string {
	dbType := kbr.DB.DataBaseType
	// refer to driver.go file where we use pgx instead of postgres
	if dbType == "pgx" {
		dbType = "postgres"
	}
	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		}
		return dsn
	}
	return "mysql://" + kbr.BuildDSN()
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// check if the current file is dir
	if fi.IsDir() {
		return nil
	}
	// only go files
	matched, err := filepath.Match("*.go", fi.Name())
	if fi.IsDir() {
		return err
	}
	// matched
	if matched {
		read, err := os.ReadFile(path)
		if err != nil {
			exitGracefully(err)
		}
		// search and replace: -1 means every occurrence
		newData := strings.Replace(string(read), "myapp", appURL, -1)
		// write changed file
		err = os.WriteFile(path, []byte(newData), 0)
		if err != nil {
			exitGracefully(err)
		}
	}
	return nil
}

func updateSource() {
	// walk entire project folder including sub-folders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		exitGracefully(err)
	}
}
