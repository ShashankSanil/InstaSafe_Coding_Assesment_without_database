package repository

import (
	"instasafe1/common"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	Transaction_Details []map[string]interface{}
	User_Details        []map[string]interface{}
}

func (tr *Repository) FindUserByEmail(Email string) []interface{} {
	var data []interface{}
	for _, user := range tr.User_Details {
		if user["Email"].(string) == Email {
			data = append(data, user)
		}
	}
	return data
}

func (tr *Repository) CreateEndUser(data map[string]interface{}) interface{} {
	uid := strings.Split(primitive.NewObjectID().String(), "\"")[1]
	data["_id"] = uid
	tr.User_Details = append(tr.User_Details, data)
	return common.GetItemById(tr.User_Details, uid)
}

func (tr *Repository) CreateTransaction(data map[string]interface{}) interface{} {
	tid := strings.Split(primitive.NewObjectID().String(), "\"")[1]
	data["_id"] = tid
	tr.Transaction_Details = append(tr.Transaction_Details, data)
	return common.GetItemById(tr.Transaction_Details, tid)
}

func (tr *Repository) GetUserByID(Id string) map[string]interface{} {
	return common.GetItemById(tr.User_Details, Id)
}

func (tr *Repository) GetAllTransactions() []map[string]interface{} {
	return tr.Transaction_Details
}

func (tr *Repository) GetTransactionByLocation(city string) []map[string]interface{} {
	output := make([]map[string]interface{}, 0)
	for _, transaction := range tr.Transaction_Details {
		if transaction["city"] == city {
			output = append(output, transaction)
		}
	}
	return output
}

func (tr *Repository) DeleteAllTransactions() []map[string]interface{} {
	emptyTransaction_Details := make([]map[string]interface{}, 0)
	tr.Transaction_Details = emptyTransaction_Details
	return emptyTransaction_Details
}

func (tr *Repository) UpdateLocation(uid string, payload map[string]interface{}) interface{} {
	userDetails := tr.GetUserByID(uid)
	userDetails["city"] = payload["city"]
	userDetails["resetLocation"] = payload["resetLocation"]
	return tr.GetUserByID(uid)
}
