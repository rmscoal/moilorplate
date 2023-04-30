// vo is the shorthand for value objects.
package vo

import (
	"encoding/base64"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserCredential struct {
	Username string
	Password string
	Tokens   UserToken
}

type UserToken struct {
	TokenID      string
	AccesssToken string
	RefreshToken string
	Version      int
	Issued       bool
	IssuedAt     time.Time
}

func (v UserCredential) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Username, validation.Required, validation.Length(3, 20)),
		validation.Field(&v.Password, validation.Required),
	)
}

func (u *UserCredential) SetEncodedPasswordFromByte(pass []byte) {
	u.Password = base64.StdEncoding.EncodeToString(pass)
}

func (u UserCredential) GetHashMixture() ([]byte, error) {
	return base64.StdEncoding.DecodeString(u.Password)
}
