package main

import rl "github.com/gen2brain/raylib-go/raylib"

type circleShape interface {
	draw()
	update()
	collidesWith(circleShape) bool
}

type player struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32
	rotation float32
}

func (p *player) draw() {
	rl.DrawCircleLinesV(p.position, p.radius, rl.White)
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	player := player{
		position: rl.Vector2{X: 400, Y: 250},
		radius:   10,
		rotation: 0,
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update

		// Draw
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.DarkGray)
			player.draw()
		}
		rl.EndDrawing()
	}
}
