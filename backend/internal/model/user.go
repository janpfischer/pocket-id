package model

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"time"
)

type User struct {
	Base

	Username  string `json:"username"`
	Email     string `json:"email" `
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAdmin   bool   `json:"isAdmin"`

	Credentials []WebauthnCredential `json:"-"`
}

func (u User) WebAuthnID() []byte { return []byte(u.ID) }

func (u User) WebAuthnName() string { return u.Username }

func (u User) WebAuthnDisplayName() string { return u.FirstName + " " + u.LastName }

func (u User) WebAuthnIcon() string { return "" }

func (u User) WebAuthnCredentials() []webauthn.Credential {
	credentials := make([]webauthn.Credential, len(u.Credentials))

	for i, credential := range u.Credentials {
		credentials[i] = webauthn.Credential{
			ID:              []byte(credential.CredentialID),
			AttestationType: credential.AttestationType,
			PublicKey:       credential.PublicKey,
			Transport:       credential.Transport,
		}

	}
	return credentials
}

func (u User) WebAuthnCredentialDescriptors() (descriptors []protocol.CredentialDescriptor) {
	credentials := u.WebAuthnCredentials()

	descriptors = make([]protocol.CredentialDescriptor, len(credentials))

	for i, credential := range credentials {
		descriptors[i] = credential.Descriptor()
	}

	return descriptors
}

type OneTimeAccessToken struct {
	Base
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`

	UserID string `json:"userId"`
	User   User
}

type OneTimeAccessTokenCreateDto struct {
	UserID    string    `json:"userId" binding:"required"`
	ExpiresAt time.Time `json:"expiresAt" binding:"required"`
}

type LoginUserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
