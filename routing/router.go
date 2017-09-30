package routing

import (
	"fmt"
	"net/http"
)

// Source for handlers/middlewares
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "INDEX PAGE!") // send data to client side
}
