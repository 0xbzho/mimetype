package matchers

import (
	"bytes"
)

// Class matches an java class file.
func Class(in []byte) bool {
	return len(in) > 4 && bytes.Equal(in[:4], []byte{0xCA, 0xFE, 0xBA, 0xBE})
}

// Swf matches an Adobe Flash swf file.
func Swf(in []byte) bool {
	return len(in) > 3 &&
		bytes.Equal(in[:3], []byte("CWS")) ||
		bytes.Equal(in[:3], []byte("FWS")) ||
		bytes.Equal(in[:3], []byte("ZWS"))
}

// Wasm matches a web assembly File Format file.
func Wasm(in []byte) bool {
	return len(in) > 4 && bytes.Equal(in[:4], []byte{0x00, 0x61, 0x73, 0x6D})
}

// Dbf matches a dBase file.
// https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm
func Dbf(in []byte) bool {
	// 3rd and 4th bytes contain the last update month and day of month
	if !(0 < in[2] && in[2] < 13 && 0 < in[3] && in[3] < 32) {
		return false
	}

	// dbf type is dictated by the first byte
	dbfTypes := []byte{
		0x02, 0x03, 0x04, 0x05, 0x30, 0x31, 0x32, 0x42, 0x62, 0x7B, 0x82,
		0x83, 0x87, 0x8A, 0x8B, 0x8E, 0xB3, 0xCB, 0xE5, 0xF5, 0xF4, 0xFB,
	}
	for _, b := range dbfTypes {
		if in[0] == b {
			return true
		}
	}

	return false
}
