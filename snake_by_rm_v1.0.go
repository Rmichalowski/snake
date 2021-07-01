package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
	"github.com/eiannone/keyboard"
)

var ekran [60][30]string
var snake [60][30]string
var ogon [60][30]int32
var dlugosc = 4 + 1
var w = len(ekran)
var h = len(ekran[0])
var headposx = w / 2
var headposy = h / 2
var acmov = 'd'
var temp byte
var run = 1
var frametime int = 100
var owoc1 = 0
var owoc2 = 0
var owoc1x = 0
var owoc1y = 0
var owoc2x = 0
var owoc2y = 0
var pause = false

func zerujekran() {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ekran[x][y] = " "
		}

	}
}

func liniagorna() {
	ekran[0][0] = "x"
	for i := 1; i < w-1; i++ {
		ekran[i][0] = "_"
	}
	ekran[w-1][0] = "x"
}

func liniadolna() {
	ekran[0][h-1] = "x"
	for i := 1; i < w-1; i++ {
		ekran[i][h-1] = "_"
	}
	ekran[w-1][h-1] = "x"
}

func liniaprawa() {
	for i := 1; i < h-1; i++ {
		ekran[w-1][i] = "|"
	}
}

func linialewa() {

	for i := 1; i < h-1; i++ {
		ekran[0][i] = "|"
	}
}

func snakehead() {
	ogon[headposx][headposy] = int32(dlugosc-1) 

	if acmov == 'w' {
		headposy--
	}
	if acmov == 'a' {
		headposx--
		headposx--
	}
	if acmov == 'd' {
		headposx++
		headposx++
	}
	if acmov == 's' {
		headposy++
	}

	
	ekran[headposx][headposy] = "O"
	if headposx == 0 || headposy == 0 || headposx == w-2 || headposy == h-1 {
		fmt.Println("PRZEGRAŁEŚ!")
		time.Sleep(time.Second * 2)
		os.Exit(3)
	}

}

func scrclr() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func drukujekran(wg *sync.WaitGroup) {

	for run == 1 {
		var wyn = ""
		zerujekran()
		snakehead()
		liniagorna()
		liniadolna()
		linialewa()
		liniaprawa()

		owocki()

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if ogon[x][y] > 0  {
					if ekran[x][y] == "O"{
						fmt.Println("PRZEGRAŁEŚ!")
						time.Sleep(time.Second * 2)
						os.Exit(3)
					}

					ekran[x][y] = "x"
				

				}
			}
		}
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if ogon[x][y] > 0 {
					ogon[x][y]--
				}
			}
		}

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				wyn = wyn + ekran[x][y]
			}
			wyn = wyn + "\n"
		}

		fmt.Print(wyn)
		fmt.Print("\n\n DEBUGER:")
		fmt.Print("\n\n headpos =", headposx, "/", headposy)
		fmt.Print("\n owoc1 =", owoc1x, "/", owoc1y)
		fmt.Print("\n owoc2 =", owoc2x, "/", owoc2y)
		fmt.Print("\n dlugosc =", dlugosc)

		time.Sleep(20 * time.Millisecond)
		scrclr()
		
		if pause == true {
			fmt.Println("Paused, press Enter to continue...")
			fmt.Scanln()
			pause = false
		}
	}
	wg.Done()
}

func owocki() {
	var maxx = len(ekran) - 1
	var maxy = len(ekran[0]) - 1
	var min = 2
	if owoc1 == 0 {
		owoc1x = (((rand.Intn(maxx-min) + min) / 2) * 2)
		owoc1y = ((rand.Intn(maxy-min) + min) / 2) * 2
		owoc1 = rand.Intn(300-40) + 40

	} else {
		ekran[owoc1x][owoc1y] = "@"
		owoc1--
		if headposx == owoc1x && headposy == owoc1y {
			dlugosc++
			owoc1 = 0
		}
	}

	if owoc2 == 0 {
		owoc2x = (((rand.Intn(maxx-min) + min) / 2) * 2)
		owoc2y = (((rand.Intn(maxy-min) + min) / 2) * 2)
		owoc2 = rand.Intn(300-40) + 40

	} else {
		ekran[owoc2x][owoc2y] = "@"
		owoc2--
		if headposx == owoc2x && headposy == owoc2y {
			dlugosc++
			owoc2 = 0
		}
	}

}

func inp(wg *sync.WaitGroup) {
	for run == 1 {
		temp, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if temp == 'w' && acmov != 's' {
			acmov = 'w'
		}
		if temp == 'a' && acmov != 'd' {
			acmov = 'a'
		}
		if temp == 'd' && acmov != 'a' {
			acmov = 'd'
		}
		if temp == 's' && acmov != 'w' {
			acmov = 's'
		}
		if temp == 'x' {
			acmov = 'x'
			run = 0
		}

		if temp == 'p' {
			pause = true
		}
	}
	wg.Done()
}

func main() {
	scrclr()
	fmt.Println("Started...")
	fmt.Println("Welcome in snake v1.0 by Rafał Miłosz Michałowski.")
	fmt.Println("Move with W,A,S,D . X to exit . Press P to Pause.  ")
	fmt.Println("Press Enter to start... ")

	fmt.Scanln()
	var wg sync.WaitGroup
	wg.Add(1)
	go inp(&wg)

	wg.Add(1)
	go drukujekran(&wg)

	wg.Wait()
	fmt.Println("KONIEC.")
}
