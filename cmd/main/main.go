// Blueprint of shortent link project

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Clear Screen Purpose

func InitClearScreen(clearMap* map[string]func()) {
	(*clearMap)["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	(*clearMap)["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear(clearMap* map[string]func()) {
	value, ok := (*clearMap)[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Terminal anda tidak bisa dihapus!")
	}
}

//

func HashFunc(valueLongURI int) string {
	var willReturn string
	willReturn = ""
	validAlphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var temp int

	for valueLongURI != 0 {
		temp = valueLongURI % 52
		valueLongURI /= 12
		willReturn += string(validAlphabet[temp]);
	}

	return willReturn
}

func InsertingHash(hashShortednedURI* map[string]string, longURI string, shortenedURI* string) {
	var found bool
	validAlphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	linearProbationInt := -1
	linearProbationString := *shortenedURI

	_, found = (*hashShortednedURI)[linearProbationString]

	if (found == true) {
		for found == true && linearProbationInt < 52{
			linearProbationInt++
			if(linearProbationInt == 52) {
				linearProbationString += string(validAlphabet[linearProbationInt - 1]);
				linearProbationInt = 0
			}
			_, found = (*hashShortednedURI)[linearProbationString+string(validAlphabet[linearProbationInt])]
		}
		if (linearProbationInt < 52) {
			linearProbationString += string(validAlphabet[linearProbationInt]);
		}
	}

	*shortenedURI = linearProbationString
	(*hashShortednedURI)[*shortenedURI] = longURI
}

func main() {
	var menu int;
	
	var (
		longURI = "";
		shortenedURI = "";
	)
	
	hashShortenedURI := make(map[string]string)
	clearMap := make(map[string]func())
	InitClearScreen(&clearMap)
	
	for true {
		fmt.Println("Kompiang Gede Sukadharma")
		fmt.Println("       2008561083")
		fmt.Println("Menu")
		fmt.Println("[1] Daftarkan URI yang ingin dipendekkan")
		fmt.Println("[2] Cari URI asli")
		fmt.Println("[3] Hapus URI yang dipendekkan")
		fmt.Println("[4] Lihat semua URI pendek yang dibuat")
		fmt.Println("[5] Exit")
		fmt.Print("Input: ");
		fmt.Scanln(&menu)

		switch menu {
			case 1:
				fmt.Print("Masukkan URI yang ingin dipendekkan : ");
				fmt.Scanln(&longURI);

				if longURI == "" { 
					fmt.Println("URI tidak boleh kosong")
					break
				}
				
				valueLongURI := 0;
				for _, v := range longURI { valueLongURI += int(v) }
				shortenedURI += HashFunc(valueLongURI)

				InsertingHash(&hashShortenedURI, longURI, &shortenedURI)
				fmt.Println("Hasil shortened URI = kompiangg.com/" + shortenedURI)
				fmt.Println("\nNOTE:")
				fmt.Println("Apabila anda mencoba link di atas pada browser")
				fmt.Println("Link tersebut tidak akan dikenali karena ini hanya simulasi")
				shortenedURI = ""
				longURI = ""
				break

			case 2:
				fmt.Print("Masukkan Short URI yang ingin dicari (kompiangg.com juga ditulis) : ");
				fmt.Scanln(&shortenedURI);
				if shortenedURI == "" {
					fmt.Println("URI tidak boleh kosong")
					break
				}

				splittedShortenedURI := strings.Split(shortenedURI, "/")

				if len(splittedShortenedURI) <= 1 {
					fmt.Println("Format Short URI anda salah")
					shortenedURI = ""
					break
				}

				if value, found := hashShortenedURI[splittedShortenedURI[1]] ; found == true{
					fmt.Println("URI yang asli adalah", value)
				} else if found == false {
					fmt.Println("Short URI tidak terdaftar", value)
				}

				shortenedURI = ""
				break

			case 3:
				fmt.Print("Masukkan short URI yang ingin dihapus (kompiangg.com juga ditulis) : ");
				fmt.Scanln(&shortenedURI);
				if shortenedURI == "" { 
					fmt.Println("URI tidak boleh kosong")
					break
				}

				splittedShortenedURI := strings.Split(shortenedURI, "/")
				if len(splittedShortenedURI) <= 1 {
					fmt.Println("Format Short URI anda salah")
					shortenedURI = ""
					break
				}
				
				delete(hashShortenedURI, splittedShortenedURI[1])
				fmt.Println("Short URI berhasil dihapus")
				shortenedURI = ""
				break

			case 4:
				fmt.Println("\nList pemetaan URI")
				for key, value := range hashShortenedURI {
					fmt.Println("kompiangg.com/"+key, "->", value);
				}
				break

			case 5:
				fmt.Println("Exit.")
				break

			default:
				fmt.Println("Menu yang diinput salah")
				break
		}
		if (menu == 5) {break} 
		menu = 0;
		// Clear Screen
		bufio.NewReader(os.Stdin).ReadString('\n')
		CallClear(&clearMap)
	}
}
