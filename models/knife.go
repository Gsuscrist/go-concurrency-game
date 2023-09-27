package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Knife struct {
	posX, PosY float32
	status     bool
	skin       *canvas.Image
	Id         int
}

func NewKnife(posX float32, posY float32, img *canvas.Image) *Knife {
	return &Knife{
		posX:   posX,
		PosY:   posY,
		status: true,
		skin:   img,
		Id:     2,
	}
}

func (knife *Knife) Run() {
	var incY float32 = 10
	knife.status = true
	for knife.status {
		if knife.PosY > 450 {
			break
		}

		knife.PosY += incY
		knife.skin.Move(fyne.NewPos(knife.posX, knife.PosY))

		time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)
	}
}

func (knife *Knife) SetStatus(status bool) {
	knife.status = status
}
