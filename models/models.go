package models

type Crypto struct {
	ID         int
	Name       string
	Difficulty float64
	Reward     float64
}

type Mining struct {
	ID          int
	Hashrate    float64
	Duration    float64
	MinedAmount float64
}

type Database struct {
	Cryptos   []Crypto
	Minings   []Mining
	LastCryptoID int
	LastMiningID int
}