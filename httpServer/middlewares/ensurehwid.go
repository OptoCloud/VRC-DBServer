package middlewares

import (
	"net/http"
	"vrcdb/core"
	"vrcdb/database"
	"vrcdb/httpServer/helpers"

	guuid "github.com/google/uuid"
)

func ValidateHwid(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		// Get authtoken header
		hardwareId := r.Header.Get(core.HttpHeaderHardwareIdKey)
		if len(hardwareId) == 0 {
			helpers.WriteGenericError(w, "general", "Missing \"Hardware-Id\" header", http.StatusBadRequest)
			return
		}
		// Check and convert authtoken header
		_, err = guuid.Parse(hardwareId)
		if err != nil {
			helpers.WriteGenericError(w, "general", "Invalid HardwareId!", http.StatusBadRequest)
			return
		}

		// Get api config
		apiConfig := database.GetApiConfig()

		// Check request size
		if r.ContentLength > apiConfig.Uploads.SizeMax {
			helpers.WriteGenericError(w, "general", "The uploaded image is too big, upload limit is specified in the API config", http.StatusBadRequest)
			return
		}

		cookie, _ := r.Cookie(core.HttpCookieAuthTokenKey)

		authValues := helpers.ContextAuthValues{
			IpAddress:  r.Header.Get(core.HttpHeaderIpAddressKey),
			IpCountry:  r.Header.Get(core.HttpHeaderIpCountryKey),
			HardwareId: hardwareId,
			Cookie:     cookie,
			Config:     apiConfig,
		}

		contextWithAuth := helpers.InsertContextAuthValues(&authValues, r.Context())

		h.ServeHTTP(w, r.WithContext(contextWithAuth))
	})
}
