package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var migrationsDir = "pkg/db/migrations/sqlite"
var DB *sql.DB

func CreateBD() *sql.DB {
	fmt.Println("start")
	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "./pkg/db/sqlite/sql.db")
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
	if err := migrate(db); err != nil {
		fmt.Println("Failed to apply migrations: ", err)
		return nil
	}
	return db
}

func migrate(db *sql.DB) error {
	// Liste des fichiers de migration
	migrations, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		fmt.Println("not found")
		return err
	}

	// Appliquer chaque migration
	for _, migration := range migrations {
		// Lire le contenu du fichier de migration
		query, err := os.ReadFile(migration)
		if err != nil {
			fmt.Println("not read")
			return err
		}

		// Exécuter la requête de migration
		if _, err := db.Exec(string(query)); err != nil {
			 fmt.Println("error")
			 return err
		}

		fmt.Println("Applied migration: ", migration)
	}
	return nil
}
