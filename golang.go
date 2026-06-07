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
	banner[2] = (`
8888888888                888     888          888    d8b                   
888                       888     888          888    Y8P                   
888                       888     888          888                          
8888888                   Y88b   d88P  .d88b.  888888 888 88888b.   .d88b.  
888                        Y88b d88P  d88""88b 888    888 888 "88b d88P"88b 
888           888888        Y88o88P   888  888 888    888 888  888 888  888 
888                          Y888P    Y88..88P Y88b.  888 888  888 Y88b 888 
8888888888                    Y8P      "Y88P"   "Y888 888 888  888  "Y88888 
                                                                        888 
                                                                   Y8b d88P 
                                                                    "Y88P"  
`)
	banner[3] = (`
.,::::::         :::      .::.  ...   ::::::::::::::::::.    :::.  .,-:::::/  
;;;;''''         ';;,   ,;;;'.;;;;;;;.;;;;;;;;'''';;;';;;;,  ';;;,;;-'````'   
 [[cccc           \[[  .[[/ ,[[     \[[,   [[     [[[  [[[[[. '[[[[[   [[[[[[/
 $$"""" cccc       Y$c.$$"  $$$,     $$$   $$     $$$  $$$ "Y$c$$"$$c.    "$$ 
 888oo,__           Y88P    "888,_ _,88P   88,    888  888    Y88 'Y8bo,,,o88o
 """"YUMMM           MP       "YMMMMMP"    MMM    MMM  MMM     YM   ''YMUP"YMM
	`)
																			   }



func main() {
	menuAwal()
}
