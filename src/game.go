/*
 * @Author: linin00
 * @Date: 2022-04-28 06:11:04
 * @LastEditTime: 2022-04-28 14:55:09
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/src/game.go
 * 天道酬勤
 */
package src

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Command_t struct {
	Snake int
	Event KeyEvent
}
type Packet_t struct {
	Buffer []termbox.Cell
	W      int
	H      int
}
type Game_t struct {
	Area    *Area_t
	Score_A int
	Score_B int
	isOver  bool

	pointc1 chan int
	pointc2 chan int
	comc1   chan KeyEvent
	comc2   chan KeyEvent
}

func NewGame(comc1, comc2 chan KeyEvent) *Game_t {
	w := 50
	h := 20
	pointc1 := make(chan int)
	pointc2 := make(chan int)
	Area := NewArea(NewSnake(RIGHT, 4, Index_t{4, 1}), NewSnake(LEFT, 4, Index_t{w - 4, h - 1}), w, h, pointc1, pointc2)
	return &Game_t{
		Score_A: 0,
		Score_B: 0,
		Area:    Area,
		isOver:  false,
		pointc1: pointc1,
		pointc2: pointc2,
		comc1:   comc1,
		comc2:   comc2,
	}
}

func (g *Game_t) retry() {
	w := 50
	h := 20
	Area := NewArea(NewSnake(RIGHT, 4, Index_t{4, 1}), NewSnake(LEFT, 4, Index_t{w - 4, h - 1}), w, h, g.pointc1, g.pointc2)
	g.Area = Area
	g.Score_A = 0
	g.Score_B = 0
	g.isOver = false
}

func (g *Game_t) restart(ret int) {
	switch ret {
	case 1:
		g.Score_A = 0
		g.Area.Snake_A = NewSnake(RIGHT, 4, Index_t{4, 1})
	case 2:
		g.Score_B = 0
		g.Area.Snake_B = NewSnake(LEFT, 4, Index_t{g.Area.Width - 4, g.Area.Height - 1})
	}
}

func (g *Game_t) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	if err := g.Render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case p1 := <-g.pointc1:
			g.Score_A += p1
		case p2 := <-g.pointc2:
			g.Score_B += p2
		case e1 := <-g.comc1:
			switch e1.Type {
			case MOVE:
				d := KeyToDirection(e1.Key)
				g.Area.Snake_A.ChangePosition(d)
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
		case e2 := <-g.comc2:
			switch e2.Type {
			case MOVE:
				d := KeyToDirection(e2.Key)
				g.Area.Snake_B.ChangePosition(d)
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if ret := g.Area.MoveSnake(); ret != 0 {
					g.restart(ret)
				}
			}
			if err := g.Render(); err != nil {
				panic(err)
			}
			time.Sleep(time.Duration(100) * time.Millisecond)
		}
	}
}
