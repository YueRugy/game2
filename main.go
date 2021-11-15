package main

import (
	"fmt"
	"image"
	"math/rand"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

const (
	successPath = "success.png"
    flagPath = "flag.png"
)

var (
	successImg image.Image
    flagImg    image.Image
    bonusImg   image.Image
)

//var count = 0
var ch = make(chan bool)

var qh = make(chan bool)

func init() {

	mat := gcv.IMRead(successPath)
	successImg, _ = gcv.MatToImg(mat)

    fMat := gcv.IMRead(flagPath)
    flagImg, _ = gcv.MatToImg(fMat)

	bMat := gcv.IMRead("bonus.png")
	bonusImg, _ = gcv.MatToImg(bMat)
}

func main() {
	time.Sleep(3 * time.Second)
	x, y := robotgo.GetMousePos()
	fmt.Println(x, y)
	robotgo.MoveMouse(x, y)
	fmt.Println(robotgo.GetMouseColor())
	//success()
	//time.Sleep(1 * time.Second)
	start()
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
	//robotgo.MoveMouse(1118, 359)
	//for color := robotgo.GetPixelColor(1117, 366); color != "2e9f1b"; {
	//	time.Sleep(3 * time.Second)
	//	color = robotgo.GetPixelColor(1117, 366)
	//}
	//robotgo.MouseClick("left", true)
	time.Sleep(2 * time.Second)

}

//number 关卡
func fight(number int) {
	var jump bool
	if 1<number&&number<5{
		go click()
		<-qh
		robotgo.MoveClick(178, 509, "left", true)
		time.Sleep(2 * time.Second)
        //点击或者进入战斗房间
       //click 函数说明已经可以点击     
		robotgo.MoveClick(1071, 658)
		time.Sleep(10 * time.Second)
        //判断是房间还是奖励项
		jump = isGo()
		fmt.Println(jump)

		if jump {
			robotgo.MoveClick(176, 509, "left", true)
			time.Sleep(3 * time.Second)
			robotgo.MoveClick(557, 332, "left", true)
			time.Sleep(3 * time.Second)
			fmt.Println("click...............")
			return
        }
	}


    //不是第一关卡和boss关卡 要执行战斗开始按钮
	if number==1||number==5 {
		//战斗开始按钮
		robotgo.MoveMouse(1071, 658)
		for color := robotgo.GetPixelColor(1071, 658); color != "d8ffff"; {
			time.Sleep(1 * time.Second)
			color = robotgo.GetPixelColor(1071, 658)
		}
		time.Sleep(2 * time.Second)
		robotgo.MouseClick("left", true)
		time.Sleep(2 * time.Second)

	}

	//hero up
	robotgo.MoveMouse(1086, 374)
	for color := robotgo.GetPixelColor(1086, 374); color != "bc911f"; {
		time.Sleep(1 * time.Second)
		color = robotgo.GetPixelColor(1086, 374)
	}
	robotgo.MoveClick(1086, 374, "left", true)
	time.Sleep(15 * time.Second)
	go success()
	for {
		skill()
		time.Sleep(25 * time.Second)
		select {
		case <-ch:
			//qh <- true
			goto Loop
		default:
			continue
		}
		//robotgo.MoveMouse(621, 767)
		//if robotgo.GetPixelColor(621, 767) != "f8f8f8" {
		//break
		//}
	}
Loop:
	fmt.Println("循环外")
	if number != 5 {
		robotgo.MoveClick(448, 354, "left", true)
		time.Sleep(3 * time.Second)
		robotgo.MoveClick(448, 354, "left", true)
		time.Sleep(18 * time.Second)
		robotgo.MoveClick(781, 369, "left")
		time.Sleep(2 * time.Second)
		robotgo.MoveClick(782, 626, "left", true)
		time.Sleep(2 * time.Second)

	}else {
		robotgo.MoveClick(648, 354, "left", true)
		time.Sleep(3 * time.Second)
		robotgo.MoveClick(648, 354, "left", true)
		time.Sleep(12 * time.Second)
		bonus()

	}
	//Look:
	//return
}

func success() {
	time.Sleep(8 * time.Second)
	ticker := time.NewTicker(8 * time.Second)
	defer ticker.Stop()
	for {
		select {
		//case <-qh:
		//goto Loop
		case <-ticker.C:
			img := robotgo.CaptureImg()
			li := gcv.FindAllImg(successImg, img)
			//fmt.Println(len(li))
			if len(li) > 0 {
				ch <- true
				return
			}
			//return
		}
	}
	//Loop:
	//return
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
func click() {
	if robotgo.GetPixelColor(1071, 658) == "d8ffff" {
		qh <- true
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
			time.Sleep(300 * time.Millisecond)
			robotgo.MoveMouse(1071, 658)
			if robotgo.GetPixelColor(1071, 658) == "d8ffff" {
				qh <- true
				return
			}
		}
	}
}

func isGo() bool {
	img := robotgo.CaptureImg()
	li := gcv.FindAllImg(flagImg, img)
	fmt.Println(len(li))
	return len(li) == 0

}

func bonus() {
	img := robotgo.CaptureImg()
	bonus := gcv.FindAllImg(bonusImg, img)
	for _, box := range bonus {
		x, y := box.Middle.X, box.Middle.Y
		robotgo.MouseClick(x, y, "left", true)
		time.Sleep(time.Second)
	}
	robotgo.MouseClick(643, 423, "left", true)
	time.Sleep(5 * time.Second)
	robotgo.MouseClick(633, 644, "left", true)
	time.Sleep(2 * time.Second)
}
