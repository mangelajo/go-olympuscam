package camera

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetParameters", func() {
	When("called", func() {
		It("should return the list of parameters", func() {
			client, mux, teardown := testSetup()
			defer teardown()

			mux.HandleFunc("/get_camprop.cgi", func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Method).To(Equal("GET"))
				ExpectHttpParam(r, "com", "desc")
				ExpectHttpParam(r, "propname", "desclist")
				fmt.Fprint(w, getCamPropDescList)
			})

			paramList, err := client.GetParameters()
			Expect(err).NotTo(HaveOccurred())
			Expect(paramList.Desc).To(HaveLen(13))
			Expect(paramList.Desc[0].Value).To(Equal("0086-0064_0466x0350"))
		})
	})
})

const getCamPropDescList = `
<?xml version="1.0" encoding="UTF-8"?>
<desclist>
   <desc>
      <propname>touchactiveframe</propname>
      <attribute>get</attribute>
      <value>0086-0064_0466x0350</value>
   </desc>
   <desc>
      <propname>takemode</propname>
      <attribute>getset</attribute>
      <value>A</value>
      <enum>iAuto P A S M ART</enum>
   </desc>
   <desc>
      <propname>noisereduction</propname>
      <attribute>get</attribute>
      <value>off</value>
      <enum>off on auto</enum>
   </desc>
   <desc>
      <propname>lowvibtime</propname>
      <attribute>get</attribute>
      <value>0</value>
      <enum>0 1 125 250 500 1000 2000 4000 8000 15000 30000</enum>
   </desc>
   <desc>
      <propname>bulbtimelimit</propname>
      <attribute>get</attribute>
      <value>2</value>
      <enum>1 2 4 8 15 20 25 30</enum>
   </desc>
   <desc>
      <propname>digitaltelecon</propname>
      <attribute>get</attribute>
      <value>off</value>
      <enum>on off</enum>
   </desc>
   <desc>
      <propname>drivemode</propname>
      <attribute>getset</attribute>
      <value>continuous-L</value>
      <enum>normal continuous-H continuous-L selftimer customselftimer</enum>
   </desc>
   <desc>
      <propname>focalvalue</propname>
      <attribute>getset</attribute>
      <value>2.0</value>
      <enum>1.0 1.1 1.2 1.4 1.6 1.8 2.0 2.2 2.5 2.8 3.2 3.5 4.0 4.5 5.0 5.6 6.3 7.1 8.0 9.0 10 11 13 14 16 18 20 22 25 29 32 36 40 45 51 57 64 72 81 91</enum>
   </desc>
   <desc>
      <propname>expcomp</propname>
      <attribute>getset</attribute>
      <value>0.0</value>
      <enum>-5.0 -4.7 -4.3 -4.0 -3.7 -3.3 -3.0 -2.7 -2.3 -2.0 -1.7 -1.3 -1.0 -0.7 -0.3 0.0 +0.3 +0.7 +1.0 +1.3 +1.7 +2.0 +2.3 +2.7 +3.0 +3.3 +3.7 +4.0 +4.3 +4.7 +5.0</enum>
   </desc>
   <desc>
      <propname>shutspeedvalue</propname>
      <attribute>getset</attribute>
      <value>250</value>
      <enum>livetime livebulb 60" 50" 40" 30" 25" 20" 15" 13" 10" 8" 6" 5" 4" 3.2" 2.5" 2" 1.6" 1.3" 1" 1.3 1.6 2 2.5 3 4 5 6 8 10 13 15 20 25 30 40 50 60 80 100 125 160 200 250 320 400 500 640 800 1000 1250 1600 2000 2500 3200 4000 5000 6400 8000</enum>
   </desc>
   <desc>
      <propname>isospeedvalue</propname>
      <attribute>getset</attribute>
      <value>800</value>
      <enum>Auto Low 200 250 320 400 500 640 800 1000 1250 1600 2000 2500 3200 4000 5000 6400 8000 10000 12800 16000 20000 25600</enum>
   </desc>
   <desc>
      <propname>wbvalue</propname>
      <attribute>getset</attribute>
      <value>0</value>
      <enum>0 18 16 17 20 35 64 23 256 257 258 259 512</enum>
   </desc>
   <desc>
      <propname>artfilter</propname>
      <attribute>getset</attribute>
      <value>ARTBKT</value>
      <enum>ART01 ART02 ART03 ART04 ART05 ART06 ART07 ART08 ART09 ART10 ART11 ART12 ARTBKT</enum>
   </desc>
</desclist>
`
