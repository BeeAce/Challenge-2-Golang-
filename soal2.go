package main

import "fmt"

// Fungsi untuk menghitung BMI
func hitungBMI(massa, tinggi float64) float64 {
	return massa / (tinggi * tinggi)
}

func main() {
	// Data Uji 1
	massaMark1, tinggiMark1 := 78.0, 1.69
	massaJohn1, tinggiJohn1 := 92.0, 1.95

	// Hitung BMI untuk Mark dan John (Data Uji 1)
	bmiMark1 := hitungBMI(massaMark1, tinggiMark1)
	bmiJohn1 := hitungBMI(massaJohn1, tinggiJohn1)

	// Menentukan apakah Mark memiliki BMI lebih tinggi dari John (Data Uji 2)
	markHigherBMI1 := bmiMark1 > bmiJohn1

	fmt.Println("=====Data 1=====")

	//Untuk menampilkan hasil BMI
	fmt.Printf("BMI Mark: %.2f\n", bmiMark1)
	fmt.Printf("BMI John: %.2f\n", bmiJohn1)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John?", markHigherBMI1)

	// Data Uji 2
	massaMark2, tinggiMark2 := 95.0, 1.88
	massaJohn2, tinggiJohn2 := 85.0, 1.76

	// Hitung BMI untuk Mark dan John (Data Uji 2)
	bmiMark2 := hitungBMI(massaMark2, tinggiMark2)
	bmiJohn2 := hitungBMI(massaJohn2, tinggiJohn2)

	// Menentukan apakah Mark memiliki BMI lebih tinggi dari John (Data Uji 2)
	markHigherBMI2 := bmiMark2 > bmiJohn2

	fmt.Println("=====Data 2=====")

	//Untuk menampilkan hasil BMI
	fmt.Printf("BMI Mark: %.2f\n", bmiMark2)
	fmt.Printf("BMI John: %.2f\n", bmiJohn2)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John?", markHigherBMI2)
}
