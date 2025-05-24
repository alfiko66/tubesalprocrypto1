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