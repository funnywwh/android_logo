package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	//	"math"
	//	"image/draw"
	"image/png"
	"os"

	"golang.org/x/image/draw"

	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("%s <xxx.png>\n", os.Args[0])
		return
	}
	fname := os.Args[1]
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	img, err := png.Decode(f)
	for _, w := range []int{192, 144, 96, 72, 48} {
		err = writeLogo(img, w)
		if err != nil {
			fmt.Println("writeLog err:%s\n", err.Error())
			return
		}
	}
}

func circleMask(dst image.Image) image.Image {
	sr := dst.Bounds()

	circle := image.NewRGBA(sr)
	gc := draw2dimg.NewGraphicContext(circle)
	gc.SetFillColor(color.Alpha{A: 0xff})
	gc.SetStrokeColor(color.Alpha{A: 0xff})
	gc.SetLineWidth(1)
	r := float64(sr.Dx() / 2)
	gc.BeginPath()

	gc.ArcTo(r, r, r, r, 0, -math.Pi*2)
	gc.Close()
	gc.Fill()
	return circle
}

func writeLogo(img image.Image, w int) (err error) {
	var path string
	switch w {
	case 48:
		path = fmt.Sprintf("./mipmap-mdpi")
	case 72:
		path = fmt.Sprintf("./mipmap-hdpi")
	case 96:
		path = fmt.Sprintf("./mipmap-xhdpi")
	case 144:
		path = fmt.Sprintf("./mipmap-xxhdpi")
	case 192:
		path = fmt.Sprintf("./mipmap-xxhdpi")
	default:
		err = fmt.Errorf("unsupport output size logo")
	}
	os.MkdirAll(path, 0777)
	imgo := image.NewRGBA(image.Rect(0, 0, w, w))
	draw.ApproxBiLinear.Scale(imgo, imgo.Bounds(), img, img.Bounds(), draw.Src, nil)
	//	draw.Draw(imgo, imgo.Bounds(), img, image.Point{X: 0, Y: 0}, draw.Src)
	f, err := os.OpenFile(fmt.Sprintf("./%s/ic_launcher.png", path), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	err = png.Encode(f, imgo)
	f.Close()
	if err != nil {
		return
	}
	imgo = image.NewRGBA(image.Rect(0, 0, w, w))
	draw.ApproxBiLinear.Scale(imgo, imgo.Bounds(), img, img.Bounds(), draw.Src,
		&draw.Options{
			SrcMask:  circleMask(img),
			SrcMaskP: image.ZP,
		})
	//	draw.Draw(imgo, imgo.Bounds(), img, image.Point{X: 0, Y: 0}, draw.Src)
	f, err = os.OpenFile(fmt.Sprintf("./%s/ic_launcher_round.png", path), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	err = png.Encode(f, imgo)
	f.Close()
	if err != nil {
		return
	}
	return
}
