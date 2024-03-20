package main

import rl "github.com/gen2brain/raylib-go/raylib"

const CursorRadius = 2

func main() {
	rl.InitWindow(1280, 720, "Electron Architect")
	defer rl.CloseWindow()

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.SetTargetFPS(244)
	rl.HideCursor()

	var cursorOffset = rl.Vector2{X: -CursorRadius, Y: -CursorRadius}
	var cursorSize = rl.Vector2{X: CursorRadius * 2, Y: CursorRadius * 2}

	var input InputHandler

	for !rl.WindowShouldClose() {

		mousePos := rl.GetMousePosition()

		if input.IsPressed(CreateNode) {
			println("Create node")
		} else if input.IsPressed(RemoveHovered) {
			println("Remove hovered")
		}

		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.Black)
			rl.DrawText("Hello world", 12, 12, 8, rl.Blue)

			rl.BeginBlendMode(rl.BlendSubtractColors)
			{
				rl.DrawRectangleV(rl.Vector2Subtract(mousePos, cursorOffset), cursorSize, rl.White)
			}
			rl.EndBlendMode()
		}
		rl.EndDrawing()
	}
}
