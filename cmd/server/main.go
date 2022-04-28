/*
 * @Author: linin00
 * @Date: 2022-04-28 08:35:22
 * @LastEditTime: 2022-04-28 14:56:00
 * @LastEditors: linin00
 * @Description:
 * @FilePath: /CSsnake/cmd/server/main.go
 * 天道酬勤
 */
package main

import (
	"log"
	"net"
	"net/rpc"
	"snakegame/src"

	"github.com/nsf/termbox-go"
	"github.com/pkg/errors"
)

var comc1 = make(chan src.KeyEvent)
var comc2 = make(chan src.KeyEvent)
var game = src.NewGame(comc1, comc2)

type SnakeService struct {
}

func NewSnakeService() *SnakeService {
	return &SnakeService{}
}
func (p *SnakeService) Command(command src.Command_t, reply *struct{}) error {
	switch command.Snake {
	case 1:
		comc1 <- command.Event
	case 2:
		comc2 <- command.Event
	default:
		return errors.New("不存在")
	}
	return nil
}
func (p *SnakeService) Watch(request struct{}, reply *src.Packet_t) error {
	w, h := termbox.Size()
	*reply = src.Packet_t{
		Buffer: termbox.CellBuffer(),
		W:      w,
		H:      h,
	}
	return nil
}
func main() {
	_ = rpc.RegisterName("SnakeService", NewSnakeService())
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go src.ListenToKeyboard(comc1)
	go src.ListenToKeyboard(comc2)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				rpc.ServeConn(conn)
				defer conn.Close()
			}()
		}
	}()
	game.Start()
}
