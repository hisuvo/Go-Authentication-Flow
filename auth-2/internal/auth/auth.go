package auth

import (
	"net/http"
	"os"
	"time"

	"gopkg.in/boj/redistore.v1"
)

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	var store, _ = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("SESSION_SECRET")))
	var users = map[string]string{"naren": "passme", "admin": "password"}

	func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "session.id")

		if (session.Values["authenticated"] != nil) && session.Values["authenticated"] != false {
			w.Write([]byte(time.Now().String()))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	}

	func LoginHandler(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session.id")

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}

		username := r.PostForm.Get("username");
		password := r.PostForm.Get("password")

		if OrginalPassowrd, ok := users[username]; ok {
			if password == OrginalPassowrd {
				session.Values["authenticated"] = true;
				session.Save(r, w)
			} else {
				http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			}

		} else {
			http.Error(w, "User is not found", http.StatusNotFound)
			return
		}

		w.Write([]byte("Login successfull"))
	}

	func LogoutHandler(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session.id")

		session.Values["authenticated"] = false
		session.Save(r, w)
		w.Write([]byte("LogOut successfully"))

	}

