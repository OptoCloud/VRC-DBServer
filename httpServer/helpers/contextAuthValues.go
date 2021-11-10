package helpers

import (
	"context"
	"net/http"
	"vrcdb/models"
	"vrcdb/models/account"
)

type contextKey int

const ContextAuthValuesKey contextKey = 69420

type ContextAuthValues struct {
	IpAddress  string
	IpCountry  string
	HardwareId string
	Cookie     *http.Cookie
	Config     models.ApiConfig
	Account    account.DbAccount
}

func GetContextAuthValues(ctx context.Context) *ContextAuthValues {

	return ctx.Value(ContextAuthValuesKey).(*ContextAuthValues)
}
func InsertContextAuthValues(authValues *ContextAuthValues, ctx context.Context) context.Context {

	//create a new request context containing the authenticated account
	return context.WithValue(ctx, ContextAuthValuesKey, authValues)
}
