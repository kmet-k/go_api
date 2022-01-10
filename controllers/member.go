package controllers

import (
	"test/models"
	"test/repository"
	"test/utils"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMember(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	members, err := repository.FindAllMember()
	if err != nil {
		json.NewEncoder(response).Encode(models.Get_data_error())
	} else {
		if members == nil {
			members = make([]models.Member, 0)
			json.NewEncoder(response).Encode(members)
		} else {
			json.NewEncoder(response).Encode(members)
		}
	}
}

func GetMember(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	objID, _ := primitive.ObjectIDFromHex(param["id"])
	member, err := repository.FindMember(objID)
	if err != nil {
		json.NewEncoder(response).Encode(models.Get_data_error())
	} else {
		json.NewEncoder(response).Encode(member)
	}
}

func PostMember(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var member models.Member
	err := json.NewDecoder(request.Body).Decode(&member)
	if err != nil {
		json.NewEncoder(response).Encode(models.Invalid_syntax())
	} else {
		// utils.Encrypt(&member.Password)
		err = repository.InsertMember(member)
		if err != nil {
			utils.SentMessage(response, models.Insert_error())
		} else {
			utils.SentMessage(response, models.Insert_success())
		}
	}
}

func PutMember(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	objID, _ := primitive.ObjectIDFromHex(param["id"])
	_, err := repository.FindMember(objID)
	if err != nil {
		message := models.ID_not_found()
		utils.SentMessage(response, message)
	} else {
		var member models.PutMember
		err = json.NewDecoder(request.Body).Decode(&member)
		if err != nil {
			utils.SentMessage(response, models.Invalid_syntax())
		} else {
			err = repository.UpdateMember(member, objID)
			if err != nil {
				json.NewEncoder(response).Encode(models.Edit_error())
			} else {
				json.NewEncoder(response).Encode(models.Edit_success())
			}
		}
	}
}

func DelMember(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	param := mux.Vars(request)
	objID, _ := primitive.ObjectIDFromHex(param["id"])
	_, err := repository.FindMember(objID)
	if err != nil {
		utils.SentMessage(response, models.ID_not_found())
	} else {
		err := repository.DeleteMember(objID)
		if err != nil {
			json.NewEncoder(response).Encode(models.Delete_error())
		} else {
			json.NewEncoder(response).Encode(models.Delete_success())

		}
	}
}
