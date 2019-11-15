package xsd

import (
	"aqwari.net/xml/xmltree"
	"aqwari.net/xml/xsdgen"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"io/ioutil"
	"os"
	"testing"
)

func tmpfile() *os.File {
	f, err := ioutil.TempFile("", "xsdgen_test")
	if err != nil {
		panic(err)
	}
	return f
}

func TestFoo(t *testing.T) {
	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.PackageName("xsd"),
	)
	if err := cfg.GenCLI(
		"schema/MD-SP-CONTENT-I02.xsd",
		"schema/MD-SP-CORE-I02.xsd",
		"schema/MD-SP-SIGNALING-I02.xsd",
		"schema/OC-SP-ESAM-API-I03-Common.xsd",
		//"schema/OC-SP-ESAM-API-I03-Manifest.xsd",
		"schema/OC-SP-ESAM-API-I03-Signal.xsd",
		"schema/SCTE35.xsd",
	); err != nil {
		t.Fatal(err)
	}
}

func TestGenerator(t *testing.T) {
	bytes2, err := ioutil.ReadFile("sample/Signal_Processing_Notification_EX.xml")
	if err != nil {
		t.Fatal(err)
	}

	buffer := bytes.NewBuffer(bytes2)
	decoder := xml.NewDecoder(buffer)

	var s SignalProcessingNotificationType
	if err := decoder.Decode(&s); err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)

	marshalIndent, err := xml.MarshalIndent(s, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	s2 := string(marshalIndent)
	fmt.Println(s2)
}

func TestAcquiredSignal_MarshalXML(t *testing.T) {
	s := SignalProcessingNotificationType{
		BatchInfo: BatchInfoType{
			Source: MovieType{
				UriId: "/mnt/support/bferrentino/teting_content/HEVC_rap",
			},
			BatchId: "1",
		},
		AcquisitionPointIdentity: "ExampleESAM",
	}

	marshalIndent, err := xml.MarshalIndent(s, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	s2 := string(marshalIndent)
	fmt.Println(s2)
}

func TestSearchFunction(t *testing.T) {
	file, err := ioutil.ReadFile("/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/ZombieDatingShow_adi.xml")
	if err != nil {
		t.Fatal(err)
	}

	root, err := xmltree.Parse(file)
	if err != nil {
		t.Fatal(err)
	}

	resut := root.SearchFunc(func(element *xmltree.Element) bool {
		return element.Attr("", "Asset_Class") == "package"
	})

	for _, el := range resut {
		//t.Logf("found %s value=%s", el.Name.Local, el.Attr("", "value"))
		for _, a := range el.StartElement.Attr {
			fmt.Println(a)
		}
	}
}

func TestMyQuery(t *testing.T) {
	result := MyQuery(linq.Range(1, 10)).GreaterThan(5).Results()
	fmt.Println(result)
}

type MyQuery linq.Query

func (q MyQuery) GreaterThan(threshold int) linq.Query {
	return linq.Query{
		Iterate: func() linq.Iterator {
			next := q.Iterate()

			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if item.(int) > threshold {
						return
					}
				}
				return
			}
		},
	}
}
