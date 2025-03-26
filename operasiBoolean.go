package main

import "fmt"

func OperasiBoolean() {
	nilaiAkhir := 80
	absensi := 95

	lulusNilaiAkhir := nilaiAkhir >= 80
	lulusAbsensi := absensi >= 90

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
