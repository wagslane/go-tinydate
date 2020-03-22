package tinydate

// GobDecode implements the gob.GobDecoder interface
func (td *TinyDate) GobDecode(data []byte) error {
	return td.UnmarshalBinary(data)
}

// GobEncode implements the gob.GobEncoder interface
func (td TinyDate) GobEncode() ([]byte, error) {
	return td.MarshalBinary()
}
