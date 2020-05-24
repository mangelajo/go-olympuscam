package camera

import (
	"encoding/xml"
)

func (c *Client) GetParameters() (*Parameters, error) {
	parameterList := Parameters{}
	err := c.ExpectOkXML( "get_camprop.cgi?com=desc&propname=desclist", &parameterList)
	if err != nil {
		return nil, err
	}
	return &parameterList, nil
}

type Parameters struct {
	XMLName xml.Name `xml:"desclist"`
	Desc    []struct {
		Propname  string `xml:"propname"`
		Attribute string `xml:"attribute"`
		Value     string `xml:"value"`
		Enum      string `xml:"enum"`
	} `xml:"desc"`
}


