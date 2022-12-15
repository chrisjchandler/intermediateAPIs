package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

type Command struct {
    Command string `json:"command"`
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
    // Parse the command from the request body
    var command Command
    err := json.NewDecoder(r.Body).Decode(&command)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Pass the command to the other API
    resp, err := http.Post("https://other-api.com/api/function", "application/json", r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Return the result of the other API's function to the original requestor
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/command", handleCommand).Methods("POST")

    http.ListenAndServe(":5000", router)
}
//This API uses the gorilla/mux package to handle HTTP requests, and the net/http package to make HTTP requests to the other API. It defines a Command struct to parse the
// command from the request body, and a handleCommand function to handle POST requests to the /api/command endpoint.
//In the handleCommand function, the API extracts the command from the request body and passes it to the other API using the http.Post function.
// It then returns the result of the other API's function to the original requestor.

//To make a curl request against this API, you can use the following command:

//curl -X POST http://localhost:5000/api/command -d '{"command": "YOUR_COMMAND"}'

//This will send a POST request to the /api/command endpoint with the specified command in the request body. The API will pass this command to the other API, and return the result to the original requestor.
// You can then use the -o or --output flag to save the result to a file, like this:

//curl -X POST http://localhost:5000/api/command -d '{"command": "YOUR_COMMAND"}' -o output.txt


//This will save the result of the other API's function to a file named output.txt.
