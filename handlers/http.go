package handlers

import (
	"net/http"

	"github.com/amirrezaask/mokhtasar/pkg"
)

type HTTPHandler struct {
	Mokhtasar *pkg.Mokhtasar
}

func (h *HTTPHandler) Long(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only GET request supported"))
		return
	}
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need a url to short"))
		return
	}
	key, err := h.Mokhtasar.GetOriginalURL(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in adding new url"))
		return
	}
	toClickURL := "localhost:8080/long?key=" + key
	w.Write([]byte(toClickURL))
}

func (h *HTTPHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need a url to short"))
		return
	}
	url, err := h.Mokhtasar.Shorten(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in finding key you gave us"))
		return
	}
	w.Write([]byte(url))
}
