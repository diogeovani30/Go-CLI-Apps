package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func dataTeman() map[int]Teman {
	data := map[int]Teman{
		1: {Nama: "Ani", Alamat: "Jl. Mangga No. 123", Pekerjaan: "Programmer", Alasan: "Menarik"},
		2: {Nama: "Budi", Alamat: "Jl. Jeruk No. 456", Pekerjaan: "Designer", Alasan: "Ingin memperdalam pemrograman"},
		3: {Nama: "Citra", Alamat: "Jl. Anggur No. 789", Pekerjaan: "Analyst", Alasan: "Menambah skill baru"},
	}
	return data
}

func getDataTeman(absen int) *Teman {
	daftarTeman := dataTeman()
	teman, exists := daftarTeman[absen]
	if !exists {
		return nil
	}
	return &teman
}

// CLI to parse user prompts in terminal
func parseArgs() (int, error) {
	args := os.Args
	if len(args) != 2 {
		return 0, fmt.Errorf("FAIL: Missing 1 Arguments!\nUsage: go run biodata.go <nomor_absen>")
	}

	absen, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, fmt.Errorf("Error: %v", err)
	}
	return absen, nil
}

func printTeman(teman *Teman) {
	if teman == nil {
		fmt.Println("Data teman tidak ditemukan.")
		return
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}

func main() {
	absen, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	teman := getDataTeman(absen)
	printTeman(teman)
}
