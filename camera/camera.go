package camera

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	DefaultIp = "192.168.0.10"
	)


type CameraMode string

const (
	ModeRec CameraMode ="rec"
	ModePlay CameraMode = "play"
	ModeShutter CameraMode ="shutter"

)

type LiveQuality string
const (
	Live240p LiveQuality = "0320x0240"
	Live480p LiveQuality = "0640x0480"
	)

type Client struct {
	clientMu sync.Mutex
	client *http.Client
	baseUrl string
	mode CameraMode
	lvqty LiveQuality
	livePreviewStopCh chan struct{}
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
		baseUrl: "http://"+DefaultIp +"/",
		mode: ModePlay,
		lvqty: Live480p,
	}
}

func (c *Client) SwitchMode(mode CameraMode, lvqty LiveQuality) error {
	resp, err := c.client.Get(c.baseUrl + "switch_cammode.cgi?mode=" + string(mode)+ "&lvqty="+ string(lvqty))
	if err != nil {
		c.lvqty = lvqty
	}
	defer resp.Body.Close()
	return ExpectOK(resp, err)
}

func (c *Client) PowerOff() error {
	resp, err := c.client.Get(c.baseUrl + "exec_pwoff.cgi")
	defer resp.Body.Close()
	return ExpectAccepted(resp, err)

}


func ExpectAccepted(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected response: %+v\n", resp)
	}
	return nil
}


func ExpectOK(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response: %+v\n", resp)
	}
	return nil
}

func ExpectOKBody(resp *http.Response, err error) ([]byte, error) {
	err = ExpectOK(resp, err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) ExpectOkXML(url string, v interface{}) error {

	response, err := c.client.Get(c.baseUrl + url)
	bytes, err := ExpectOKBody(response, err)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, v)
}


