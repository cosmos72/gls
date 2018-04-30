// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
  to use compile this file, run the command:

	go tool compile -o runtime.a -pack -std -+ id.go

  which tell gc compiler that it's part of the package "runtime" in the standard library
  If you are not fleeing in panic and screaming already, you should.

  Then install it manually with the commands:

	DEST="$(go env GOPATH)/pkg/$(go env GOOS)_$(go env GOARCH)/github.com/cosmos72/gls"
	mkdir -p     "$DEST"
	mv runtime.a "$DEST"
*/

package runtime

import (
	"unsafe"
)

// getg returns the pointer to the current g.
// The compiler rewrites calls to this function into instructions
// that fetch the g directly (from TLS or from the dedicated register).
func getg() *struct{}

//go:nosplit
func GoID() uintptr {
	return uintptr(unsafe.Pointer(getg()))
}
