package middlewares

import (
	"net/http"
	"vrcdb/database"
	"vrcdb/httpServer/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authenticator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var authTokenId primitive.ObjectID

		requestContext := r.Context()
		authValues := helpers.GetContextAuthValues(requestContext)

		// Get authtoken
		if authValues.Cookie == nil {
			helpers.WriteGenericError(w, "general", "Missing \"AuthToken\" cookie", http.StatusBadRequest)
			return
		}

		authTokenId, err = primitive.ObjectIDFromHex(authValues.Cookie.Value)
		if err != nil {
			helpers.WriteGenericError(w, "general", "Invalid AuthToken format", http.StatusBadRequest)
			return
		}

		// Get account
		authValues.Account, err = database.GetApiAccountByAuthtoken(requestContext, authTokenId)
		if err != nil {
			helpers.WriteGenericError(w, "general", "No account connected to AuthToken", http.StatusInternalServerError)
			return
		}

		h.ServeHTTP(w, r)
	})
}
