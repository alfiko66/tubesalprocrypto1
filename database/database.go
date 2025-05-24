package database

import "alprocrypto/models"

var DB models.Database

func InitDB() {
	cryptos := [3]models.Crypto{
		{ID: 1, Name: "Bitcoin", Difficulty: 100, Reward: 6.25},
		{ID: 2, Name: "Ethereum", Difficulty: 50, Reward: 2},
		{ID: 3, Name: "Litecoin", Difficulty: 30, Reward: 12.5},
	}

	minings := [1]models.Mining{
		{ID: 1, Hashrate: 100, Duration: 24, MinedAmount: 0.001},
	}

	for i := 0; i < 3; i++ {
		DB.Cryptos[i] = cryptos[i]
	}
	for i := 0; i < 1; i++ {
		DB.Minings[i] = minings[i]
	}

	DB.NCrypto = 3
	DB.NMining = 1
	DB.LastCryptoID = 3
	DB.LastMiningID = 1
}