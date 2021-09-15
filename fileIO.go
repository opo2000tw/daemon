package main

import (
	"errors"
	"log"
	"os"
	"reflect"
	"unsafe"

	"github.com/q191201771/naza/pkg/nazalog"
)

const (
	create = iota
	update
	retrieve
	delete
	write
	read
)

var errFile = errors.New("FileWriter: fuck")

type FileWriter struct {
	str string
	fp  *os.File
}

func (f *FileWriter) Open(filename string) (err error) {
	f.fp, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	return err
}

func (f *FileWriter) Create(filename string) (err error) {
	f.fp, err = os.Create(filename)
	return err
}

func (f *FileWriter) Remove(filename string) (err error) {
	return os.Remove(filename)
}

func (f *FileWriter) Dispose() (err error) {
	if f.fp == nil {
		return errFile
	}
	return f.fp.Close()
}

func (f *FileWriter) RW(s int, filename string, data string) (err error) {
	if _, err := os.Stat(filename); err != nil {
		if err = f.Create(filename); err != nil {
			return err
		}
	} else {
		if err = f.Open(filename); err != nil {
			return err
		}
	}
	defer f.fp.Close()
	switch s {
	case delete:
		if err := f.Dispose(); err != nil {
			return err
		}
		if err := f.Remove(filename); err != nil {
			return err
		}
	case read:
		f.ReadString()
	case write:
		f.WriteString(data)
	default:
		nazalog.Debugf("can't find this option")
	}
	return nil
}

func (f *FileWriter) Name() string {
	if f.fp == nil {
		return ""
	}
	return f.fp.Name()
}

func (f *FileWriter) WriteString(b string) (err error) {
	if f.fp == nil {
		return errFile
	}
	_, err = f.fp.WriteString(b)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (f *FileWriter) WriteByte(b []byte) (err error) {
	if f.fp == nil {
		return errFile
	}
	_, err = f.fp.Write(b)
	return err
}

func (f *FileWriter) ReadString() (err error) {
	if f.fp == nil {
		return errFile
	}
	data := make([]byte, 100)
	count, err := f.fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	nazalog.Debugf("read %d bytes: %q\n", count, data[:count])
	return nil
}

func string2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
