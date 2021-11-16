package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	returnFlag  = 0
	cotinueFlag = 1
)

var (
	clickCh         = make(chan bool)
	listenCh        = make(chan int)
	clickOnTickerCh = make(chan bool)
)

var bonusList []*Point

type Point struct {
	X int
	Y int
}

func init() {
	bonusList = make([]*Point, 0, 4)
	bonusList = append(bonusList, &Point{
		X: 519,
		Y: 220},
		&Point{
			X: 917,
			Y: 305,
		},
		&Point{
			X: 400,
			Y: 599,
		},
		&Point{
			X: 810,
			Y: 664,
		})
	//fMat := gcv.IMRead(flagPath)
	//flagImg, _ = gcv.MatToImg(fMat)

	//bMat := gcv.IMRead(bonusPath)
	//bonusImg, _ = gcv.MatToImg(bMat)
}

func main() {
	for {
		<-time.After(3 * time.Second)
		start()
	}
	//success()
	//time.Sleep(1 * time.Second)
	//opencv()
}

func start() {
	//关卡选择
	robotgo.MoveClick(1037, 626, "left", true)
	<-time.After(time.Second)
	//队伍选择
	for robotgo.GetPixelColor(977, 699) != "3ecffd" {
		//fmt.Println(robotgo.GetMouseColor())
		<-time.After(time.Second)
	}
	robotgo.MoveClick(977, 699)
	<-time.After(time.Second)
	//fight
	for i := 0; i < 5; i++ {
		fight(i + 1)
	}
}

func skill() {
	//第个技能
	<-time.After(time.Second)
	robotgo.MoveClick(503, 342, "left", true)
	<-time.After(2 * time.Second)
	//第二个技能
	robotgo.MoveClick(648, 354, "left", true)
	<-time.After(2 * time.Second)
	//第三个技能
	robotgo.MoveClick(508, 346, "left", true)
	<-time.After(2 * time.Second)
	robotgo.MoveClick(1118, 359)
	<-time.After(2 * time.Second)
}

//number 关卡
func fight(number int) {
	if 1 < number && number < 5 {
		go click()
		<-clickCh
		//点击或者进入战斗房间
		//click 函数说明已经可以点击
		robotgo.MoveClick(1071, 658)
		go clickByTicker()
		go listen()
		select {
		case flag := <-listenCh:
			clickOnTickerCh <- true
			if flag == returnFlag {
				return
			} else {
				fighting(number)
			}
		case <-time.After(time.Second * 28):
			clickOnTickerCh <- true
			return
		}
	} else {
		fighting(number)
	}
}
func fighting(number int) {
	for robotgo.GetPixelColor(1071, 658) != "d8ffff" {
		<-time.After(time.Second)
	}
	<-time.After(time.Second)
	robotgo.MoveClick(1071, 658)
	<-time.After(time.Second)
	for robotgo.GetPixelColor(1081, 375) != "bc911f" {
		<-time.After(time.Second)
	}
	<-time.After(time.Second)
	robotgo.MoveClick(1081, 375, "left", true)
	<-time.After(time.Second)
	go success()
	for {
		for !strings.HasPrefix(robotgo.GetPixelColor(1055, 360), "c") &&
			!strings.HasPrefix(robotgo.GetPixelColor(1103, 302), "c") {
			<-time.After(time.Second)
		}
		skill()
		select {
		case flag := <-listenCh:
			if flag == returnFlag {
				fightEnd(number)
				return
			} else {
				continue
			}
		}
	}

}
func fightEnd(number int) {
	//fmt.Println("------------------------")
	if number != 5 {
		fmt.Println("------------")
		robotgo.MoveClick(448, 354, "left", true)
		<-time.After(3 * time.Second)
		robotgo.MoveClick(448, 354, "left", true)
		<-time.After(18 * time.Second)
		robotgo.MoveClick(100, 200, "left", true)
		<-time.After(3 * time.Second)
		robotgo.MoveClick(781, 369, "left")
		<-time.After(3 * time.Second)
		robotgo.MoveClick(782, 626, "left", true)
		<-time.After(3 * time.Second)
	} else {
		robotgo.MoveClick(648, 354, "left", true)
		<-time.After(3 * time.Second)
		robotgo.MoveClick(648, 354, "left", true)
		<-time.After(18 * time.Second)
		bonus()
	}
}

func success() {
	<-time.After(18 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {

		select {
		case <-ticker.C:
			//if strings.HasSuffix(robotgo.GetPixelColor(657, 716), "f") ||
			//	strings.HasPrefix(robotgo.GetPixelColor(657, 716), "f") ||
			//	strings.HasPrefix(robotgo.GetPixelColor(657, 716), "0") {
			if robotgo.GetPixelColor(312, 616) == "726e6b" {
				listenCh <- returnFlag
				return
			}
			if robotgo.GetPixelColor(1115, 359) == "c9a523" {
				listenCh <- cotinueFlag
				return
			}
		}
	}
}

func click() {
	if robotgo.GetPixelColor(1071, 658) == "d8ffff" {
		clickCh <- true
	}
	fixX, fixY := 137, 385
	constAdd := 55
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fixX = fixX + constAdd
			robotgo.MoveClick(fixX, fixY, "left", true)
			<-time.After(500 * time.Millisecond)
			robotgo.MoveMouse(1071, 658)
			if robotgo.GetPixelColor(1071, 658) == "d8ffff" {
				clickCh <- true
				return
			}
		}
	}
}

func bonus() {
	for _, box := range bonusList {
		robotgo.MoveClick(box.X, box.Y, "left", true)
		<-time.After(time.Second)

	}
	robotgo.MoveClick(643, 423, "left", true)
	<-time.After(5 * time.Second)
	robotgo.MoveClick(633, 644, "left", true)
	<-time.After(2 * time.Second)
}

func listen() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			if robotgo.GetPixelColor(1109, 352) == "c9a124" {
				listenCh <- cotinueFlag
				return
			}
			if robotgo.GetPixelColor(1062, 628) == "7a4125" {
				listenCh <- returnFlag
				return
			}
		}
	}
}

func clickByTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-clickOnTickerCh:
			return
		case <-ticker.C:
			robotgo.MoveClick(80, 80, "left")
		}
	}
}
