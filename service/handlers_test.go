package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mailing-service/cmd/mailing-service/config"
	"mailing-service/service"
	"mailing-service/service/servicefakes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateMailingDetailsHandler(t *testing.T) {
	_, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	srv := service.NewService(&servicefakes.FakeDBInterface{})

	router := service.NewRouter(&config.RouterConfig{Port: "8999"})
	router.RegisterHandlers(srv)

	tests := []struct {
		name           string
		request        func(t *testing.T) *http.Request
		expectedStatus int
	}{
		{
			name: "success",
			request: func(*testing.T) *http.Request {

				body := &service.CreateMailingDetailsRequest{
					Email:      "test@test.com",
					Title:      "title",
					Content:    "content",
					MailingID:  1,
					InsertTime: time.Now().Format(time.RFC3339),
				}
				bodyBytes, err := json.Marshal(body)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}
				req, err := http.NewRequest("POST", "/api/messages", bytes.NewBuffer(bodyBytes))
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				return req
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "malformed reqeust",
			request: func(*testing.T) *http.Request {
				body := &struct {
					MailingID string `json:"mailing_id"`
				}{
					MailingID: "malformed mailing_id",
				}
				bodyBytes, err := json.Marshal(body)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}
				req, err := http.NewRequest("POST", "/api/messages", bytes.NewBuffer(bodyBytes))
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				return req
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		tc := tc
		req := tc.request(t)
		resp := httptest.NewRecorder()
		router.Router.ServeHTTP(resp, req)

		if resp.Result().StatusCode != tc.expectedStatus {
			t.Fatalf("incorrect status code")
		}
	}
}
