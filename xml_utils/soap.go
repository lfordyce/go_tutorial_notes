package xml_utils

import (
	"encoding/xml"
	"strconv"
)

type Envelope struct {
	XMLName xml.Name  `xml:"Envelope"`
	Soapenv XmlnsAttr `xml:"soapenv,attr"`
	XSI     XmlnsAttr `xml:"xsi,attr"`
	XSD     XmlnsAttr `xml:"xsd,attr"`
	SER     XmlnsAttr `xml:"ser,attr"`
	Header  *Header   `xml:"Header"`
	Body    *Body     `xml:"Body"`
}

func (env *Envelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "soapenv:" + start.Name.Local
	return e.EncodeElement(*env, start)
}

type Header struct {
	XMLName xml.Name `xml:"Header"`
}

func (h *Header) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "soapenv:" + start.Name.Local
	return e.EncodeElement(*h, start)
}

type Body struct {
	XMLName        xml.Name `xml:"Body"`
	WSGetInfosLink *WSGetInfosLink
}

func (b *Body) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "soapenv:" + start.Name.Local
	return e.EncodeElement(*b, start)
}

type WSGetInfosLink struct {
	XMLName       xml.Name      `xml:"wsGetInfosLink"`
	EncodingStyle SoapenvAttr   `xml:"encodingStyle,attr"`
	NCataId       XSIIntType    `xml:"nCataId"`
	BRef          XSIIntType    `xml:"bRef"`
	NStart        XSIIntType    `xml:"nStart"`
	NEnd          XSIIntType    `xml:"nEnd"`
	BAsc          XSIIntType    `xml:"bAsc"`
	StrStartCTime XSIStringType `xml:"strStartCTime"`
	StrEndCTime   XSIStringType `xml:"strEndCTime"`
	StrLoginId    XSIStringType `xml:"strLoginId"`
	StrPwd        XSIStringType `xml:"strPwd"`
	StrKey        XSIStringType `xml:"strKey"`
}

func (ws *WSGetInfosLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "ser:" + start.Name.Local
	return e.EncodeElement(*ws, start)
}

type XSIIntType struct {
	XSIIntAttr XSIIntAttr `xml:",attr"`
	Value      string     `xml:",innerxml"`
}

func (xsi XSIIntType) GetInt() int {
	i, err := strconv.Atoi(xsi.Value)
	if err != nil {
		panic(err)
	}
	return i
}

func (xsi *XSIIntType) SetInt(i int) {
	xsi.Value = strconv.Itoa(i)
}

type XSIIntAttr struct{}

func (XSIIntAttr) MarshalXMLAttr(xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "xsi:type"},
		Value: "xsd:int",
	}, nil
}

type XSIStringType struct {
	XSIStringAttr XSIStringAttr `xml:",attr"`
	Value         string        `xml:",innerxml"`
}

type XSIStringAttr struct{}

func (XSIStringAttr) MarshalXMLAttr(xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "xsi:type"},
		Value: "xsd:string",
	}, nil
}

type SoapenvAttr string

func (a SoapenvAttr) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "soapenv:" + n.Local},
		Value: string(a),
	}, nil
}
