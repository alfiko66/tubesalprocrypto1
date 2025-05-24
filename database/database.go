package database

import "alprocrypto/models"

var DB *models.Database

func InitDB() *models.Database {
	DB = &models.Database{
		Cryptos: []models.Crypto{
			{ID: 1, Name: "Bitcoin", Difficulty: 100, Reward: 6.25},
			{ID: 2, Name: "Ethereum", Difficulty: 50, Reward: 2},
			{ID: 3, Name: "Litecoin", Difficulty: 30, Reward: 12.5},
		},
		Minings: []models.Mining{
			{ID: 1, Hashrate: 100, Duration: 24, MinedAmount: 0.001},
		},
		LastCryptoID: 3,
		LastMiningID: 1,
	}
	return DB
}