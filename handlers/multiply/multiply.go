package multiply

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var operation MultiplyDto

	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := operation.Number1 * operation.Number2

	fmt.Fprintf(w, "%d", product)
}
