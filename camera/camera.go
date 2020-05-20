package camera

import (
	"fmt"
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

type Client struct {
	clientMu sync.Mutex
	client *http.Client
	baseUrl string
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
		baseUrl: "http://"+DefaultIp +"/",
	}
}

func (c *Client) PowerOff() error {
	resp, err := c.client.Get(c.baseUrl + "exec_pwoff.cgi")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		fmt.Printf("unexpected poweroff response: %+v\n", resp)
	}
	return nil
}





