package lodestone

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// GetCharacter fetches a Character by ID.
func GetCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Fatalln("Missing ID parameter")
	}

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Fatalln(err)
	}

	c, err := scraper.FetchCharacter(uint32(uid))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Fatalln(err)
	}

	cJSON, err := json.Marshal(c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Fatalln(err)
	}

	fmt.Fprint(w, string(cJSON))
}
