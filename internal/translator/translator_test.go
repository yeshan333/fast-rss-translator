package translator_test

import (
	"bytes" // Added bytes import
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/yeshan333/fast-rss-translator/internal/translator"
)

func TestExecute(t *testing.T) {
	// This test makes actual network calls, consider mocking if needed for CI/CD
	// For now, we'll keep it as an integration test for google translate

	// Skip this test in CI environments or if proxy is not available
	if os.Getenv("CI") != "" || os.Getenv("HTTP_PROXY") == "" {
		t.Skip("Skipping Google Translate integration test in CI or when HTTP_PROXY is not set")
	}

	trans := &translator.Translator{
		HttpProxy: os.Getenv("HTTP_PROXY"), // Assumes HTTP_PROXY is set for testing
		Feed: translator.Feed{
			Name:            "feed_test_google.xml",
			Url:             "https://grafana.com/blog/index.xml", // Using a real RSS feed for testing
			TargetLanguage:  "zh",
			TranslateMode:   "origin",
			TranslateEngine: "google",
			MaxPost:         1,
		},
	}
	translateContent := "Hello World"
	translated := trans.DoTranslate(translateContent)
	if translated == "" || translated == translateContent {
		t.Errorf("Google Translate failed, expected translation for '%s', got '%s'", translateContent, translated)
	}
	t.Logf("Google Translate: '%s' -> '%s'", translateContent, translated)
}

func TestTranslateWithCloudflare(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header 'Bearer test-api-key', got '%s'", authHeader)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check content type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type header 'application/json', got '%s'", contentType)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Log received request at mock server
		t.Logf("Mock server received request URI: %s", r.RequestURI)
		t.Logf("Mock server received request URL Path: %s", r.URL.Path)
		t.Logf("Mock server received request method: %s", r.Method)
		// t.Logf("Mock server received request headers: %#v", r.Header) // Can be verbose

		// Check request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Mock server failed to read request body: %v", err)
		}
		expectedRequestBody := `"content":"你是一个专业的翻译助手，可以将用户输入的内容翻译成双语展现的形式，使用【】包裹原文，然后再跟译文，例如：Hello World，处理后为：【Hello World】你好世界。注意返回不要夹带任何信息除了译文和原文外的任何信息。翻译：Elixir with AI"`
		if !strings.Contains(string(body), expectedRequestBody) {
			t.Errorf("Mock Server: Expected request body to contain '%s', got '%s'", expectedRequestBody, string(body))
		}
		if !strings.Contains(string(body), `"model":"@cf/google/gemma-3-12b-it"`) {
			t.Errorf("Mock Server: Expected request body to contain model info, got '%s'", string(body))
		}

		// Send mock response
		mockResponse := `{
			"id": "id-1751744438633",
			"object": "chat.completion",
			"created": 1751744438,
			"model": "@cf/google/gemma-3-12b-it",
			"choices": [
				{
					"index": 0,
					"message": {
						"role": "assistant",
						"content": "【Elixir with AI】Elixir 与人工智能"
					},
					"logprobs": null,
					"finish_reason": "stop"
				}
			],
			"usage": {
				"prompt_tokens": 75,
				"completion_tokens": 11,
				"total_tokens": 86
			}
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, mockResponse)
	}))
	defer server.Close()

	// Override the actual Cloudflare API endpoint with the mock server's URL for this test
	// This requires a way to inject the URL, which is not directly available in the current structure.
	// For a real unit test, you'd refactor translateWithCloudflare to accept an http.Client and the URL,
	// or use a global variable for the endpoint that can be changed in tests.
	//
	// As a workaround for now, we will test the DoTranslate method which calls translateWithCloudflare internally.
	// We will set the CloudflareAccountID to a value that allows us to identify the mock server.
	// The actual constant `cloudflareAPIEndpoint` in `translator.go` will be used, so we need to ensure our mock server
	// URL structure matches what Sprintf expects if we were to modify that constant.
	// However, since we can't modify the constant for testing without changing the source,
	// we rely on the test setup to ensure the correct mock server is called if we were to intercept DNS or similar.
	//
	// A better approach would be to pass the API endpoint URL to the translateWithCloudflare function or the Translator struct.
	// For this exercise, we'll assume the `cloudflareAPIEndpoint` constant is modified for testing or we use environment variables.

	// Let's simulate setting the endpoint for testing if we could:
	// originalCloudflareAPIEndpoint := translator.CloudflareAPIEndpoint // Assuming it's a public var
	// translator.CloudflareAPIEndpoint = server.URL + "/client/v4/accounts/%s/ai/v1/chat/completions" // This is not possible with current const
	// defer func() { translator.CloudflareAPIEndpoint = originalCloudflareAPIEndpoint }()

	// Since we can't easily change the URL constant, we'll proceed with the test
	// and it will make a real HTTP request if not properly mocked at a lower level (e.g. http client injection).
	// For the purpose of this test, we will assume that if we set CloudflareAccountID to "test-account-id",
	// and if we could intercept HTTP calls, we would direct calls to `https://api.cloudflare.com/client/v4/accounts/test-account-id/...`
	// to our `server.URL`.
	//
	// The current `translateWithCloudflare` uses a global `cloudflareAPIEndpoint` constant.
	// To properly test this without network calls, `http.DefaultClient` could be mocked,
	// or `translateWithCloudflare` should accept an `*http.Client`.

	// We will proceed by creating a custom HTTP client and transport for this test.
	// This is a common way to mock HTTP clients in Go.
	originalClient := http.DefaultClient
	http.DefaultClient = &http.Client{
		Transport: &customTransport{serverURL: server.URL, t: t},
	}
	defer func() { http.DefaultClient = originalClient }()


	trans := translator.Translator{
		Feed: translator.Feed{
			Name:                "feed_cloudflare_test.xml",
			Url:                 "http://example.com/rss.xml", // Dummy URL, not used by DoTranslate directly for this test
			TargetLanguage:      "zh",                         // Not directly used by Cloudflare prompt but good to have
			TranslateMode:       "bilingual",                  // To ensure DoTranslate is called
			TranslateEngine:     "cloudflare",
			CloudflareAccountID: "test-account-id", // This will be part of the URL
			CloudflareApiKey:    "test-api-key",
		},
	}

	inputText := "Elixir with AI"
	expectedOutput := "【Elixir with AI】Elixir 与人工智能"
	actualOutput := trans.DoTranslate(inputText)

	if actualOutput != expectedOutput {
		t.Errorf("TranslateWithCloudflare: expected '%s', got '%s'", expectedOutput, actualOutput)
	}

	// Test case for when API key or Account ID is missing
	trans.Feed.CloudflareApiKey = ""
	actualOutput = trans.DoTranslate(inputText)
	if actualOutput != inputText {
		t.Errorf("TranslateWithCloudflare (missing API key): expected original text '%s', got '%s'", inputText, actualOutput)
	}
}

// customTransport to redirect requests to the mock server
type customTransport struct {
	serverURL string
	t         *testing.T
}

func (ct *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Check if the request is intended for Cloudflare
	// This is a simplified check. A more robust check might involve parsing the URL.
	if strings.Contains(req.URL.Host, "api.cloudflare.com") && strings.Contains(req.URL.Path, "/ai/v1/chat/completions") {
		// Redirect to mock server
		newURL := ct.serverURL + req.URL.Path
		if req.URL.RawQuery != "" {
			newURL += "?" + req.URL.RawQuery
		}

		// Create a new request to the mock server
		// We need to be careful about the host in the URL for the httptest server
		// The server expects requests to its own host.

		// Read the original request's body because it can only be read once.
		var bodyBytes []byte
		var err error
		if req.Body != nil {
			bodyBytes, err = io.ReadAll(req.Body)
			if err != nil {
				ct.t.Fatalf("Failed to read original request body: %v", err)
				return nil, err
			}
			// It's important to close the original body.
			req.Body.Close()
		}

		// Create a new request to the mock server.
		// The URL for NewRequest should be the mock server's root URL, as its handler is registered there.
		// Pass the original body content using a new buffer.
		mockReq, err := http.NewRequest(req.Method, ct.serverURL, bytes.NewBuffer(bodyBytes))
		if err != nil {
			ct.t.Fatalf("Failed to create request to mock server: %v", err)
			return nil, err
		}

		// Copy all headers from the original request to the new request.
		mockReq.Header = make(http.Header)
		for k, v := range req.Header {
			mockReq.Header[k] = v
		}

		// The httptest server's handler is at its root.
		// Send the request to the mock server.
		// Use a new client or http.DefaultClient if it's not the one being modified.
		// Since we are modifying http.DefaultClient's transport, we should use a specific client here
		// to avoid recursion if this transport was, by mistake, assigned to the client making this call.
		// However, the client making the *original* call is http.DefaultClient.
		// The client *inside* this RoundTrip should be a fresh one or one that doesn't use this transport.
		client := &http.Client{}

		ct.t.Logf("CustomTransport: Sending mock request to URL: %s", mockReq.URL.String())
		ct.t.Logf("CustomTransport: Mock request method: %s", mockReq.Method)
		// ct.t.Logf("CustomTransport: Mock request headers: %#v", mockReq.Header) // Verbose
		if mockReq.Body != nil {
			ct.t.Logf("CustomTransport: Mock request body is not nil")
		} else {
			ct.t.Logf("CustomTransport: Mock request body is nil")
		}

		return client.Do(mockReq)

	}
	// For other requests, use the default transport (which would be http.DefaultTransport)
	return http.DefaultTransport.RoundTrip(req)
}
