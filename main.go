package main

import (
	"fmt"
)

func main() {
	var pabrik, pasar, hasil int
	fmt.Println("Metode North West Corner")
	fmt.Print("Masukkan jumlah pabrik: ")
	fmt.Scan(&pabrik)
	fmt.Print("Masukkan jumlah pasar: ")
	fmt.Scan(&pasar)

	biaya := make([][]int, pabrik)
	supply := make([]int, pabrik)
	demand := make([]int, pasar)
	bil := make([][]int, pabrik)

	fmt.Println("Masukkan biaya dari setiap pabrik dan pasar:")
	for i := 0; i < pabrik; i++ {
		biaya[i] = make([]int, pasar)
		for j := 0; j < pasar; j++ {
			fmt.Printf("Biaya[%d][%d]: ", i+1, j+1)
			fmt.Scan(&biaya[i][j])
		}
	}

	fmt.Println("\nMasukkan penawaran dari setiap pabrik:")
	for i := 0; i < pabrik; i++ {
		fmt.Printf("Penawaran[%d]: ", i+1)
		fmt.Scan(&supply[i])
	}

	fmt.Println("\nMasukkan permintaan dari setiap pasar:")
	for j := 0; j < pasar; j++ {
		fmt.Printf("Permintaan[%d]: ", j+1)
		fmt.Scan(&demand[j])
	}

	fmt.Println("\nHasil dalam tabel transportasi:")
	for i := 0; i < pabrik; i++ {
		bil[i] = make([]int, pasar)
		for j := 0; j < pasar; j++ {
			if i == 0 && j == 0 {
				if supply[i] <= demand[j] {
					bil[i][j] = supply[i]
					supply[i] -= bil[i][j]
					demand[j] -= bil[i][j]
				} else {
					bil[i][j] = demand[j]
					supply[i] -= bil[i][j]
					demand[j] -= bil[i][j]
				}
			} else {
				if supply[i] <= demand[j] {
					bil[i][j] = supply[i]
					supply[i] -= bil[i][j]
					demand[j] -= bil[i][j]
				} else {
					bil[i][j] = demand[j]
					supply[i] -= bil[i][j]
					demand[j] -= bil[i][j]
				}
			}
		}
	}

	for i := 0; i < pabrik; i++ {
		for j := 0; j < pasar; j++ {
			fmt.Printf("%d   ", bil[i][j])
		}
		fmt.Println()
	}

	hasil = 0
	for i := 0; i < pabrik; i++ {
		for j := 0; j < pasar; j++ {
			hasil += biaya[i][j] * bil[i][j]
		}
	}

	fmt.Println("Biaya minimum transportasi:", hasil)
}
