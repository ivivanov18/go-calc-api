package add

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	var operation AddDto

	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := operation.Number1 + operation.Number2

	fmt.Fprintf(w, "%d", sum)
}
