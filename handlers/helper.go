package handlers

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	genModels "github.com/wasimkhan042/ustore-auth/gen/models"
	"github.com/wasimkhan042/ustore-auth/models"
)

func toItemGen(item *models.Subscription) *genModels.UserProduct {
	itm := models.Subscription{}
	days := itm.EndTime.Sub(itm.StartTime).Hours() / 24
	return &genModels.UserProduct{
		ItemName:            item.ItemName,
		SubsPrice:           item.SubscriptionPrice,
		ExpiryDate:          strfmt.DateTime(item.EndTime),
		RemainingTimeInDays: fmt.Sprintf("%v", int(days)),
	}
}
