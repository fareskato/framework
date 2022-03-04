package kabarda

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/pop"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

////////////////////
// USE POP MIGRATION
////////////////////

// PopConnect connect to database via pop will return transactions and error(if exists)
func (k *Kabarda) PopConnect() (*pop.Connection, error) {
	tx, err := pop.Connect("development")
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreatePopMigrations create up and down migrations
func (k *Kabarda) CreatePopMigrations(up, down []byte, migrationName, migrationType string) error {
	migrationPath := k.RootPath + "/migrations"
	err := pop.MigrationCreate(migrationPath, migrationName, migrationType, up, down)
	if err != nil {
		return err
	}
	return nil
}

// RunPopMigrations run migrations up
func (k *Kabarda) RunPopMigrations(conn *pop.Connection) error {
	// path to migrations folder
	migrationPath := k.RootPath + "/migrations"
	// create file migrator
	fileMigrator, err := pop.NewFileMigrator(migrationPath, conn)
	if err != nil {
		return err
	}
	// Run migration up
	err = fileMigrator.Up()
	if err != nil {
		return err
	}
	return nil
}

// PopMigrateDown run migrations down
func (k *Kabarda) PopMigrateDown(conn *pop.Connection, steps ...int) error {
	// path to migrations folder
	migrationPath := k.RootPath + "/migrations"
	// default step is 1
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	// create file migrator
	fileMigrator, err := pop.NewFileMigrator(migrationPath, conn)
	if err != nil {
		return err
	}
	// Run migration down
	err = fileMigrator.Down(step)
	if err != nil {
		return err
	}
	return nil
}

// PopMigrationsReset reset all migrations
func (k *Kabarda) PopMigrationsReset(conn *pop.Connection) error {
	// path to migrations folder
	migrationPath := k.RootPath + "/migrations"
	// create file migrator
	fileMigrator, err := pop.NewFileMigrator(migrationPath, conn)
	if err != nil {
		return err
	}
	// reset migrations
	err = fileMigrator.Reset()
	if err != nil {
		return err
	}
	return nil
}

/////////////////////////
// USE PURE SQL MIGRATION
/////////////////////////

func (k *Kabarda) MigrateUp(dsn string) error {
	mgr, err := migrate.New("file://"+k.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer mgr.Close()
	if err := mgr.Up(); err != nil {
		log.Println("error run migration")
		return err
	}
	return nil
}

func (k *Kabarda) MigrateDownAll(dsn string) error {
	mgr, err := migrate.New("file://"+k.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer mgr.Close()
	if err := mgr.Down(); err != nil {
		return err
	}
	return nil

}

func (k *Kabarda) MigrateSteps(n int, dsn string) error {
	mgr, err := migrate.New("file://"+k.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer mgr.Close()
	if err := mgr.Steps(n); err != nil {
		return err
	}
	return nil
}

func (k *Kabarda) MigrateForce(dsn string) error {
	mgr, err := migrate.New("file://"+k.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer mgr.Close()
	if err := mgr.Force(-1); err != nil {
		return err
	}
	return nil
}
