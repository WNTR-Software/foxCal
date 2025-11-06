package web

// TODO: This is a sample for rika, remove later

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	webutils "git.mstar.dev/mstar/goutils/http"
	"github.com/rs/zerolog/hlog"
	"gorm.io/gorm"

	"github.com/WNTR-Software/foxcal/backend/storage/dbgen"
	"github.com/WNTR-Software/foxcal/backend/storage/models"
)

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Create or update a key-value pair
//
// @Summary Create or update a key-value pair
// @Description create or update key-value pair
// @ID create-or-update-kv
// @Accept json
// @Produce json
// @Param input body web.KV true "Key value data"
// @Success 200 {string} string "No response body"
// @Failure 500 {object} web.Rfc9457Placeholder "RFC9457 problem details"
// @Router /sampleKv [post]
func handleSampleKvPost(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// As per recent cloudflare report, this can prevent a bug with HTTP/2
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	}()

	log := hlog.FromRequest(r)
	data := KV{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		_ = webutils.ProblemDetailsStatusOnly(w, http.StatusBadRequest)
		log.Warn().Err(err).Msg("Failed to parse request body as json")
		// Equals
		// webutils.ProblemDetails(
		// 	w,
		// 	http.StatusBadRequest, // Status code
		// 	"about:blank", // Error type
		// 	http.StatusText(http.StatusBadRequest), // Error title
		// 	nil, // Error details (string pointer)
		// 	map[string]any{
		// 		"reference": "RFC 9457",
		// 	}, // Extra values
		// )
		return
	}

	kv, err := dbgen.SampleKv.Where(dbgen.SampleKv.Key.Eq(data.Key)).First()
	switch err {
	case nil:
		kv.Value = data.Value
		_, err := dbgen.SampleKv.Where(dbgen.SampleKv.ID.Eq(kv.ID)).
			UpdateSimple(dbgen.SampleKv.Value.Value(data.Value))
		if err != nil {
			log.Error().
				Err(err).
				Str("key", data.Key).
				Str("value", data.Value).
				Msg("Failed to update value in db")
			_ = webutils.ProblemDetailsStatusOnly(w, http.StatusInternalServerError)
		}
	case gorm.ErrRecordNotFound:
		err := dbgen.SampleKv.Create(&models.SampleKv{
			Key:   data.Key,
			Value: data.Value,
		})
		if err != nil {
			log.Error().
				Err(err).
				Str("key", data.Key).
				Str("value", data.Value).
				Msg("Failed to store key-value in db")
			_ = webutils.ProblemDetailsStatusOnly(w, http.StatusInternalServerError)
		}
	default:
		log.Error().Err(err).Str("key", data.Key).Msg("Failed to check if key is already known")
		_ = webutils.ProblemDetailsStatusOnly(w, http.StatusInternalServerError)
	}
}

// Get the value for a key-value pair
//
// @Summary Get the value for a key
// @Description get value of key
// @ID get-kv-value
// @Produce json
// @Param key path string true "Key to get"
// @Success 200 {object} web.KV "The key value pair found"
// @Failure 404 {object} web.Rfc9457Placeholder "No key-value pair found"
// @Failure 500 {object} web.Rfc9457Placeholder "RFC9457 problem details"
// @Router /sampleKv/{id} [get]
func handleSampleKvGet(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	key := r.PathValue("id")
	kv, err := dbgen.SampleKv.Where(dbgen.SampleKv.Key.Eq(key)).First()
	switch err {
	case nil:
		rawData, _ := json.Marshal(&KV{
			Key:   kv.Key,
			Value: kv.Value,
		})
		_, _ = fmt.Fprint(w, string(rawData))
	case gorm.ErrRecordNotFound:
		_ = webutils.ProblemDetailsStatusOnly(w, http.StatusNotFound)
	default:
		log.Error().Err(err).Str("key", key).Msg("Failed to search for key in db")
		_ = webutils.ProblemDetailsStatusOnly(w, http.StatusInternalServerError)
	}
}
