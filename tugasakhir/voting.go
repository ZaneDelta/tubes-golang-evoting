package main

import (
	"bufio"
	"fmt"
	"time"
)

func prosesVoting(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("VOTING")

	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong, voting belum bisa dilakukan.")
		pause(reader)
		return
	}

	if jumlahPemilih == 0 {
		fmt.Println("Data pemilih masih kosong, voting belum bisa dilakukan.")
		pause(reader)
		return
	}

	tampilkanPemilih()

	id := inputIntMin(reader, "Masukkan ID pemilih: ", 1)
	indexPemilih := sequentialSearchPemilih(id)

	if indexPemilih == -1 {
		fmt.Println("ID pemilih tidak ditemukan.")
		pause(reader)
		return
	}

	if dataPemilih[indexPemilih].sudahMemilih {
		fmt.Println("Pemilih ini sudah menggunakan hak suara.")
		pause(reader)
		return
	}

	fmt.Println()
	tampilkanKandidat()

	nomor := inputIntMin(reader, "Pilih nomor urut kandidat: ", 1)
	indexKandidat := sequentialSearchKandidat(nomor)

	if indexKandidat == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
		pause(reader)
		return
	}

	dataKandidat[indexKandidat].jumlahSuara++
	dataPemilih[indexPemilih].sudahMemilih = true

	waktuVoting := time.Now().Format("02-01-2006 15:04:05")

	fmt.Println("Voting berhasil.")
	fmt.Printf("Pemilih      : %s\n", dataPemilih[indexPemilih].nama)
	fmt.Printf("Kandidat     : %s\n", dataKandidat[indexKandidat].nama)
	fmt.Printf("Waktu voting : %s\n", waktuVoting)

	pause(reader)
}