package world

import (
	"github.com/nath-ellis/AmongSus/space"
	"github.com/solarlune/resolv"
)

var Speed float64 = 2

func Init() {
	NewObject(0, 476, "platform")
	NewObject(124, 476, "platform")
	NewObject(248, 476, "platform")
	NewObject(372, 476, "platform")
	NewObject(496, 476, "platform")
	NewObject(620, 476, "platform")
	NewObject(744, 476, "platform")
	NewObject(868, 476, "platform")
	NewObject(992, 476, "platform")
	NewObject(1116, 476, "platform")
	NewObject(1240, 476, "platform")

	space.Space.Add(resolv.NewObject(0, 476, 1200, 124, "object"))
}
