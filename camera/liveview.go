package camera

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)


func (c *Client) StartLivePreview() error {

	// E M-10 live view seems to work only in rec mode, exec_shuttermisc ignores the request
	if c.mode != ModeRec {
		err := c.SwitchMode(ModeRec, c.lvqty)
		if err != nil {
			return err
		}
	}

	c.livePreviewStopCh = make(chan struct{})

	pc, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return err
	}
	addr := strings.Split(pc.LocalAddr().String(), ":")
	port := addr[len(addr)-1]

	fmt.Printf("livepreview listening on %s\n", pc.LocalAddr().String())

	go c.livePreviewListener(pc, c.livePreviewStopCh)

	resp, err := c.client.Get(c.baseUrl + "exec_takemisc.cgi?com=startliveview&port="+port)
	defer resp.Body.Close()
	return ExpectOK(resp, err)
}

func (c *Client) StopLivePreview() error {
	close(c.livePreviewStopCh)
	resp, err := c.client.Get(c.baseUrl + "exec_takemisc.cgi?com=stopliveview")

	defer resp.Body.Close()
	return ExpectOK(resp, err)
}



type PictureReceiver struct {
	frameId uint32
	chunk uint16
	streamId uint32
	picData []byte
}

func (r *PictureReceiver) processData(buf []byte, n int){

	packetType := binary.BigEndian.Uint16(buf[0:2])
	chunk := binary.BigEndian.Uint16(buf[2:4])
	frameId := binary.BigEndian.Uint32(buf[4:8])
	streamId := binary.BigEndian.Uint32(buf[8:12])
	sod := binary.BigEndian.Uint16(buf[0xcc:0xce])



	switch packetType {
	case 0x9060:

		r.frameId = frameId
		r.chunk = chunk
		r.streamId = streamId
		r.picData = buf[0xcc:n]

		if sod == 0xffd8 {
			fmt.Printf("Start of frame found\n")
		} else {
			fmt.Printf("sod: %04x\n")
		}
	case 0x8060:
		if chunk == r.chunk + 1 && r.frameId == frameId {
			r.picData = append(r.picData, buf[12:n]...)
			r.chunk = chunk
		} else {
			fmt.Printf("Unexpected chunk number %04x vs %04x\n", chunk, r.chunk+1)
		}
	case 0x80e0:
		if chunk == r.chunk +1 && r.frameId == frameId {
			r.picData = append(r.picData, buf[12:n]...)
			r.chunk = chunk
			ioutil.WriteFile(fmt.Sprintf("frame_%08x.jpeg", r.frameId), r.picData, 0444)
			fmt.Printf("Frame done, len=%d %02x %02x %02x %02x\n", len(r.picData),
				r.picData[0],r.picData[1],r.picData[2],r.picData[3])
			r.picData = nil
		} else {
			fmt.Printf("Unexpected chunk number %04x vs %04x\n", chunk, r.chunk+1)
		}
	}

}

func (c *Client) livePreviewListener(pc net.PacketConn, stopCh chan struct{}) {
	buf := make([]byte, 66000)
	r := PictureReceiver{}
	fmt.Printf("liveview started\n")
	for {
		select {
		case  <-stopCh:
			fmt.Printf("liveview stopped\n")

			return
		default:
			n, _, err := pc.ReadFrom(buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "livePreviewListener receiver error: %s", err)
				return
			}
			bufCopy := make([]byte, len(buf))
			copy(bufCopy, buf)
			r.processData(bufCopy, n)

		}
	}
}
