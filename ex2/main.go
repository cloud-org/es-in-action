package main

//import (
//	"es-in-action/common"
//	"github.com/olivere/elastic/v7"
//	"log"
//)
//
//func main() {
//
//	client, err := elastic.NewClient(
//		elastic.SetURL(common.Endpoint),
//		elastic.SetSniff(false),
//	)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	version, err := client.ElasticsearchVersion(common.Endpoint)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	log.Printf("version: %v\n", version)
//	//res, err := client.Get().Do(context.Background())
//	//if err != nil {
//	//	log.Println(err)
//	//	return
//	//}
//	//
//	//log.Printf("%+v\n", res)
//
//	return
//}
