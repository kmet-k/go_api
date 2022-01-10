package controllers

import (
	"encoding/json"
	"test/models"
	"test/repository"
	"test/utils"

	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var auth models.AuthenMember
	err := json.NewDecoder(request.Body).Decode(&auth)
	if err != nil {
		json.NewEncoder(response).Encode(models.Invalid_syntax())
	} else {
		member, err := repository.FindMemberUsernamePassword(auth)
		if err != nil {
			json.NewEncoder(response).Encode(models.Authen_invalid())
		} else {
			at, rt := utils.GenerateToken(member)
			json.NewEncoder(response).Encode(models.RespondAuthen{
				Result:  true,
				Message: "Login success",
				Data: models.Authen{
					Firstname:    member.Firstname,
					Lastname:     member.Lastname,
					Username:     member.Username,
					AccessToken:  at,
					RefreshToken: rt,
				},
			})
		}
	}
}

func VerifyAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		authHeader := request.Header.Get("Authorization")
		tokenString := authHeader[len("Bearer "):]
		if len(authHeader) <= 0 {
			response.WriteHeader(401)
			json.NewEncoder(response).Encode(models.Invalid_token())
		} else {
			tokenValid, _ := utils.ValidAccessToken(tokenString)
			if !tokenValid.Valid {
				response.WriteHeader(401)
				json.NewEncoder(response).Encode(models.Invalid_token())
			} else {
				next.ServeHTTP(response, request)
			}
		}
	})
}

func GetNewToken(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var auth token
	err := json.NewDecoder(request.Body).Decode(&auth)
	if err != nil {
		utils.SentMessage(response, models.Invalid_syntax())
	} else {
		jwt, err := utils.ValidRefreshToken(auth.RefreshToken)
		if !jwt.Valid {
			if err.Error() == "Token is expired" {
				response.WriteHeader(403)
				json.NewEncoder(response).Encode(models.Invalid_token())
			} else {
				response.WriteHeader(401)
				json.NewEncoder(response).Encode(models.Invalid_token())
			}
		} else {
			id := utils.ClaimsID(auth.RefreshToken)
			objID, _ := primitive.ObjectIDFromHex(id)
			member, err := repository.FindMember(objID)
			if err != nil {
				json.NewEncoder(response).Encode(models.Username_Not_Found())
			} else {
				at, rt := utils.GenerateToken(member)
				json.NewEncoder(response).Encode(token{
					AccessToken:  at,
					RefreshToken: rt,
				})
			}
		}
	}
}

type token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func Logout(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var auth token
	err := json.NewDecoder(request.Body).Decode(&auth)
	if err != nil {
		utils.SentMessage(response, models.Invalid_syntax())
	} else {
		tokenValid, _ := utils.ValidRefreshToken(auth.RefreshToken)
		if !tokenValid.Valid {
			utils.SentMessage(response, models.Invalid_token())
		} else {
			utils.SentMessage(response, models.Logout_success())
		}
	}
}
