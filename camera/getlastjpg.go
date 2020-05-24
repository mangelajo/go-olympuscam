package camera

func (c *Client) GetLastJpeg() ([]byte, error) {
	response, err := c.client.Get(c.baseUrl + "exec_takemisc.cgi?com=getlastjpg")
	return ExpectOKBody(response, err)
}

