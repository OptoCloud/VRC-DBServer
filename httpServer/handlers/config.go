package handlers

import (
	"net/http"
	"vrcdb/httpServer/helpers"
)

func ConfigGet(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJson(w, helpers.GetContextAuthValues(r.Context()).Config, http.StatusOK)
}
