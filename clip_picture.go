package main

import (
	"fmt"
	"github.com/fogleman/gg"
)

func main() {
	// 加载图片，这里路径换成自己的
	im, err := gg.LoadImage("example.jpg")
	if err != nil {
		panic(err)
	}

	// 获取图片的宽度和高度
	w := im.Bounds().Size().X // 1920
	h := im.Bounds().Size().Y // 1040

	fmt.Println(w, h)

	dc := gg.NewContext(w, h)

	// 画矩形
	dc.DrawRectangle(
		float64(w/3),
		float64(h/2),
		float64(200), // 宽
		float64(300)) // 高

	// 对画布进行裁剪
	dc.Clip()
	// 加载图片
	dc.DrawImage(im, 0, 0)
	err = dc.SavePNG("out.png")
	if err != nil {
		return
	}
}
