package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Cow struct {
	PosX, PosY float32
	status     bool
	skin       *canvas.Image
	id         int
}

func NewCow(posX float32, posY float32, img *canvas.Image, id int) *Cow {
	return &Cow{
		PosX:   posX,
		PosY:   posY,
		status: true,
		skin:   img,
		id:     id,
	}
}

func (cow *Cow) run() {
	var incY float32 = 10
	cow.status = true
	for cow.status {
		if cow.PosY > 400 {
			break
		}
		cow.PosY += incY
		cow.skin.Move(fyne.NewPos(cow.PosX, cow.PosY))

		time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)
	}
}

func (cow *Cow) SetStatus(status bool) {
	cow.status = status

}
