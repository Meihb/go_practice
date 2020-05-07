package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	//图片大小
	const size = 300
	//根据大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	//从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		//让Sin值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size;
		//Sin的幅度为一半的像素,向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2;
		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//创建文件
	//    ./是你当前的工程目录，并不是该go文件所对应的目录。
	//    比如myProject/src/main/main.go
	//    main.go里使用./,其路径不是myProject/src/main/，而是myProject/
	file, err := os.Create("pkg//sin.png")
	if err != nil {
		log.Fatal(err)
	}
	//使用Png格式把数据写入文件
	png.Encode(file, pic) //将image信息写入文件中

	file.Close()
}
