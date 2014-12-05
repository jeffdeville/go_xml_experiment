package parse_xml

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Account struct {
	Id int
}

func main() {
	account := Account{Id: 123}

	tmplBytes, err := ioutil.ReadFile("./retrieveAccount.xml.tmpl")
	CheckError(err)

	tmplText := string(tmplBytes)

	tmpl, err := template.New("retrieveAccount").Parse(tmplText)
	CheckError(err)

	err = tmpl.Execute(os.Stdout, account)
	CheckError(err)
}
