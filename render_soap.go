package parse_xml

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Account struct {
	Id int
}

// func checkError(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	account := Account{Id: 123}

	tmplBytes, err := ioutil.ReadFile("./retrieveAccount.xml.tmpl")
	checkError(err)

	tmplText := string(tmplBytes)

	tmpl, err := template.New("retrieveAccount").Parse(tmplText)
	checkError(err)

	err = tmpl.Execute(os.Stdout, account)
	checkError(err)
}
