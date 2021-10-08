package indexies

import "github.com/hachi-n/full-text-search-engine/internal/model/documents"

type NgramIndex struct {
	Word          string
	DocumentIdMap NgramIndexMap
}

type NgramIndexMap struct {
	DocumentId int
	WordPoints []int
}

func NewNgramIndex(id documents.ID) *NgramIndex {
	return nil
}

func (ngramIndex *NgramIndex) CreateIndex() error {
	return nil
}
