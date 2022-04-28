/*
 * @Author: linin00
 * @Date: 2022-04-28 06:13:31
 * @LastEditTime: 2022-04-28 13:16:23
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/src/screen.go
 * 天道酬勤
 */
package src

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor1  = termbox.ColorGreen
	snakeColor2  = termbox.ColorRed
)

func (g *Game_t) Render() error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		left   = (w - g.Area.Width) / 2
		right  = (w + g.Area.Width) / 2
		top    = midY - (g.Area.Height / 2)
		bottom = midY + (g.Area.Height / 2) + 1
	)

	renderTitle(left, top)
	renderArea(g.Area, top, bottom, left)
	renderSnake1(left, bottom, g.Area.Snake_A)
	renderSnake2(left, bottom, g.Area.Snake_B)
	renderFood(left, bottom, g.Area.Food)
	renderScore(left, bottom, g.Score_A, g.Score_B)
	renderQuitMessage(right, bottom)
	return termbox.Flush()
}

func renderSnake1(left, bottom int, s *Snake_t) {
	for _, b := range s.Body {
		termbox.SetCell(left+b.X, bottom-b.Y, ' ', snakeColor1, snakeColor1)
	}
}

func renderSnake2(left, bottom int, s *Snake_t) {
	for _, b := range s.Body {
		termbox.SetCell(left+b.X, bottom-b.Y, ' ', snakeColor2, snakeColor2)
	}
}

func renderFood(left, bottom int, f *Food_t) {
	termbox.SetCell(left+f.Pos.X, bottom-f.Pos.Y, f.Emoji, defaultColor, bgColor)
}

func renderArea(a *Area_t, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+a.Width, i, '│', defaultColor, bgColor)
	}
	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+a.Width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+a.Width, bottom, '┘', defaultColor, bgColor)
	fill(left, top, a.Width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, a.Width, 1, termbox.Cell{Ch: '─'})
}

func renderScore(left, bottom, s1, s2 int) {
	score := fmt.Sprintf("Score of green snake: %v", s1)
	tbprint(left, bottom+1, defaultColor, defaultColor, score)
	score = fmt.Sprintf("Score of red snake: %v", s2)
	tbprint(left, bottom+2, defaultColor, defaultColor, score)
}

func renderQuitMessage(right, bottom int) {
	m := "Press ESC to quit"
	tbprint(right-len(m), bottom+1, defaultColor, defaultColor, m)
}

func renderTitle(left, top int) {
	tbprint(left, top-1, defaultColor, defaultColor, "Snake Game")
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func tbprint(x int, y int, fg termbox.Attribute, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
