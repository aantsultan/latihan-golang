package main

import "fmt"

// defer -> memanggil function di akhir
func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("run application")
}

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()
	if error {
		// panic -> membuat program berhenti
		panic("Uups Error")
	} else {
		fmt.Println("Run App")
	}

}

func endAppRecover() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("error : ", message)
}
func runAppRecover(error bool) {
	defer endAppRecover()
	if error {
		// panic -> membuat program berhenti
		panic("Uups Error Recover")
	} else {
		fmt.Println("Run App Recover")
	}

}

func main() {
	runApplication()

	//runApp(true)
	runApp(false)

	runAppRecover(true)

	fmt.Println("aant sultan")

}
