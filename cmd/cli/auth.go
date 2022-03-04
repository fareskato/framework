package main

import (
	"github.com/fatih/color"
	"os"
)

func doAuth() error {
	// ensure database is set
	checkForDB()
	// create migrations
	dbType := kbr.DB.DataBaseType

	// connect to DB via pop
	tx, err := kbr.PopConnect()
	if err != nil {
		exitGracefully(err)
	}
	defer tx.Close()

	upBytes, err := templateFS.ReadFile("templates/migrations/auth_tables." + dbType + ".sql")
	if err != nil {
		exitGracefully(err)
	}
	downBytes := []byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;")

	err = kbr.CreatePopMigrations(upBytes, downBytes, "auth", "sql")
	if err != nil {
		exitGracefully(err)
	}
	// run migrations
	err = kbr.RunPopMigrations(tx)
	if err != nil {
		exitGracefully(err)
	}

	// copy all auth functionality related files(users and token models)
	err = copyFileFromTemplate("templates/data/user.go.txt", kbr.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/data/token.go.txt", kbr.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}
	// copy remember token
	err = copyFileFromTemplate("templates/data/remember_token.go.txt", kbr.RootPath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy auth middlewares
	err = copyFileFromTemplate("templates/middlewares/auth.go.txt", kbr.RootPath+"/middlewares/auth.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/middlewares/auth-token.go.txt", kbr.RootPath+"/middlewares/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/middlewares/remember_token.go.txt", kbr.RootPath+"/middlewares/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}
	// copy handlers
	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", kbr.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}
	// copy views and mailer views
	err = copyFileFromTemplate("templates/mailer/password-reset.html.gohtml", kbr.RootPath+"/mail/password-reset.html.gohtml")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/mailer/password-reset.plain.gohtml", kbr.RootPath+"/mail/password-reset.plain.gohtml")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/forgot.jet", kbr.RootPath+"/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/reset-password.jet", kbr.RootPath+"/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/login.jet", kbr.RootPath+"/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/register.jet", kbr.RootPath+"/views/register.jet")
	if err != nil {
		exitGracefully(err)
	}
	// correct the imports on auth
	appURL = os.Getenv("APP_NAME")
	updateSource()

	// inform developer
	color.Yellow("")
	color.Yellow(" - users, tokens and remember_tokens migrations created in migrations directory and ran")
	color.Yellow(" - user and token models created in data directory")
	color.Yellow(" - auth middlewares created in middlewares directory")
	color.Yellow("")
	color.Yellow("auth routes added to your routes.go")
	color.Yellow("")
	color.Green(" - Please uncomment Users and Tokens models in data/models.go, and add appropriate middlewares to your routes")
	color.Yellow("")
	color.Yellow("Authentication routes are:")
	color.Green("/users/login Get/Post")
	color.Green("/users/register Get/Post")
	color.Green("/users/logout Post")
	color.Green("/users/forgot-password Get/Post")
	color.Green("/users/reset-password Get/Post")
	return nil
}
