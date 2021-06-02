package data

import (
	"time"
)

func AddArchiveToMongo(d *CreateArchiveRequest) *CreateArchiveResponse{
	
	var response CreateArchiveResponse

	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= CreateArchiveResponse{
			Response:"Error! User not registered!",
			Status:403,
		}
	} else{
		
		//Add date and ranking time to the object
		t2e2 := time.Now()
		d.RankingIndex = t2e2.UnixNano()
		d.TransactionDate = t2e2.Format("2006-01-02")

		err:= AddArchiveToMongoDocument(d)
		if err!=nil{
			response= CreateArchiveResponse{
				Response:"Error! Some error occured!",
				Status:501,
			}
		} else{
			response= CreateArchiveResponse{
				Response:"success",
				Status:200,
			}
		}
	}
	return &response
}