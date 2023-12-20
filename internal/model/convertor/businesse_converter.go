package converter

import (
	"be/internal/entity"
	"be/internal/model"
)

func BusinesseToResponse(businesse *entity.Businesses) *model.BusinesseResponse {
	return &model.BusinesseResponse{
		ID:                  businesse.ID,
		Location:            businesse.Location,
		Latitude:            businesse.Latitude,
		Longitude:           businesse.Longitude,
		Term:                businesse.Term,
		Radius:              businesse.Radius,
		Categories:          businesse.Categories,
		Locale:              businesse.Locale,
		Price:               businesse.Price,
		OpenNow:             businesse.OpenNow,
		OpenAt:              businesse.OpenAt,
		Attributes:          businesse.Attributes,
		SortBy:              businesse.SortBy,
		DevicePlatform:      businesse.DevicePlatform,
		ReservationDate:     businesse.ReservationDate,
		ReservationTime:     businesse.ReservationTime,
		ReservationCover:    businesse.ReservationCover,
		MatchPartySizeParam: businesse.MatchPartySizeParam,
		Limit:               businesse.Limit,
		Offset:              businesse.Offset,
	}
}
