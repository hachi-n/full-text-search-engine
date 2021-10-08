package indexing

import (
	"fmt"
	"github.com/hachi-n/full-text-search-engine/internal/architecture"
	"github.com/hachi-n/full-text-search-engine/internal/model/documents"
	"github.com/hachi-n/full-text-search-engine/internal/model/indexies"
)

func Apply(numeric int, textFilePath string) error {
	arch := new(architecture.Architecture)
	if err := arch.Validate(); err != nil {
		if err := arch.Create(); err != nil {
			return err
		}
	}

	ngramDocument := documents.NewNgramDocument()
	if err := ngramDocument.Add(documents.TextFilePath(textFilePath)); err != nil {
		return err
	}
	if err := ngramDocument.Save(); err != nil {
		return err
	}



	ngramIndex := indexies.NewNgramIndex(ngramDocument.LastDocumentId())
	if err := ngramIndex.CreateIndex(); err != nil {
		return nil
	}

	fmt.Println(ngramDocument)
	return nil
}
