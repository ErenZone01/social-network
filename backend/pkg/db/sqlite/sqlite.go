package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var migrationsDir = "pkg/db/migrations/sqlite"
var fileDB = "./pkg/db/sqlite/sql.db"
var DB *sql.DB

func CheckDB() *sql.DB {
	var _, err = os.ReadFile(fileDB)
	if err != nil {
		return CreateBD()
	}
	db, err := sql.Open("sqlite3", fileDB)
	if err != nil {
		fmt.Println("err : ", err)
		return nil
	}
	DB = db
	return db
}

func CreateBD() *sql.DB {
	// Ouvrir la connexion à la base de données

	db, err := sql.Open("sqlite3", fileDB)
	if err != nil {
		fmt.Println("err : ", err)
		return nil
	}
	DB = db

	// Créer le répertoire des migrations s'il n'existe pas
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		fmt.Println("err : ", err)
		return nil
	}
	fmt.Println("Connection bd etablishing")
	// Appliquer les migrations
	if err := migration(db); err != nil {
		fmt.Println("Failed to apply migrations: ", err)
		return nil
	}
	return db
}

var db *sql.DB

func migration(db *sql.DB) error {
	// Création de l'instance de migration
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		fmt.Println("Erreur lors de la création de l'instance de migration:", err)
		return err
	}

	m, err1 := migrate.NewWithDatabaseInstance("file://"+migrationsDir, "sqlite", driver)
	if err1 != nil {
		fmt.Println("Erreur lors de l'exécution des migrations:", err1)
		return err1
	}

	// Exécution des migrations vers la dernière version disponible
	if err2 := m.Up(); err2 != nil && err2 != migrate.ErrNoChange {
		fmt.Println("Erreur lors de l'exécution des migrations:", err2)
		return err2
	}

	fmt.Println("Migrations appliquées avec succès")
	return nil
}
