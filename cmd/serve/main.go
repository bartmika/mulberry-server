// github.com/bartmika/mulberry-server/cmd/serve/main.go
package main

import (
    "fmt"
    "net/http"

    sqldb "github.com/bartmika/mulberry-server/pkg/db"
    "github.com/bartmika/mulberry-server/internal/repositories"
    "github.com/bartmika/mulberry-server/internal/controllers"
)

func main() {
    db := sqldb.ConnectDB() //TODO: Implement in the future.

    userRepo := repositories.NewUserRepo(db)

    c := controllers.NewBaseHandler(userRepo)

    router := http.NewServeMux()
    router.HandleFunc("/", c.HandleRequests)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
        Handler: router,
	}

    if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        panic(err)
    }
}
