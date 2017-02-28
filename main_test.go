package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigAndRun(t *testing.T) {
	err := configAndRun()
	if err != errFileNotDefined {
		t.Fatal("errFileNotDefined expected")
	}
}

func TestRun(t *testing.T) {
	cfg.File = ""
	_, err := run()
	if err != errFileNotDefined {
		t.Fatal("errFileNotDefined expected")
	}

	cfg.File = "error file"
	_, err = run()
	if err == nil {
		t.Fatal("error expected")
	}

	cfg.Decode = false
	cfg.File = "./testdata/test.txt"
	var outBuff []byte
	outBuff, err = run()
	if err != nil {
		t.Fatal(err)
	}

	if string(outBuff) != "anVzdCBhIHRlc3Qgc3RyaW5n" {
		t.Fatal("error encode string to base64")
	}

	cfg.Decode = true
	cfg.File = "./testdata/test.b64"
	outBuff, err = run()
	if err != nil {
		t.Fatal(err)
	}
	if string(outBuff) != "just a test string" {
		t.Fatal("error decode base64 to string")
	}

	tmpfile, err := ioutil.TempFile("", "B64TEST")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	cfg.Decode = true
	cfg.File = "./testdata/test.b64"
	cfg.Output = tmpfile.Name()
	outBuff, err = run()
	if err != nil {
		t.Fatal(err)
	}
	if string(outBuff) != "just a test string" {
		t.Fatal("error decode base64 to string")
	}

	b, _ := ioutil.ReadFile(tmpfile.Name())
	if string(b) != "just a test string" {
		t.Fatal("error decode base64 to file")
	}

}
