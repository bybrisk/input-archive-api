package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/input-archive-api/data"
)

/*func TestAddArchiveToMongo(t *testing.T){
	var emptyString string
	payload := &data.CreateArchiveRequest{
		UserID: "6083deb86fcd474489784fee",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
		TransactionObject: []data.TransactionBlock{
			data.TransactionBlock{
				Bot: "Which transaction is this?",
				User: "latest",
			},
			data.TransactionBlock{
				Bot: "When do you want it?",
				User: "02/05/2021",
			},
			data.TransactionBlock{
				Bot: "Do you want packing?",
				User: "Yes",
			},
			data.TransactionBlock{
				Bot: "Thanks for ordering",
				User: emptyString,
			},
		},
	}
	res := data.AddArchiveToMongo(payload)
	fmt.Println(res)
}*/

func TestGetArchiveFromMongo (t *testing.T) {
	payLoad := &data.GetArchiveRequestPagewise{
		UserID: "6083deb86fcd474489784fee",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
		Page: 5,
	}

	res := data.GetArchiveFromMongo(payLoad)
	fmt.Println(res)
}
