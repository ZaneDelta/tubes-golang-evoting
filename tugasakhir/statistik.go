package main

import (
	"bufio"
	"fmt"
)

func tampilkanStatistik(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("STATISTIK HASIL VOTING")

	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong.")
		pause(reader)
		return
	}

	totalSuara := hitungTotalSuara()

	fmt.Printf("Total suara masuk: %d\n", totalSuara)
	fmt.Println("------------------------------------------------------------")

	for i := 0; i < jumlahKandidat; i++ {
		persentase := 0.0

		if totalSuara > 0 {
			persentase = float64(dataKandidat[i].jumlahSuara) / float64(totalSuara) * 100
		}

		fmt.Printf("Kandidat %d - %-22s : %.2f%% (%d suara)\n",
			dataKandidat[i].nomorUrut,
			potongTeks(dataKandidat[i].nama, 22),
			persentase,
			dataKandidat[i].jumlahSuara)
	}

	fmt.Println("------------------------------------------------------------")

	if totalSuara == 0 {
		fmt.Println("Belum ada suara masuk.")
	} else {
		suaraTertinggi := cariSuaraTerbanyak()

		fmt.Println("Kandidat dengan suara terbanyak:")
		for i := 0; i < jumlahKandidat; i++ {
			if dataKandidat[i].jumlahSuara == suaraTertinggi {
				fmt.Printf("- Nomor %d, %s dengan %d suara\n",
					dataKandidat[i].nomorUrut,
					dataKandidat[i].nama,
					dataKandidat[i].jumlahSuara)
			}
		}
	}

	pause(reader)
}

func hitungTotalSuara() int {
	total := 0

	for i := 0; i < jumlahKandidat; i++ {
		total += dataKandidat[i].jumlahSuara
	}

	return total
}

func cariSuaraTerbanyak() int {
	if jumlahKandidat == 0 {
		return 0
	}

	maks := dataKandidat[0].jumlahSuara

	for i := 1; i < jumlahKandidat; i++ {
		if dataKandidat[i].jumlahSuara > maks {
			maks = dataKandidat[i].jumlahSuara
		}
	}

	return maks
}