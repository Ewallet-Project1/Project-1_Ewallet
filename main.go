package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)
type AppConfig struct{
	DB_USERNAME string
	DB_PASSWORD string
 	DB_HOST		string
 	DB_PORT		string
 	DB_NAME		string
}

func InitDB()(*sql.DB, error){
	// Capture connection properties.
	// var connectionString = "root:0000@tcp(localhost:3306)/ewallet"
	var cfg = AppConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
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

func main(){
	db, _ := InitDB()
	
	defer db.Close()
	
		fmt.Println("Pilih menu:")
		fmt.Println("[0]: Keluar Aplikasi")
		fmt.Println("[1]: Register")
		fmt.Println("[2]: Login")
		fmt.Println("[3]: Lihat Profil Akun")
		fmt.Println("[4]: Edit Akun")
		fmt.Println("[5]: Delete Akun")
		fmt.Println("[6]: Top-up")
		fmt.Println("[7]: Transaction")
		fmt.Println("[8]: History Top-up")
		fmt.Println("[9]: History Transaction")
		fmt.Println("[10]: Lihat Profil pengguna lain")
		var pilihan int
		fmt.Println("Masukkan angka sesuai pilihan menu:")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 0:
			fmt.Println("Exit From System")
		case 1:
			fmt.Println("register")
		case 2:

		case 3:
			
		case 4:

		case 5:

		case 6:

		case 7:
			
		case 8:

		case 9:

		case 10:
			
		}
	}