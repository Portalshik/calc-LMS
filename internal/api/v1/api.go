package v1

import (
	"calc-lms/internal/calculator"
	"encoding/json"
	"fmt"
	"net/http"
)


type Request struct {
	Expression string `json:"expression"`
}

// Ответ с результатом
type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}


func Calculate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Метод запроса:", r.Method)
	fmt.Println("Заголовки:", r.Header)
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Expression == "" {
		http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		return
	}

	response := Response{}
	result, err := calculator.Calc(req.Expression)
	fmt.Fprint(w, result, err)

	if err != nil {
		response.Error = fmt.Sprintf("%v", err)
	} else {
		response.Result = result
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}