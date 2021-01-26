// github.com/bartmika/cmd/serve.go
package main

import (
	"net/http"

	"github.com/bartmika/internal/manager"
)

func main() {
	manager := manager.NewManager()

    // Integrate this program with the signal handler provided by the operating
	// system through the go routine.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // Create a new go routine which will become our main runtime loop because
	// we will halt this current main runtime loop with a go channel.
	go func() {
		if execErr := manager.RunMainRuntimeLoop(); execErr != nil {
			log.Fatalf("main error: %s\n", execErr)
		}
	}()

	// Halt the main runtime loop until the program receives a shutdown signal.
	<-done

    // Execute the graceful shutdown sub-routine which will terminate any
	// active connections and reject any new connections.
	manager.StopMainRuntimeLoop()
}
