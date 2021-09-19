package lodestone

import (
	"encoding/json"
	"log"
    "strconv"
    "net/http"
    "fmt"
)

// LodestoneCharacter fetches a Character by ID.
func LodestoneCharacter(w http.ResponseWriter, r *http.Request) {
    // Comment this to disable CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Methods", "GET")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }

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
