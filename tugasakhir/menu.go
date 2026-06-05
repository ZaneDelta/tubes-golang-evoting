package main

import (
	"bufio"
	"fmt"
)

func menuAwal(reader *bufio.Reader) bool {
	for {
		clearScreen()
		fmt.Println("============================================================")
		fmt.Println("+++ SETUP BANNER APLIKASI E-VOTING +++")
		fmt.Println("============================================================")
		fmt.Println("1. Pilih Banner")
		fmt.Println("0. Keluar")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			if pilihBanner(reader) {
				fmt.Println("Banner berhasil dipilih.")
				pause(reader)
				return true
			}
		case 0:
			fmt.Println("Program selesai.")
			return false
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func pilihBanner(reader *bufio.Reader) bool {
	banner := [2]string{
		`
    ______            _    ______  ___________   ________
   / ____/           | |  / / __ \/_  __/  _/ | / / ____/
  / __/    ______    | | / / / / / / /  / //  |/ / / __  
 / /___   /_____/    | |/ / /_/ / / / _/ // /|  / /_/ /  
/_____/              |___/\____/ /_/ /___/_/ |_/\____/   
`,
		`                                                                                    
  ▄▄▄▄▄▄▄             ▄▄▄          ▄▄▄▄      ▄▄▄▄▄▄▄    ▄▄▄▄▄▄   ▄▄     ▄▄▄  ▄   ▄▄▄▄ 
 █▀██▀▀▀             █▀██  ██▀▀  ▄█▀▀████▄  █▀▀██▀▀▀▀  █▀ ██     ██▄   ██▀   ▀██████▀ 
   ██                  ██  ██    ██    ██      ██         ██     ███▄  ██      ██   ▄ 
   ████                ██  ██    ██    ██      ██         ██     ██ ▀█▄██      ██  ██ 
   ██        ▀▀▀▀      ██▄ ██    ██    ██      ██         ██     ██   ▀██      ██  ██ 
   ▀█████               ▀███▀     ▀████▀       ▀██▄     ▄▄██▄▄ ▀██▀    ██      ▀█████ 
                                                                               ▄   ██ 
                                                                               ▀████▀ 
`,
	}

	for {
		clearScreen()
		fmt.Println("============================================================")
		fmt.Println("+++ PILIH BANNER E-VOTING +++")
		fmt.Println("============================================================")

		fmt.Println("1. Banner 1")
		fmt.Print(banner[0])

		fmt.Println("2. Banner 2")
		fmt.Print(banner[1])

		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih banner: ")

		switch pilihan {
		case 1:
			bannerDipilih = banner[0]
			return true
		case 2:
			bannerDipilih = banner[1]
			return true
		case 0:
			return false
		default:
			fmt.Println("Pilihan banner tidak tersedia.")
			pause(reader)
		}
	}
}

func menuUtama(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("1. Kelola Kandidat")
		fmt.Println("2. Kelola Pemilih")
		fmt.Println("3. Voting")
		fmt.Println("4. Statistik Hasil Voting")
		fmt.Println("5. Sorting Data Kandidat")
		fmt.Println("6. Searching Kandidat")
		fmt.Println("0. Keluar")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			menuKelolaKandidat(reader)
		case 2:
			menuKelolaPemilih(reader)
		case 3:
			prosesVoting(reader)
		case 4:
			tampilkanStatistik(reader)
		case 5:
			menuSorting(reader)
		case 6:
			menuSearching(reader)
		case 0:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuKelolaKandidat(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("MENU KELOLA KANDIDAT")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Edit Kandidat")
		fmt.Println("4. Hapus Kandidat")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			tambahKandidat(reader)
		case 2:
			tampilkanKandidat()
			pause(reader)
		case 3:
			editKandidat(reader)
		case 4:
			hapusKandidat(reader)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuKelolaPemilih(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("MENU KELOLA PEMILIH")
		fmt.Println("1. Tambah Pemilih")
		fmt.Println("2. Tampilkan Pemilih")
		fmt.Println("3. Edit Pemilih")
		fmt.Println("4. Hapus Pemilih")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			tambahPemilih(reader)
		case 2:
			tampilkanPemilih()
			pause(reader)
		case 3:
			editPemilih(reader)
		case 4:
			hapusPemilih(reader)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuSorting(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("SORTING DATA KANDIDAT")
		fmt.Println("1. Selection Sort berdasarkan jumlah suara")
		fmt.Println("2. Insertion Sort berdasarkan nomor urut")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		if pilihan == 0 {
			return
		}

		if jumlahKandidat == 0 {
			fmt.Println("Data kandidat masih kosong.")
			pause(reader)
			continue
		}

		switch pilihan {
		case 1:
			ascending := pilihUrutanAscending(reader)
			selectionSortJumlahSuara(ascending)
			fmt.Println("Data berhasil diurutkan dengan Selection Sort.")
			tampilkanKandidat()
			pause(reader)
		case 2:
			ascending := pilihUrutanAscending(reader)
			insertionSortNomorUrut(ascending)
			fmt.Println("Data berhasil diurutkan dengan Insertion Sort.")
			tampilkanKandidat()
			pause(reader)
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuSearching(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("SEARCHING KANDIDAT")
		fmt.Println("1. Sequential Search berdasarkan nomor urut")
		fmt.Println("2. Binary Search berdasarkan nomor urut")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		if pilihan == 0 {
			return
		}

		if jumlahKandidat == 0 {
			fmt.Println("Data kandidat masih kosong.")
			pause(reader)
			continue
		}

		nomor := inputIntMin(reader, "Masukkan nomor urut kandidat: ", 1)

		switch pilihan {
		case 1:
			index := sequentialSearchKandidat(nomor)
			if index == -1 {
				fmt.Println("Kandidat tidak ditemukan.")
			} else {
				tampilkanDetailKandidat(index)
			}
			pause(reader)
		case 2:
			insertionSortNomorUrut(true)
			fmt.Println("Data diurutkan ascending terlebih dahulu sebelum Binary Search.")
			index := binarySearchKandidat(nomor)
			if index == -1 {
				fmt.Println("Kandidat tidak ditemukan.")
			} else {
				tampilkanDetailKandidat(index)
			}
			pause(reader)
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}