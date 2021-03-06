package yaml_test

import (
	"testing"

	yaml "github.com/go-faster/yamlx"
)

func FuzzUnmarshal(f *testing.F) {
	addFuzzingCorpus(f)

	// TODO(tdakkota): move to addFuzzingCorpus, currently DecodeEncodeDecode fuzzing fails
	//  due to some marshaling issues
	for _, file := range readSuite(f) {
		for _, test := range file.Tests {
			f.Add([]byte(test.YAML))
			if test.JSON != "" {
				f.Add([]byte(test.JSON))
			}
		}
	}

	f.Fuzz(func(t *testing.T, input []byte) {
		var v interface{}
		_ = yaml.Unmarshal(input, &v)
	})
}
