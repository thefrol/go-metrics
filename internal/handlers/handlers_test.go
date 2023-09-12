package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	type result struct {
		code int
		body string
	}

	tt := []struct {
		name     string
		path     string
		expected result
	}{
		{
			name: "Push counter",
			path: "/update/counter/PollCount/10",
			expected: result{
				code: http.StatusOK,
			},
		},
		{
			name: "Push gauge",
			path: "/update/gauge/Alloc/13.123",
			expected: result{
				code: http.StatusOK,
			},
		},
		{
			name: "Push unknown metric kind",
			path: "/update/unknown/Alloc/12.123",
			expected: result{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "Push without name metric",
			path: "/update/Alloc/12.123",
			expected: result{
				code: http.StatusNotFound,
			},
		},
		{
			name: "Push counter with invalid name",
			path: "/update/counter/Alloc/18446744073709551617",
			expected: result{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "Push counter with invalid value",
			path: "/update/gauge/PollCount/10\\.0",

			expected: result{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "Push method get",
			path: "/",

			expected: result{
				code: http.StatusMethodNotAllowed,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//	s := storage.New()
			//	e := echo.New()
			assert := assert.New(t)
			//	Webhook(s)
			//	e.Start("8080")
			res, _ := http.Post("http://127.0.0.1:8080"+tc.path, "text/plain", nil)
			assert.Equal(tc.expected.code, res.StatusCode)
			defer res.Body.Close()
		})
	}
}

func TestAllMetrics(t *testing.T) {
	type result struct {
		code int
		body string
	}

	tt := []struct {
		name     string
		path     string
		expected result
	}{
		{
			name: "Push method POST",
			path: "/update/gauge/Alloc/13.123",

			expected: result{
				code: http.StatusMethodNotAllowed,
			},
		},
		{
			name: "Push counter",
			path: "/",
			expected: result{
				code: http.StatusOK,
				body: "{}",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//s := storage.New()
			e := echo.New()
			assert := assert.New(t)
			e.Start("8080")
			//AllMetrics(s)
			res, _ := http.Get("http://127.0.0.1:8080" + tc.path)
			assert.Equal(tc.expected.code, res.StatusCode)

			if tc.expected.code == http.StatusOK {
				respBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				assert.NotEmpty(len(respBody))
				defer res.Body.Close()
			}
		})
	}
}

func TestGetMetric(t *testing.T) {
	type result struct {
		code int
		body string
	}

	tt := []struct {
		name     string
		path     string
		expected result
	}{
		{
			name: "Get counter",
			path: "/value/counter/PollCount",
			expected: result{
				code: http.StatusOK,
				body: "9601",
			},
		},
		{
			name: "Get gauge",
			path: "/value/gauge/Alloc",
			expected: result{
				code: http.StatusOK,
				body: "11.345",
			},
		},
		{
			name: "Get unknown metric kind",
			path: "/value/unknown/Alloc",
			expected: result{
				code: http.StatusNotFound,
			},
		},
		{
			name: "Get unknown counter",
			path: "/value/counter/unknown",
			expected: result{
				code: http.StatusNotFound,
			},
		},
		{
			name: "Get unknown gauge",
			path: "/value/gauge/unknown",
			expected: result{
				code: http.StatusNotFound,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//	s := storage.New()
			e := echo.New()
			assert := assert.New(t)
			//	ValueMetrics(s)
			e.Start("8080")
			res, _ := http.Get("http://127.0.0.1:8080" + tc.path)
			assert.Equal(tc.expected.code, res.StatusCode)

			if tc.expected.code == http.StatusOK {
				respBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				defer res.Body.Close()
				assert.NotEmpty(string(respBody))
			}
		})
	}
}
