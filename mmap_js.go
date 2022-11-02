// Copyright 2020 Evan Shaw. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmap

import "syscall"

var (
	EJS  = syscall.NewError("not supported by js")
)

func mmap(len int, inprot, inflags, fd uintptr, off int64) ([]byte, error) {
	return nil, EJS
}

func (m MMap) flush() error {
	return EJS
}

func (m MMap) lock() error {
	return EJS
}

func (m MMap) unlock() error {
	return EJS
}

func (m MMap) unmap() error {
	return EJS
}
