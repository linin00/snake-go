/*
 * @Author: linin00
 * @Date: 2022-04-28 07:41:58
 * @LastEditTime: 2022-04-28 14:55:06
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/src/keyboard.go
 * 天道酬勤
 */
package src

import (
	"github.com/nsf/termbox-go"
)

type eventType int

const (
	MOVE eventType = 1 + iota
	RETRY
	END
)

type KeyEvent struct {
	Type eventType
	Key  termbox.Key
}

func KeyToDirection(key termbox.Key) direction {
	switch key {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	case termbox.KeyArrowDown:
		return DOWN
	default:
		return 0
	}
}

func ListenToKeyboard(evChan chan KeyEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- KeyEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyArrowDown:
				evChan <- KeyEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyArrowRight:
				evChan <- KeyEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyArrowUp:
				evChan <- KeyEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyEsc:
				evChan <- KeyEvent{Type: END, Key: 0}
			default:
				if ev.Ch == 'r' {
					evChan <- KeyEvent{Type: RETRY, Key: 0}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
