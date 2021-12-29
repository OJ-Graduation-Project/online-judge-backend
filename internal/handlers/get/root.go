package get

import(
	"net/http"
	"fmt"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome\n")
}
