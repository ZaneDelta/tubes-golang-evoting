package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func pilihUrutanAscending(reader *bufio.Reader) bool {
	for {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")

		pilihan := inputInt(reader, "Pilih urutan: ")

		if pilihan == 1 {
			return true
		}

		if pilihan == 2 {
			return false
		}

		fmt.Println("Pilihan urutan tidak tersedia.")
	}
}

func tampilkanHeader() {
	fmt.Println("============================================================")

	if bannerDipilih != "" {
		fmt.Print(bannerDipilih)
	} else {
		fmt.Println("+++ APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL (E-VOTING) +++")
	}

	fmt.Printf("Mode: %s | Kode Project: JEBB-%d\n", eel, jebb24)
	fmt.Println("============================================================")
}

func clearScreen() {
	for i := 0; i < 25; i++ {
		fmt.Println()
	}
}

func pause(reader *bufio.Reader) {
	fmt.Print("\nTekan ENTER untuk melanjutkan...")
	reader.ReadString('\n')
}

func inputString(reader *bufio.Reader, pesan string) string {
	for {
		fmt.Print(pesan)

		teks, _ := reader.ReadString('\n')
		teks = strings.TrimSpace(teks)

		if teks == "" {
			fmt.Println("Input tidak boleh kosong.")
		} else {
			return teks
		}
	}
}

func inputStringBolehKosong(reader *bufio.Reader, pesan string) string {
	fmt.Print(pesan)

	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}

func inputInt(reader *bufio.Reader, pesan string) int {
	for {
		fmt.Print(pesan)

		teks, _ := reader.ReadString('\n')
		teks = strings.TrimSpace(teks)

		angka, err := strconv.Atoi(teks)
		if err != nil {
			fmt.Println("Input harus berupa angka.")
		} else {
			return angka
		}
	}
}

func inputIntMin(reader *bufio.Reader, pesan string, minimal int) int {
	for {
		angka := inputInt(reader, pesan)

		if angka < minimal {
			fmt.Printf("Input minimal adalah %d.\n", minimal)
		} else {
			return angka
		}
	}
}

func inputKonfirmasi(reader *bufio.Reader, pesan string) bool {
	for {
		jawaban := strings.ToLower(inputString(reader, pesan))

		if jawaban == "y" || jawaban == "ya" {
			return true
		}

		if jawaban == "n" || jawaban == "tidak" {
			return false
		}

		fmt.Println("Masukkan y atau n.")
	}
}

func potongTeks(teks string, batas int) string {
	if len(teks) <= batas {
		return teks
	}

	if batas <= 3 {
		return teks[:batas]
	}

	return teks[:batas-3] + "..."
}