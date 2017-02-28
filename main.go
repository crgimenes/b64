package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goConfig"
)

type config struct {
	File   string `cfg:"file"`
	Output string `cfg:"output" cfgDefault:"-"`
	Decode bool   `cfg:"decode" cfgDefault:"false"`
}

var errFileNotDefined = errors.New("Input file not defined")
var cfg config

func run() (err error) {
	if cfg.File == "" {
		err = errFileNotDefined
		return
	}

	var buff []byte
	buff, err = ioutil.ReadFile(cfg.File)
	if err != nil {
		return
	}

	var outBuff []byte
	if cfg.Decode {
		outBuff, err = base64.StdEncoding.DecodeString(string(buff))
		if err != nil {
			return
		}
	} else {
		outBuff = make([]byte, base64.StdEncoding.EncodedLen(len(buff)))
		base64.StdEncoding.Encode(outBuff, buff)
	}

	if cfg.Output == "-" {
		fmt.Println(string(outBuff))
		return nil
	}
	return ioutil.WriteFile(cfg.Output, outBuff, 0644)
}

func configAndRun() error {
	goConfig.PrefixEnv = "BASE64"
	if err := goConfig.Parse(&cfg); err != nil {
		return err
	}
	return run()
}

func main() {
	err := configAndRun()
	if err != nil {
		println(err.Error())
		goConfig.Usage()
		os.Exit(1)
	}
}
