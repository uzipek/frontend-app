package frontend_app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{}  `json:"data"`
}

func GetJSONResponse(code int, message string, data interface{}) string {
	jsonResponse := JSONResponse{Code: code, Message: message, Data: data}
	jsonBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func GetErrorJSONResponse(code int, message string) string {
	return GetJSONResponse(code, message, map[string]string{})
}

func GetSuccessJSONResponse(data interface{}) string {
	return GetJSONResponse(http.StatusOK, "Success", data)
}

func GetErrorStatusCode(code int) int {
	switch code {
	case http.StatusBadRequest:
		return http.StatusBadRequest
	case http.StatusInternalServerError:
		return http.StatusInternalServerError
	case http.StatusUnauthorized:
		return http.StatusUnauthorized
	case http.StatusNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func GetRandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBit = 6
		letterIdxMax = 1 << letterIdxBit
	)

	var src = rand.NewSource(time.Now().UnixNano())
	var letters = strings.Repeat(letterBytes, letterIdxMax)
	var b []byte

	for i := 0; i < length; i++ {
		b = append(b, letters[src.Int63()&letterIdxMax])
	}

	return string(b)
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}