package cbrotli_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/portablecloud/brotli"
)

const dataPath = "./test_data"

var testData = []struct {
	InputFile  string
	OutputFile string
}{
	{"case_1.bin.br", "case_1.bin"},
	{"case_2.bin.br", "case_2.bin"},
	{"case_3.html.br", "case_3.html"},
}

func TestReader(t *testing.T) {
	for _, pair := range testData {
		input, err := os.OpenFile(filepath.Join(dataPath, pair.InputFile), os.O_RDONLY, 0600)
		if err != nil {
			t.Fatalf("error during input file opening, %s", err)
			return
		}
		defer input.Close()
		output, err := ioutil.ReadFile(filepath.Join(dataPath, pair.OutputFile))
		if err != nil {
			t.Fatalf("error during output file opening, %s", err)
			return
		}

		reader := cbrotli.NewReader(input)
		defer reader.Close()

		decoded, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Fatalf("error during decompression, %s", err)
			return
		}

		if bytes.Compare(decoded, output) != 0 {
			t.Fatal("bytes are not matching")
			return
		}
	}
}

func TestDecode(t *testing.T) {
	for _, pair := range testData {
		input, err := ioutil.ReadFile(filepath.Join(dataPath, pair.InputFile))
		if err != nil {
			t.Fatalf("error during input file opening, %s", err)
			return
		}
		output, err := ioutil.ReadFile(filepath.Join(dataPath, pair.OutputFile))
		if err != nil {
			t.Fatalf("error during output file opening, %s", err)
			return
		}

		decoded, err := cbrotli.Decode(input)
		if err != nil {
			t.Fatalf("error during decompression, %s", err)
			return
		}

		if bytes.Compare(decoded, output) != 0 {
			t.Fatal("bytes are not matching")
			return
		}
	}
}
