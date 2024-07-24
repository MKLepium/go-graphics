package graphicswrapper

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	imageloader "github.com/mklepium/go-graphics/internal/imageLoader"
)

func Start() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	img, err := imageloader.DecodeWebP("assets/Dirt.webp")
	if err != nil {
		fmt.Println("Error decoding WebP file:", err)
		return
	}
	rl_image := rl.NewImageFromImage(img)
	rl_texture := rl.LoadTextureFromImage(rl_image)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.DrawTexture(rl_texture, 0, 0, rl.White)

		rl.EndDrawing()
	}
}
