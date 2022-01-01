package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zigapk/prpo-chargers/internal/logger"
	"github.com/zigapk/prpo-chargers/internal/models/charger"
	"net/http"
	"strconv"
)

// GetChargersHandle  @Summary      Get the list of chargers.
// @Description       Get the list of chargers.
// @Produce           json
// @Param             offset  query     int  false  "Query offset."
// @Param             limit   query     int  false  "Query limit."
// @Success           200     {object}  []charger.Charger
// @Failure           500     {object}  errors.ResponseError
// @Router            /chargers/ [get]
func GetChargersHandle(w http.ResponseWriter, r *http.Request) {
	var offset = 0
	var limit = 3
	if r.URL.Query().Has("offset") {
		fmt.Sscanf(r.URL.Query().Get("offset"), "%d", &offset)
	}

	if r.URL.Query().Has("limit") {
		fmt.Sscanf(r.URL.Query().Get("limit"), "%d", &limit)
	}

	chargers, err := charger.Page(offset, limit)

	if err != nil {
		logger.Log.Warn().Err(err).Send()
	}

	// Write response.
	res, _ := json.Marshal(chargers)
	_, _ = w.Write(res)
}

// GetReservationsHandle  @Summary      Get the list of reservations for a particular charger.
// @Description           Get the list of reservations for a particular charger
// @Produce               json
// @Param                 id      path      string  true   "Id of the charger."
// @Param                 offset  query     int     false  "Query offset."
// @Param                 limit   query     int     false  "Query limit."
// @Success               200     {object}  []reservation.Reservation
// @Failure               500     {object}  errors.ResponseError
// @Router                /chargers/{id}/reservations [get]
func GetReservationsHandle(w http.ResponseWriter, r *http.Request) {
	chargerId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	c, err := charger.ForID(chargerId)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var offset = 0
	var limit = 3
	if r.URL.Query().Has("offset") {
		fmt.Sscanf(r.URL.Query().Get("offset"), "%d", &offset)
	}

	if r.URL.Query().Has("limit") {
		fmt.Sscanf(r.URL.Query().Get("limit"), "%d", &limit)
	}

	reservations, err := c.ReservationsPage(offset, limit)

	if err != nil {
		logger.Log.Warn().Err(err).Send()
	}

	// Write response.
	res, _ := json.Marshal(reservations)
	_, _ = w.Write(res)
}
