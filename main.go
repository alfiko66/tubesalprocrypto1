package main

import (
	"alprocrypto/database"
	"alprocrypto/handlers"
	"alprocrypto/utils"
	"fmt"
)
func main() {
	database.InitDB()

	for {
		fmt.Print(`
Simulasi Penambangan Crypto
1. Lihat Daftar Crypto
2. Tambah Crypto
3. Tambang
4. Laporan
0. Keluar
Pilihan Anda: `)

		switch utils.ScanNumber("") {
		case 1:
			handlers.AddCrypto()
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
