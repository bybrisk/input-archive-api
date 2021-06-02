package data

import (
	"encoding/json"
	"io"
)	

func (d *CreateArchiveResponse) CreateArchiveResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *CreateArchiveRequest) FromJSONToCreateArchiveRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}