package controllers

import (
    "os"
    "context"
    "strings"
    "log"
    "net/http"

    "github.com/bartmika/mulberry-server/pkg/utils"
)

// Middleware will split the full URL path into slash-sperated parts and save to
// the context to flow downstream in the app for this particular request.
func URLProcessorMiddleware(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Split path into slash-separated parts, for example, path "/foo/bar"
        // gives p==["foo", "bar"] and path "/" gives p==[""]. Our API starts with
        // "/api/v1", as a result we will start the array slice at "3".
        p := strings.Split(r.URL.Path, "/")[3:]

        // log.Println(p) // For debugging purposes only.

        // Open our program's context based on the request and save the
        // slash-seperated array from our URL path.
        ctx := r.Context()
        ctx = context.WithValue(ctx, "url_split", p)

        // Flow to the next middleware.
        fn(w, r.WithContext(ctx))
    }
}

func JWTProcessorMiddleware(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        // Read our application's signing key and attach it to the application
        // context so it can flow downstream in all our applications.
        mySigningKey := []byte(os.Getenv("MULBERRY_APP_SIGNING_KEY"))
        ctx = context.WithValue(ctx, "jwt_signing_key", mySigningKey)

        reqToken := r.Header.Get("Authorization")

        if reqToken != "" {
            // Special thanks to "poise" via https://stackoverflow.com/a/44700761
            splitToken := strings.Split(reqToken, "Bearer ")
            reqToken = splitToken[1]

            // log.Println(reqToken) // For debugging purposes only.

            m, err := utils.ProcessJWT(mySigningKey, reqToken)
            if err == nil {
                ctx = context.WithValue(ctx, "is_authorized", true)
                ctx = context.WithValue(ctx, "session_uuid", m["session_uuid"])
                ctx = context.WithValue(ctx, "client_uuid", m["client_uuid"])

                // Flow to the next middleware with our JWT token saved.
                fn(w, r.WithContext(ctx))
                return
            }
            log.Println("JWTProcessorMiddleware | ProcessJWT | err", err)
        }

        // Flow to the next middleware without anything done.
        ctx = context.WithValue(ctx, "is_authorized", false)
        fn(w, r.WithContext(ctx))
    }
}

func Middleware(fn http.HandlerFunc) http.HandlerFunc {
    // Attach our middleware
    fn = URLProcessorMiddleware(fn)
    fn = JWTProcessorMiddleware(fn)
    return func(w http.ResponseWriter, r *http.Request) {
        // Flow to the next middleware.
        fn(w, r)
    }
}
