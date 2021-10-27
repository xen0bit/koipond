package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1000
	screenHeight = 1000
	cols         = 1000
	rows         = 1000
	damping      = float32(0.99)
)

var (
	current  = [cols][rows]float32{}
	previous = [cols][rows]float32{}
)

type Game struct {
	noiseImage *image.RGBA
}

func (g *Game) Init() {
	// for i := 0; i < cols; i++ {
	// 	for j := 0; j < rows; j++ {
	// 		current[i][j] = 0
	// 		previous[i][i] = 0
	// 	}
	// }
	//previous[100][100] = 255
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if (0 <= mx && mx < screenWidth) && (0 <= my && my < screenHeight) {
			previous[mx][my] = 255
		}
	}
	// Generate the noise with random RGB values.
	for i := 1; i < cols-1; i++ {
		for j := 1; j < rows-1; j++ {
			current[i][j] = (previous[i-1][j]+previous[i+1][j]+
				previous[i][j-1]+previous[i][j+1]+
				previous[i-1][j-1]+previous[i-1][j+1]+
				previous[i+1][j-1]+previous[i+1][j+1])/4 - current[i][j]

			current[i][j] = current[i][j] * damping
			//current[i][j] = 0.5
			index := (i + j*cols) * 4
			//fmt.Println(index)

			g.noiseImage.Pix[index+0] = 0 //uint8(current[i][j] * 255)
			g.noiseImage.Pix[index+1] = 0 //uint8(current[i][j] * 255)
			g.noiseImage.Pix[index+2] = uint8(current[i][j] * 255)
			g.noiseImage.Pix[index+3] = 255
		}
	}
	//fmt.Println(current)
	//swap
	temp := previous
	previous = current
	current = temp
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.noiseImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Noise (Ebiten Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	g.Init()
	//ebiten.SetMaxTPS(5)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
