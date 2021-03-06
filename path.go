package main

import (
	"github.com/mitchellh/go-homedir"
	"github.com/q191201771/naza/pkg/nazalog"
)

func DirPath(expaned string) string {
	var path string
	if expaned == "" {
		path = dirHome()
	} else {
		expaned = "~/" + expaned
		path = dirHomeExpand(expaned)
	}
	return path
}

func dirHome() string {
	dirHome, err := homedir.Dir()
	if err != nil {
		nazalog.Fatal(err)
	}
	return dirHome
}

func dirHomeExpand(expaned string) string {
	dirHomeExpand, err := homedir.Expand(expaned)
	if err != nil {
		nazalog.Fatal(err)
	}
	return dirHomeExpand
}
