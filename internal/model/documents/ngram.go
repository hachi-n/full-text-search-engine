package documents

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/full-text-search-engine/internal/architecture"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ID int
type TextFilePath string
type NgramDocument map[ID]TextFilePath

const (
	documentMapFileName = "document.map"
)

var (
	documentMapFilePath = filepath.Join(architecture.DatabaseDir, documentMapFileName)
)

func NewNgramDocument() NgramDocument {
	var document = make(NgramDocument)
	if _, err := os.Stat(documentMapFilePath); err != nil {
		return document
	}
	f, err := os.Open(documentMapFilePath)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	if err := json.Unmarshal(b, &document); err != nil {
		fmt.Errorf("error: %v", err)
	}
	return document
}

func (document NgramDocument) Save() error {
	jsonByte, err := json.MarshalIndent(document, "", strings.Repeat(" ", 4))
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(documentMapFilePath, jsonByte, 0644); err != nil {
		return err
	}
	return nil
}

func (document NgramDocument) Add(textFilePath TextFilePath) error {
	textFilePathStr := string(textFilePath)
	documentDataFilePathStr := filepath.Join(architecture.DataDir, filepath.Base(textFilePathStr))
	documentDataFilePath := TextFilePath(documentDataFilePathStr)

	id := document.lastDocumentId()
	id++
	document[id] = documentDataFilePath
	lastDocumentId = id

	reader, err := os.Open(textFilePathStr)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(documentDataFilePathStr)
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

var lastDocumentId ID
func (document NgramDocument) LastDocumentId() ID {
	if lastDocumentId != 0 {
		return lastDocumentId
	}
	var max ID
	for key := range document {
		if key > max {
			max = key
		}
	}
	return max
}
