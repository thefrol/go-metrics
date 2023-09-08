package handlers

//
//import (
//	"github.com/Sofja96/go-metrics.git/internal/storage"
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
////type want struct {
////	code        int
////	response    string
////	contentType string
////}
//
//func TestWebhook(t *testing.T) {
//	type args struct {
//		s *storage.MemStorage
//	}
//	s := storage.New()
//	e := echo.New()
//	// описываем ожидаемое тело ответа при успешном запросе
//	successBody := s.AllMetrics()
//	//`{
//	//    "response": {
//	//        "text": "Извините, я пока ничего не умею"
//	//    },
//	//    "version": "1.0"
//	//}`
//
//	// описываем набор данных: метод запроса, ожидаемый код ответа, ожидаемое тело
//	testCases := []struct {
//		method       string
//		expectedCode int
//		expectedBody string
//	}{
//		//{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
//		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
//		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
//		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: successBody},
//		{method: http.MethodGet, expectedCode: http.StatusOK, expectedBody: successBody},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.method, func(t *testing.T) {
//			//	r := httptest.NewRequest(tc.method, "/", nil)
//			w := httptest.NewRecorder()
//
//			// вызовем хендлер как обычную функцию, без запуска самого сервера
//			//	Webhook(w, r)
//			//	Webhook(s)
//			//AllMetrics(s)
//			e.GET("/", AllMetrics(s))
//
//			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
//			// проверим корректность полученного тела ответа, если мы его ожидаем
//			if tc.expectedBody != "" {
//				// assert.JSONEq помогает сравнить две JSON-строки
//				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
//			}
//		})
//	}
//}
//
////func TestWebhook(t *testing.T) {
////	//e := echo.New()
////	//rec := httptest.NewRecorder()
////	//c := e.NewContext(req, rec)
////	type args struct {
////		s *storage.MemStorage
////	}
////	type want struct {
////		code int
////		//	response    string
////		contentType string
////	}
////	tests := []struct {
////		name string
////		args args
////		want want
////		//want echo.HandlerFunc
////	}{
////
////		{
////			name: "positive test #1",
////			want: want{
////				code: 200,
////				//	response:    `{"":""}`,
////				contentType: "text/plain;",
////			},
////		},
////		// TODO: Add test cases.
////	}
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			if got := Webhook(tt.args.s); !reflect.DeepEqual(got, tt.want) {
////				t.Errorf("Webhook() = %v, want %v", got, tt.want)
////			}
////		})
////	}
////}
