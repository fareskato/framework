package main

import (
	"fmt"
	"time"
)

// doSessionTable store session in db
func doSessionTable() error {
	// database type used(postgres or mysql)
	dbType := kbr.DB.DataBaseType
	// just in case user typed postgresql instead of postgres and mariadb instead of mysql
	if dbType == "postgresql" {
		dbType = "postgres"
	}
	if dbType == "mariadb" {
		dbType = "mysql"
	}
	// create migrations files
	fileName := fmt.Sprintf("%d_create_sessions_tables", time.Now().UnixMicro())
	upFile := kbr.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := kbr.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		exitGracefully(err)
	}
	err = copyDataToFile([]byte("drop table if exists sessions cascade;"), downFile)
	if err != nil {
		exitGracefully(err)
	}
	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}
	// all good
	return nil
}
