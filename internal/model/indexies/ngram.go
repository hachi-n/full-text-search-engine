package indexies

import (
	"fmt"
	"github.com/hachi-n/full-text-search-engine/internal/architecture"
	"github.com/hachi-n/full-text-search-engine/internal/model/documents"
	"github.com/hachi-n/full-text-search-engine/internal/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type SplitedWord = string
type WordPoints = []int
type NgramIndexMap map[SplitedWord]NgramPointMap
type NgramPointMap map[documents.ID]WordPoints

const (
	indexMapFileName = "index.map"
)

var (
	indexMapFilePath = filepath.Join(architecture.DatabaseDir, indexMapFileName)
)

func NewNgramIndexMap() NgramIndexMap {
	index := NgramIndexMap{"": NgramPointMap{}}
	if _, err := os.Stat(indexMapFilePath); err != nil {
		return index
	}

	if err := util.LoadJsonFile(indexMapFilePath, &index); err != nil {
		fmt.Errorf("error: %v", err)
	}

	return index
}

func (index NgramIndexMap) Add(id documents.ID, path documents.TextFilePath, numeric int) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := index.addWithDividing(id, b, numeric); err != nil {
		return err
	}

	return nil
}
func (index NgramIndexMap) addWithDividing(id documents.ID, b []byte, numeric int) error {
	words, err := index.divideString(string(b), numeric)
	if err != nil {
		return err
	}
	// add
	for point, word := range words {
		if ngramPointMap, ok := index[word]; ok {
			ngramPointMap[id] = append(ngramPointMap[id], point)
			index[word] = ngramPointMap
		} else {
			ngramPointMap = map[documents.ID]WordPoints{
				id: {point},
			}
			index[word] = ngramPointMap
		}
	}
	return nil
}

func (index NgramIndexMap) divideString(str string, numeric int) ([]string, error) {
	words := []string{}
	// unigram
	if numeric == 1 {
		words = strings.Split(str, "")
	} else {
		return nil, fmt.Errorf("numeric 1 only")
	}
	return words, nil
}

func (index NgramIndexMap) Save() error {
	jsonByte, err := util.PrettyJson(index)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(indexMapFilePath, jsonByte, 0644); err != nil {
		return err
	}

	return nil
}

func (index NgramIndexMap) Search(searchString string, numeric int) ([]documents.ID, error) {
	words, err := index.divideString(searchString, numeric)
	if err != nil {
		return nil, err
	}

	// TODO
	// Too messy.Toooooooo fxxxxxxxxxxxxck!!!!
	// I'm writing something that works for the time being because I didn't have time
	currentPointMap := make(map[documents.ID]WordPoints)
	i := 0
	ngramPointMap := index[words[i]]
	currentPointMap = ngramPointMap
	i++

	var status bool
	if len(words) < i {
		_, status = index[words[i]]
	} else {
		status = false
	}

	for status {
		ngramPointMap := index[words[i]]

		for id, points := range currentPointMap {
			points := func() []int {
				var incrPoints []int
				for _, p := range points {
					p++
					incrPoints = append(incrPoints, p)
				}
				return incrPoints
			}()

			ngramPoints := ngramPointMap[id]

			results := func(t, s []int) []int {
				var results []int
				for _, v1 := range t {
					for _, v2 := range s {
						if v1 == v2 {
							results = append(results, v1)
						}
					}
				}
				return results
			}(points, ngramPoints)

			currentPointMap[id] = results
		}
		i++
		if len(words) < i {
			_, status = index[words[i]]
		} else {
			status = false
		}
	}

	var documentIds []int
	for k, _ := range currentPointMap {
		documentIds = append(documentIds, k)
	}
	return documentIds, nil
}
