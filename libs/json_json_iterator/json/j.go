package json

import (
	sysJson "encoding/json"
	"io"

	jsoNiter "github.com/json-iterator/go"
)

type RawMessage = sysJson.RawMessage

var instance = jsoNiter.ConfigCompatibleWithStandardLibrary

func Marshal(obj interface{}) ([]byte, error) {
	return instance.Marshal(obj)
}

func Unmarshal(data []byte, v interface{}) error {
	return instance.Unmarshal(data, v)
}

func NewDecoder(reader io.Reader) *jsoNiter.Decoder {
	return jsoNiter.NewDecoder(reader)
}

func NewEncoder(writer io.Writer) *jsoNiter.Encoder {
	return jsoNiter.NewEncoder(writer)
}
