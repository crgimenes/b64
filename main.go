package main

import (
	"encoding/base64"
	"errors"
	"flag"
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

func run() (outBuff []byte, err error) {
	if cfg.File == "" {
		lastPar := flag.NArg() - 1
		cfg.File = flag.Arg(lastPar)
		if cfg.File == "" {
			err = errFileNotDefined
			return
		}
	}

	var buff []byte
	buff, err = ioutil.ReadFile(cfg.File)

	if err != nil {
		return
	}

	if cfg.Decode {
		outBuff = make([]byte, base64.StdEncoding.DecodedLen(len(buff)))
		_, err = base64.StdEncoding.Decode(outBuff, buff)
		if err != nil {
			return
		}
	} else {
		outBuff = make([]byte, base64.StdEncoding.EncodedLen(len(buff)))
		base64.StdEncoding.Encode(outBuff, buff)
	}

	if cfg.Output == "-" {
		fmt.Println(string(outBuff))
		return
	}
	err = ioutil.WriteFile(cfg.Output, outBuff, 0644)

	return
}

func configAndRun() (err error) {
	goConfig.PrefixEnv = "BASE64"
	err = goConfig.Parse(&cfg)
	if err != nil {
		return
	}
	_, err = run()
	return
}

func main() {
	if err := configAndRun(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
