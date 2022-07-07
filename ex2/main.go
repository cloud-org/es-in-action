package main

import (
	"context"
	"es-in-action/common"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {

	client, err := elastic.NewClient(
		elastic.SetURL(common.Endpoint),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		return
	}

	version, err := client.ElasticsearchVersion(common.Endpoint)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("version: %v\n", version)

	//res, err := client.Get().Do(context.Background())
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//log.Printf("%+v\n", res)
	item := common.JobRet{
		HostId: "***REMOVED***",
		CronId: "***REMOVED***",
		Ctime:  "2022-07-07 22:33:33",
	}
	resp, err := client.Index().Index("cron").BodyJson(&item).Do(context.Background())
	if err != nil {
		log.Println("do err:", err)
		return
	}

	log.Printf("resp is %+v\n", resp)

	return
}
