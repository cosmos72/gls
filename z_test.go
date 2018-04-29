// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gls

import (
	"testing"
)

var verbose bool = false

func AsyncGoId() <-chan uintptr {
	ch := make(chan uintptr)
	go func() {
		ch <- GoId()
	}()
	return ch
}

func TestGoId(t *testing.T) {
	id1 := GoId()
	id2 := GoId()
	if id1 == id2 {
		if verbose {
			t.Logf("TestGoId: 0x%x == 0x%x", id1, id2)
		}
	} else {
		t.Errorf("TestGoId: 0x%x != 0x%x", id1, id2)
	}
}

func TestAsyncGoId1(t *testing.T) {
	id1 := GoId()
	id2 := <-AsyncGoId()
	if id1 != id2 {
		if verbose {
			t.Logf("TestAsyncGoId1: 0x%x != 0x%x", id1, id2)
		}
	} else {
		t.Errorf("TestAsyncGoId1: 0x%x == 0x%x", id1, id2)
	}
}

func TestAsyncGoId2(t *testing.T) {
	ch1 := AsyncGoId()
	ch2 := AsyncGoId()
	id1 := <-ch1
	id2 := <-ch2
	if id1 != id2 {
		if verbose {
			t.Logf("TestAsyncGoId2: 0x%x != 0x%x", id1, id2)
		}
	} else {
		t.Errorf("TestAsyncGoId2: 0x%x == 0x%x", id1, id2)
	}
}

// check that Get() returns repeteable results
func TestMap1(t *testing.T) {
	Set(0, '#')
	v, ok := Get(0)
	if ok && v == '#' {
		if verbose {
			t.Logf("TestMap1: expecting (#, true) found (%v, %v)", ok, v)
		}
	} else {
		t.Errorf("TestMap1: expecting (#, true) found (%v, %v)", ok, v)
	}
}

// check that changes to the map returned by GetAll()
// are visible in subsequent calls to Get() and GetAll()
func TestMap2(t *testing.T) {
	m := GetAll()
	m[1] = 2
	v, ok := Get(1)
	if ok && v == 2 {
		if verbose {
			t.Logf("TestMap2: expecting (2, true) found (%v, %v)", ok, v)
		}
	} else {
		t.Errorf("TestMap2: expecting (2, true) found (%v, %v)", ok, v)
	}
}

// check that changes to the map passed to SetAll()
// are visible in subsequent Get() and GetAll()
func TestMap3(t *testing.T) {
	m := make(Map)
	SetAll(m)
	m["a"] = "b"
	v, ok := Get("a")
	if ok && v == "b" {
		if verbose {
			t.Logf("TestMap3: expecting (b, true) found (%v, %v)", ok, v)
		}
	} else {
		t.Errorf("TestMap3: expecting (b, true) found (%v, %v)", ok, v)
	}
}

// check that different goroutines get independent maps
func TestMap4(t *testing.T) {
	Set(1.0, 2.0)
	m1 := GetAll()
	ch := make(chan Map)
	go func() {
		ch <- GetAll()
		DelAll()
	}()
	m2 := <-ch
	if len(m1) != len(m2) {
		if verbose {
			t.Logf("TestMap4: len(m1) != len(m2)")
		}
	} else {
		t.Errorf("TestMap4: len(m1) == len(m2)")
	}
}
