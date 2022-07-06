package main

// github.com/shurcooL/graphql
import (
	"context"
	"github.com/machinebox/graphql"
	"log"
)

type Counters struct {
	TotalItems         int64
	TotalPages         int64
	RequestedNrOfItems int64
	RequestedPageNr    int64
	ReceivedNrOfItems  int64
}

type Schedulers struct {
	Counters Counters
}

func main() {

	domainName := "http://localhost:30080/"
	graphQLReqPath := "api/worker/private/v1/"
	client := graphql.NewClient(domainName + graphQLReqPath)
	req := graphql.NewRequest(`
		{
			   schedulers{
				   Counters{
						TotalItems
						TotalPages
						RequestedNrOfItems
						RequestedPageNr
						ReceivedNrOfItems
				   }
				   Items{
					   ID
					   WorkerID
					   Worker{
						   Name
					   }
					   MiningConfigID
					   MiningConfig{
							ID
							Name
							Description
			
							MiningAppTypeID
							MiningAppType{
								ID
								Name
							}
							MiningAppVersionID
							MiningAppVersion{
									ID
									Description
									MiningAppTypeID
									MiningAppType{
										ID
										Name
									}
									Description
									Version
									Config
			
									IsActive
			
									FileID
									File{
										ID
										CategoryID
										Size
										Name
										FullName
										OriginalName
										Extension
									}
			
									CreatedAt
									UpdatedAt
			
								
							}
							Config
						   
					   }
					   From
					   To
			
					   AllDay
					   OnAlways
			
					   OnMonday
					   OnTuesday
					   OnWednesday
					   OnThursday
					   OnFriday
					   OnSaturday
					   OnSunday
			
					   CreatedAt
					   UpdatedAt
				   }
			   }
			}
	`)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Auth-Token", "c451ae9f2a27e9c4ff684444b2a8d5be6073dd9d2810b0f93f2613ee42c29dc0")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	//var respData ResponseStruct
	//var respData schedulers
	//var respData Schedulers
	var respData interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData)
}
