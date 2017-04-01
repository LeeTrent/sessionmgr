package sessionmgr

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
)

type Session struct {
	SessionId, UserName string
}

type SessionMgr struct {
	sMap map[string]string // Session ID, User Name
}

func NewSessionMgr() *SessionMgr {
	sm := SessionMgr{
		sMap: make(map[string]string),
	}
	return &sm
}

func (sm *SessionMgr) GetUserName(sessionId string) string {
	return sm.sMap[sessionId]
}

func (sm *SessionMgr) GetSession(sessionId string) (Session, error) {
	if userName, ok := sm.sMap[sessionId]; ok {
		return Session{sessionId, userName}, nil
	}
	errMsg := fmt.Sprintf("Session for Session ID '%s' NOT FOUND", sessionId)
	return Session{}, errors.New(errMsg)
}

func (sm *SessionMgr) RemoveSession(sessionId string) {
	delete(sm.sMap, sessionId)
}

func (sm *SessionMgr) CreateSession(userName string) Session {
	sessionId := uuid.NewV4().String()
	sm.sMap[sessionId] = userName
	return Session{sessionId, userName}
}

func (sm *SessionMgr) CreateSessionId() string {
	return uuid.NewV4().String()
}
