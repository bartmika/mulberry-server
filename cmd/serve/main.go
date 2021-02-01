// github.com/bartmika/mulberry-server/cmd/serve/main.go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
	"syscall"
	"time"

    sqldb "github.com/bartmika/mulberry-server/pkg/db"
    "github.com/bartmika/mulberry-server/internal/repositories"
    "github.com/bartmika/mulberry-server/internal/controllers"
)

func main() {
    // Get our environment variables which will used to configure our application.
    databaseHost := os.Getenv("MULBERRY_DB_HOST")
    databasePort := os.Getenv("MULBERRY_DB_PORT")
    databaseUser := os.Getenv("MULBERRY_DB_USER")
    databasePassword := os.Getenv("MULBERRY_DB_PASSWORD")
    databaseName := os.Getenv("MULBERRY_DB_NAME")

    db, err := sqldb.ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    userRepo := repositories.NewUserRepo(db)
    tsdRepo := repositories.NewTimeSeriesDatumRepo(db)

    c := controllers.NewBaseHandler(userRepo, tsdRepo)

    router := http.NewServeMux()
    router.HandleFunc("/", controllers.ChainMiddleware(c.HandleRequests))

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
        Handler: router,
	}

    done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    go runMainRuntimeLoop(srv)

	log.Print("Server Started")

	// Run the main loop blocking code.
	<-done

    stopMainRuntimeLoop(srv)
}

func runMainRuntimeLoop(srv *http.Server) {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("listen: %s\n", err)
    }
}

func stopMainRuntimeLoop(srv *http.Server) {
    log.Printf("Starting graceful shutdown now...")

    // Execute the graceful shutdown sub-routine which will terminate any
	// active connections and reject any new connections.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
    log.Printf("Graceful shutdown finished.")
    log.Print("Server Exited")
}
