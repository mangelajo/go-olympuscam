package camera

func (c *Client) GetRecView() ([]byte, error) {
	response, err := c.client.Get(c.baseUrl + "exec_takemisc.cgi?com=getrecview")
	return ExpectOKBody(response, err)
}

