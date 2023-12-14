package main

import (
	"database/sql"
	"ewallets-tim-1/controllers"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func InitDB() (*sql.DB, error) {
	// Capture connection properties.
	// var connectionString = "root:0000@tcp(localhost:3306)/ewallet"
	var cfg = AppConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error connection to database:", err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("error ping connection:", pingErr)
		return nil, pingErr
	}
	fmt.Println("Connected!")
	return db, nil
}

func main() {
	db, _ := InitDB()

	defer db.Close()

	fmt.Println("Pilih menu:")
	fmt.Println("[0]: Keluar Aplikasi")
	fmt.Println("[1]: Register")
	fmt.Println("[2]: Login")
	var pilihan int
	fmt.Println("Masukkan angka sesuai pilihan menu:")
	fmt.Scanln(&pilihan)

	// var noTelp string
	switch pilihan {
	case 0:
		fmt.Println("Exit From System")
	case 1:
		controllers.Register(db)
	case 2:
		noTelp := controllers.Login(db)
		fmt.Println("Login Berhasil")
		fmt.Println("Masukkan angka sesuai pilihan menu:")
		fmt.Println("[3]: Lihat Profil Akun")
		fmt.Println("[4]: Edit Akun")
		fmt.Println("[5]: Delete Akun")
		fmt.Println("[6]: Top-up")
		fmt.Println("[7]: Transaction")
		var pilihanLogin int
		fmt.Scanln(&pilihanLogin)
		// fmt.Println("[8]: History Top-up")
		// fmt.Println("[9]: History Transaction")
		// fmt.Println("[10]: Lihat Profil pengguna lain")
		switch pilihanLogin {
		case 3:
			controllers.ReadProfile(db, noTelp)
		case 4:
			controllers.EditProfile(db, noTelp)
		case 5:
			controllers.Delete(db, noTelp)
		case 6:
			controllers.TopUpSaldo(db, noTelp)
		case 7:
			controllers.Transaction(db, noTelp)
		// case 8:

		// case 9:

		// case 10:

		}

	}
}
