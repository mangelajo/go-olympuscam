package camera

import (
	"encoding/xml"
)


func (c *Client) GetCommandList() (*CommandList, error) {
	commandList := CommandList{}
	err := c.ExpectOkXML( "get_commandlist.cgi", &commandList)
	if err != nil {
		return nil, err
	}
	return &commandList, nil
}



type CommandList struct {
	XMLName xml.Name `xml:"oishare"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version"`
	Support []struct {
		Text string `xml:",chardata"`
		Func string `xml:"func,attr"`
	} `xml:"support"`
	Cgi []struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		HTTPMethod struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Cmd1 struct {
				Text   string `xml:",chardata"`
				Name   string `xml:"name,attr"`
				Param1 []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Cmd2 struct {
						Text   string `xml:",chardata"`
						Name   string `xml:"name,attr"`
						Param2 []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"param2"`
						Cmd3 []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"cmd3"`
					} `xml:"cmd2"`
				} `xml:"param1"`
				Cmd2 struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"cmd2"`
			} `xml:"cmd1"`
		} `xml:"http_method"`
	} `xml:"cgi"`
}


