package reservation

import (
	"time"

	"github.com/zigapk/prpo-chargers/internal/database"
	"github.com/zigapk/prpo-chargers/internal/logger"
)

type Reservation struct {
	ID int `json:"id" db:"id"`

	UserId    string `json:"user_id" db:"user_id"`
	ChargerId string `json:"charger_id" db:"charger_id"`

	TimeFrom time.Time `json:"time_from" db:"time_from"`
	TimeUtil time.Time `json:"time_until" db:"time_until"`
}

// New charger with a given username and password.
func New(userUid string, chargerId int, timeFrom time.Time, timeUntil time.Time) (*Reservation, error) {
	// Insert charger to database.
	r := &Reservation{}
	insert := `INSERT INTO reservations (user_id, charger_id, time_from, time_until) VALUES ($1, $2, $3, $4) RETURNING *`
	err := database.DB.Get(r, insert, userUid, chargerId, timeFrom, timeUntil)

	// Handle errors.
	if err != nil {
		return nil, err
	}

	return r, nil
}

// ForID queries charger for specified id.
func ForID(id int) (*Reservation, error) {
	r := &Reservation{}

	query := `SELECT * FROM reservations WHERE id=$1`
	err := database.DB.Get(r, query, id)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Reservation) Update(timeFrom time.Time, timeUntil time.Time) error {
	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}

	// Set new name.
	update := `UPDATE reservations SET time_from=$2, time_until=$3  WHERE id=$1 RETURNING *`
	err = tx.Get(r, update, r.ID, timeFrom, timeUntil)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	// Commit.
	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	return nil
}

func (r *Reservation) Delete() error {
	del := "DELETE FROM reservations WHERE ID=$1"
	_, err := database.DB.Exec(del, r.ID)
	return err
}

func CollidingReservations(chargerId int, timeFrom time.Time, timeUntil time.Time) ([]*Reservation, error) {
	var reservations []*Reservation

	query := `SELECT * FROM reservations WHERE charger_id = $3 AND NOT (($1 >= time_until and $2 >= time_until) OR ($1 <= time_from and $2 <= time_from))`
	err := database.DB.Select(&reservations, query, timeFrom, timeUntil, chargerId)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}
