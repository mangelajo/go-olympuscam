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
				keys, ok := r.URL.Query()["com"]
				Expect(ok).To(BeTrue())
				Expect(keys).To(HaveLen(1))
				Expect(keys[0]).To(Equal("getlastjpg"))
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
