package main

import "fmt"

func main() {
	
	var ujian = 90
	var absensi = 80

	var lulusNilaiAkhir = ujian > 75
	var lulusAbsensi = absensi > 70
	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(ujian >= 75 && absensi >= 70)
}