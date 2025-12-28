package controller

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func (c *Controller) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := c.GoogleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (c *Controller) GoogleAuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := c.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := c.GoogleOauthConfig.Client(context.Background(), token)

	// Get user info
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	// io.Copy(w, resp.Body)
	http.Redirect(w, r, "http://localhost:8000", http.StatusTemporaryRedirect)
}
