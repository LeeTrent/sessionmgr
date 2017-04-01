package sessionmgr

import (
	"github.com/satori/go.uuid"
)

type Session struct {
	SessionId, UserName string
}

type SessionMgr struct {
	sMap map[string]Session // Session ID, User Name
}

func NewSessionMgr() *SessionMgr {
	sm := SessionMgr{
		sMap: make(map[string]Session),
	}
	return &sm
}

func (sm *SessionMgr) GetUserName(sessionId string) string {
	return sm.sMap[sessionId].UserName
}

func (sm *SessionMgr) GetSession(sessionId string) (Session, bool) {
	session, found := sm.sMap[sessionId]
	return session, found
}

func (sm *SessionMgr) RemoveSession(sessionId string) {
	delete(sm.sMap, sessionId)
}

func (sm *SessionMgr) CreateSession(userName string) Session {
	sessionId := uuid.NewV4().String()
	sm.sMap[sessionId] = Session{sessionId, userName}
	return sm.sMap[sessionId]
}

func (sm *SessionMgr) CreateSessionId() string {
	return uuid.NewV4().String()
}