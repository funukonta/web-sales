package pkg

import (
	"encoding/json"
	"net/http"
)

func GetJsonBody(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	defer r.Body.Close()
	return err
}

type JsonBod struct {
	code    int `json:"code"`
	Data    any `json:"data,omitempty"`
	Message any `json:"pesan,omitempty"`
}

func Response(code int, data *JsonBod) *JsonBod {
	res := new(JsonBod)

	res.code = code
	if data.Data != nil {
		res.Data = data.Data
	}

	if data.Message != nil {
		res.Message = data.Message
	}

	return res
}

func (r *JsonBod) Send(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.code)

	json.NewEncoder(w).Encode(r)
}
