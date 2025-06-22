package index

import (
	"fmt"
	//"log"

	"github.com/blevesearch/bleve"
)
var(
    title string
    short_description string
    image_url string
    class string
)
func BSearch(val string, Res bleve.Index) (map[string]interface{}, error) {
    query := bleve.NewMatchQuery(val)
    searchRequest := bleve.NewSearchRequest(query)
    searchResult, err := Res.Search(searchRequest)
    if err != nil {
        return nil, fmt.Errorf("error searching index: %v", err)
    }

    results := make(map[string]interface{})
    for _, hit := range searchResult.Hits {
        doc, err := Res.Document(hit.ID)
        if err != nil {
            return nil, fmt.Errorf("error retrieving document: %v", err)
        }

        docData := make(map[string]interface{})
        for _, field := range doc.Fields {
            docData[field.Name()] = string(field.Value())
        }

        results[hit.ID] = docData
    }
    //log.Printf("the results %v", results)
    return results, nil
}
