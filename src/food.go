/*
 * @Author: linin00
 * @Date: 2022-04-28 06:10:33
 * @LastEditTime: 2022-04-28 06:29:51
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/src/food.go
 * 天道酬勤
 */
package src

import (
	"math/rand"
	"os"
	"strings"
)

type Food_t struct {
	Pos    Index_t
	Points int
	Emoji  rune
}

func NewFood(x, y int) *Food_t {
	return &Food_t{
		Points: rand.Intn(10) + 1,
		Emoji:  GetFoodEmoji(),
		Pos: Index_t{
			X: x,
			Y: y,
		},
	}
}

func GetFoodEmoji() rune {
	if HasUnicodeSupport() {
		return RandomFoodEmoji()
	}

	return '@'
}

func RandomFoodEmoji() rune {
	f := []rune{
		'🍒',
		'🍍',
		'🍑',
		'🍇',
		'🍏',
		'🍌',
		'🍫',
		'🍭',
		'🍕',
		'🍩',
		'🍗',
		'🍖',
		'🍬',
		'🍤',
		'🍪',
	}

	return f[rand.Intn(len(f))]
}

func HasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
