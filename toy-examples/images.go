package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
    Width, Height int
}

func (a Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (a Image) At(x,y int) color.Color {
    clr := uint8 (x*y)
    return color.RGBA{clr,clr,255,255}
}

func (a Image) Bounds() image.Rectangle {
    return image.Rect(0,0,a.Width,a.Height)
}

func main() {
    m := Image{128,128}
    pic.ShowImage(m)
}
