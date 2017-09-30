package routing

import (
	"net/http"
)

// Source for routing paths
func LoadRoutes() {
	http.HandleFunc("/", index)
}
