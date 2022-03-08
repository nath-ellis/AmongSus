package space

import "github.com/solarlune/resolv"

var Space *resolv.Space

func Init(width int, height int) {
	Space = resolv.NewSpace(width, height, 1, 1)
}