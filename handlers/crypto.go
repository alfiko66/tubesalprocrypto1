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

func DisplayCryptos() {
	if database.DB.NCrypto == 0 {
		fmt.Println("Tidak ada crypto yang tersedia.")
		return
	}

	fmt.Println("\nCrypto yang tersedia:")
	for i := 0; i < database.DB.NCrypto; i++ {
		c := database.DB.Cryptos[i]
		fmt.Printf("%d. %s (Difficulty: %.1f, Reward: %.4f)\n", 
			c.ID, c.Name, c.Difficulty, c.Reward)
	}
}