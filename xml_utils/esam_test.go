package xml_utils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"
)

type Element struct {
	XMLName    xml.Name
	Attributes []xml.Attr
	E1         string `xml:"ELEM1"`
}

func TestDynamicAttr(t *testing.T) {
	element := Element{
		XMLName: xml.Name{
			Space: "",
			Local: "SignalProcessingNotification",
		},
		Attributes: []xml.Attr{
			{Name: xml.Name{
				Space: "",
				Local: "xmlns:sig",
			}, Value: "urn:cablelabs:md:xsd:signaling:3.0"},
			{Name: xml.Name{
				Space: "",
				Local: "xmlns:common",
			}, Value: "urn:cablelabs:iptvservices:esam:xsd:common:1"},
			{Name: xml.Name{
				Space: "",
				Local: "acquisitionPointIdentity",
			}, Value: "ExampleESAM"},
		},
		E1: "bar",
	}

	indent, err := xml.MarshalIndent(element, " ", "    ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(indent))
}

func TestXmlBuilder(t *testing.T) {
	p := Plan{PlanCode: "test"}
	p.UnitAmountInCents.AddCurrency("USD", 4000)
	p.UnitAmountInCents.AddCurrency("USD", 5000)
	if xmlstring, err := xml.MarshalIndent(p, "", "    "); err == nil {
		xmlstring = []byte(xml.Header + string(xmlstring))
		fmt.Printf("%s\n", xmlstring)
	}
}

type CurrencyArray struct {
	CurrencyList []Currency
}

func (c *CurrencyArray) AddCurrency(currency string, amount int) {
	newc := Currency{Amount: amount}
	newc.XMLName.Local = currency
	c.CurrencyList = append(c.CurrencyList, newc)
}

type Currency struct {
	XMLName xml.Name
	Amount  int `xml:",innerxml"`
}

type Plan struct {
	XMLName           xml.Name      `xml:"plan"`
	PlanCode          string        `xml:"plan_code,omitempty"`
	CreatedAt         *time.Time    `xml:"created_at,omitempty"`
	UnitAmountInCents CurrencyArray `xml:"unit_amount_in_cents"`
	SetupFeeInCents   CurrencyArray `xml:"setup_in_cents"`
}

func TestBuildSignal(t *testing.T) {

	esam := &ESAM{
		Namespace:                "urn:cablelabs:iptvservices:esam:xsd:signal:1",
		SignalingNamespace:       "urn:cablelabs:md:xsd:signaling:3.0",
		CommonNamespace:          "urn:cablelabs:iptvservices:esam:xsd:common:1",
		XsiNamespace:             "http://www.w3.org/2001/XMLSchema-instance",
		AcquisitionPointIdentity: "ExampleESAM",
		Batch: &CommonBatchInfo{
			BatchId: "dc44d5ad-fef2-4a4c-9a56-e6c5f565dfcb",
			Source: &CommonBatchSource{
				Type:  "content:MovieType",
				UriId: "comcast.com-TEST2015052210096017",
			},
		},
	}

	indent, err := xml.MarshalIndent(esam, " ", "    ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(indent))
}

type ESAM struct {
	XMLName                  xml.Name `xml:"SignalProcessingNotification"`
	Namespace                string   `xml:"namespace,attr"`
	SignalingNamespace       string   `xml:"xmlns:sig,attr"`
	CommonNamespace          string   `xml:"xmlns:common,attr"`
	XsiNamespace             string   `xml:"xmlns:xsi,attr"`
	AcquisitionPointIdentity string   `xml:"acquisitionPointIdentity,attr"`
	Batch                    *CommonBatchInfo
}

type CommonBatchInfo struct {
	XMLName xml.Name `xml:"common:BatchInfo"`
	BatchId string   `xml:"batchId,attr"`
	Source  *CommonBatchSource
}

type CommonBatchSource struct {
	XMLName xml.Name `xml:"common:Source"`
	Type    string   `xml:"xsi:type,attr"`
	UriId   string   `xml:"uriId,attr"`
}

func TestTestBuildSignalWithEtree(t *testing.T) {
	document := etree.NewDocument()
	spn := document.CreateElement("SignalProcessingNotification")
	spn.CreateAttr("xmlns", "")

}

func TestReadXmlandParse(t *testing.T) {
	file, err := ioutil.ReadFile("/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/ZombieDatingShow_adi.xml")
	if err != nil {
		t.Fatal(err)
	}

	buffer := bytes.NewBuffer(file)
	decoder := xml.NewDecoder(buffer)
	var n Nodes
	if err := decoder.Decode(&n); err != nil {
		t.Fatal(err)
	}

	walk([]Nodes{n}, func(nodes Nodes) bool {
		if len(nodes.Attrs) > 0 {

			sort.Slice(nodes.Attrs, func(i, j int) bool {
				return nodes.Attrs[i].Value <= nodes.Attrs[j].Value
			})

			idx := sort.Search(len(nodes.Attrs), func(i int) bool {
				return strings.Contains(nodes.Attrs[i].Value, "LocalAd")
				//return nodes.Attrs[i].Value >= "LocalAdBreak_1"
			})

			//if idx < len(nodes.Attrs) && nodes.Attrs[idx].Value == "LocalAdBreak_1" {
			//	fmt.Printf("found \"%s\" at files[%d]\n", nodes.Attrs[idx], idx)
			//}

			if idx < len(nodes.Attrs) && strings.Contains(nodes.Attrs[idx].Value, "LocalAd") {
				fmt.Printf("found \"%s\" at files[%d]\n", nodes.Attrs[idx], idx)
			}

			//for _, att := range nodes.Attrs {
			//	//fmt.Println("attribute: ", att)
			//	//fmt.Printf("Local: %s, Space: %s, Value: %s\n", att.Name.Local, att.Name.Space, att.Value)
			//	if att.Value == "LocalAdBreak_1" {
			//		fmt.Println("attribute: ", att)
			//		//fmt.Println("Local Ad Break 1 Value: ", att.Value)
			//	}
			//}

		}
		return true
	})
}

func TestMatcher(t *testing.T) {
	m := make(map[ResponseSignalAttr]string)

	m[AcquisitionSignalID] = "1"
	m[AcquisitionPointIdentity] = "TEST_POINT"
	m[SignalPointID] = "1"
	m[Action] = "create"
	var r ResponseSignal
	matchFieldToKey(m, &r)

	r.XMLName.Local = "ResponseSignal"
	fmt.Println(r)
}

func TestConstructSignalProcessingNotification(t *testing.T) {

	m := make(map[ResponseSignalAttr]string)

	m[AcquisitionSignalID] = "1"
	m[AcquisitionPointIdentity] = "TEST_POINT"
	m[SignalPointID] = "1367.166"
	m[Action] = "create"

	b := &BatchInfo{
		BatchId: "1",
		Source: &Source{
			Type:  "content:MovieType",
			UriId: "/mnt/support/bferrentino/teting_content/HEVC_rap",
		},
	}

	info := &SegmentationDescriptorInfo{
		SegmentEventId:   "1",
		SegmentTypeId:    "52",
		UpidType:         "9",
		Upid:             "1",
		Duration:         "60.026",
		SegmentNum:       "1",
		SegmentsExpected: "1",
	}

	descriptor := &SCTE35PointDescriptor{
		SpliceCommandType: "06",
		Segmentation:      info,
	}

	point := &NPTPoint{
		NptPoint:   "1367.166",
		Descriptor: descriptor,
	}

	s := &SignalProcessingNotification{
		Xmlns:   string(CableLabsSignalUrn),
		Xsi:     CableLabsSchemaInstanceXsi,
		Common:  CableLabsCommonUrn,
		Content: "",
		Core:    "",
		Sig:     CableLabsSignalingUrn,
		Batch:   b,
	}

	s.AddResponseSignal(m, point)
	s.AddConditioningInfo("PT1367.166S", "1", "PT60.026S")

	marshalIndent, err := xml.MarshalIndent(s, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	s2 := string(marshalIndent)
	fmt.Println(s2)
	//replace := strings.Replace(s2, "></sig:SegmentationDescriptorInfo>", " />", -1)
	//s3 := strings.Replace(replace, "></common:Source>", " />", -1)
	//
	//fmt.Println(s3)

	//r := regexp.MustCompile("></[[:alnum]]*>")

	//r, err := regexp.Compile(`><(\\|\/)([a-zA-Z0-9_]*)>`)
	r, err := regexp.Compile(`><(\\|\/)([a-zA-Z0-9.,:_]*)>`)
	if err != nil {
		t.Fatal(err)
	}
	matches := r.FindAllString(s2, -1)
	if len(matches) > 0 {
		fmt.Println(matches)
	}
	//ns := r.ReplaceAllString(string(marshalIndent), "/>")
	//fmt.Println("ns:", ns)

	//if xmlstring, err := xml.MarshalIndent(s, "", "    "); err == nil {
	//	xmlstring = []byte(xml.Header + string(xmlstring))
	//	fmt.Printf("%s\n", xmlstring)
	//}

}
