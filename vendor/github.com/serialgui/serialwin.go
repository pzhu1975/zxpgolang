package serialgui

import (
	"fmt"
	//	"log"
	//	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Iomode byte

type Tooldatabind struct {
	Deviceid       int
	Daudratevalue  int
	Datawidthvalue int
	Parityvalue    int
	Stopbitvalue   int
	Receivemode    Iomode
	Sendmode       Iomode
	AutoSend       bool
	SendText       string
}

const (
	Inputascii Iomode = 1 + iota
	Inputbinary
)

const (
	Outputscii Iomode = 1 + iota
	Outputbinary
)

type RateData struct {
	Rate     int
	Ratedata string
}

func BaudRate() []*RateData {
	return []*RateData{
		{9600, "9600"},
		{14400, "14400"},
		{19200, "19200"},
		{38400, "38400"},
		{115200, "115200"},
		{230400, "230400"},
		{380400, "380400"},
		{460800, "460800"},
		{921600, "921600"},
	}
}

type DeviceData struct {
	Device     int
	Devicename string
}

func DeviceID() []*DeviceData {
	return []*DeviceData{
		{1, "COM1"},
		{2, "COM2"},
		{3, "COM3"},
		{4, "COM4"},
		{5, "COM5"},
		{6, "COM6"},
		{7, "COM7"},
		{8, "COM8"},
		{9, "COM9"},
	}
}

type WidthData struct {
	Width     int
	Widthdata string
}

func DataWidth() []*WidthData {
	return []*WidthData{
		{1, "5"},
		{2, "6"},
		{3, "7"},
		{4, "8"},
	}
}

type ParityData struct {
	Parity     int
	Paritydata string
}

func ParityD() []*ParityData {
	return []*ParityData{
		{1, "无校验"},
		{2, "奇校验"},
		{3, "偶校验"},
	}
}

type StopData struct {
	Stop     int
	Stopdata string
}

func StopBit() []*StopData {
	return []*StopData{
		{1, "一位停止位"},
		{2, "一位半停止位"},
		{2, "二位停止位"},
	}
}
func SerialWindow() {
	//	mw.Close()
	//	var tooldatas *ToolDataS
	var tooldatas *Tooldatabind
	tooldatas = new(Tooldatabind)
	listtext := "Hello Wold"
	tooldatas = &Tooldatabind{1, 115200, 4, 1, 1, Inputascii, Outputscii, false, "1000"}
	if err := (MainWindow{
		AssignTo: &mw, //.MainWindow,
		Title:    "串口调试器",
		DataBinder: DataBinder{
			AssignTo:       &db,
			AutoSubmit:     true,
			DataSource:     tooldatas,
			ErrorPresenter: ErrorPresenterRef{&ep},
		},
		MinSize: Size{600, 400},
		Layout:  VBox{},
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
										Value:         Bind("Deviceid", SelRequired{}), //
										BindingMember: "Device",
										DisplayMember: "Devicename",
										Model:         DeviceID(),
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
										Value:         Bind("Daudratevalue", SelRequired{}),
										BindingMember: "Rate",
										DisplayMember: "Ratedata",
										Model:         BaudRate(),
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
										Value:         Bind("Datawidthvalue", SelRequired{}),
										BindingMember: "Width",
										DisplayMember: "Widthdata",
										Model:         DataWidth(),
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
										Value:         Bind("Parityvalue", SelRequired{}),
										BindingMember: "Parity",
										DisplayMember: "Paritydata",
										Model:         ParityD(),
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
										Value:         Bind("Stopbitvalue", SelRequired{}),
										BindingMember: "Stop",
										DisplayMember: "Stopdata",
										Model:         StopBit(),
									},
								},
							},

							RadioButtonGroupBox{
								ColumnSpan: 2,
								Title:      "接收设置",
								Layout:     HBox{},
								DataMember: "Receivemode",
								Buttons: []RadioButton{
									{Text: "ACSII", Value: Inputascii},
									{Text: "二进制", Value: Inputbinary},
								},
							},
							RadioButtonGroupBox{
								ColumnSpan: 2,
								Title:      "发送设置",
								Layout:     HBox{},
								DataMember: "Sendmode",
								Buttons: []RadioButton{
									{Text: "ACSII", Value: Outputscii},
									{Text: "二进制", Value: Outputbinary},
								},
							},
							GroupBox{
								Name:   "发送方式设置",
								Title:  "发送方式设置",
								Layout: HBox{}, //Grid{Columns: 2}, //
								Children: []Widget{
									CheckBox{
										Name:    "自动",
										Text:    "自动",
										Checked: Bind("AutoSend"),
									},
									LineEdit{
										Name: "初始值",
										Text: Bind("SendText"),
									},
								},
							},
						},
					},
					ListBox{
						Name:       "串口数据",
						Column:     2,
						DataMember: listtext,
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
								OnClicked: func() {
									//									fmt.Println("TEST", tooldatas.Deviceid)
									fmt.Printf("%+v\n", tooldatas)
									//									tooldatas.Deviceid = 2
									//									deviceid = "COM4"
								},
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
