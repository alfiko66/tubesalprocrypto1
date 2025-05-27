package utils

import (
	"alprocrypto/database"
	"fmt"
)

func DisplayMainMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|               SIMULASI CRYPTO MINING SEDERHANA            |
+-----------------------------------------------------------+
| 1. MANAGE CRYPTO                                           |
| 2. DISPLAY ALL CRYPTOS                                     |
| 3. ESTIMATE MINING                                         |
| 4. SEARCH CRYPTO                                           |
| 5. SORT CRYPTOS                                            |
| 6. MINING REPORT                                           |
| 7. EXIT                                                    |
+-----------------------------------------------------------+
`)
}

func DisplayManageCryptoMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                    MANAGE CRYPTO MENU                     |
+-----------------------------------------------------------+
| 1. ADD CRYPTO                                              |
| 2. MODIFY CRYPTO                                           |
| 3. DELETE CRYPTO                                           |
| 4. BACK                                                    |
+-----------------------------------------------------------+
`)
}

func DisplaySortMenu() {
	fmt.Printf(`
+-----------------------------------------------------------+
|                   SORT CRYPTOS BY FIELD                   |
+-----------------------------------------------------------+
| 1. BY DIFFICULTY  - Selection Sort                   			|
| 2. BY REWARD      - Insertion Sort                   			|
| 3. BACK                                                   |
+-----------------------------------------------------------+
`)
}

func DisplayCryptoList() {
	fmt.Printf(`
+------------------------------------------------------------------+
|                    DAFTAR ASET CRYPTO                     			|
+------------------------------------------------------------------+
`)

	for i := 0; i < database.DB.NCrypto; i++ {
		crypto := database.DB.Cryptos[i]
		fmt.Printf("| ID: %-3d Nama: %-15s Difficulty: %-6.2f Reward: %-7.2f |\n",
			crypto.ID, crypto.Name, crypto.Difficulty, crypto.Reward)
	}
	fmt.Println("+------------------------------------------------------------------+")
}

func DisplayMiningList() {
	fmt.Printf(`
+------------------------------------------------------------------+
|                    DAFTAR ASET CRYPTO                     			|
+------------------------------------------------------------------+
`)

	for i := 0; i < database.DB.NMining; i++ {
		mining := database.DB.Minings[i]
		fmt.Printf("| ID: %-3d Nama: %-15f Difficulty: %-6.2f Reward: %-7.2f |\n",
			mining.ID, mining.Hashrate, mining.Duration, mining.MinedAmount)
	}
	fmt.Println("+------------------------------------------------------------------+")
}

func DisplayOneCrypto(i int) {
	fmt.Printf(`
+------------------------------------------------------------------+
|                    DAFTAR ASET CRYPTO                     			|
+------------------------------------------------------------------+
`)
		crypto := database.DB.Cryptos[i]
		fmt.Printf("| ID: %-3d Nama: %-15s Difficulty: %-6.2f Reward: %-7.2f |\n",
			crypto.ID, crypto.Name, crypto.Difficulty, crypto.Reward)
	fmt.Println("+------------------------------------------------------------------+")
}

