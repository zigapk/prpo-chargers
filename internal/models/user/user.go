package user

import (
	"github.com/zigapk/prpo-chargers/internal/util/key"
)

type User struct {
	UID string
}

func FromToken(accessToken string) (*User, error) {
	claims, err := key.Validate(accessToken)
	if err != nil {
		return nil, err
	}

	u := &User{
		UID: claims.Subject,
	}
	return u, nil
}
