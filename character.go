package lodestone

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Character(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	c, err := scraper.FetchCharacter(uint32(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Fatalln(err)
	}

	cJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(w, string(cJSON))
}
