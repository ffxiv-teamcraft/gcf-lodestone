package lodestone

import (
	"encoding/json"
	"log"
    "net/http"
    "fmt"

	"github.com/xivapi/godestone/v2"
)

// SearchCharacter searches a character by name and returns a list of results.
func LodestoneCharacterSearch(w http.ResponseWriter, r *http.Request) {
    // Comment this to disable CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Methods", "GET")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }

	name := r.URL.Query().Get("name")
	server := r.URL.Query().Get("server")
	if name == "" || server == "" {
    	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

    opts := godestone.CharacterOptions{
        Name:  name,
        World: server,
    }

    characters := scraper.SearchCharacters(opts)

	cJSON, err := json.Marshal(characters)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(w, string(cJSON))
}
