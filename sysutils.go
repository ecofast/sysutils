// Copyright 2016~2017 ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package sysutils implements some useful system utility functions
// in the way of which Delphi(2007) RTL has done.
package sysutils

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StrToInt(s string) int {
	if ret, err := strconv.Atoi(s); err == nil {
		return ret
	}
	panic(fmt.Sprintf("Cannot convert %s to int!", s))
}

func StrToIntDef(s string, defaultvalue int) int {
	if ret, err := strconv.Atoi(s); err == nil {
		return ret
	}
	return defaultvalue
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

func BytesToStr(bs []byte) string {
	for i := 0; i < len(bs); i++ {
		if bs[i] == 0 {
			return string(bs[0:i])
		}
	}
	return string(bs)
}

func BytesToUInt8(bs []byte) uint8 {
	return uint8(bs[0])
}

func BytesToUInt16(bs []byte) uint16 {
	return binary.LittleEndian.Uint16(bs)
}

func BytesToUInt32(bs []byte) uint32 {
	return binary.LittleEndian.Uint32(bs)
}

func UInt16ToBytes(v uint16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, v)
	return bs
}

func Int16ToBytes(v int16) []byte {
	return UInt16ToBytes(uint16(v))
}

func UInt32ToBytes(v uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, v)
	return bs
}

func Int32ToBytes(v int32) []byte {
	return UInt32ToBytes(uint32(v))
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

func ChangeFileExt(filename, extension string) string {
	i := strings.LastIndex(filename, ".")
	if i > 0 {
		ret := filename[:i]
		return ret + extension
	}
	return filename + extension
}

func DateTimeToStr(dt time.Time) string {
	year, month, day := dt.Date()
	hour, min, sec := dt.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, int(month), day, hour, min, sec)
}

func DateToStr(dt time.Time) string {
	year, month, day := dt.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day)
}

func TimeToStr(dt time.Time) string {
	hour, min, sec := dt.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

func MilliSecondsBetween(now, then time.Time) int64 {
	return int64(now.Sub(then) / time.Millisecond)
}

func SecondsBetween(now, then time.Time) int64 {
	return MilliSecondsBetween(now, then) / 1000
}
