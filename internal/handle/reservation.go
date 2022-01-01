package handle

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zigapk/prpo-chargers/internal/handle/errors"
	"github.com/zigapk/prpo-chargers/internal/logger"
	middleware "github.com/zigapk/prpo-chargers/internal/middleware"
	"github.com/zigapk/prpo-chargers/internal/models/reservation"
	charger "github.com/zigapk/prpo-chargers/internal/models/reservation"
	"github.com/zigapk/prpo-chargers/internal/util"
	"net/http"
	"strconv"
	"time"
)

type newReservationRequest struct {
	ChargerId int       `json:"charger_id"`
	TimeFrom  time.Time `json:"time_from"`
	TimeUntil time.Time `json:"time_until"`
}

// DeleteReservationRequest  @Summary      Delete a single reservation.
// @Description              Delete a single reservation
// @Param                    id  path  string  true  "Id of the reservation."
// @Success                  204
// @Failure                  500  {object}  errors.ResponseError
// @Router                   /reservations/{id} [delete]
func DeleteReservationRequest(w http.ResponseWriter, r *http.Request) {
	user := middleware.UserFromRequest(r)

	reservationId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	existingReservation, err := reservation.ForID(reservationId)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if user.UID != existingReservation.UserId {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	err = existingReservation.Delete()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateReservationRequest  @Summary      Update a single reservation
// @Description              Update a single reservation
// @Produce                  json
// @Param                    charger_id  body      int     true  "Charger ID."
// @Param                    time_from   body      string  true  "Start time."
// @Param                    time_until  body      string  true  "End time."
// @Success                  200         {object}  reservation.Reservation
// @Failure                  500         {object}  errors.ResponseError
// @Router                   /reservations [post]
func CreateReservationRequest(w http.ResponseWriter, r *http.Request) {
	user := middleware.UserFromRequest(r)

	data := &newReservationRequest{}
	ok := util.ParseJSON(w, r, data)
	if !ok {
		return
	}

	// Check whether there is space for reservation.
	ress, err := charger.CollidingReservations(data.ChargerId, data.TimeFrom, data.TimeUntil)

	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if len(ress) > 0 {
		errors.Response(w, errors.CollidingReservation)
		return
	}

	newReservation, err := reservation.New(user.UID, data.ChargerId, data.TimeFrom, data.TimeUntil)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}

	// Write response.
	res, _ := json.Marshal(newReservation)
	_, _ = w.Write(res)

}

// UpdateReservationRequest  @Summary      Update a single reservation
// @Description              Update a single reservation
// @Produce                  json
// @Param                    id          path      string  true  "Id of the reservation."
// @Param                    charger_id  body      int     true  "Charger ID."
// @Param                    time_from   body      string  true  "Start time."
// @Param                    time_until  body      string  true  "End time."
// @Success                  200         {object}  reservation.Reservation
// @Failure                  500         {object}  errors.ResponseError
// @Router                   /reservations/{id} [put]
func UpdateReservationRequest(w http.ResponseWriter, r *http.Request) {
	user := middleware.UserFromRequest(r)

	reservationId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	existingReservation, err := reservation.ForID(reservationId)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if user.UID != existingReservation.UserId {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	data := &newReservationRequest{}
	ok := util.ParseJSON(w, r, data)
	if !ok {
		return
	}

	// Check whether there is space for reservation.
	ress, err := charger.CollidingReservations(data.ChargerId, data.TimeFrom, data.TimeUntil)

	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if len(ress) > 1 || (len(ress) > 0 && ress[0].ID != existingReservation.ID) {
		errors.Response(w, errors.CollidingReservation)
		return
	}

	err = existingReservation.Update(data.TimeFrom, data.TimeUntil)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}

	// Write response.
	res, _ := json.Marshal(existingReservation)
	_, _ = w.Write(res)

}
