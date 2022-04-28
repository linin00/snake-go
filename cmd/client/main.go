/*
 * @Author: linin00
 * @Date: 2022-04-28 11:26:28
 * @LastEditTime: 2022-04-28 14:56:26
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/cmd/client/main.go
 * 天道酬勤
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"snakegame/src"

	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor1  = termbox.ColorGreen
	snakeColor2  = termbox.ColorRed
)

func main() {
	fmt.Println("输入远程地址")
	var addr string
	fmt.Scanln(&addr)
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	recive := make(chan src.Packet_t)
	go func() {
		for {
			var board src.Packet_t
			err = client.Call("SnakeService.Watch", struct{}{}, &board)
			if err != nil {
				log.Fatal(err)
			}
			recive <- board
		}
	}()
	event := make(chan src.KeyEvent)
	go src.ListenToKeyboard(event)
	termbox.Init()
	defer termbox.Close()
	for {
		select {
		case board := <-recive:
			w, h := board.W, board.H
			termbox.Clear(defaultColor, defaultColor)
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					cell := board.Buffer[w*i+j]
					termbox.SetCell(j, i, cell.Ch, cell.Fg, cell.Bg)
				}
			}
			termbox.Flush()
		case e := <-event:
			err = client.Call("SnakeService.Command", src.Command_t{
				Snake: 2,
				Event: e,
			}, new(struct{}))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
