package helper

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

func ParseTime(timeString string) (time.Time, error) {
	resp, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}

	return resp, err
}

func ParseToStruct(data interface{}, m protoiface.MessageV1) error {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	js, err := jspbMarshal.MarshalToString(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(js), data)
	return err
}

func ParseToProto(m protoiface.MessageV1, data interface{}) error {
	var jspbUnmarshal jsonpb.Unmarshaler
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = jspbUnmarshal.Unmarshal(bytes.NewBuffer(js), m)

	return err
}
