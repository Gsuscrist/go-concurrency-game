package scenes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type LoserScene struct {
	window fyne.Window
}

func NewLoserScene(window fyne.Window) *LoserScene {
	return &LoserScene{window: window}
}

func (s *LoserScene) Show() {
	bg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Background-lost.png"))
	bg.Resize(fyne.NewSize(900, 600))
	bg.Move(fyne.NewPos(0, 0))

	btnReset := widget.NewButton("Reset Game", func() {
		fmt.Println("new game")
		mainScene := NewMainMenuScene(s.window)
		mainScene.Show()
	})
	btnReset.Resize(fyne.NewSize(150, 50))
	btnReset.Move(fyne.NewPos(500, 170))

	s.window.SetContent(container.NewWithoutLayout(bg, btnReset))

}
