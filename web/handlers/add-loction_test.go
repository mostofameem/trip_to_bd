package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"post-service/config"
	"post-service/location"
	"post-service/mongodb"
	"post-service/web/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLocationService is a mock for the LocationService
type MockLocationService struct {
	mock.Mock
}

// MockLocationService is a mock for the LocationService
type MockHandlers struct {
	mock.Mock
}

func (m *MockLocationService) AddLocation(ctx context.Context, loc *location.Location) error {
	args := m.Called(ctx, loc)
	return args.Error(0)
}

func (m *MockLocationService) GetLocation(ctx context.Context, title string) (*mongodb.Location, error) {
	args := m.Called(ctx)
	locations := args.Get(0).(*mongodb.Location) // Cast the first returned value to *[]mongodb.Location
	err := args.Error(1)                         // Get the second returned value as an error
	return locations, err
}

func TestAddLocation_TableDriven(t *testing.T) {
	mockLocSvc := new(MockLocationService)

	tests := []struct {
		name           string
		input          handlers.LocationReq
		mockErr        error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			input: handlers.LocationReq{
				Title:        "Beautiful Park",
				Descriptions: "A serene park in the city center",
				BestTime:     "Morning",
				PictureUrl:   "http://example.com/park.jpg",
			},
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":"Location addedd Successfully","message":"Success","status":true}`,
		},

		{
			name: "Invalid Input (missing title)",
			input: handlers.LocationReq{
				Title:        "", // Invalid, required field
				Descriptions: "A serene park in the city center",
				BestTime:     "Morning",
				PictureUrl:   "http://example.com/park.jpg",
			},
			mockErr:        errors.New("bad request"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"invaild request body","status":false}`, // Corrected JSON format
		},
		{
			name: "Service Error",
			input: handlers.LocationReq{
				Title:        "Beautiful Park",
				Descriptions: "A serene park in the city center",
				BestTime:     "Morning",
				PictureUrl:   "http://example.com/park.jpg",
			},
			mockErr:        errors.New("service error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLocSvc.ExpectedCalls = nil
			body, _ := json.Marshal(tt.input)

			req, err := http.NewRequest("POST", "/location", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			h := handlers.NewHandlers(&config.Config{}, mockLocSvc)

			rr := httptest.NewRecorder()

			if tt.mockErr == nil {
				mockLocSvc.On("AddLocation", mock.Anything, &location.Location{
					Title:        tt.input.Title,
					Descriptions: tt.input.Descriptions,
					BestTime:     tt.input.BestTime,
					PictureUrl:   tt.input.PictureUrl,
				}).Return(nil).Once()
			} else {
				mockLocSvc.On("AddLocation", mock.Anything, &location.Location{
					Title:        tt.input.Title,
					Descriptions: tt.input.Descriptions,
					BestTime:     tt.input.BestTime,
					PictureUrl:   tt.input.PictureUrl,
				}).Return(tt.mockErr).Once()
			}

			h.AddLocation(rr, req)

			// Check the response status code
			assert.Equal(t, tt.expectedStatus, rr.Code)

			// Check the response body
			assert.Contains(t, rr.Body.String(), tt.expectedBody)

			// Verify that AddLocation was called with the correct parameters
			if tt.mockErr == nil {
				mockLocSvc.AssertCalled(t, "AddLocation", mock.Anything, &location.Location{
					Title:        tt.input.Title,
					Descriptions: tt.input.Descriptions,
					BestTime:     tt.input.BestTime,
					PictureUrl:   tt.input.PictureUrl,
				})
			}
		})
	}
}
