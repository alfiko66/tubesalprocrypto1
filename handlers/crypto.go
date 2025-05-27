package handlers

import (
	"alprocrypto/database"
	"alprocrypto/models"
	"alprocrypto/utils"
	"fmt"
)

func AddCrypto() {

	name := utils.ScanString("Masukkan nama crypto: ")
	difficulty := utils.ScanFloat("Masukkan tingkat kesulitan: ")
	reward := utils.ScanFloat("Masukkan reward: ")
	
	database.DB.LastCryptoID++
	crypto := models.Crypto{
		ID:         database.DB.LastCryptoID,
		Name:       name,
		Difficulty: difficulty,
		Reward:     reward,
	}

	database.DB.Cryptos[database.DB.NCrypto] = crypto
	database.DB.NCrypto++
	fmt.Println("Crypto berhasil ditambahkan")
}

func EditCrypto() {
	utils.DisplayCryptoList()
	id := utils.ScanNumber("Masukkan ID crypto yang ingin diubah: ")

	var index = -1
	for i := 0; i < database.DB.NCrypto; i++ {
		if database.DB.Cryptos[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Crypto tidak ditemukan.")
		return
	}

	name := utils.ScanString("Masukkan nama baru: ")
	difficulty := utils.ScanFloat("Masukkan tingkat kesulitan baru: ")
	reward := utils.ScanFloat("Masukkan reward baru: ")

	database.DB.Cryptos[index].Name = name
	database.DB.Cryptos[index].Difficulty = difficulty
	database.DB.Cryptos[index].Reward = reward

	fmt.Println("Crypto berhasil diubah.")
}

func DeleteCrypto() {
	utils.DisplayCryptoList()
	id := utils.ScanNumber("Masukkan ID crypto yang ingin dihapus: ")

	index := -1
	for i := 0; i < database.DB.NCrypto; i++ {
		if database.DB.Cryptos[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Crypto tidak ditemukan.")
		return
	}

	for i := index; i < database.DB.NCrypto-1; i++ {
		database.DB.Cryptos[i] = database.DB.Cryptos[i+1]
	}
	database.DB.NCrypto--

	fmt.Println("Crypto berhasil dihapus.")
}

func SearchCrypto() {
	name := utils.ScanString("Masukkan nama crypto yang dicari: ")
	found := false

	fmt.Println("Hasil pencarian:")
	for i := 0; i < database.DB.NCrypto; i++ {
		if database.DB.Cryptos[i].Name == name {
			utils.DisplayOneCrypto(i)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada crypto yang cocok.")
	}
	utils.ScanString("Tekan Enter untuk melanjutkan...")
}



