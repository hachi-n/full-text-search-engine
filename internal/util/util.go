package util

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type filePath = string

func FileCopy(dst, src filePath) error {
	reader, err := os.Open(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer writer.Close()
	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}
	return nil
}

func PrettyJson(v interface{}) ([]byte, error) {
	const count = 4
	return json.MarshalIndent(v, "", strings.Repeat(" ", count))
}

func LoadJsonFile(src filePath, strct interface{}) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, strct); err != nil {
		return err
	}
	return nil
}
