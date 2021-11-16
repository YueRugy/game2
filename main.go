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
	time.Sleep(3 * time.Second)
	for {
		time.Sleep(3 * time.Second)
		start()
	}
	//success()
	//time.Sleep(1 * time.Second)
	//opencv()
}

func start() {
	//关卡选择
	robotgo.MoveClick(1037, 626, "left")
	time.Sleep(2 * time.Second)
	//队伍选择
	robotgo.MoveMouse(984, 701)
	for color := robotgo.GetPixelColor(984, 701); !strings.HasSuffix(color, "ffff"); {
		fmt.Println(color)
		time.Sleep(1 * time.Second)
		color = robotgo.GetPixelColor(984, 701)
	}
	robotgo.MouseClick("left", true)
	time.Sleep(2 * time.Second)
	//fight
	for i := 0; i < 5; i++ {
		fight(i + 1)
	}
}

func skill() {
	//第个技能
	robotgo.MoveClick(503, 342, "left", true)
	time.Sleep(2 * time.Second)
	//第二个技能
	robotgo.MoveClick(648, 354, "left", true)
	time.Sleep(2 * time.Second)
	//第三个技能
	robotgo.MoveClick(508, 346, "left", true)
	time.Sleep(2 * time.Second)
	robotgo.MoveClick(1118, 359)
	time.Sleep(2 * time.Second)

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
				goto Start
			}
		case <-time.After(time.Second * 18):
			clickOnTickerCh <- true
			return
		}
	}
Start:
	robotgo.MoveClick(1071, 658)
	<-time.After(time.Second)
	for color := robotgo.GetPixelColor(1086, 374); color != "bc911f"; {
		<-time.After(time.Second)
		color = robotgo.GetPixelColor(1086, 374)
	}
	robotgo.MoveClick(1086, 374, "left", true)
	go success()
	for {
		for color := robotgo.GetPixelColor(1086, 374); color != "bc911f"; {
			<-time.After(time.Second)
			color = robotgo.GetPixelColor(1086, 374)
		}
		skill()
		//time.Sleep(25 * time.Second)
		select {
		case flag := <-listenCh:
			if flag == cotinueFlag {
				fightEnd(number)
				return
			} else {
				continue
			}
		}
	}
}

func fightEnd(number int) {
	if number != 5 {
		robotgo.MoveClick(448, 354, "left", true)
		time.Sleep(3 * time.Second)
		robotgo.MoveClick(448, 354, "left", true)
		time.Sleep(18 * time.Second)
		robotgo.MoveClick(100, 200, "left", true)
		time.Sleep(3 * time.Second)
		robotgo.MoveClick(781, 369, "left")
		time.Sleep(3 * time.Second)
		robotgo.MoveClick(782, 626, "left", true)
		time.Sleep(3 * time.Second)
	} else {
		robotgo.MoveClick(648, 354, "left", true)
		time.Sleep(3 * time.Second)
		robotgo.MoveClick(648, 354, "left", true)
		time.Sleep(18 * time.Second)
		bonus()
	}
}

func success() {
	time.Sleep(10 * time.Second)
	ticker := time.NewTicker(8 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if robotgo.GetPixelColor(1086, 374) == "bc911f" {
				listenCh <- returnFlag
				return
			}
			if robotgo.GetPixelColor(1086, 374) == "bc911f" {
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
			time.Sleep(500 * time.Millisecond)
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
		time.Sleep(time.Second)

	}
	robotgo.MoveClick(643, 423, "left", true)
	time.Sleep(5 * time.Second)
	robotgo.MoveClick(633, 644, "left", true)
	time.Sleep(2 * time.Second)
}

func listen() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			if robotgo.GetPixelColor(1071, 658) == "d8ffff" {
				listenCh <- returnFlag
				return
			}
			if robotgo.GetPixelColor(1071, 658) == "d8ffff" {
				listenCh <- cotinueFlag
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
