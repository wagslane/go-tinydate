package tinydate

import (
	"errors"
)

const tinyDateBinaryVersion byte = 1

// MarshalBinary implements the encoding.BinaryMarshaler interface
func (td TinyDate) MarshalBinary() ([]byte, error) {
	enc := []byte{
		tinyDateBinaryVersion, // byte 0 : version
		byte(td.year >> 8),    // byte 1-2: year
		byte(td.year),
		byte(td.month), // byte 3: month
		byte(td.day),   // byte 4: day
	}
	return enc, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface
func (td *TinyDate) UnmarshalBinary(data []byte) error {
	buf := data
	if len(buf) == 0 {
		return errors.New("tinydate: no data")
	}

	if buf[0] != tinyDateBinaryVersion {
		return errors.New("tinydate: unsupported version")
	}

	if len(buf) != /*version*/ 1+ /*year*/ 2+ /*month*/ 1+ /*day*/ 1 {
		return errors.New("tinydate: invalid length")
	}

	td.year = uint16(buf[2]) | uint16(buf[1])<<8
	td.month = uint8(buf[3])
	td.day = uint8(buf[4])

	return nil
}
