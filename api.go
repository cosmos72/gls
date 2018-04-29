// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gls

import "sync"

// return the current goroutine ID.
//
// note that the returned value is DIFFERENT from most other goroutine libraries:
// this GoId() returns the address, converted to uintptr, of the runtime.g struct.
// NOT the runtime.g.goid field returned by most other libraries.
func GoId() uintptr

// map of goroutine-local variables.
type Map map[interface{}]interface{}

var (
	table = make(map[uintptr]Map)
	lock  sync.RWMutex
)

// delete all goroutine-local variables.
// if a goroutine used *any* function from this package, except GoId() and DelAll(),
// it MUST invoke DelAll() before such goroutine exits, otherwise it will leak memory.
func DelAll() {
	id := GoId()
	lock.Lock()
	delete(table, id)
	lock.Unlock()
}

// get all goroutine-local variables. the returned Map is a mutable reference,
// i.e. changes to it are visible in subsequent calls to GetAll() and Get()
// from the same goroutine, until either DelAll() or SetAll() are invoked
func GetAll() Map {
	id := GoId()
	lock.Lock()
	m := table[id]
	if m == nil {
		m = make(Map)
		table[id] = m
	}
	lock.Unlock()
	return m
}

// set all goroutine-local variables. the Map argument is a mutable reference,
// i.e. if it's modified *after* the call to SetAll(),
// changes to it are visible in subsequent calls to GetAll() and Get()
// from the same goroutine, until either DelAll() or SetAll() are invoked
func SetAll(m Map) {
	if m == nil {
		m = make(Map)
	}
	id := GoId()
	lock.Lock()
	table[id] = m
	lock.Unlock()
}

// get a single goroutine-local variable.
// slightly faster than the equivalent GetAll()[key]
func Get(key interface{}) (interface{}, bool) {
	id := GoId()
	lock.RLock()
	m, ok := table[id]
	lock.RUnlock()
	val, ok := m[key]
	return val, ok
}

// set a single goroutine-local variable.
// equivalent to GetAll()[key] = val
func Set(key, val interface{}) {
	GetAll()[key] = val
}

// delete a single goroutine-local variable.
// slightly faster than the equivalent delete(GetAll(), key)
func Del(key interface{}) {
	id := GoId()
	lock.Lock()
	m, ok := table[id]
	lock.Unlock()
	if ok {
		delete(m, key)
	}
}
