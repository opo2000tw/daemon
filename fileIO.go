package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"unsafe"
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

type fileGroup struct {
	PID *FileWriter
}

type FileWriter struct {
	fp *os.File
}

func (f *FileWriter) open(filename string) (err error) {
	f.fp, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	return err
}

func (f *FileWriter) create(filename string) (err error) {
	f.fp, err = os.Create(filename)
	return err
}

func (f *FileWriter) remove(filename string) (err error) {
	return os.Remove(filename)
}

func (f *FileWriter) dispose() (err error) {
	return f.fp.Close()
}

func (f *FileWriter) RW(s int, filename string, data string) (err error) {
	if _, err := os.Stat(filename); err != nil {
		if err = f.create(filename); err != nil {
			return err
		}
	} else {
		if err = f.open(filename); err != nil {
			return err
		}
	}
	defer f.dispose()
	switch s {
	case delete:
		if err := f.remove(filename); err != nil {
			return err
		}
	case read:
		if err := f.readString(); err != nil {
			return err
		}
	case write:
		if err := f.writeString(data); err != nil {
			return err
		}
	default:
		return errFile
	}
	return nil
}

func (f *FileWriter) Name() string {
	return f.fp.Name()
}

func (f *FileWriter) writeString(b string) (err error) {
	_, err = f.fp.WriteString(b)
	return err
}

func (f *FileWriter) writeByte(b []byte) (err error) {
	_, err = f.fp.Write(b)
	return err
}

func (f *FileWriter) readString() (err error) {
	lineReader := bufio.NewReader(f.fp)
	for {
		// 相同使用場景下可以採用的方法
		// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		// func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
		// func (b *Reader) ReadString(delim byte) (line string, err error)
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		// 如下是某些業務邏輯操作
		// 如下程式碼列印每次讀取的檔案行內容
		fmt.Println(string(line))
	}
	return err
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
