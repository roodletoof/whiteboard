package main

import (
	"errors"
	color_lib "image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Vec2 struct {
    x, y float32
}

func cursorPosition() Vec2 {
    x, y := ebiten.CursorPosition()
    return Vec2{
        x: float32(x),
        y: float32(y),
    }
}

type stroke struct {
    points []Vec2
    width float32
    color color_lib.Color
}

func (s *stroke) appendIfMoved(point Vec2) {
    if len(s.points) != 0 {
        lastPoint := s.points[len(s.points)-1]
        if lastPoint.x == point.x && lastPoint.y == point.y {
            return
        }

    }
    s.points = append(s.points, point)
}

func (s *stroke) drawLast(screen *ebiten.Image) error {
    if s.color == nil {
        return errors.New("stroke has nil color")
    }

    if len(s.points) == 0 {
        return nil
    }

    from := s.points[len(s.points)-1]
    vector.DrawFilledCircle(screen, from.x, from.y, s.width/2, s.color, true)

    if len(s.points) == 1 {
        return nil
    }

    to := s.points[len(s.points)-2]
    vector.StrokeLine(
        screen,
        from.x,
        from.y,
        to.x,
        to.y,
        s.width,
        s.color,
        true,
    )
    vector.DrawFilledCircle(screen, to.x, to.y, s.width/2, s.color, true)

    return nil
}

func (s *stroke) draw(screen *ebiten.Image) error {
    if s.color == nil {
        return errors.New("stroke has nil color")
    }

    from := s.points[0]
    vector.DrawFilledCircle(screen, from.x, from.y, s.width/2, s.color, true)
    for _, to := range s.points[1:] {
        vector.StrokeLine(
            screen,
            from.x,
            from.y,
            to.x,
            to.y,
            s.width,
            s.color,
            true,
        )
        vector.DrawFilledCircle(screen, to.x, to.y, s.width/2, s.color, true)
        from = to
    }

    return nil
}
