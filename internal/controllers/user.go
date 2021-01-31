// FILE LOCATION: github.com/bartmika/mulberry-server/internal/controllers/user.go
package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/google/uuid"

    "github.com/bartmika/mulberry-server/pkg/models"
    "github.com/bartmika/mulberry-server/pkg/utils"
)

// To run this API, try running in your console:
// $ http post 127.0.0.1:5000/api/v1/register email="fherbert@dune.com" password="the-spice-must-flow" name="Frank Herbert"
func (h *BaseHandler) postRegister(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Initialize our array which will store all the results from the remote server.
    var requestData models.RegisterRequest

    // Read the JSON string and convert it into our golang stuct else we need
    // to send a `400 Bad Request` errror message back to the client,
    err := json.NewDecoder(r.Body).Decode(&requestData) // [1]
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // For debugging purposes, print our output so you can see the code working.
    fmt.Println(requestData.Name)
    fmt.Println(requestData.Email)
    fmt.Println(requestData.Password)

    // Lookup the email and if it is not unique we need to generate a `400 Bad Request` response.
    if userFound, _ := h.UserRepo.FindByEmail(ctx, requestData.Email); userFound != nil {
        http.Error(w, "Email alread exists", http.StatusBadRequest)
        return
    }

    // Generate a `UUID` for our record.
    uid := uuid.New().String()

    // Secure our password.
    passwordHash, err := utils.HashPassword(requestData.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Save our new user account.
    if err := h.UserRepo.Create(ctx, uid, requestData.Name, requestData.Email, passwordHash); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Generate our response.
    responseData := models.RegisterResponse{
        Message: "You have successfully registered an account.",
        Uuid: uid,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {  // [2]
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// To run this API, try running in your console:
// $ http post 127.0.0.1:5000/api/v1/login email="fherbert@dune.com" password="the-spice-must-flow"
func (h *BaseHandler) postLogin(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    var requestData models.LoginRequest

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // For debugging purposes, print our output so you can see the code working.
    fmt.Println(requestData.Email)
    fmt.Println(requestData.Password)

    // Lookup the user in our database, else return a `400 Bad Request` error.
    user, err := h.UserRepo.FindByEmail(ctx, requestData.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if user == nil {
        http.Error(w, "Email does not exist", http.StatusBadRequest)
        return
    }

    // Verify the inputted password and hashed password match.
    passwordMatch := utils.CheckPasswordHash(requestData.Password, user.PasswordHash)
    if passwordMatch == false {
        http.Error(w, "Incorrect password", http.StatusBadRequest)
        return
    }

    // Finally return success.
    responseData := models.LoginResponse{
        AccessToken: "TODO: WE WILL FIGURE OUT HOW TO DO THIS IN ANOTHER ARTICLE!",
        Uuid: user.Uuid,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// SPECIAL THANKS:
// [1][2]: Learned from:
// a. https://blog.golang.org/json
// b. https://stackoverflow.com/questions/21197239/decoding-json-using-json-unmarshal-vs-json-newdecoder-decode
