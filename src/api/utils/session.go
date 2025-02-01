package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/valkey-io/valkey-go"
)

var Sessions *SessionManager

type SessionManager struct {
	Valkey valkey.Client
	Ctx    context.Context
}

// generateSessionID is a function to generate a session ID.
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// NewSessionManager is a function to create a new session manager.
func NewSessionManager() {
	Sessions = &SessionManager{
		Valkey: InitValkeyClient(),
		Ctx:    context.Background(),
	}
}

func (sm *SessionManager) CreateSession(userID string, metadata map[string]interface{}) (string, error) {
	sessionId, err := generateSessionID()
	if err != nil {
		return "", err
	}
	sessionData := map[string]interface{}{"userId": userID, "metadata": metadata}
	sessionJson, err := json.Marshal(sessionData)
	if err != nil {
		return "", err
	}
	err = sm.Valkey.Do(
		sm.Ctx,
		sm.Valkey.B().Set().Key(
			"session:"+sessionId).
			Value(valkey.BinaryString(sessionJson)).
			Build()).
		Error()
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

// GetSession is a function to get a session.
func (sm *SessionManager) GetSession(sessionId string) (map[string]interface{}, error) {
	sessionJson, err := sm.Valkey.Do(
		sm.Ctx,
		sm.Valkey.B().Get().Key("session:"+sessionId).Build(),
	).AsBytes()
	if err != nil {
		return nil, err
	}
	var sessionData map[string]interface{}
	err = json.Unmarshal(sessionJson, &sessionData)
	if err != nil {
		return nil, err
	}
	return sessionData, nil
}

func (sm *SessionManager) DeleteSession(sessionId string) error {
	err := sm.Valkey.Do(
		sm.Ctx,
		sm.Valkey.B().Del().Key("session:"+sessionId).Build(),
	).Error()
	if err != nil {
		return err
	}
	return nil
}
