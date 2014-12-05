package parse_xml_test

import (
	. "github.com/jeffdeville/parse_xml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseXml", func() {
	var retrieveAccount RetrieveAccount
	BeforeEach(func() {
		retrieveAccount = Test()
	})

	Describe("Parsing Xml", func() {
		It("Should populate the RetrieveAccount", func() {
			Expect(retrieveAccount.AccountId).To(Equal("644405"))
			Expect(retrieveAccount.Version).To(Equal("293"))
		})
	})
})
