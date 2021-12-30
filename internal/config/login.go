package config

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zigapk/prpo-chargers/internal/util/key"
	"io/ioutil"
	"net/http"
)

type login struct {
	SigningPublicKey *rsa.PublicKey
}

// LoadKeys loads keys used for signing jwt tokens.
func (l *login) LoadKeys() error {
	url := fmt.Sprintf("%s/signing_key", Server.AuthServiceUrl())
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse public key.
	l.SigningPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(body)
	if err != nil {
		return err
	}

	key.SetKey(l.SigningPublicKey)

	return nil
}
