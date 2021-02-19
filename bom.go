package csvreader

import (
	"bytes"
)

var BOM_UTF8 = []byte{239, 187, 191}

func bomCheck(data []byte) []byte {
	if bytes.Equal(data[:3], BOM_UTF8) {
		return data[3:]
	}
	return data
}
