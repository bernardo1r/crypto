package crypto

import (
	"bytes"
	"errors"
)

func Pad(data *[]byte, n uint8) []byte {
	pad := int(n) - (len(*data) % int(n))
	*data = append(*data, bytes.Repeat([]byte{byte(pad)}, pad)...)
	return *data
}

func Unpad(data *[]byte) ([]byte, error) {
	pad := int((*data)[len(*data)-1])
	if len(*data) < pad {
		return nil, errors.New("padding too large for data")
	}

	if bytes.Count((*data)[len(*data)-pad:], []byte{byte(pad)}) != pad {
		return nil, errors.New("padding too small for data")
	}

	*data = (*data)[:len(*data)-pad]
	return *data, nil
}
