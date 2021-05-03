package handlers

import (
	"net/http"

	"github.com/amirrezaask/mokhtasar/pkg"
	"go.uber.org/zap"
)

type HTTPHandler struct {
	Mokhtasar *pkg.Mokhtasar
	Logger    *zap.SugaredLogger
}

func (h *HTTPHandler) Long(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only GET request supported"))
		h.Logger.Errorf("method is not supported: %s", r.Method)
		return
	}
	key := r.URL.Query().Get("url")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need a url to short"))
		h.Logger.Errorf("key is not provided")
		return
	}
	url, err := h.Mokhtasar.GetOriginalURL(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in adding new url"))
		h.Logger.Errorf("error in expanding given key: %v", err)
		return
	}
	w.Write([]byte(url))
}

func (h *HTTPHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only GET request supported"))
		h.Logger.Errorf("method is not supported: %s", r.Method)
		return
	}

	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need a url to short"))
		h.Logger.Errorf("no url provided")
		return
	}
	url, err := h.Mokhtasar.Shorten(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in finding url you gave us"))
		h.Logger.Errorf("error in shortening given url: %v", err)
		return
	}
	w.Write([]byte(url))
}
