package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
)

// DrawRectangle 画图像矩形, 标注, 颜色; 返回保存图像路径
// 使用绝对坐标值
// ref: https://github.com/fogleman/gg/blob/master/examples/rotated-image.go
func DrawRectangle(im_path, save_path string, x, y, w, h float64) error {
	// 读取图片
	im, err := gg.LoadImage(im_path)
	if err != nil {
		return err
	}
	fmt.Println("宽度:", im.Bounds().Max.X)
	fmt.Println("高度:", im.Bounds().Max.Y)

	// 创建画布
	dc := gg.NewContextForImage(im)

	// 绘制矩形框
	dc.SetHexColor("#FF0000")    // 颜色
	dc.SetLineWidth(10)          // 线宽
	dc.DrawRectangle(x, y, w, h) // 坐标
	//dc.DrawRectangle(x+10, y+10, w+10, h+10) // 坐标
	//dc.DrawRectangle(x+20, y+20, w+20, h+20) // 坐标
	dc.Stroke()

	// 保存图片
	if err := dc.SavePNG(save_path); err != nil {
		return err
	}

	return nil
}

// DrawRectangleWithNormalized 使用归一化的坐标
func DrawRectangleWithNormalized(im_path, save_path string, xN1, yN1, xN2, yN2 float64) error {
	// 读取图片
	im, err := gg.LoadImage(im_path)
	if err != nil {
		return err
	}
	fmt.Println("宽度:", im.Bounds().Max.X)
	fmt.Println("高度:", im.Bounds().Max.Y)

	width := im.Bounds().Max.X
	high := im.Bounds().Max.Y

	// 创建画布
	dc := gg.NewContextForImage(im)

	// 绘制矩形框
	dc.SetHexColor("#FF0000") // 颜色
	dc.SetLineWidth(1)        // 线宽
	dc.DrawRectangle(
		float64(width)*xN1,
		float64(high)*yN1,
		float64(width)*(xN2-xN1),
		float64(high)*(yN2-yN1)) // 坐标
	//dc.DrawRectangle(x+10, y+10, w+10, h+10) // 坐标
	//dc.DrawRectangle(x+20, y+20, w+20, h+20) // 坐标
	dc.Stroke()

	// 保存图片
	if err := dc.SavePNG(save_path); err != nil {
		return err
	}

	return nil
}

func main() {
	imPath := "example_with_box.png"
	savePath := "example_with_box2.png"
	var x, y, w, h float64 = 500, 100, 200, 100
	if err := DrawRectangle(imPath, savePath, x, y, w, h); err != nil {
		logrus.Error(err)
		return
	}

	imPath2 := "example_with_box.png"
	savePath2 := "example_with_ormalized_box.png"
	var x1, y1, x2, y2 float64 = 0.3663233995437622, 0.22497340308295355, 0.5041877269744873, 0.783307139078776
	if err := DrawRectangleWithNormalized(imPath2, savePath2, x1, y1, x2, y2); err != nil {
		logrus.Error(err)
		return
	}

	// 删除图片
	//if err := os.Remove(savePath); err != nil {
	//	logrus.Error(err)
	//	return
	//}
}
