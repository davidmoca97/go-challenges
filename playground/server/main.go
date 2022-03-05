package main

import "net/http"

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "<<token>>",
		MaxAge: 120,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/cookie", setCookie)
	http.ListenAndServe(":7777", router)
}
