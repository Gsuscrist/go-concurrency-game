package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Axe struct {
	posX, PosY float32
	status     bool
	skin       *canvas.Image
	Id         int
}

func NewAxe(posX float32, posY float32, img *canvas.Image) *Axe {
	return &Axe{
		posX:   posX,
		PosY:   posY,
		status: true,
		skin:   img,
		Id:     1,
	}
}

func (axe *Axe) Run() {
	var incY float32 = 10
	axe.status = true
	for axe.status {
		if axe.PosY > 450 {
			break
		}

		axe.PosY += incY
		axe.skin.Move(fyne.NewPos(axe.posX, axe.PosY))

		time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)
	}
}

func (axe *Axe) SetStatus(status bool) {
	axe.status = status
}
