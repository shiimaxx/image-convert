package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var cases = []struct {
	args    string
	outputs []string
}{
	{
		args: "image-convert testdata",
		outputs: []string{
			"testdata/icon-001.png",
			"testdata/dir-001/icon-002.png",
			"testdata/dir-002/icon-003.png",
			"testdata/dir-002/dir-002-001/icon-004.png",
		},
	},
	{
		args: "image-convert -s jpg -d gif testdata",
		outputs: []string{
			"testdata/icon-001.gif",
			"testdata/dir-001/icon-002.gif",
			"testdata/dir-002/icon-003.gif",
			"testdata/dir-002/dir-002-001/icon-004.gif",
		},
	},
}

func TestRun_ImageConvert(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	for _, c := range cases {
		args := strings.Split(c.args, " ")
		status := cli.Run(args)
		if status != ExitCodeOK {
			t.Errorf("expected %d to eq %d", status, ExitCodeOK)
		}
		for _, o := range c.outputs {
			_, err := os.Stat(o)
			if os.IsNotExist(err) {
				t.Error(err)
			}
		}
	}

}

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("image-convert -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("image-convert version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_noArguments(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := []string{"image-convert"}

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := "Missing arguments\n"
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_fileNotExists(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("image-convert dummy_file", " ")

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := "dummy_file: No such file or directory\n"
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_isNotDir(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	tempfile, _ := ioutil.TempFile("", "temp")
	defer os.Remove(tempfile.Name())

	args := strings.Split(fmt.Sprintf("image-convert %s", tempfile.Name()), " ")

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := fmt.Sprintf("%s: Is a not directory\n", tempfile.Name())
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}
