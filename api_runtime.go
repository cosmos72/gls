// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !386,!amd64,!arm,!arm64,!mips,!mips64,!ppc64,!s390x

// gls/runtime is deceivingly simple...
// but requires a lot of trickery to compile and install.
// prefer assembler implementations id_*.s

package gls

import (
	"github.com/cosmos72/gls/runtime"
)

// return the current goroutine ID.
//
// note that the returned value is DIFFERENT from most other goroutine libraries:
// this GoID() returns the address, converted to uintptr, of the runtime.g struct.
// NOT the runtime.g.goid field returned by most other libraries.
func GoID() uintptr {
	return runtime.GoID()
}
