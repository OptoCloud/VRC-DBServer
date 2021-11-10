package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"vrcdb/core"
	"vrcdb/database"
	"vrcdb/httpServer/helpers"
	"vrcdb/httpServer/requestModels"
	"vrcdb/models/account"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthGetAccountKeyFromRegKey(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	authValues := helpers.GetContextAuthValues(requestContext)

	if authValues.Cookie == nil {
		helpers.WriteGenericError(w, "general", "AuthToken cookie is populated", http.StatusBadRequest)
		return
	}

	var requestData requestModels.AuthLoginGetRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		helpers.WriteGenericError(w, "general", "Invalid json data", http.StatusBadRequest)
		return
	}

	authValues.Account, err = database.GetApiAccountByUsername(requestContext, requestData.Username)
	if err != nil {
		helpers.WriteGenericError(w, "general", "Invalid username/password", http.StatusUnauthorized)
		return
	}

	if authValues.Account.Password.Hash != core.ComputeServerhashSimple(requestData.PasswordHash, requestData.Username) {
		helpers.WriteGenericError(w, "general", "Invalid username/password", http.StatusUnauthorized)
		return
	}

	var found bool = false
	var authToken account.DbAuthToken
	for _, authToken = range authValues.Account.AuthTokens {
		if authToken.IpAddress == authValues.IpAddress && authToken.HardwareId == authValues.HardwareId {
			found = true
			break
		}
	}

	if !found {
		utcNow := time.Now().UTC()

		authToken := account.DbAuthToken{
			Id:         primitive.NewObjectID(),
			IpAddress:  authValues.IpAddress,
			IpCountry:  authValues.IpCountry,
			HardwareId: authValues.HardwareId,
			ExpiresAt:  utcNow.AddDate(core.AuthTokenDurationYears, core.AuthTokenDurationMonths, core.AuthTokenDurationDays),
		}

		query := bson.M{"_id": authValues.Account.Id}
		update := bson.M{
			"$set": bson.M{
				"last_login": utcNow,
				"last_seen":  utcNow,
			},
			"$addToSet": bson.M{
				"authtokens": authToken,
			},
		}

		_, err = database.CollectionAccounts.UpdateOne(requestContext, query, update)
		if err != nil {
			log.Printf("Failed to insert authtoken: %s\n", err.Error())
			helpers.WriteGenericError(w, "general", "Server accounted an error", http.StatusInternalServerError)
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     core.HttpCookieAuthTokenKey,
		Value:    authToken.Id.Hex(),
		Path:     "/",
		Expires:  authToken.ExpiresAt,
		Secure:   true,
		HttpOnly: true,
	})

	helpers.WriteJson(w, authValues.Account.ToHttp(), http.StatusOK)
}
func AuthLogoutPut(w http.ResponseWriter, r *http.Request) {
	var err error

	requestContext := r.Context()
	authValues := helpers.GetContextAuthValues(requestContext)

	utcNow := time.Now().UTC()

	query := bson.M{"_id": authValues.Account.Id}
	update := bson.M{
		"$set": bson.M{
			"last_seen": utcNow,
		},
		"$pull": bson.M{
			"authtokens": bson.M{
				"$elemMatch": bson.M{
					"hwid": authValues.HardwareId,
				},
			},
		},
	}

	_, err = database.CollectionAccounts.UpdateOne(requestContext, query, update)
	if err != nil {
		log.Printf("Failed to insert authtoken: %s\n", err.Error())
		helpers.WriteGenericError(w, "general", "Server accounted an error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     core.HttpCookieAuthTokenKey,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	})

	helpers.WriteGenericOk(w)
}
func AuthParametersGet(w http.ResponseWriter, r *http.Request) {
	accountName := r.URL.Query().Get("username")

	if len(accountName) == 0 {
		helpers.WriteGenericError(w, "general", "Missing \"username\" query parameter", http.StatusBadRequest)
		return
	}

	dbAccount, err := database.GetApiAccountByUsername(r.Context(), accountName)

	var parameters account.PasswordParams
	if err == nil {
		parameters = dbAccount.Password.ToParams()
	} else {
		parameters = account.PasswordParams{
			Salt:      core.ComputePasswordSalt(accountName),
			CpuLimit:  core.PasswordHashDefaultCpuLimit,
			RamLimit:  core.PasswordHashDefaultRamLimit,
			Algorithm: core.PasswordHashDefaultAlgorithm,
		}
	}

	helpers.WriteJson(w, parameters, http.StatusOK)
}
