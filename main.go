package main

import (
	"flag"
	"log"
	"time"
	utils "github.com/aumkar-007/text_search_engine/utils"
)

func main(){
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p","enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "monkeys", "search query")
	flag.Parse()
	log.Println("Full text search is in progress")
	start :=time.Now()
	docs,err := utils.LoadDocuments(dumpPath)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents %v ", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents %v ", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents %v ", len(matchedIDs), time.Since(start))
	for _, id :=range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}