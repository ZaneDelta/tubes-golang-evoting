package main

// proses selection sort manual berdasarkan jumlah suara (gabriel edbert)
func selectionSortJumlahSuara(ascending bool) {
	for i := 0; i < jumlahKandidat-1; i++ {
		posisi := i

		for j := i + 1; j < jumlahKandidat; j++ {
			if ascending {
				if dataKandidat[j].jumlahSuara < dataKandidat[posisi].jumlahSuara {
					posisi = j
				}
			} else {
				if dataKandidat[j].jumlahSuara > dataKandidat[posisi].jumlahSuara {
					posisi = j
				}
			}
		}

		if posisi != i {
			temp := dataKandidat[i]
			dataKandidat[i] = dataKandidat[posisi]
			dataKandidat[posisi] = temp
		}
	}
}

// proses insertion sort manual berdasarkan nomor urut (gabriel edbert)
func insertionSortNomorUrut(ascending bool) {
	for i := 1; i < jumlahKandidat; i++ {
		key := dataKandidat[i]
		j := i - 1

		if ascending {
			for j >= 0 && dataKandidat[j].nomorUrut > key.nomorUrut {
				dataKandidat[j+1] = dataKandidat[j]
				j--
			}
		} else {
			for j >= 0 && dataKandidat[j].nomorUrut < key.nomorUrut {
				dataKandidat[j+1] = dataKandidat[j]
				j--
			}
		}

		dataKandidat[j+1] = key
	}
}

// proses mencari kandidat berdasarkan nomor urut secara sequential (gabriel edbert)
func sequentialSearchKandidat(nomor int) int {
	for i := 0; i < jumlahKandidat; i++ {
		if dataKandidat[i].nomorUrut == nomor {
			return i
		}
	}

	return -1
}

// proses binary search setelah data kandidat diurutkan ascending (gabriel edbert)
func binarySearchKandidat(nomor int) int {
	kiri := 0
	kanan := jumlahKandidat - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if dataKandidat[tengah].nomorUrut == nomor {
			return tengah
		} else if dataKandidat[tengah].nomorUrut < nomor {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

// proses mencari pemilih berdasarkan id secara sequential (gabriel edbert)
func sequentialSearchPemilih(id int) int {
	for i := 0; i < jumlahPemilih; i++ {
		if dataPemilih[i].idPemilih == id {
			return i
		}
	}

	return -1
}