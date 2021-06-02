
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-archive-api/data"
)

// swagger:route POST /archive/get/page archive getArchivePagewise
// Get transaction archive pagewise (per transaction)
//
// responses:
//	200: getArchiveResponsePagewise
//  422: errorValidation
//  501: errorResponses

func (p *Input_Archive) Get_Archive_Pagewise (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-archive-api Module")
	request := &data.GetArchiveRequestPagewise{}

	err:=request.FromJSONToGetArchiveRequestPagewise(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = request.ValidateGetArchiveRequestPagewise()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-archive-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//get archive from mongoDB pagewie
	response := data.GetArchiveFromMongo(request)

	//writing to the io.Writer
	w.Header().Set("Content-Type", "application/json")
	
	err = response.GetArchiveResponsePagewiseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}