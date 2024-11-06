package main

import (
	"fmt"
)

type Celcius struct {
	suhu float64
}

type Fahrenheit struct {
	suhu float64
}

type Kelvin struct {
	suhu float64
}

func (c Celcius) toCelcius() float64 {
	return c.suhu
}

func (c Celcius) toFahrenheit() float64 {
	suhuFahrenheit := ((9.0 / 5.0) * c.suhu) + 32
	return suhuFahrenheit
}

func (c Celcius) toKelvin() float64 {
	suhuKelvin := c.suhu + 273.15
	return suhuKelvin
}

func (f Fahrenheit) toFahrenheit() float64 {
	return f.suhu
}

func (f Fahrenheit) toCelcius() float64 {
	suhuCelcius := (f.suhu - 32) * (5.0 / 9.0)
	return suhuCelcius
}

func (f Fahrenheit) toKelvin() float64 {
	suhuKelvin := (f.suhu + 459.67) * (5.0 / 9.0)
	return suhuKelvin
}

func (k Kelvin) toKelvin() float64 {
	return k.suhu
}

func (k Kelvin) toCelcius() float64 {
	suhuCelcius := k.suhu - 273.15
	return suhuCelcius
}

func (k Kelvin) toFahrenheit() float64 {
	suhuFahrenheit := (k.suhu * (9.0 / 5.0)) - 459.67
	return suhuFahrenheit
}

type hitungSuhu interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {

	/*
	*
	* CARA 1 :
	*
	 */
	//fmt.Println("Masukkan kalulator suhu yang ingin dipakai")
	//fmt.Println("1. Celcius ke Fahrenheit")
	//fmt.Println("2. Celcius ke Kelvin")
	//fmt.Println("3. Fahrenheit ke Celcius")
	//fmt.Println("4. Fahrenheit ke Kelvin")
	//fmt.Println("5. Kelvin ke Celcius")
	//fmt.Println("6. Kelvin ke Fahrenheit")
	//fmt.Println("masukkan pilihan Anda: ")
	//
	//var pilihan int
	//fmt.Scanf("%d", &pilihan)
	//for pilihan < 1 || pilihan > 6 {
	//	fmt.Println("kalkulator tidak tersedia, masukkan kalkulator pilihan Anda: ")
	//	fmt.Scanf("%d", &pilihan)
	//}
	//
	//var suhu float64
	//fmt.Println("masukkan suhu: ")
	//fmt.Scanf("%f", &suhu)
	//
	//var suhuAkhir float64
	//if pilihan == 1 {
	//	suhuAkhir = CelciusToFahrenheit(suhu)
	//} else if pilihan == 2 {
	//	suhuAkhir = CelciusToKelvin(suhu)
	//} else if pilihan == 3 {
	//	suhuAkhir = FahrenheitToCelcius(suhu)
	//} else if pilihan == 4 {
	//	suhuAkhir = FahrenheitToKelvin(suhu)
	//} else if pilihan == 5 {
	//	suhuAkhir = KelvinToCelcius(suhu)
	//} else {
	//	suhuAkhir = KelvinToFahrenheit(suhu)
	//}
	//
	//fmt.Printf("Suhu akhir hasil konversi adalah: %.2f\n", suhuAkhir)

	fmt.Println("Pilihan suhu awal")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Masukkan suhu awal yang diinginkan: ")

	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)
	for suhuAwal < 1 || suhuAwal > 3 {
		fmt.Println("suhu awal tidak valid, Masukkan suhu awal yang diinginkan: ")
		fmt.Scanf("%d", &suhuAwal)
	}

	fmt.Println("Pilihan suhu akhir")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Masukkan suhu akhir yang diinginkan: ")

	var suhuAkhir int
	fmt.Scanf("%d", &suhuAkhir)
	for suhuAkhir < 1 || suhuAkhir > 3 {
		fmt.Println("suhu akhir tidak valid, Masukkan suhu akhir yang diinginkan: ")
		fmt.Scanf("%d", &suhuAkhir)
	}

	var suhu float64
	fmt.Println("Masukkan suhu: ")
	fmt.Scanf("%d", &suhu)

	var interfaceSuhu hitungSuhu
	switch suhuAwal {
	case 1:
		interfaceSuhu = Celcius{suhu: suhu}
		break
	case 2:
		interfaceSuhu = Fahrenheit{suhu: suhu}
		break
	case 3:
		interfaceSuhu = Kelvin{suhu: suhu}
		break
	}

	var suhuFinal float64
	switch suhuAkhir {
	case 1:
		suhuFinal = interfaceSuhu.toCelcius()
		break
	case 2:
		suhuFinal = interfaceSuhu.toFahrenheit()
		break
	case 3:
		suhuFinal = interfaceSuhu.toKelvin()
		break
	}

	fmt.Printf("suhu akhir yang di dapat adalah: %.2f", suhuFinal)
}

func CelciusToFahrenheit(suhuCelcius float64) float64 {
	suhuFahrenheit := ((9.0 / 5.0) * suhuCelcius) + 32
	return suhuFahrenheit
}

func CelciusToKelvin(suhuCelcius float64) float64 {
	suhuKelvin := suhuCelcius + 273.15
	return suhuKelvin
}

func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	suhuCelcius := (suhuFahrenheit - 32) * (5.0 / 9.0)
	return suhuCelcius
}

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	suhuKelvin := (suhuFahrenheit + 459.67) * (5.0 / 9.0)
	return suhuKelvin
}

func KelvinToCelcius(suhuKelvin float64) float64 {
	suhuCelcius := suhuKelvin - 273.15
	return suhuCelcius
}

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	suhuFahrenheit := (suhuKelvin * (9.0 / 5.0)) - 459.67
	return suhuFahrenheit
}
