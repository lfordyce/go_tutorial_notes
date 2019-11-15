package xml_utils

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var data1 = []byte(`
<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.receive.appservice.jcms.hanweb.com">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:wsGetInfosLink soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <nCataId xsi:type="xsd:int">3</nCataId>
         <bRef xsi:type="xsd:int">3</bRef>
         <nStart xsi:type="xsd:int">3</nStart>
         <nEnd xsi:type="xsd:int">3</nEnd>
         <bAsc xsi:type="xsd:int">3</bAsc>
         <strStartCTime xsi:type="xsd:string">gero et</strStartCTime>
         <strEndCTime xsi:type="xsd:string">sonoras imperio</strEndCTime>
         <strLoginId xsi:type="xsd:string">quae divum incedo</strLoginId>
         <strPwd xsi:type="xsd:string">verrantque per auras</strPwd>
         <strKey xsi:type="xsd:string">per auras</strKey>
      </ser:wsGetInfosLink>
   </soapenv:Body>
</soapenv:Envelope>`)

func TestEnvelope_MarshalXML(t *testing.T) {
	env := new(Envelope)

	// unmarshal xml data into env
	if err := xml.Unmarshal(data1, env); err != nil {
		panic(err)
	}
	fmt.Println(env)
	fmt.Println()

	// modify env
	env.Body.WSGetInfosLink.NStart.SetInt(123)
	env.Body.WSGetInfosLink.NEnd.SetInt(321)
	env.Body.WSGetInfosLink.StrLoginId.Value = "John Doe"

	// marshal modified env back into xml
	b, err := xml.MarshalIndent(env, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
