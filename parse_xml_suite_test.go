package parse_xml_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParseXml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParseXml Suite")
}
