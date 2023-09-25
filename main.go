package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"strconv"
	"strings"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

var (
	running  = true
	bkgColor = rl.NewColor(147, 211, 196, 255)

	texture      = rl.Texture2D{}
	grassSprite  = rl.Texture2D{}
	hillSprite   = rl.Texture2D{}
	fenceSprite  = rl.Texture2D{}
	houseSprite  = rl.Texture2D{}
	waterSprite  = rl.Texture2D{}
	tilledSprite = rl.Texture2D{}

	playerSprite = rl.Texture2D{}

	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerMoving                                  bool
	playerDir                                     int
	playerUp, playerDown, playerRight, playerLeft bool

	playerFrame int
	frameCount  int

	tileDest   rl.Rectangle
	tileSrc    rl.Rectangle
	tileMap    []int
	srcMap     []string
	mapW, mapH int

	playerSpeed float32 = 1.4

	cam rl.Camera2D
)

func drawScene() {
	//rl.DrawTexture(grassSprite, 100, 50, rl.White)
	for i := 0; i < len(tileMap); i++ {
		if tileMap[i] != 0 {
			tileDest.X = tileDest.Width * float32(i%mapW)
			tileDest.Y = tileDest.Height * float32(i/mapW)

			if srcMap[i] == "g" {
				texture = grassSprite
			}
			if srcMap[i] == "l" {
				texture = hillSprite
			}
			if srcMap[i] == "f" {
				texture = fenceSprite
			}
			if srcMap[i] == "h" {
				texture = houseSprite
			}
			if srcMap[i] == "w" {
				texture = waterSprite
			}
			if srcMap[i] == "t" {
				texture = tilledSprite
			}

			if srcMap[i] == "h" || srcMap[i] == "f" {
				tileSrc.X = 0
				tileSrc.Y = 0
				rl.DrawTexturePro(grassSprite, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
			}

			tileSrc.X = tileSrc.Width * float32((tileMap[i]-1)%int(texture.Width/int32(tileSrc.Width)))
			tileSrc.Y = tileSrc.Height * float32((tileMap[i]-1)/int(texture.Width/int32(tileSrc.Width)))

			rl.DrawTexturePro(texture, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
		}
	}
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyComma) {
		playerMoving = true
		playerDir = 1
		playerUp = true

	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyO) {
		playerMoving = true
		playerDir = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyA) {
		playerMoving = true
		playerDir = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyE) {
		playerMoving = true
		playerDir = 3
		playerRight = true
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		playerSpeed = 3
	}
	if !rl.IsKeyDown(rl.KeyLeftShift) {
		playerSpeed = 1.4
	}
}

func update() {
	running = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
	} else if frameCount%45 == 1 {
		playerFrame++

	}
	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}
	if !playerMoving && playerFrame > 1 {
		playerFrame = 0
	}
	playerSrc.X = playerSrc.Width * float32(playerFrame)

	playerSrc.Y = playerSrc.Height * float32(playerDir)

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)

	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func loadMap(mapFile string) {
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println("the error is: ", err)
		os.Exit(1)
	}

	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	mapW = -1
	mapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if mapW == -1 {
			mapW = 26
		} else if mapH == -1 {
			mapH = 16
		} else if i < mapW*mapH+2 {
			tileMap = append(tileMap, m)
		} else {
			srcMap = append(srcMap, sliced[i])
		}
	}

	if len(tileMap) > mapW*mapH {
		tileMap = tileMap[:len(tileMap)-1]
	}

	/* for test
	mapW = 5
	mapH = 5
	for i := 0; i < (mapH * mapW); i++ {
		tileMap = append(tileMap, 1)
	}
	*/
}

func initializing() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] - concurrencia")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Tilesets/Grass.png")
	hillSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Tilesets/Hills.png")
	fenceSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Tilesets/Fences.png")
	houseSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Tilesets/Wooden House.png")
	waterSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Tilesets/Water.png")
	tilledSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Tilesets/Tilled Dirt.png")

	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc = rl.NewRectangle(0, 0, 16, 16)

	playerSprite = rl.LoadTexture("src/assets/Sprout Lands - Sprites - Basic pack/Characters/Basic Charakter Spritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(100, 100, 60, 60)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0.0, 3)

	loadMap("src/assets/maps/one.map")
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func main() {
	initializing()

	for running {
		input()
		update()
		render()
	}
	quit()
}
