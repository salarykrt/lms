package main

import (
	"salarykart/handlers"
	"github.com/gin-gonic/gin"
	"github.com/nsf/jsondiff"
	sloggin "github.com/samber/slog-gin"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func testLogger() gin.HandlerFunc {
	return sloggin.New(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
}

func readFromJsonFile(loc string) string {
	f, err := os.ReadFile(loc)
	if err != nil {
		panic(err)
	}
	return string(f)
}

func TestInsertLead(t *testing.T) {

	// TODO: probably abstract out this reading from file and checking the response into a separate method
	mockResponse := readFromJsonFile("test_data/insert_lead_response.json")
	conn, err := handlers.InitDB()
	if err != nil {
		return
	}

	r, err := InitRouter(GetRouter(testLogger()), handlers.CreateApp(conn))
	if err != nil {
		panic(err)
		return
	}

	req, _ := http.NewRequest("POST", "/api/v1/lead", strings.NewReader(readFromJsonFile("test_data/insert_lead_request.json")))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	diff, _ := jsondiff.CompareStreams(strings.NewReader(mockResponse), w.Body, &jsondiff.Options{})
	assert.Equal(t, diff, jsondiff.FullMatch)
	if diff != jsondiff.FullMatch {
		t.Log(diff)
		t.Log(w.Body)
	}
	assert.Equal(t, http.StatusOK, w.Code)
}
