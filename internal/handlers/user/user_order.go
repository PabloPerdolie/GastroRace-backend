package userhand

import (
	"net/http"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {
	//ctx := cont.Get(r, "user")
	//user, ok := ctx.(middleware.UserData)
	//
	//var order models.Order
	//json.NewDecoder(r.Body).Decode(&order)
	//
	//if !ok {
	//	http.Error(w, "Failed to transform", http.StatusBadRequest)
	//	return
	//}
	// todo insert
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {

}

//todo status change
