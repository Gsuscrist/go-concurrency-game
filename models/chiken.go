package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Chicken struct {
	PosX, PosY float32
	Status     bool
	Skin       *canvas.Image
	Id         int
}

func NewChicken(posX float32, posY float32, img *canvas.Image) *Chicken {
	return &Chicken{
		PosX:   posX,
		PosY:   posY,
		Status: true,
		Skin:   img,
		Id:     7,
	}
}

func (chicken *Chicken) Run() {
	var speed float32 = 10
	chicken.Status = true
	for chicken.Status {
		if chicken.PosX > 780 {
			chicken.PosX = 50
		}
		chicken.PosX += speed
		chicken.Skin.Move(fyne.NewPos(chicken.PosX, chicken.PosY))

		time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)

	}
}

func (chicken *Chicken) SetStatus(status bool) {
	chicken.Status = status
}
