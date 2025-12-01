package main // mendefinisikan package utama agar bisa dieksekusi

import "fmt" // mengimpor package fmt untuk output ke terminal

func main() { // fungsi utama tempat eksekusi program dimulai

	fmt.Println("===FUNDAMENTAL===")
	fmt.Println("Hello World")     // mencetak teks "Hello World"
	fmt.Println("Angka satu :", 1) // mencetak teks dan angka
	fmt.Println("Benar :", true)   // mencetak teks dan nilai boolean
	fmt.Println("Ryan Besto")      // mencetak nama kamu
	fmt.Println()

	// menghitung panjang string
	fmt.Println("===STRING===")
	fmt.Println(len("Halo, nama saya besto")) // spasi juga dihitung
	fmt.Println("Nama saya besto"[0])         // mengambil byte dari huruf pertama
	fmt.Println("Nama saya besto"[1])         // mengambil byte dari huruf kedua
	fmt.Println()

	// deklarasi + assignment digabung sesuai saran (S1021)
	fmt.Println("===VARIABEL===")
	// var name string
	// name = "besto saragih"
	// fmt.Println(name)
	// penulisan di atas bisa disingkat jadi

	name := "besto saragih"
	fmt.Println(name)

	name = "Ryan Besto"
	fmt.Println(name)
	fmt.Println()

	var (
		firstname = "Ryan"
		middlename = "Besto"
		lastname = "Saragih"
	)
	
	fmt.Println(firstname,middlename,lastname) //Ryan Besto Saragih
	fmt.Println()

	//Constant

}
