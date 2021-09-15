package main

import (
	"errors"
	"log"
	"os"

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
	fileName string
	fp       *os.File
}

func (f *FileWriter) Open(filename string) (err error) {
	f.fp, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
	if f.fp == nil {
		err = f.Create(filename)
		if err != nil {
			return err
		}
		err = f.Open(filename)
		if err != nil {
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
	case retrieve:
		nazalog.Debugf("%+v", f.Name())
	case read:
		//f.fp.ReadString
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

func (f *FileWriter) WriteRaw(b []byte) (err error) {
	if f.fp == nil {
		return errFile
	}
	_, err = f.fp.Write(b)
	return err
}

func (f *FileWriter) ReadString(b []byte) (err error) {
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
