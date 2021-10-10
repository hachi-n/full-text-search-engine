package search

import (
	"fmt"
	"github.com/hachi-n/full-text-search-engine/internal/architecture"
	"github.com/hachi-n/full-text-search-engine/internal/model/documents"
	"github.com/hachi-n/full-text-search-engine/internal/model/indexies"
	"io/ioutil"
)

func Apply(numeric int, value string) error {
	arch := new(architecture.Architecture)
	if err := arch.Validate(); err != nil {
		return err
	}

	ngramDocument := documents.NewNgramDocumentMap()

	ngramIndex := indexies.NewNgramIndexMap()
	documentIds, err := ngramIndex.Search(value, numeric)
	if err != nil {
		return err
	}

	documentPaths := ngramDocument.Search(documentIds)

	for _, documentPath := range documentPaths {
		fmt.Println(documentPath)
		b, err := ioutil.ReadFile(documentPath)
		if err != nil {
			fmt.Println("can't open file")
		}
		fmt.Println(string(b))
	}
	return nil
}
