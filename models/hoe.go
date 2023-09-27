package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Hoe struct {
	posX, PosY float32
	status     bool
	skin       *canvas.Image
	Id         int
}

func NewHoe(posX float32, posY float32, img *canvas.Image) *Hoe {
	return &Hoe{
		posX:   posX,
		PosY:   posY,
		status: true,
		skin:   img,
		Id:     3,
	}
}

func (hoe *Hoe) Run() {
	var incY float32 = 10
	hoe.status = true
	for hoe.status {
		if hoe.PosY > 450 {
			break
		}

		hoe.PosY += incY
		hoe.skin.Move(fyne.NewPos(hoe.posX, hoe.PosY))

		time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)
	}
}

func (hoe *Hoe) SetStatus(status bool) {
	hoe.status = status
}
