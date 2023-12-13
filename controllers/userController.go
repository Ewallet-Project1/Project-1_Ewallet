package controllers

import (
	"database/sql"
	"ewallets-tim-1/entities"
	"fmt"
	"log"
)

// 1. fitur register
func Register(db *sql.DB) {
	var newUser entities.User
	fmt.Println("Masukkan Nama")
	fmt.Scanln(&newUser.FullName)

	fmt.Println("Masukkan No Telepon")
	fmt.Scanln(&newUser.Phone)

	fmt.Println("Masukkan Password")
	fmt.Scanln(&newUser.Password)

	fmt.Println("Masukkan Alamat")
	fmt.Scanln(&newUser.Address)

	fmt.Println("Masukkan Saldo Awal")
	fmt.Scanln(&newUser.Balance)

	result, errInsert := db.Exec("INSERT INTO users (full_name, phone, password, address, balance) VALUES (?, ?, ?, ?, ?)", &newUser.FullName, &newUser.Phone, &newUser.Password, &newUser.Address, &newUser.Balance)
	if errInsert != nil {
		log.Fatal("error insert data", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Register Success")
		} else {
			fmt.Println("Register Failed")
		}
	}
}

// 2. Login
func Login(db *sql.DB) (telp string) {
	var noTelp string
	var password string
	var newUser entities.User
	fmt.Println("Masukkan No Telepon")
	fmt.Scanln(&newUser.Phone)

	fmt.Println("Masukkan Password")
	fmt.Scanln(&newUser.Password)
	noTelp = *&newUser.Phone
	password = *&newUser.Password

	var user entities.User
	row := db.QueryRow("SELECT phone, password FROM users WHERE phone = ?", noTelp)
	if err := row.Scan(&user.Phone, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No telp: %s tidak terdaftar", noTelp)
		}
		fmt.Println(err)
	}
	if password != user.Password {
		log.Fatal("Password Salah !")
	}
	fmt.Println("Login Berhasil")
	fmt.Println("Masukkan angka sesuai pilihan menu:")
	fmt.Println("[3]: Lihat Profil Akun")
	fmt.Println("[4]: Edit Akun")
	fmt.Println("[5]: Delete Akun")
	// fmt.Println("[6]: Top-up")
	// fmt.Println("[7]: Transaction")
	// fmt.Println("[8]: History Top-up")
	// fmt.Println("[9]: History Transaction")
	// fmt.Println("[10]: Lihat Profil pengguna lain")

	return noTelp
}
