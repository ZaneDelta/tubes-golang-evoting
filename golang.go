package main 
import (
	"fmt"

)

/*
-Voting
-Kelola Kandidat
	-Menambah Kandidat
	-Mengedit Kandidat
	-Menghapus Kandidat
-Urutkan Kandidat
	-Selection Sort
	-Insertion Sort
-Cari Kandidat
	-Sequential Search
	-Binary Search
-Statistik Suara
-Keluar
*/

type Kandidat struct {
	nomorUrut   int
	nama        string
	visiMisi    string
	jumlahSuara int
}

// var dataKandidat [MaksData]Kandidat
// var jumlahKandidat int

/* func datadummy() {
	dataKandidat[0] = Kandidat{
		nomorUrut:   1,
		nama:          "Ilman Baruna Sigma",
		visiMisi:      "Kudus Kota Kretek",
		jumlahSuara: 0,
	}
	dataKandidat[1] = Kandidat{
		nomorUrut:   2,
		nama:          "Arjuna Muhammad Arby Sigma",
		visiMisi:      "Nganjuk Kota Kecil",
		jumlahSuara: 0,
	}
	dataKandidat[2] = Kandidat{
		nomorUrut:   3,
		nama:          "Fiandra Lazuart Sigma",
		visiMisi:      "Situbondo Kota Bahari",
		jumlahSuara: 0,
	}
}
*/
func menuAwal(){
	var banner[] string
	banner[0] = (`
    ______            _    ______  ___________   ________
   / ____/           | |  / / __ \/_  __/  _/ | / / ____/
  / __/    ______    | | / / / / / / /  / //  |/ / / __  
 / /___   /_____/    | |/ / /_/ / / / _/ // /|  / /_/ /  
/_____/              |___/\____/ /_/ /___/_/ |_/\____/   
	`)
	banner[1] = (`                                                                                    
  ▄▄▄▄▄▄▄             ▄▄▄          ▄▄▄▄      ▄▄▄▄▄▄▄    ▄▄▄▄▄▄   ▄▄     ▄▄▄  ▄   ▄▄▄▄ 
 █▀██▀▀▀             █▀██  ██▀▀  ▄█▀▀████▄  █▀▀██▀▀▀▀  █▀ ██     ██▄   ██▀   ▀██████▀ 
   ██                  ██  ██    ██    ██      ██         ██     ███▄  ██      ██   ▄ 
   ████                ██  ██    ██    ██      ██         ██     ██ ▀█▄██      ██  ██ 
   ██        ▀▀▀▀      ██▄ ██    ██    ██      ██         ██     ██   ▀██      ██  ██ 
   ▀█████               ▀███▀     ▀████▀       ▀██▄     ▄▄██▄▄ ▀██▀    ██      ▀█████ 
                                                                               ▄   ██ 
                                                                               ▀████▀ 
`)
																			   }



func main() {
	menuAwal()
}







