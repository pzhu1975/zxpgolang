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

	"github.com/lxn/walk"
	//	. "github.com/lxn/walk/declarative"
	//	"github.com/lxn/win"
	"serialgui"
)

var (
	uiflag *string //environment parament
	inTE   *walk.TextEdit

//	mycanvas *walk.Canvas
//	mytangle walk.Rectangle
)

type MyMainWindow struct {
	*walk.MainWindow
	tabWidget *walk.TabWidget
	//	tabPage   *walk.TabPage
	//	prevFilePath string
}

func main() {
	uiflag = flag.String("_ui", "gui", "   <选择操作界面> [web] [gui] [cmd]")
	flag.Parse()
	fmt.Println(*uiflag)
	serialgui.SerialWindow()
	fmt.Println("maiwindow test")
}
