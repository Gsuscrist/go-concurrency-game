package scenes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type WinnerScene struct {
	window fyne.Window
}

func NewWinnerScene(window fyne.Window) *WinnerScene {
	return &WinnerScene{window: window}
}

func (s *WinnerScene) Show() {
	bg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Background-win.png"))
	bg.Resize(fyne.NewSize(900, 600))
	bg.Move(fyne.NewPos(0, 0))

	btnReset := widget.NewButton("Reset", func() {
		fmt.Println("new game")
		mainScene := NewMainMenuScene(s.window)
		mainScene.Show()
	})
	btnReset.Resize(fyne.NewSize(150, 50))
	btnReset.Move(fyne.NewPos(650, 400))

	s.window.SetContent(container.NewWithoutLayout(bg, btnReset))

}
