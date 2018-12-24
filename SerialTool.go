package main

//go build -ldflags="-H windowsgui"  for hidden the cmd windows

import (
	"flag"
	"fmt"

	//	"math"
	//	"os"
	//	"path"
	//	"strings"
	//	"unicode/utf8"

	//	"github.com/lxn/walk"
	//	. "github.com/lxn/walk/declarative"
	//	"github.com/lxn/win"
	"github.com/serialgui"
)

var (
	uiflag *string //environment parament
//	mycanvas *walk.Canvas
//	mytangle walk.Rectangle
)

func main() {
	uiflag = flag.String("_ui", "gui", "   <选择操作界面> [web] [gui] [cmd]")
	flag.Parse()
	fmt.Println(*uiflag)
	serialgui.SerialWindow()
	fmt.Println("maiwindow test")
}
