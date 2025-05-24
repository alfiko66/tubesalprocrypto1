package models

const NMAX = 1000

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
	Cryptos      [NMAX]Crypto
	Minings      [NMAX]Mining
	NCrypto      int
	NMining      int
	LastCryptoID int
	LastMiningID int
}