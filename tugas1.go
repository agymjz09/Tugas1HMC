package main

import "fmt"

func kurang (d, c int) int {
	return d - c
}
func tambah (d, c int) int {
	return d + c
}
func bagi (d, c float64) float64 {
	return d / c
}
func kali (d, c int) int {
	return d * c
}
func modul (d, c int) int {
	return d % c
} 
func volumeSilinder (d, c float64) float64 {
	return 3.14 * d * d * c
}
func volumeBola (d float64) float64 {
	return 1.33 * 3.14 * d * d * d
}
func kelilingPersegiPanjang (d, c int) int {
	return 2 * (d + c)
}
func main() {
	result1 := kurang (9, 4)
	result2 := tambah (5, 11)
	result3 := bagi (21, 7)
	result4 := kali (3, 10)
	result5 := modul (16, 3) 
	result6 := volumeSilinder(8, 5)
	result7 := volumeBola(7)
	result8 := kelilingPersegiPanjang(10, 1)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)	
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
	fmt.Println(result8)
}