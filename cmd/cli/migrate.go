package main

/*
	All needed logic for handling migrations
*/

// doMigrate run migrations commands
func doMigrate(arg2, arg3 string) error {
	// ensure that database is set
	checkForDB()
	// connect to DB via pop
	tx, err := kbr.PopConnect()
	if err != nil {
		exitGracefully(err)
	}
	defer tx.Close()

	// run migration commands
	switch arg2 {
	case "up":
		//err := kbr.MigrateUp(dsn)
		err := kbr.RunPopMigrations(tx)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := kbr.PopMigrateDown(tx, -1)
			if err != nil {
				return err
			}
		} else {
			err := kbr.PopMigrateDown(tx, 1)
			if err != nil {
				return err
			}
		}
	// reset DB: run all migrations down then run up again
	case "reset":
		err := kbr.PopMigrationsReset(tx)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}
	return nil
}
