package utils

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/kynmh69/study-passkey/logger"
	"go.uber.org/zap"
)

var WebAuthn *webauthn.WebAuthn

func InitWebAuthn(displayName, rpid string, origins []string) {
	var err error
	WebAuthn, err = webauthn.New(&webauthn.Config{
		RPDisplayName: displayName,
		RPID:          rpid,
		RPOrigins:     origins,
	})
	if err != nil {
		logger.Logger.Error("failed to create WebAuthn", zap.Error(err))
	}
}
