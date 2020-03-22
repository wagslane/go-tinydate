package tinydate

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (td TinyDate) MarshalText() ([]byte, error) {
	return []byte(td.Format(iso8601Date)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (td *TinyDate) UnmarshalText(data []byte) error {
	newTD, err := Parse(iso8601Date, string(data))
	if err != nil {
		return err
	}
	td = &newTD
	return nil
}
