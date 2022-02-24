// Copyright 2022 Mark Stahl. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the BSD-LICENSE file.
package main

import (
	"testing"
)

func Test_NoCommand_ErrorUsage(t *testing.T) {
	m := NewMain()
	if err := m.Run(); err != ErrorUsage {
		t.Fatalf(err.Error())
	}
}
