package internal

import (
	"encoding/json"
	"os"
)

func ReadJSON(filename string, data interface{}) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err1 := f.Close(); err == nil {
			err = err1
		}
	}()
	dec := json.NewDecoder(f)
	if err := dec.Decode(data); err != nil {
		return err
	}
	return
}

func WriteJSON(filename string, data interface{}) (err error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err1 := f.Close(); err == nil {
			err = err1
		}
	}()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return
}
