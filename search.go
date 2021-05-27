package lodestone

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/xivapi/godestone/v2"
)

func SearchCharacter(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	server := r.URL.Query().Get("server")
	if name == "" || server == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	opts := godestone.CharacterOptions{
		Name:  name,
		World: server,
	}

	var characters []godestone.CharacterSearchResult

	for character := range scraper.SearchCharacters(opts) {
		if character.Error != nil {
			http.Error(w, http.StatusText(StatusInternalServerError), StatusInternalServerError)
			log.Fatalln(character.Error)
		}

		characters = append(characters, *character)
	}

	cJSON, err := json.Marshal(characters)
	if err != nil {
		http.Error(w, http.StatusText(StatusInternalServerError), StatusInternalServerError)
		log.Fatalln(err)
	}

	fmt.Fprint(w, string(cJSON))
}
