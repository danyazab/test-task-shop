package authenticator

import "net/http"

type Credential struct {
	Name     string `env:"USER_ADMIN"`
	Password string `env:"PASSWORD_ADMIN"`
}

type Authenticator struct {
	admins []Credential
}

func NewAuthenticator(admins ...Credential) *Authenticator {
	return &Authenticator{
		admins: admins,
	}
}

func (a Authenticator) acesAllowed(user, pass string) bool {
	for _, u := range a.admins {
		if u.Password == pass && u.Name == user {
			return true
		}
	}
	return false
}

func (a Authenticator) BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	user, pass, ok := r.BasicAuth()
	if !ok || !a.acesAllowed(user, pass) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized.\n"))
		return false
	}
	return true
}
