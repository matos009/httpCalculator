package main

import (
	"encoding/json"
	"fmt"
	"httpCalculator/calculate"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}

type BadResponse struct {
	Error string `json:"error"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeBadResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeBadResponse(w, "Not valid expression", http.StatusUnprocessableEntity)
		return
	}

	result, err := calculate.Calc(req.Expression)
	if err != nil {
		writeBadResponse(w, "Something wrong in expression!", http.StatusUnprocessableEntity)
		return
	}

	writeGoodResponse(w, result)

}

func writeBadResponse(w http.ResponseWriter, err string, errCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(BadResponse{Error: err})
}
func writeGoodResponse(w http.ResponseWriter, result float64) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: result})
}

func main() {
	http.HandleFunc("/api/v1/calculate", CalculateHandler)
	fmt.Print("Starting a server on port:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print("Failed to start server:", err)

	}
}
