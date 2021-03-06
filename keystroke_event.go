// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

type KeyStrokeEvent struct {
	Character rune
	Modifier  KeyboardModifier
}

func (e KeyStrokeEvent) String() string {
	return concat(e.Modifier.String(), string(e.Character))
}
