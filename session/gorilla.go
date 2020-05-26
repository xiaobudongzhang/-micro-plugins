package session

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

var (
	sessionIdNamePrefix = "session-id-"
	store               *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("OnNUU5RUr6Ii2HMI0d6E54bXTS52tCCL"))
}

func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	var sId string

	for _, c := range r.Cookies() {
		if strings.Index(c.Name, sessionIdNamePrefix) == 0 {
			sId = c.Name
			break
		}
	}

	if sId == "" {
		sId = sessionIdNamePrefix + uuid.New().String()
	}
	ses, _ := store.Get(r, sId)
	if ses.ID == "" {
		cookie := &http.Cookie{Name: sId, Value: sId, Path: "/", Expires: time.Now().Add(30 * time.Second), MaxAge: 0}
		http.SetCookie(w, cookie)

		ses.ID = sId
		ses.Save(r, w)
	}
	return ses
}
