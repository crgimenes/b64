package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/crgimenes/goConfig"
)

type config struct {
	File   string `cfg:"file"`
	Decode bool   `cfg:"decode" cfgDefault:"false"`
}

func main() {
	cfg := config{}

	goConfig.PrefixEnv = "BASE64"
	err := goConfig.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("%#v\n", cfg)

	if cfg.File == "" {
		goConfig.Usage()
		return
	}

	buff, err := ioutil.ReadFile(cfg.File)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.Decode {
		decode, err := base64.StdEncoding.DecodeString(string(buff))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(decode)
	} else {
		encoded := base64.StdEncoding.EncodeToString(buff)
		fmt.Println(encoded)
	}

}
