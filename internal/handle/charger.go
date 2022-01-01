package handle

import (
	"encoding/json"
	"fmt"
	"github.com/zigapk/prpo-chargers/internal/logger"
	"github.com/zigapk/prpo-chargers/internal/models/charger"
	"net/http"
)

type newReservationRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// GetChargersHandle  @Summary      Get the list of chargers.
// @Description    Get the list of chargers.
// @Produce       json
// @Success        200  {object}  []Charger
// @Failure       500       {object}  errors.ResponseError
// @Router         /chargers/ [get]
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
