package main

import "fmt"

// Fungsi untuk menghitung skor rata-rata
func hitungRataRata(skor []int) float64 {
	total := 0
	for _, nilai := range skor {
		total += nilai
	}
	return float64(total) / float64(len(skor))
}

func main() {
	// Skor untuk setiap tim (Data 1)
	skorLumbaLumba := []int{96, 108, 89}
	skorKoala := []int{88, 91, 110}

	// Hitung skor rata-rata
	rataRataLumbaLumba := hitungRataRata(skorLumbaLumba)
	rataRataKoala := hitungRataRata(skorKoala)

	// Skor minimum
	skorMinimum := 100

	fmt.Println("=====Data 1=====")

	// Untuk menampilkan skor rata-rata dari setiap tim
	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", rataRataLumbaLumba)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", rataRataKoala)

	// Cek pemenang (Data 1)
	if rataRataLumbaLumba > rataRataKoala && rataRataLumbaLumba >= float64(skorMinimum) {
		fmt.Println("Pemenang pada Data 1: Tim Lumba-lumba")
	} else if rataRataKoala > rataRataLumbaLumba && rataRataKoala >= float64(skorMinimum) {
		fmt.Println("Pemenang pada Data 1: Tim Koala")
	} else if rataRataLumbaLumba == rataRataKoala && rataRataLumbaLumba >= float64(skorMinimum) && rataRataKoala >= float64(skorMinimum) {
		fmt.Println("Seri pada Data 1: Kedua tim memiliki skor rata-rata yang sama")
	} else {
		fmt.Println("Tidak ada Pemenang pada Data 1: Skor rata-rata di bawah minimum")
	}

	// Skor untuk setiap tim (Data Bonus 1)
	skorLumbaLumbaBonus1 := []int{97, 112, 101}
	skorKoalaBonus1 := []int{109, 95, 123}

	// Hitung skor rata-rata (Data Bonus 1)
	rataRataLumbaLumbaBonus1 := hitungRataRata(skorLumbaLumbaBonus1)
	rataRataKoalaBonus1 := hitungRataRata(skorKoalaBonus1)

	fmt.Println("=====Data Bonus 1=====")

	// Untuk menampilkan skor rata-rata dari setiap tim
	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", rataRataLumbaLumbaBonus1)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", rataRataKoalaBonus1)

	// Cek pemenang (Data Bonus 1)
	if rataRataLumbaLumbaBonus1 > rataRataKoalaBonus1 && rataRataLumbaLumbaBonus1 >= float64(skorMinimum) {
		fmt.Println("Pemenang pada Data Bonus 1: Tim Lumba-lumba")
	} else if rataRataKoalaBonus1 > rataRataLumbaLumbaBonus1 && rataRataKoalaBonus1 >= float64(skorMinimum) {
		fmt.Println("Pemenang pada Data Bonus 1: Tim Koala")
	} else if rataRataLumbaLumbaBonus1 == rataRataKoalaBonus1 && rataRataLumbaLumbaBonus1 >= float64(skorMinimum) && rataRataKoalaBonus1 >= float64(skorMinimum) {
		fmt.Println("Seri pada Data Bonus 1: Kedua tim memiliki skor rata-rata yang sama")
	} else {
		fmt.Println("Tidak ada Pemenang pada Data Bonus 1: Skor rata-rata di bawah minimum")
	}

	// Skor untuk setiap tim (Data Bonus 2)
	skorLumbaLumbaBonus2 := []int{97, 112, 101}
	skorKoalaBonus2 := []int{109, 95, 106}

	// Hitung skor rata-rata (Data Bonus 2)
	rataRataLumbaLumbaBonus2 := hitungRataRata(skorLumbaLumbaBonus2)
	rataRataKoalaBonus2 := hitungRataRata(skorKoalaBonus2)

	fmt.Println("=====Data Bonus 2=====")

	// Untuk menampilkan skor rata-rata dari setiap tim
	fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", rataRataLumbaLumbaBonus2)
	fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", rataRataKoalaBonus2)

	// Cek pemenang (Data Bonus 2)
	if rataRataLumbaLumbaBonus2 > rataRataKoalaBonus2 && rataRataLumbaLumbaBonus2 >= float64(skorMinimum) {
		fmt.Println("Pemenang pada Data Bonus 2: Tim Lumba-lumba")
	} else if rataRataKoalaBonus2 > rataRataLumbaLumbaBonus2 && rataRataKoalaBonus2 >= float64(skorMinimum) {
		fmt.Println("Pemenang pada Data Bonus 2: Tim Koala")
	} else if rataRataLumbaLumbaBonus2 == rataRataKoalaBonus2 && rataRataLumbaLumbaBonus2 >= float64(skorMinimum) && rataRataKoalaBonus2 >= float64(skorMinimum) {
		fmt.Println("Seri pada Data Bonus 2: Kedua tim memiliki skor rata-rata yang sama")
	} else {
		fmt.Println("Tidak ada Pemenang pada Data Bonus 2: Skor rata-rata di bawah minimum")
	}
}
