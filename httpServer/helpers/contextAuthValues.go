package helpers

import (
	"context"
	"vrcdb/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type contextKey int

const ContextAuthValuesKey contextKey = 69420

type ContextAuthValues struct {
	IpAddress string
	IpCountry string
	Config    models.ApiConfig
	ClientKey primitive.ObjectID
}

func GetContextAuthValues(ctx context.Context) *ContextAuthValues {

	return ctx.Value(ContextAuthValuesKey).(*ContextAuthValues)
}
func InsertContextAuthValues(authValues *ContextAuthValues, ctx context.Context) context.Context {

	//create a new request context containing the authenticated account
	return context.WithValue(ctx, ContextAuthValuesKey, authValues)
}
