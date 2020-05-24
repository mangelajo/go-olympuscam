package camera

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Camera clientf", func() {
	When("SwitchOff is called", func() {
		It("should return no err", func() {
			client, mux, teardown := testSetup()
			defer teardown()

			mux.HandleFunc("/exec_pwoff.cgi", func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Method).To(Equal("GET"))
				w.WriteHeader(http.StatusAccepted)
			})

			err := client.PowerOff()
			Expect(err).NotTo(HaveOccurred())
		})
	})


	When("SwitchMode is called", func() {
		It("should return no err", func() {
			client, mux, teardown := testSetup()
			defer teardown()

			mux.HandleFunc("/switch_cammode.cgi", func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Method).To(Equal("GET"))
				w.WriteHeader(http.StatusOK)
			})

			err := client.SwitchMode(ModeShutter, Live480p)
			Expect(err).NotTo(HaveOccurred())
		})
	})

})


func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testSetup() (client *Client, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewClient()
	client.baseUrl = server.URL + "/"

	return client, mux, server.Close
}


func ExpectHttpParam(r *http.Request, param, value string) {
	keys, ok := r.URL.Query()[param]
	Expect(ok).To(BeTrue())
	Expect(keys).To(HaveLen(1))
	Expect(keys[0]).To(Equal(value))
}
