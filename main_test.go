package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
)

var ascii string

func init() {
	aciiBytes := make([]byte, 94)
	for i := range aciiBytes {
		aciiBytes[i] = byte(i + 33)
	}
	ascii = string(aciiBytes)
}

func String(t *testing.T, length int, chars string) (string, error) {
	t.Helper()
	result := make([]rune, length)
	runes := []rune(chars)
	x := int64(len(runes))
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(x))
		if err != nil {
			return "", fmt.Errorf("String: %w", err)
		}
		result[i] = runes[num.Int64()]
	}
	return string(result), nil
}

func ASCII(t *testing.T, length int) (string, error) {
	t.Helper()
	return String(t, length, ascii)
}

func Kana(t *testing.T, length int) (string, error) {
	t.Helper()
	return String(t, length, "いろはにほへどちりぬるをわがよたれぞつねならんういのおくやまきょうこえてあさきゆめみじよいもせず")
}

func TestSaveHandler(t *testing.T) {
	alnum, err := ASCII(t, 5)
	if err != nil {
		t.Fatal(err)
	}

	kana, err := Kana(t, 3)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		target string
		body   string
	}
	cases := map[string]struct {
		args       args
		wantStatus int
	}{
		"success": {
			args: args{
				target: "/save/TestSaveHandler0Success",
				body:   "body",
			},
			wantStatus: http.StatusFound,
		},
		fmt.Sprintf("%s is too long", alnum): {
			args: args{
				target: "/save/TestSaveHandler0alnum",
				body:   alnum,
			},
			wantStatus: http.StatusBadRequest,
		},
		fmt.Sprintf("%s should be accepted", kana): {
			args: args{
				target: "/save/TestSaveHandler0iroha",
				body:   kana,
			},
			wantStatus: http.StatusFound,
		},
	}

	for name, tc := range cases {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			values := url.Values{}
			values.Set("body", tc.args.body)

			request := httptest.NewRequest(http.MethodPost, tc.args.target, strings.NewReader(values.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			responseRecorder := httptest.NewRecorder()

			handler := makeHandler(saveHandler)
			handler.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.wantStatus {
				t.Errorf("got status %d, want %d", responseRecorder.Code, tc.wantStatus)
				return
			}
		})
	}
}

func TestSaveHandlerIsGoroutineSafe(t *testing.T) {
	var saveCalled int

	f := func(t *testing.T) {
		t.Helper()

		values := url.Values{}
		values.Set("body", "body")

		request := httptest.NewRequest(http.MethodPost, "/save/TestSaveHandlerIsGoroutineSafe", strings.NewReader(values.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		responseRecorder := httptest.NewRecorder()

		handler := makeHandler(saveHandler)
		handler.ServeHTTP(responseRecorder, request)

		wantStatus := http.StatusFound
		if responseRecorder.Code != wantStatus {
			t.Logf("got status %d, want %d", responseRecorder.Code, wantStatus)
			return
		}

		saveCalled++
	}

	wantCalled := 10
	var wg sync.WaitGroup
	for i := 0; i < wantCalled; i++ {
		wg.Add(1)
		go func() {
			f(t)
			wg.Done()
		}()
	}
	wg.Wait()

	if saveCalled != wantCalled {
		t.Errorf("got called %d, want %d", saveCalled, wantCalled)
	}
}
