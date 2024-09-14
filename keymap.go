package main

import "github.com/hajimehoshi/ebiten/v2"


type keymap []ebiten.Key
var resize = keymap{
    ebiten.KeyD,
    ebiten.KeyK,
}
var drawBlack = keymap{
    ebiten.KeyF,
    ebiten.KeyJ,
}
var drawWhite = keymap{
    ebiten.KeyS,
    ebiten.KeyL,
}
var undo = keymap{
    ebiten.KeyZ,
    ebiten.KeyU,
}
var redo = keymap{
    ebiten.KeyX,
    ebiten.KeyR,
}

func (k keymap) check(checker func(ebiten.Key) bool) bool {
    for _, key := range k {
        if checker(key) {
            return true
        }
    }
    return false
}
