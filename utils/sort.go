package utils

import (
	"alprocrypto/database"
	"fmt"
)

func SortByDifficulty() {
	n := database.DB.NCrypto
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if database.DB.Cryptos[j].Difficulty < database.DB.Cryptos[minIdx].Difficulty {
				minIdx = j
			}
		}
		database.DB.Cryptos[i], database.DB.Cryptos[minIdx] = database.DB.Cryptos[minIdx], database.DB.Cryptos[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Difficulty (Selection Sort).")
}

func SortByReward() {
	n := database.DB.NCrypto
	for i := 1; i < n; i++ {
		key := database.DB.Cryptos[i]
		j := i - 1
		for j >= 0 && database.DB.Cryptos[j].Reward > key.Reward {
			database.DB.Cryptos[j+1] = database.DB.Cryptos[j]
			j--
		}
		database.DB.Cryptos[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Reward (Insertion Sort).")
}
