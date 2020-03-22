package tinydate

// MarshalJSON implements the json.Marshaler interface
func (td TinyDate) MarshalJSON() ([]byte, error) {
	return []byte(td.Format(`"` + iso8601Date + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (td *TinyDate) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	var err error
	*td, err = Parse(`"`+iso8601Date+`"`, string(data))
	return err
}
