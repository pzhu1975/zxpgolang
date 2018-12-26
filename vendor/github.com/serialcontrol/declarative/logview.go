// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package declarative

import (
	"github.com/lxn/walk"
)

type LogView struct {
	// Window

	Background         Brush
	ContextMenuItems   []MenuItem
	Enabled            Property
	Font               Font
	MaxSize            Size
	MinSize            Size
	Name               string
	OnBoundsChanged    walk.EventHandler
	OnKeyDown          walk.KeyEventHandler
	OnKeyPress         walk.KeyEventHandler
	OnKeyUp            walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Persistent         bool
	RightToLeftReading bool
	ToolTipText        Property
	Visible            Property

	// Widget

	// LogView
	AssignTo **walk.LogView
	Value    Property
}

func (lv LogView) Create(builder *Builder) error {
	// var style uint32
	// if lv.HScroll {
	// 	style |= win.WS_HSCROLL
	// }
	// if lv.VScroll {
	// 	style |= win.WS_VSCROLL
	// }

	w, err := walk.NewLogView(builder.Parent())
	if err != nil {
		return err
	}

	if lv.AssignTo != nil {
		*lv.AssignTo = w
	}
	return builder.InitWidget(lv, w, func() error {
		return nil
	})
}
