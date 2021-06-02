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

func GetArchiveFromMongo(d *GetArchiveRequestPagewise) *GetArchiveResponsePagewise{

	var response GetArchiveResponsePagewise
	var emptyDataStruct PagewiseArchiveReponseFromDB

	//page is less than zero error
	if (d.Page < 1) {
		response= GetArchiveResponsePagewise{
			Data: emptyDataStruct,
			Response:"Error! Request error!",
			Status:401,
		}
		return &response
	}

	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= GetArchiveResponsePagewise{
			Data: emptyDataStruct,
			Response:"Error! User not registered!",
			Status:403,
		}
	} else{

		data,err := GetArchiveFromMongoDocument(d)
	
		if err!=nil {
			response= GetArchiveResponsePagewise{
				Data: emptyDataStruct,
				Response:"Error! Some error occured!",
				Status:501,
			}
		} else if data.RankingIndex == 0 {
			response= GetArchiveResponsePagewise{
				Data: data,
				Response:"Error! Page out of range!",
				Status:200,
			}
		} else{
			response= GetArchiveResponsePagewise{
				Data: data,
				Response:"success",
				Status:200,
			}
		}
	}
	return &response
}