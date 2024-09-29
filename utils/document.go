package utils

import (
	"compress/gzip"
	"os"
	"encoding/xml"
)

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([]document, error) {
	f, err := os.Open(path)
	if err!=nil {
		return nil, err
	}
	defer f.Close()
	gz, error := gzip.NewReader(f)
	if error!=nil {
		return nil, error
	}
	defer gz.Close()
	dec := xml.NewDecoder(gz)
	dump := struct{
		Documents [] document `xml:"doc"`
	}{}
	if  error := dec.Decode(&dump); error!=nil{
		return nil,err
	}
	docs := dump.Documents
	for i:= range docs{
		docs[i].ID = i
	}
	return docs, nil
}