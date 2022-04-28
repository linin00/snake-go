/*
 * @Author: linin00
 * @Date: 2022-04-28 06:10:52
 * @LastEditTime: 2022-04-28 10:41:20
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/src/snake.go
 * 天道酬勤
 */
package src

type direction int

const (
	LEFT direction = 1 + iota
	RIGHT
	UP
	DOWN
)

type Snake_t struct {
	Body      []Index_t
	Direction direction
	Length    int
}

func NewSnake(dir direction, len int, head Index_t) *Snake_t {
	var res Snake_t
	res.Body = make([]Index_t, len)
	res.Body[len-1] = head
	for i := len - 2; i >= 0; i-- {
		switch dir {
		case LEFT:
			res.Body[i].X = res.Body[i+1].X + 1
			res.Body[i].Y = res.Body[i+1].Y
		case RIGHT:
			res.Body[i].X = res.Body[i+1].X - 1
			res.Body[i].Y = res.Body[i+1].Y
		case UP:
			res.Body[i].X = res.Body[i+1].X
			res.Body[i].Y = res.Body[i+1].Y - 1
		case DOWN:
			res.Body[i].X = res.Body[i+1].X
			res.Body[i].Y = res.Body[i+1].Y + 1
		}
	}
	res.Direction = dir
	res.Length = len
	return &res
}

func (snake *Snake_t) ChangePosition(newDir direction) {
	opposite := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}
	if o := opposite[newDir]; o != 0 && o != snake.Direction {
		snake.Direction = newDir
	}
}

func (snake *Snake_t) head() Index_t {
	return snake.Body[len(snake.Body)-1]
}

func (snake *Snake_t) Move() {
	h := snake.head()
	c := Index_t{
		X: h.X,
		Y: h.Y,
	}
	switch snake.Direction {
	case RIGHT:
		c.X++
	case LEFT:
		c.X--
	case UP:
		c.Y++
	case DOWN:
		c.Y--
	}
	if snake.Length > len(snake.Body) {
		snake.Body = append(snake.Body, c)
	} else {
		snake.Body = append(snake.Body[1:], c)
	}
}

func (snake *Snake_t) isOnPosition(idx Index_t) bool {
	for _, t := range snake.Body {
		if t.X == idx.X && t.Y == idx.Y {
			return true
		}
	}

	return false
}
