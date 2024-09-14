package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const width = 512
const height = 512

type appOption func(a *app)

type app struct {
    strokes []stroke
    width float32
    canvas *ebiten.Image
    sizeAnchorPos Vec2
    historyIndex uintptr
}

func (a *app) Draw(screen *ebiten.Image) {
    screen.Fill(color.White)

    opts := ebiten.DrawImageOptions{}
    screen.DrawImage(a.canvas, &opts)

    cursor := cursorPosition()
    vector.StrokeCircle(
        screen,
        cursor.x,
        cursor.y,
        a.width/2,
        3,
        color.White,
        true,
    )
    vector.StrokeCircle(
        screen,
        cursor.x,
        cursor.y,
        a.width/2,
        1,
        color.Black,
        true,
    )
}

func (a *app) Layout(_, _ int) (screenWidth int, screenHeight int) {
    return width, height
}

func (a *app) Update() error {

    check_any_drawing_input := func(checker func(ebiten.Key) bool) bool {
        return drawWhite.check(checker) || drawBlack.check(checker)
    }

    is_drawing := check_any_drawing_input(ebiten.IsKeyPressed)
    just_started_drawing := check_any_drawing_input(inpututil.IsKeyJustPressed)

    var clr color.Color
    if drawBlack.check(ebiten.IsKeyPressed) {
        clr = color.Black
    } else if drawWhite.check(ebiten.IsKeyPressed) {
        clr = color.White
    }

    if just_started_drawing {
        if clr == nil {
            log.Fatal("drawing color in update is nil") // should never happen
        }
        a.strokes = append(a.strokes, stroke{
        	points: []Vec2{},
        	width:  a.width,
        	color:  clr,
        })
    }
    // TODO current history pointer is not updated as it should
    // TODO tests

    if len(a.strokes) > 0 {
        currentStroke := &a.strokes[len(a.strokes)-1]
        if is_drawing {
            currentStroke.appendIfMoved(cursorPosition())
            currentStroke.drawLast(a.canvas)
        }
    }

    if resize.check(inpututil.IsKeyJustPressed) {
        a.sizeAnchorPos = cursorPosition()
    }

    if resize.check(ebiten.IsKeyPressed) {
        prevPos := a.sizeAnchorPos
        currPos := cursorPosition()
        reltivePos := Vec2{
            prevPos.x - currPos.x,
            prevPos.y - currPos.y,
        }
        length := float32(math.Sqrt(
            float64(reltivePos.x) * float64(reltivePos.x) + float64(reltivePos.y) * float64(reltivePos.y),
        ))
        a.width = length * 2.0
    }

    return nil
}

func newApp( opts ...appOption ) app {
    a := app{}
    for _, opt := range opts {
        opt(&a)
    }
	return a
}

func main() {

    ebiten.SetWindowTitle("whiteboard")
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    ebiten.CursorPosition()
	whiteboardApp := newApp( func(a *app) {
        a.width = 10
        a.canvas = ebiten.NewImage(width, height)
        a.canvas.Fill(color.White)
    })

	ebiten.RunGame(&whiteboardApp)
}
