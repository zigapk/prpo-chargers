package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/zigapk/prpo-chargers/internal/logger"
	"github.com/zigapk/prpo-chargers/internal/models/charger"
)

const (
	flagName = "name"
	flagLat  = "lat"
	flagLon  = "lon"
)

var CreateCharger = &cli.Command{
	Name:  "createcharger",
	Usage: "Create new charger.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     flagName,
			Usage:    "Name for new charger.",
			Required: true,
		},
		&cli.Float64Flag{
			Name:     flagLat,
			Usage:    "Lat of the new charger.",
			Required: true,
		},

		&cli.Float64Flag{
			Name:     flagLon,
			Usage:    "Lon of the new charger.",
			Required: true,
		},
	},
	Action: createCharger,
}

func createCharger(c *cli.Context) error {
	ch, err := charger.New(c.String(flagName), c.Float64(flagLat), c.Float64(flagLon))
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}

	fmt.Printf("Successfully create new charger with ID \"%d\".\n", ch.ID)

	return nil
}
