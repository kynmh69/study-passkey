package domain

import "github.com/go-webauthn/webauthn/webauthn"

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
