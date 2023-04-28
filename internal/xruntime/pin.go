// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xruntime

import (
    _ "runtime"
    _ "unsafe"
)

//go:linkname RuntimeProcPin runtime.procPin
func RuntimeProcPin() int

//go:linkname RuntimeProcUnpin runtime.procUnpin
func RuntimeProcUnpin()
