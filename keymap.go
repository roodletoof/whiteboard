package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)


type keymap struct{
	keys []ebiten.Key
	mouseButtons []ebiten.MouseButton
}

var resize = keymap{
	keys: []ebiten.Key{ebiten.KeyD, ebiten.KeyK,},
	mouseButtons: []ebiten.MouseButton{ebiten.MouseButtonMiddle,},
}
var drawBlack = keymap{
	keys: []ebiten.Key{ebiten.KeyF, ebiten.KeyJ,},
	mouseButtons: []ebiten.MouseButton{ebiten.MouseButtonLeft,},
}
var drawWhite = keymap{
	keys: []ebiten.Key{ebiten.KeyS, ebiten.KeyL,},
	mouseButtons: []ebiten.MouseButton{ebiten.MouseButtonRight,},
}
var undo = keymap{
    keys: []ebiten.Key{ebiten.KeyZ, ebiten.KeyU,},
}
var redo = keymap{
    keys: []ebiten.Key{ebiten.KeyX, ebiten.KeyR,},
}

func (k keymap) isJustPressed() bool {
	for _, key := range k.keys {
		if inpututil.IsKeyJustPressed(key) {
			return true
		}
	}

	for _, mouseButton := range k.mouseButtons {
		if inpututil.IsMouseButtonJustPressed(mouseButton) {
			return true
		}
	}

	return false
}

func (k keymap) isJustReleased() bool {
	for _, key := range k.keys {
		if inpututil.IsKeyJustReleased(key) {
			return true
		}
	}

	for _, mouseButton := range k.mouseButtons {
		if inpututil.IsMouseButtonJustReleased(mouseButton) {
			return true
		}
	}

	return false
}

func (k keymap) isPressed() bool {
	for _, key := range k.keys {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}

	for _, mouseButton := range k.mouseButtons {
		if ebiten.IsMouseButtonPressed(mouseButton) {
			return true
		}
	}

	return false
}
