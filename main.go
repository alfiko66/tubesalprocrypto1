package main

import (
	"alprocrypto/database"
	"alprocrypto/handlers"
	"alprocrypto/utils"
)

func main() {
	database.InitDB()

	for {
		utils.DisplayMainMenu()
		choice := utils.ScanNumber("Pilih menu: ")

		switch choice {
		case 1:
			showManageCryptoMenu()
		case 2:
			utils.DisplayCryptoList()
		case 3:
			handlers.EstimateMining()
		case 4:
			handlers.SearchCrypto()
		case 5:
			showSortMenu()
		case 6:
			utils.DisplayMiningList()
		case 7:
			println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			println("Pilihan tidak valid.")
		}
	}
}

func showManageCryptoMenu() {
	for {
		utils.DisplayManageCryptoMenu()
		choice := utils.ScanNumber("Pilih menu: ")
		switch choice {
		case 1:
			handlers.AddCrypto()
		case 2:
			handlers.EditCrypto()
		case 3:
			handlers.DeleteCrypto()
		case 4:
			return
		default:
			println("Pilihan tidak valid.")
		}
	}
}

func showSortMenu() {
	for {
		utils.DisplaySortMenu()
		choice := utils.ScanNumber("Pilih menu: ")
		switch choice {
		case 1:
			utils.SortByDifficulty()
		case 2:
			utils.SortByReward()
		case 3:
			return
		default:
			println("Pilihan tidak valid.")
		}
	}
}
