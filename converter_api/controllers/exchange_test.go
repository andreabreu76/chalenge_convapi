package controllers

import (
	"github.com/andreabreu76/converter_api/entities"
	"github.com/andreabreu76/converter_api/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExchangeHandlers_ConvertCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockExchangeService(ctrl)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		service   services.ExchangeService
		args      args
		expected  int
		setupMock func()
	}{
		{
			name:    "successful conversion",
			service: mockService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8000/exchange/452.90/BRL/USD/4.50", nil),
			},
			expected: http.StatusOK,
			setupMock: func() {
				mockService.EXPECT().ConvertCurrency(gomock.Any(), 452.90, "BRL", "USD", 4.50).Return(&entities.ConversionResult{ConvertedValue: 2048.05}, nil)
			},
		},
		{
			name:    "invalid amount",
			service: mockService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8000/exchange/invalid/BRL/USD/4.50", nil),
			},
			expected:  http.StatusBadRequest,
			setupMock: func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			h := &ExchangeHandlers{
				service: tt.service,
			}

			h.ConvertCurrency(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder)
			assert.Equal(t, tt.expected, resp.Code)
		})
	}
}
