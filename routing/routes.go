package routing

import (
	"net/http"
)

// Source for routing paths
func LoadRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user/create", createUser)
}
