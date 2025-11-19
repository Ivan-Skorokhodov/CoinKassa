package cookie

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, cookieValue string) {
	expiration := time.Now().AddDate(1, 0, 0)
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    cookieValue,
		Expires:  expiration,
		HttpOnly: true,
		//SameSite: http.SameSiteLaxMode,
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
	}
	http.SetCookie(w, &cookie)
}
