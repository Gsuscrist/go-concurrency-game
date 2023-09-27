package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"pelota/scenes"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Le Cousine Concurrente")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(900, 600))

	//Cargar y mostrar la escena principal
	mainMenuScene := scenes.NewMainMenuScene(myWindow)

	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
