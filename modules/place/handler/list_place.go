package placehandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) ListAllPlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter iomodel.Filter

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		lat := ctx.Query("lat")
		lng := ctx.Query("lng")
		if lat != "" && lng != "" {
			lat, _ := strconv.ParseFloat(lat, 64)
			lng, _ := strconv.ParseFloat(lng, 64)
			filter.Lat = &lat
			filter.Lng = &lng
		}

		guest := ctx.Query("guest")
		if guest != "" {
			guest, _ := strconv.Atoi(guest)
			filter.Guest = &guest
		}

		numBed := ctx.Query("num_bed")
		if numBed != "" {
			numBed, _ := strconv.Atoi(numBed)
			filter.NumBed = &numBed
		}

		bedRoom := ctx.Query("bed_room")
		if bedRoom != "" {
			bedRoom, _ := strconv.Atoi(bedRoom)
			filter.Bedroom = &bedRoom
		}

		priceFrom := ctx.Query("price_from")
		if priceFrom != "" {
			priceFrom, _ := strconv.Atoi(priceFrom)
			filter.PriceFrom = &priceFrom
		}

		priceTo := ctx.Query("price_to")
		if priceTo != "" {
			priceTo, _ := strconv.Atoi(priceTo)
			filter.PriceTo = &priceTo
		}

		dateFrom := ctx.Query("date_from")
		if dateFrom != "" {
			filter.DateFrom = &dateFrom
		}
		dateTo := ctx.Query("date_to")
		if dateTo != "" {
			filter.DateTo = &dateTo
		}

		var listPlaceReq iomodel.ListPlaceReq
		if err := ctx.ShouldBind(&listPlaceReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		places, err := hdl.placeUC.ListAllPlace(ctx.Request.Context(), &paging, &filter, listPlaceReq.UserEmail)
		if err != nil {
			panic(err)
		}

		res := convert.ConvertPlaceToListModel(places, &paging)
		ctx.JSON(http.StatusOK, res)
	}
}
