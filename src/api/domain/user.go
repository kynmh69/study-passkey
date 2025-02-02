package domain

import (
	"encoding/json"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/kynmh69/study-passkey/prisma/db"
)

type User struct {
	ID          string                `json:"id,omitempty"`
	Name        string                `json:"name,omitempty"`
	Email       string                `json:"email,omitempty"`
	Credentials []webauthn.Credential `json:"credentials,omitempty"`
}

func (u *User) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.Email
}

func (u *User) WebAuthnIcon() string {
	return ""
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

type CredentialsList struct {
	Credentials []webauthn.Credential `json:"credentials"`
}

// NewWebAuthnUser is a function that creates a new user for WebAuthn.
func NewWebAuthnUser(dbUser *db.UserModel) (*User, error) {
	var credList CredentialsList
	if err := json.Unmarshal(dbUser.Credentials, &credList); err != nil {
		return nil, err
	}

	return &User{
		ID:          dbUser.ID,
		Name:        dbUser.Username,
		Email:       dbUser.Email,
		Credentials: credList.Credentials,
	}, nil
}
