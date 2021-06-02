
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-archive-api/data"
)

// swagger:route POST /archive/create archive createNewArchive
// Save a transaction to archive for future reference.
//
// responses:
//	200: createArchiveResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Archive) Create_Archive (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-archive-api Module")
	request := &data.CreateArchiveRequest{}

	err:=request.FromJSONToCreateArchiveRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = request.ValidateCreateArchiveRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-archive-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add archive to mongoDB
	response := data.AddArchiveToMongo(request)

	//writing to the io.Writer
	w.Header().Set("Content-Type", "application/json")
	
	err = response.CreateArchiveResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}