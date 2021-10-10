package indexing

import (
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

	ngramDocument := documents.NewNgramDocumentMap()
	if err := ngramDocument.Add(textFilePath); err != nil {
		return err
	}
	if err := ngramDocument.Save(); err != nil {
		return err
	}

	ngramIndex := indexies.NewNgramIndexMap()
	if err := ngramIndex.Add(ngramDocument.LastDocumentId(), textFilePath, numeric); err != nil {
		return err
	}
	if err := ngramIndex.Save(); err != nil {
		return err
	}

	return nil
}
