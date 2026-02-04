package main

import (
	"embed"
	"fmt"
	"iter"

	ecs "github.com/bernhardfritz/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const maxBunnies = 50000
const maxBatchElements = 8192

type Position struct {
	rl.Vector2
}

func (position Position) GetComponent() ecs.AnyComponent {
	return positionComponent
}

type Speed struct {
	rl.Vector2
}

func (speed Speed) GetComponent() ecs.AnyComponent {
	return speedComponent
}

type Color struct {
	rl.Color
}

func (color Color) GetComponent() ecs.AnyComponent {
	return colorComponent
}

var positionComponent ecs.Component[Position]
var speedComponent ecs.Component[Speed]
var colorComponent ecs.Component[Color]

var texBunny rl.Texture2D

//go:embed resources/raybunny.png
var ASSETS embed.FS

func init() {
	rl.AddFileSystem(ASSETS)
}

func main() {
	const screenWidth = 800
	const screenHeight = 450

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - bunnymark")
	defer rl.CloseWindow()

	texBunny = rl.LoadTexture("resources/raybunny.png")
	defer rl.UnloadTexture(texBunny)

	world := ecs.NewWorld(maxBunnies)
	world.RegisterComponent(&positionComponent)
	world.RegisterComponent(&speedComponent)
	world.RegisterComponent(&colorComponent)
	movementSystem := world.CreateSystem(updateBunnies, positionComponent, speedComponent)
	drawSystem := world.CreateSystem(drawBunnies, positionComponent, colorComponent)

	bunniesCount := 0

	rl.SetTargetFPS(60)

	var update = func() {
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			// Create more bunnies
			for i := 0; i < 100; i++ {
				if bunniesCount < maxBunnies {
					mousePosition := rl.GetMousePosition()
					world.CreateEntity(
						Position{rl.Vector2{X: mousePosition.X, Y: mousePosition.Y}},
						Speed{rl.Vector2{X: float32(rl.GetRandomValue(-250, 250) / 60), Y: float32(rl.GetRandomValue(-250, 250) / 60)}},
						Color{rl.Color{R: uint8(rl.GetRandomValue(50, 240)), G: uint8(rl.GetRandomValue(80, 240)), B: uint8(rl.GetRandomValue(100, 240)), A: 255}},
					)
					bunniesCount++
				}
			}
		}

		// Update bunnies
		movementSystem()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		drawSystem()

		rl.DrawRectangle(0, 0, screenWidth, 40, rl.Black)
		rl.DrawText(fmt.Sprintf("bunnies: %d", bunniesCount), 120, 10, 20, rl.Green)
		rl.DrawText(fmt.Sprintf("batched draw calls: %d", 1+bunniesCount/maxBatchElements), 320, 10, 20, rl.Maroon)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.SetMainLoop(update)
}

func updateBunnies(w *ecs.World, bunnies iter.Seq[ecs.Entity]) {
	for bunny := range bunnies {
		position, speed := positionComponent.Get(w, bunny), speedComponent.Get(w, bunny)
		position.X += speed.X
		position.Y += speed.Y

		if position.X+float32(texBunny.Width)/2 > float32(rl.GetScreenWidth()) || position.X+float32(texBunny.Width)/2 < 0 {
			speed.X *= -1
		}
		if position.Y+float32(texBunny.Height)/2 > float32(rl.GetScreenHeight()) || position.Y+float32(texBunny.Height)/2-40 < 0 {
			speed.Y *= -1
		}
	}
}

func drawBunnies(w *ecs.World, bunnies iter.Seq[ecs.Entity]) {
	for bunny := range bunnies {
		position, color := positionComponent.Get(w, bunny), colorComponent.Get(w, bunny)
		rl.DrawTexture(texBunny, int32(position.X), int32(position.Y), rl.Color{R: color.R, G: color.G, B: color.B, A: color.A})
	}
}
