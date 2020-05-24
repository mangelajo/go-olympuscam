package camera

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetLastJpeg", func() {
	When("called", func() {
		It("Should use the right parameters and return data", func() {
			client, mux, teardown := testSetup()
			defer teardown()

			mux.HandleFunc("/exec_takemisc.cgi", func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Method).To(Equal("GET"))
				ExpectHttpParam(r, "com", "getlastjpg")
				fmt.Fprint(w, getLastJpegOutput)
			})

			data, err := client.GetLastJpeg()
			Expect(err).NotTo(HaveOccurred())

			dataStr := string(data)
			Expect(dataStr).To(Equal(getLastJpegOutput))
		})
	})
})

const getLastJpegOutput = "I'm a JPEG, aren't I?"
