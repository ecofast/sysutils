// Copyright 2016~2017 ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package sysutils implements some useful system utility functions
// in the way of which Delphi(2007) RTL has done.
package sysutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func BoolToStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func StrToBool(s string) bool {
	if ret, err := strconv.ParseBool(s); err == nil {
		return ret
	}
	return false
}

func FloatToStr(f float64) string {
	return fmt.Sprintf("%g", f)
}

func GetApplicationPath() string {
	path := filepath.Dir(os.Args[0])
	return path + string(os.PathSeparator)
}

func DirectoryExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		return true
	}
	return false
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func CreateFile(filename string) bool {
	// os.MkdirAll(path.Dir(filename))
	_, err := os.Create(filename)
	if err == nil {
		return true
	}
	return false
}

func IncludeTrailingBackslash(path string) string {
	if !strings.HasSuffix(path, string(os.PathSeparator)) {
		return path + string(os.PathSeparator)
	}
	return path
}
