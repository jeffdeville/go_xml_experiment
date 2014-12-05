package parse_xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type RetrieveAccount struct {
	AccountId string `xml:"Body>retrieveAccountResponse>accountNumber"`
	Version   string `xml:"Body>retrieveAccountResponse>version"`
}

func Test() RetrieveAccount {
	xmlData, err := ioutil.ReadFile("./retrieveAccountResponse.xml")
	CheckError(e)

	var resp RetrieveAccount
	xml.Unmarshal(xmlData, &resp)
	fmt.Println("AccountId:", resp.AccountId, "Version:", resp.Version)
	return resp
}
