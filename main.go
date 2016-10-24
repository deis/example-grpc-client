package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/kmala/example-grpc-server/_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// main starts an HTTP server listening on $PORT which dispatches to request handlers.
func main() {
	http.Handle("/healthz", http.HandlerFunc(healthcheckHandler))
	// wrap the poweredByHandler with logging middleware
	http.Handle("/", logRequestMiddleware(http.HandlerFunc(poweredByHandler)))
	port := getenv("PORT", "5000")
	log.Printf("listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// poweredByHandler writes a response by querying the grpc server.
func poweredByHandler(w http.ResponseWriter, r *http.Request) {
	powered := getenv("POWERED_BY", "Deis")
	server := getenv("SERVER_NAME", "")
	if server == "" {
		log.Println("Server not configured")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}
	address := fmt.Sprintf("%s.%s.svc.cluster.local:80", server, server)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to the server: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}
	defer conn.Close()
	c := pb.NewPoweredByClient(conn)
	resp, err := c.PoweredBy(context.Background(), &pb.Request{Name: powered})
	if err != nil {
		log.Printf("Error sending the request: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	// print the string to the ResponseWriter
	fmt.Fprintf(w, resp.Message)
}

// healthcheckHandler returns 200 for kubernetes healthchecks.
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}

// logRequestMiddleware writes out HTTP request information before passing to the next handler.
func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remote := r.RemoteAddr
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			remote = forwardedFor
		}
		log.Printf("%s %s %s %s", remote, r.Method, r.URL, r.Header.Get("X-Forwarded-Proto"))
		// pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func getenv(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}
