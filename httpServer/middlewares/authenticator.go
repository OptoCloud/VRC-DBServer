package middlewares

import (
	"net/http"
	"vrcdb/core"
	"vrcdb/database"
	"vrcdb/httpServer/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authenticator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		requestContext := r.Context()

		// Get clientkey header
		clientKeyStr := r.Header.Get(core.HttpClientKeyKey)
		if len(clientKeyStr) == 0 {
			helpers.WriteGenericError(w, "general", "Missing \"Client-Key\" header", http.StatusBadRequest)
			return
		}

		clientKey, err := primitive.ObjectIDFromHex(clientKeyStr)
		if err != nil {
			helpers.WriteGenericError(w, "general", "Invalid client key format", http.StatusBadRequest)
			return
		}

		// Get account
		result := database.CollectionUploaders.FindOne(requestContext, bson.M{"_id": clientKey})
		err = result.Err()
		if err != nil {
			helpers.WriteGenericError(w, "general", "Invalid client key", http.StatusInternalServerError)
			return
		}

		// Get api config
		apiConfig := database.GetApiConfig()

		authValues := helpers.ContextAuthValues{
			IpAddress: r.Header.Get(core.HttpHeaderIpAddressKey),
			IpCountry: r.Header.Get(core.HttpHeaderIpCountryKey),
			Config:    apiConfig,
			ClientKey: clientKey,
		}

		contextWithAuth := helpers.InsertContextAuthValues(&authValues, requestContext)

		h.ServeHTTP(w, r.WithContext(contextWithAuth))
	})
}
