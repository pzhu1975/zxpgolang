package serialgui

import (
	//	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type RateData struct {
	Rate     int
	Ratedata string
}

func BaudRate() []*RateData {
	return []*RateData{
		{9600, "9600"},
		{115200, "115200"},
	}
}

type DeviceData struct {
	Device     int
	Devicedata string
}

func DeviceID() []*DeviceData {
	return []*DeviceData{
		{0, "COM1"},
		{1, "COM2"},
		{2, "COM3"},
		{3, "COM4"},
		{4, "COM5"},
		{5, "COM6"},
		{6, "COM7"},
		{7, "COM8"},
		{8, "COM9"},
	}
}

type WidthData struct {
	Width     int
	Widthdata string
}

func DataWidth() []*WidthData {
	return []*WidthData{
		{0, "7"},
		{1, "8"},
		{1, "9"},
	}
}

type ParityData struct {
	Parity     int
	Paritydata string
}

func Parity() []*ParityData {
	return []*ParityData{
		{0, "奇校验"},
		{1, "偶校验"},
	}
}

type StopData struct {
	Stop     int
	Stopdata string
}

func StopBit() []*StopData {
	return []*StopData{
		{0, "有停止位"},
		{1, "无停止位"},
	}
}
func SerialWindow() {
	//	mw.Close()
	if err := (MainWindow{
		AssignTo: &mw, //.MainWindow,
		Title:    "串口调试器",
		MinSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Name:    "Lay window",
				MaxSize: Size{200, 400},
				Children: []Widget{
					VSplitter{
						Name: "Lay window",
						Children: []Widget{
							GroupBox{
								Name:   "设备号设置",
								Title:  "设备号设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									Label{
										Text: "设备号",
									},
									ComboBox{
										CurrentIndex:  0,
										BindingMember: "Device",
										DisplayMember: "Devicedata",
										Model:         DeviceID(),
										MaxLength:     8,
										//										MinSize:   Size{30, 10},
									},
								},
							},
							GroupBox{
								Name:   "波特率设置",
								Title:  "波特率设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									Label{
										Text: "波特率",
									},
									ComboBox{
										CurrentIndex:  0,
										BindingMember: "Rate",
										DisplayMember: "Ratedata",
										Model:         BaudRate(),
										MaxLength:     8,
										//										MinSize:   Size{30, 10},
									},
								},
							},
							GroupBox{
								Name:   "数据位设置",
								Title:  "数据位设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									Label{
										Text: "数据位",
									},
									ComboBox{
										CurrentIndex:  0,
										BindingMember: "Width",
										DisplayMember: "Widthdata",
										Model:         DataWidth(),
										MaxLength:     8,
										//										MinSize:   Size{30, 10},
									},
								},
							},
							GroupBox{
								Name:   "校验位设置",
								Title:  "校验位设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									Label{
										Text: "校验位",
									},
									ComboBox{
										CurrentIndex:  0,
										BindingMember: "Parity",
										DisplayMember: "Paritydata",
										Model:         Parity(),
										MaxLength:     8,
										//MinSize:       Size{30, 10},
									},
								},
							},
							GroupBox{
								Name:   "停止位设置",
								Title:  "停止位设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									Label{
										Text: "停止位",
									},
									ComboBox{
										CurrentIndex:  0,
										BindingMember: "Stop",
										DisplayMember: "Stopdata",
										Model:         StopBit(),
										MaxLength:     8,
										//MinSize:   Size{30, 10},
									},
								},
							},
							GroupBox{
								Name:   "接收设置",
								Title:  "接收设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									RadioButton{
										Name: "ACSII",
										Text: "ACSII",
									},
									RadioButton{
										Name: "二进制",
										Text: "二进制",
									},
								},
							},
							GroupBox{
								Name:   "发送设置",
								Title:  "发送设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									RadioButton{
										Name: "ACSII",
										Text: "ACSII",
									},
									RadioButton{
										Name: "二进制",
										Text: "二进制",
									},
								},
							},
							GroupBox{
								Name:   "发送方式设置",
								Title:  "发送方式设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									CheckBox{
										Name: "自动",
										Text: "自动",
									},
									LineEdit{
										Name: "初始值",
										Text: "1000",
									},
								},
							},
						},
					},
					ListBox{
						Name:   "串口数据",
						Column: 2,
					},
				},
			},
			Composite{
				Name:   "串口数据",
				Layout: HBox{}, //Grid{Columns: 2},
				Children: []Widget{
					TextEdit{Name: "串口输入"},
					GroupBox{
						Name:   "发送方式设置",
						Layout: VBox{}, //Grid{Columns: 2}, //
						Children: []Widget{
							PushButton{
								Name: "打开串口",
								Text: "打开串口",
							},
							PushButton{
								Name: "清除输入",
								Text: "清除输入",
							},
							PushButton{
								Name: "清除输出",
								Text: "清除输出",
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		panic(err)
	}
	// 运行窗体程序
	mw.Run()
}

// 点击开始事件
// func clientStart() {

// 	if runStopBtn.Text() == "重新连接服务器" {
// 		runStopBtn.SetEnabled(false)
// 		runStopBtn.SetText("正在连接服务器…")
// 		clientStop()
// 		return
// 	}

// 	runStopBtn.SetText("断开服务器连接")

// }
