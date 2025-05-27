package handlers

import (
	"alprocrypto/database"
	"alprocrypto/models"
	"alprocrypto/utils"
	"fmt"
)

func EstimateMining() {
	if database.DB.NCrypto == 0 {
		fmt.Println("Tidak ada aset crypto untuk ditambang.")
		return
	}

	utils.DisplayCryptoList()

	id := utils.ScanNumber("Masukkan ID crypto yang ingin ditambang: ")
	hashrate := utils.ScanFloat("Masukkan hashrate (misal: 100): ")
	duration := utils.ScanFloat("Masukkan durasi (jam): ")

	var selected *models.Crypto
	for i := 0; i < database.DB.NCrypto; i++ {
		if database.DB.Cryptos[i].ID == id {
			selected = &database.DB.Cryptos[i]
			break
		}
	}

	if selected == nil {
		fmt.Println("Crypto tidak ditemukan.")
		return
	}

	minedAmount := (selected.Reward * hashrate / selected.Difficulty) * duration

	database.DB.LastMiningID++
	database.DB.Minings[database.DB.NMining] = models.Mining{
		ID:          database.DB.LastMiningID,
		Hashrate:    hashrate,
		Duration:    duration,
		MinedAmount: minedAmount,
	}
	database.DB.NMining++

	fmt.Printf("Estimasi hasil mining: %.6f %s\n", minedAmount, selected.Name)
}

func ShowMiningReport() {
	if database.DB.NMining == 0 {
		fmt.Println("Belum ada data mining.")
		return
	}

	fmt.Println("\nLaporan Mining:")
	for i := 0; i < database.DB.NMining; i++ {
		m := database.DB.Minings[i]
		fmt.Printf("ID: %d, Hashrate: %.2f, Duration: %.2f jam, Total Mined: %.6f\n",
			m.ID, m.Hashrate, m.Duration, m.MinedAmount)
	}
	utils.ScanString("Tekan Enter untuk melanjutkan...")
}

