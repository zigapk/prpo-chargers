package charger

import (
	"time"

	"github.com/zigapk/prpo-chargers/internal/database"
	"github.com/zigapk/prpo-chargers/internal/logger"
)

type Charger struct {
	ID int `db:"id"`

	Name string `db:"name"`

	Lat float64 `db:"lat"`
	Lon float64 `db:"lon"`

	DateCreated time.Time `db:"date_created"`
}

// New charger with a given username and password.
func New(name string, lat float64, lon float64) (*Charger, error) {
	// Insert charger to database.
	c := &Charger{}
	insert := `INSERT INTO chargers (name, lat, lon) VALUES ($1, $2, $3) RETURNING *`
	err := database.DB.Get(c, insert, name, lat, lon)

	// Handle errors.
	if err != nil {
		return nil, err
	}

	return c, nil
}

// ForID queries charger for specified id.
func ForID(id int) (*Charger, error) {
	c := &Charger{}

	query := `SELECT * FROM chargers WHERE id=$1`
	err := database.DB.Get(c, query, id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Charger) Update(name string, lat float64, lon float64) error {
	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}

	// Set new name.
	update := `UPDATE chargers SET name=$2, lat=$3, lon=$4 WHERE id=$1 RETURNING *`
	err = tx.Get(c, update, c.ID, name, lat, lon)
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
