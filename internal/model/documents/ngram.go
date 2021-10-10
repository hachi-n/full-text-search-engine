package documents

import (
	"fmt"
	"github.com/hachi-n/full-text-search-engine/internal/architecture"
	"github.com/hachi-n/full-text-search-engine/internal/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ID = int
type TextFilePath = string

type NgramDocumentMap map[ID]TextFilePath

const (
	documentMapFileName = "document.map"
)

var (
	documentMapFilePath = filepath.Join(architecture.DatabaseDir, documentMapFileName)
)

func NewNgramDocumentMap() NgramDocumentMap {
	var document = make(NgramDocumentMap)
	if _, err := os.Stat(documentMapFilePath); err != nil {
		return document
	}

	if err := util.LoadJsonFile(documentMapFilePath, &document); err != nil {
		fmt.Errorf("error: %v", err)
	}

	return document
}

func (document NgramDocumentMap) Save() error {
	jsonByte, err := util.PrettyJson(document)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(documentMapFilePath, jsonByte, 0644); err != nil {
		return err
	}
	return nil
}

func (document NgramDocumentMap) Add(textFilePath TextFilePath) error {
	documentDataFilePath := filepath.Join(architecture.DataDir, filepath.Base(textFilePath))
	id := document.LastDocumentId()
	id++
	document[id] = documentDataFilePath
	lastDocumentId = id

	if err := util.FileCopy(documentDataFilePath, textFilePath); err != nil {
		return err
	}

	return nil
}

var lastDocumentId ID

func (document NgramDocumentMap) LastDocumentId() ID {
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

func (document NgramDocumentMap) Search(ids []ID) []TextFilePath {
	var paths []TextFilePath
	for _, id := range ids {
		paths = append(paths, document[id])
	}
	return paths
}
