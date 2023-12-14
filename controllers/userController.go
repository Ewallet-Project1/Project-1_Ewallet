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
	return noTelp
}

// 3. Read Data / lihat profil akun
func ReadProfile(db *sql.DB, noTelp string) {

	var user entities.User
	row := db.QueryRow("SELECT full_name, phone, address, balance, created_at FROM users WHERE phone = ?", noTelp)
	if err := row.Scan(&user.FullName, &user.Phone, &user.Address, &user.Balance, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			errorRead := fmt.Errorf("No telp : %s tidak terdaftar", noTelp)
			fmt.Println(errorRead)
		} else {
			// note, printf tidak keluar
			fmt.Printf("Nama Pengguna: %s\nNomor Telepon: %s\nAlamat : %s\nSaldo : %d", user.FullName, user.Phone, user.Address, user.Balance)
		}
		fmt.Print(err)
	}
	fmt.Printf("Nama Pengguna: %s\nNomor Telepon: %s\nAlamat : %s\nSaldo : %d", user.FullName, user.Phone, user.Address, user.Balance)
}

// 4. Upadate user
func EditProfile(db *sql.DB, noTelp string) {
	var pilihAngka int
	var pilihEdit string
	var hasilhEdit string
	var newUser entities.User
	fmt.Println("Pilih data yang ingin diubah:")
	fmt.Println("[1] Nama")
	fmt.Println("[2] No Telepon")
	fmt.Println("[3] Password")
	fmt.Println("[4] Alamat")
	fmt.Scanln(&pilihAngka)

	if pilihAngka == 1 {
		pilihEdit = "full_name"
		fmt.Scanln(&newUser.FullName)
		hasilhEdit = *&newUser.FullName
	} else if pilihAngka == 2 {
		pilihEdit = "phone"
		fmt.Scanln(&newUser.Phone)
		hasilhEdit = *&newUser.Phone
	} else if pilihAngka == 3 {
		pilihEdit = "password"
		fmt.Scanln(&newUser.Password)
		hasilhEdit = *&newUser.Password
	} else if pilihAngka == 4 {
		pilihEdit = "address"
		fmt.Scanln(&newUser.Address)
		hasilhEdit = *&newUser.Address
	} else {
		log.Fatal("Pilihan tidak ada, Keluar Dari Sistem")
	}

	queryResult := fmt.Sprintf("UPDATE users SET %s = '%s' WHERE phone = %s", pilihEdit, hasilhEdit, noTelp)
	result, errInsert := db.Exec(queryResult)
	if errInsert != nil {
		log.Fatal("error update data", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Update Success")
		} else {
			fmt.Println("Update Failed")
		}
	}
}

// 5. Delete Users
func Delete(db *sql.DB, noTelp string) {

	result, errDelete := db.Exec("DELETE FROM users WHERE phone=?", noTelp)
	if errDelete != nil {
		log.Fatal("error Edit :", errDelete.Error())
	} else {
		rows, _ := result.RowsAffected()
		if rows > 0 {
			fmt.Println("success deleted:)")
		} else {
			fmt.Println("failed deleted :(")
		}
	}
}

// 6. Top up
func TopUpSaldo(db *sql.DB, noTelp string) {
	var user entities.User
	var jumlahTopUp uint64
	var status string
	var addMoney uint64
	fmt.Println("Masukkan jumlah uang yang ingin ditbamahkan ke saldo:")
	fmt.Scanln(&jumlahTopUp)

	if jumlahTopUp < 1000 {
		fmt.Println("miskin jgn belagu, kerja lagi")
	} else {

		row := db.QueryRow("SELECT id, phone, balance FROM users WHERE phone = ?", noTelp)
		if err := row.Scan(&user.ID, &user.Phone, &user.Balance); err != nil {
			if err == sql.ErrNoRows {
				errorRead := fmt.Errorf("Id dengan : %s tidak terdaftar", noTelp)
				fmt.Println(errorRead)
			}
			fmt.Print(err)
		}

		addMoney = user.Balance + jumlahTopUp
		result, errTopUp := db.Exec("UPDATE users SET balance = ? WHERE phone = ?", addMoney, noTelp)
		if errTopUp != nil {
			log.Fatal("Error Top up", errTopUp.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Printf("Top Up sebanyak Rp %d telah berhasil\n", jumlahTopUp)
				fmt.Printf("Saldo saat ini : %d\n", addMoney)
				status = "Berhasil"
			} else {
				fmt.Println("Top Up Failed")
			}
		}

		resultTopUp, errorTopUp := db.Exec("INSERT INTO top_up (user_id, amount, status) VALUES (?, ?, ?)", user.ID, jumlahTopUp, status)
		if errorTopUp != nil {
			log.Fatal("error insert top up data", errorTopUp.Error())
		} else {
			row, _ := resultTopUp.RowsAffected()
			if row > 0 {
				fmt.Println("Insert Top Up data Success")
			} else {
				fmt.Println("Insert Top Up data Failed")
			}
		}
	}

}

// 7.transaksi
func Transaction(db *sql.DB, noTelp string) {
	var user entities.User
	var userPenerima entities.User
	var jumlahTransfer uint64
	var reduceMoney uint64
	var addaMoney uint64
	var statusTransfer string
	fmt.Println("masukkan nomor tujuan:")
	fmt.Scanln(&userPenerima.Phone)

	// scan pengirim
	rows := db.QueryRow("SELECT id, full_name, phone, balance FROM users WHERE phone = ?", noTelp)
	if errPengirim := rows.Scan(&user.ID, &user.FullName, &user.Phone, &user.Balance); errPengirim != nil {
		if errPengirim == sql.ErrNoRows {
			errorRead := fmt.Errorf("Id dengan : %s tidak terdaftar", noTelp)
			fmt.Println(errorRead)
		}
	}
	if user.Balance < 5000 {
		log.Fatal("TOP UP DULU, ANDA MISKIN !")
	}
	if userPenerima.Phone == user.Phone {
		log.Fatal("Maaf tidak bisa mengirim uang ke diri sendiri")
	} else {
		// scan penerima
		row := db.QueryRow("SELECT id, full_name, phone, balance FROM users WHERE phone =?", userPenerima.Phone)
		if errPenerima := row.Scan(&userPenerima.ID, &userPenerima.FullName, &userPenerima.Phone, &userPenerima.Balance); errPenerima != nil {
			if errPenerima == sql.ErrNoRows {
				errorRead := fmt.Errorf("Id dengan : %s tidak terdaftar", userPenerima.Phone)
				fmt.Println(errorRead)
			}
			fmt.Println(errPenerima)
		}

		fmt.Println("pengirim :", user.FullName,"\n", "Penerima :", userPenerima.FullName, "\nSaldo Anda saat ini:", user.Balance)
		fmt.Println("Masukkan Jumlah Transfer :")
		fmt.Scanln(&jumlahTransfer)

		if jumlahTransfer < 1000 {
			log.Fatal("MINIMAL TRANSFER 1000 !!!")
		} else {
			reduceMoney = user.Balance - jumlahTransfer
			result, errTransfer := db.Exec("UPDATE users SET balance = ? WHERE Phone = ?", reduceMoney, user.Phone)
			if errTransfer != nil {
				log.Fatal("ERROR Transfer : ", errTransfer.Error())
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Printf("Transfer sebanyak Rp.%d ke %s telah Berhasil! \n", jumlahTransfer, userPenerima.FullName)
					statusTransfer = "Hore ! Transaksi anda Berhasil !"
				} else {
					fmt.Println("Transaksi Gagal !")
				}
			}

			addaMoney = userPenerima.Balance + jumlahTransfer
			result, errrTransfer := db.Exec("UPDATE users SET balance =? WHERE phone = ?", addaMoney, userPenerima.Phone)
			if errrTransfer != nil {
				log.Fatal("Error Transfer", errrTransfer.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Printf("Transfer sebanyak Rp.%d dari %s, telah Diterima oleh  %s", jumlahTransfer, user.FullName, userPenerima.FullName)
				} else {
					fmt.Println("Yaah, Maaf Transaksi anda Gagal")
				}
			}

			resultTransfer, errorTransfer := db.Exec("INSERT INTO transactions (from_user_id, to_user_id, amount, status) VALUES (?,?,?,?)", user.ID, userPenerima.ID, jumlahTransfer, statusTransfer)
			if errorTransfer != nil {
				log.Fatal("Error Insert Transaction Data", errorTransfer.Error())
			} else {
				row, _ := resultTransfer.RowsAffected()
				if row > 0 {
					fmt.Println("Transfer Success!")
					fmt.Printf("Saldo Anda saat ini : %d", reduceMoney)
				} else {
					fmt.Println("Transfer Failed!!!")
				}
			}
		}
	}
}
