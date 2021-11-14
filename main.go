package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

var count = 0

func main() {
	time.Sleep(3 * time.Second)
	//x, y := robotgo.GetMousePos()
	//fmt.Println(x, y)
	//fmt.Println(robotgo.GetPixelColor(x, y))
	//time.Sleep(1 * time.Second)
	start()
	//opencv()
}

func start() {
	//关卡选择
	robotgo.MoveClick(1037, 626, "left")
	time.Sleep(1 * time.Second)
	//队伍选择
	robotgo.MoveMouse(984, 701)
	for color := robotgo.GetPixelColor(984, 701); !strings.HasSuffix(color, "ffff"); {
		fmt.Println(color)
		time.Sleep(1 * time.Second)
		color = robotgo.GetPixelColor(984, 701)
	}
	fmt.Println("-------")
	robotgo.MouseClick("left", true)
	time.Sleep(2 * time.Second)
	fight()
	//战斗结算
	robotgo.MoveClick(648, 354, "left", true)
	time.Sleep(5 * time.Second)
	robotgo.MoveClick(648, 354, "left", true)
	time.Sleep(5 * time.Second)
	//end
	robotgo.MoveClick(781, 369, "left")
	time.Sleep(2 * time.Second)
	robotgo.MoveClick(782, 626, "left", true)
	time.Sleep(2 * time.Second)
	robotgo.MoveClick(516, 746, "left", true)
	time.Sleep(2 * time.Second)

	robotgo.MoveClick(736, 587, "left", true)
	time.Sleep(2 * time.Second)
	robotgo.MoveClick(543, 451, "left", true)
	time.Sleep(2 * time.Second)
	robotgo.MoveClick(697, 454, "left", true)
	time.Sleep(8 * time.Second)
	robotgo.MoveClick(679, 566, "left", true)
	time.Sleep(8 * time.Second)
	//robotgo.MoveMouse(1086, 374)
	//for color := robotgo.GetPixelColor(1086, 374); !strings.HasPrefix(color, "bc"); {
	//	fmt.Println(color)
	//	time.Sleep(1 * time.Second)
	//	color = robotgo.GetPixelColor(1086, 374)
	//}
	//robotgo.MouseClick("left", true)
	//time.Sleep(2 * time.Second)
	//x,y = robotgo.GetMousePos()

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
	time.Sleep(2 * time.Second)
	robotgo.MoveMouse(1117, 366)
	for color := robotgo.GetPixelColor(1117, 366); color != "41a328"; {
		time.Sleep(3 * time.Second)
		color = robotgo.GetPixelColor(1117, 366)
	}
	robotgo.MouseClick("left", true)
	time.Sleep(2 * time.Second)

}

func fight() {
	//战斗开始按钮
	robotgo.MoveMouse(1071, 658)
	for color := robotgo.GetPixelColor(1071, 658); color != "d8ffff"; {
		fmt.Println(color)
		time.Sleep(1 * time.Second)
		color = robotgo.GetPixelColor(1071, 658)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("-------")
	robotgo.MouseClick("left", true)
	time.Sleep(1 * time.Second)
	//hero up
	robotgo.MoveMouse(1086, 374)
	for color := robotgo.GetPixelColor(1086, 374); color != "bc911f"; {
		fmt.Println(color)
		time.Sleep(1 * time.Second)
		color = robotgo.GetPixelColor(1086, 374)
	}
	robotgo.MouseClick("left", true)
	time.Sleep(8 * time.Second)
	for {
		skill()
		robotgo.MoveMouse(621, 767)
		if robotgo.GetPixelColor(621, 767) == "f8f8f8" {
			break
		}
	}
}

func Opencv() {
	name := "test.png"
	name1 := "test_001.png"
	robotgo.SaveCapture(name1, 10, 10, 30, 30)
	robotgo.SaveCapture(name)

	fmt.Print("gcv find image: ")
	fmt.Println(gcv.FindImgFile(name1, name))
	fmt.Println(gcv.FindAllImgFile(name1, name))
	//gcv.FindAllImg()

	bit := robotgo.OpenBitmap(name1)
	defer robotgo.FindBitmap(bit)
	fmt.Print("find bitmap: ")
	fmt.Println(robotgo.FindBitmap(bit))

	// bit0 := robotgo.CaptureScreen()
	// img := robotgo.ToImage(bit0)
	// bit1 := robotgo.CaptureScreen(10, 10, 30, 30)
	// img1 := robotgo.ToImage(bit1)
	// defer robotgo.FreeBitmapArr(bit0, bit1)
	img := robotgo.CaptureImg()
	mat := gcv.IMRead("test_001.png")
	img1, err := gcv.MatToImg(mat)
	if err != nil {
		fmt.Println(err)
	}
	//img1 := robotgo.CaptureImg(10, 10, 30, 30)

	fmt.Print("gcv find image: ")
	fmt.Println(gcv.FindImg(img1, img))
	fmt.Println()
	x1, y1 := gcv.FindImgXY(img1, img)
	robotgo.MoveClick(x1, y1, "left", true)

	res := gcv.FindAllImg(img1, img)
	fmt.Println(res[0].TopLeft.Y, res[0].Rects.TopLeft.X, res)
	x, y := res[0].TopLeft.X, res[0].TopLeft.Y
	robotgo.Move(x, y-rand.Intn(5))
	robotgo.MilliSleep(100)
	robotgo.Click()

}

//621 767 f8f8f8
