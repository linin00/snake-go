/*
 * @Author: linin00
 * @Date: 2022-04-28 06:10:41
 * @LastEditTime: 2022-04-28 10:47:14
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/src/area.go
 * 天道酬勤
 */
package src

import (
	"math/rand"
	"time"
)

type Area_t struct {
	Snake_A *Snake_t
	Snake_B *Snake_t
	Food    *Food_t
	Width   int
	Height  int
	pointc1 chan int
	pointc2 chan (int)
}

func NewArea(s1 *Snake_t, s2 *Snake_t, w int, h int, c1 chan (int), c2 chan (int)) *Area_t {
	rand.Seed(time.Now().UnixNano())
	res := &Area_t{
		Snake_A: s1,
		Snake_B: s2,
		Width:   w,
		Height:  h,
		pointc1: c1,
		pointc2: c2,
	}
	res.PlaceFood()
	return res
}

func (a *Area_t) MoveSnake() int {
	a.Snake_A.Move()
	a.Snake_B.Move()
	if ret := a.isSnakeLeftArea(); ret != 0 {
		return ret
	}

	if a.HasFood(a.Snake_A.head()) {
		go func() {
			a.pointc1 <- a.Food.Points
		}()
		a.Snake_A.Length++
		a.PlaceFood()
	} else if a.HasFood(a.Snake_B.head()) {
		go func() {
			a.pointc2 <- a.Food.Points
		}()
		a.Snake_B.Length++
		a.PlaceFood()
	}
	return 0
}

func (a *Area_t) HasFood(idx Index_t) bool {
	return a.Food.Pos.X == idx.X && a.Food.Pos.Y == idx.Y
}

func (a *Area_t) PlaceFood() {
	var idx Index_t
	for {
		idx = Index_t{
			X: rand.Intn(a.Width),
			Y: rand.Intn(a.Height),
		}
		if !a.isOccupied(idx) {
			break
		}
	}
	a.Food = NewFood(idx.X, idx.Y)
}

func (a *Area_t) isSnakeLeftArea() int {
	h1 := a.Snake_A.head()
	h2 := a.Snake_B.head()
	if h1.X > a.Width || h1.Y > a.Height || h1.X < 0 || h1.Y < 0 {
		return 1
	} else if h2.X > a.Width || h2.Y > a.Height || h2.X < 0 || h2.Y < 0 {
		return 2
	}
	return 0
}

func (a *Area_t) isOccupied(idx Index_t) bool {
	return a.Snake_A.isOnPosition(idx) || a.Snake_B.isOnPosition(idx)
}
