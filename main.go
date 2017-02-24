package main

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/crgimenes/goConfig"
)

type config struct {
	File   string `cfg:"file"`
	output string `cfg:"output" cfgDefault:"-"`
	Decode bool   `cfg:"decode" cfgDefault:"false"`
}

func main() {
	cfg := config{}

	goConfig.PrefixEnv = "BASE64"
	err := goConfig.Parse(&cfg)
	if err != nil {
		println(err)
		return
	}

	if cfg.File == "" {
		goConfig.Usage()
		return
	}

	buff, err := ioutil.ReadFile(cfg.File)
	if err != nil {
		println(err)
		return
	}

	var outBuff []byte
	if cfg.Decode {
		outBuff, err = base64.StdEncoding.DecodeString(string(buff))
		if err != nil {
			println(err)
			return
		}
		if cfg.output == "-" {
			println(outBuff)
		} else {
			err = saveFile(cfg.output, outBuff)
			if err != nil {
				println(err)
				return
			}
		}
	} else {
		base64.StdEncoding.Encode(outBuff, buff)
		if cfg.output == "-" {
			println(outBuff)
		} else {
			err = saveFile(cfg.output, outBuff)
			if err != nil {
				println(err)
				return
			}
		}
	}
}

func saveFile(fineName string, value []byte) (err error) {
	err = ioutil.WriteFile(fineName, value, 0644)
	return
}
