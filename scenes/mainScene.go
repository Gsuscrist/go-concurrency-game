package scenes

import (
	_ "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"pelota/models"
)

type MainMenuScene struct {
	window fyne.Window
}

var axe *models.Axe
var knife *models.Knife
var hoe *models.Hoe

var chicken *models.Chicken

var cow1 *models.Cow
var cow2 *models.Cow
var cow3 *models.Cow

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window}
}

func (s *MainMenuScene) Show() {
	bg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Background-kitchen.png"))
	bg.Resize(fyne.NewSize(900, 600))
	bg.Move(fyne.NewPos(0, 0))

	axeImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Axe-Sprite.png"))
	axeImg.Resize(fyne.NewSize(70, 70))
	axeImg.Move(fyne.NewPos(100, 10))

	axe = models.NewAxe(100, 10, axeImg)

	knifeImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Log-Sprite.png"))
	knifeImg.Resize(fyne.NewSize(70, 70))
	knifeImg.Move(fyne.NewPos(400, 10))

	knife = models.NewKnife(400, 10, knifeImg)

	HoeImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Hoe-Sprite.png"))
	HoeImg.Resize(fyne.NewSize(70, 70))
	HoeImg.Move(fyne.NewPos(700, 10))

	hoe = models.NewHoe(700, 10, HoeImg)

	ChickenImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/Chicken-Sprite.png"))
	ChickenImg.Resize(fyne.NewSize(50, 50))
	ChickenImg.Move(fyne.NewPos(30, 465))

	chicken = models.NewChicken(30, 465, ChickenImg)

	CowImg1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/Cow-Sprites.png"))
	CowImg1.Resize(fyne.NewSize(120, 120))
	CowImg1.Move(fyne.NewPos(100, 420))

	CowImg2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/Cow-Sprites.png"))
	CowImg2.Resize(fyne.NewSize(120, 120))
	CowImg2.Move(fyne.NewPos(400, 420))

	CowImg3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/Cow-Sprites.png"))
	CowImg3.Resize(fyne.NewSize(120, 120))
	CowImg3.Move(fyne.NewPos(700, 420))

	cow1 = models.NewCow(100, 600, CowImg1, 5)
	cow2 = models.NewCow(100, 600, CowImg2, 5)
	cow3 = models.NewCow(100, 600, CowImg1, 5)

	btnSelectWeapon1 := widget.NewButton("Vaca #1", func() {
		s.StartGame(1)
	})
	btnSelectWeapon1.Resize(fyne.NewSize(100, 20))
	btnSelectWeapon1.Move(fyne.NewPos(100, 550))

	btnSelectWeapon2 := widget.NewButton("Vaca #2", func() {
		s.StartGame(2)
	})
	btnSelectWeapon2.Resize(fyne.NewSize(100, 20))
	btnSelectWeapon2.Move(fyne.NewPos(400, 550))

	btnSelectWeapon3 := widget.NewButton("Vaca #3", func() {
		s.StartGame(3)
	})
	btnSelectWeapon3.Resize(fyne.NewSize(100, 20))
	btnSelectWeapon3.Move(fyne.NewPos(700, 550))

	s.window.SetContent(container.NewWithoutLayout(bg,
		axeImg, HoeImg, knifeImg,
		ChickenImg,
		CowImg1, CowImg2, CowImg3,
		btnSelectWeapon1, btnSelectWeapon2, btnSelectWeapon3))

	go chicken.Run()

}

func (s *MainMenuScene) GetMurderWeapon(weapon int) {
	for {
		if knife.PosY > 420 {
			s.StopGame()
			s.ShowMessage(weapon, knife.Id)
		}
		if axe.PosY > 420 {
			s.StopGame()
			s.ShowMessage(weapon, axe.Id)
		}
		if hoe.PosY > 420 {
			s.StopGame()
			s.ShowMessage(weapon, hoe.Id)
		}
	}
}

func (s *MainMenuScene) ShowMessage(weapon int, murderWeapon int) {
	if weapon == murderWeapon {

		winScene := NewWinnerScene(s.window)
		winScene.Show()
	} else {

		loseScene := NewLoserScene(s.window)
		loseScene.Show()
	}
}

func (s *MainMenuScene) StartGame(weapon int) {
	go axe.Run()
	go knife.Run()
	go hoe.Run()

	go s.GetMurderWeapon(weapon)

}

func (s *MainMenuScene) StopGame() {
	axe.SetStatus(false)
	knife.SetStatus(false)
	hoe.SetStatus(false)
	chicken.SetStatus(false)
}
