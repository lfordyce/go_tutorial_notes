package xml_utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	"regexp"
	"strings"
	"testing"
)

func removal(x string) {
	r := regexp.MustCompile(`><(\\|\/)([a-zA-Z0-9]*)>`)
	allString := r.FindAllString(string(x), -1)
	fmt.Println(allString)
	ns := r.ReplaceAllString(string(x), "/>")
	fmt.Println(ns)
}

func removeEmptyTag(sXml string) {
	// <[A-Za-z_:]+.*?>
	r, err := regexp.Compile(`<[A-Za-z_:]+.*?><(\\|\/)([a-zA-Z0-9_]*)>`)
	checkErr(err)
	matches := r.FindAllString(sXml, -1)

	fmt.Println(sXml)

	if len(matches) > 0 {
		// <[A-Za-z_:]+.*?>
		r, err = regexp.Compile(`<[A-Za-z_:]+.*?>`)
		for i := 0; i < len(matches); i++ {

			xmlTag := r.FindString(matches[i])
			xmlTag = strings.Replace(xmlTag, "<", "", -1)
			xmlTag = strings.Replace(xmlTag, ">", "", -1)
			sXml = strings.Replace(sXml, matches[i], "<"+xmlTag+" />", -1)

		}
	}

	fmt.Println("")
	fmt.Println(sXml)
}

type ParseXML struct {
	Person struct {
		Name     string `xml:"Name"`
		LastName string `xml:"LastName"`
		Test     string `xml:"Abc"`
	} `xml:"Person"`
}

func TestParseXML(t *testing.T) {
	var err error
	var newPerson ParseXML

	newPerson.Person.Name = "Boot"
	newPerson.Person.LastName = "Testing"

	var bXml []byte
	var sXml string
	bXml, err = xml.Marshal(newPerson)
	checkErr(err)

	sXml = string(bXml)

	r, err := regexp.Compile(`<([a-zA-Z0-9]*)><(\\|\/)([a-zA-Z0-9]*)>`)
	checkErr(err)
	matches := r.FindAllString(sXml, -1)

	fmt.Println(sXml)

	if len(matches) > 0 {
		r, err = regexp.Compile("<([a-zA-Z0-9]*)>")
		for i := 0; i < len(matches); i++ {

			xmlTag := r.FindString(matches[i])
			xmlTag = strings.Replace(xmlTag, "<", "", -1)
			xmlTag = strings.Replace(xmlTag, ">", "", -1)
			sXml = strings.Replace(sXml, matches[i], "<"+xmlTag+"/>", -1)

		}
	}

	fmt.Println("")
	fmt.Println(sXml)
}

func checkErr(chk error) {
	if chk != nil {
		panic(chk)
	}
}

type myStruct struct {
	XMLName xml.Name `xml:"a"`
	Id      string   `xml:"id,attr"`
	Id2     string   `xml:"another_attr,attr"`
}

func TestCloseTags(t *testing.T) {
	s := &myStruct{Id: "an_id", Id2: "some_value"}

	x, err := xml.Marshal(s)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("x:", string(x))

	r := regexp.MustCompile("></[[:alnum]]*>")
	allString := r.FindAllString(string(x), -1)
	fmt.Println(allString)
	ns := r.ReplaceAllString(string(x), "/>")
	fmt.Println("ns:", ns)
}

func TestCloseAppData(t *testing.T) {
	appData := AppData{App: "SVOD", Name: "Year", Value: "2019"}
	marshal, err := xml.Marshal(appData)
	if err != nil {
		t.Fatal(err)
	}

	//errgroup.WithContext(ctx)

	fmt.Println(string(marshal))

	//r := regexp.MustCompile("></[[:alnum]]*>")
	//ns := r.ReplaceAllString(string(marshal), "/>")
	//fmt.Println("ns:", ns)

	r, err := regexp.Compile(`><(\\|\/)([a-zA-Z0-9_]*)>`)
	checkErr(err)
	matches := r.FindAllString(string(marshal), -1)
	if len(matches) > 0 {
		fmt.Println(matches)
	}
	ns := r.ReplaceAllString(string(marshal), "/>")
	fmt.Println("ns:", ns)
}

type Nodes struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Nodes    `xml:",any"`
}

func (n *Nodes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Nodes

	return d.DecodeElement((*node)(n), &start)
}

func walk(nodes []Nodes, f func(Nodes) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}

var data = []byte(`
<content>
    <p class="foo">this is content area</p>
    <animal>
        <p>This id dog</p>
        <dog>
           <p>tommy</p>
        </dog>
    </animal>
    <birds>
        <p class="bar">this is birds</p>
        <p>this is birds</p>
    </birds>
    <animal>
        <p>this is animals</p>
    </animal>
</content>
`)

func TestDecoder(t *testing.T) {
	buf := bytes.NewBuffer(data)
	dec := xml.NewDecoder(buf)

	var n Nodes
	err := dec.Decode(&n)
	if err != nil {
		panic(err)
	}

	walk([]Nodes{n}, func(n Nodes) bool {
		if len(n.Attrs) > 0 {
			for _, att := range n.Attrs {
				fmt.Println("attribute: ", att)
			}
		}
		//if n.XMLName.Local == "p" {
		//	fmt.Println(string(n.Content))
		//	fmt.Println(n.Attrs)
		//}
		return true
	})
}

//func TestCustomEncoding(t *testing.T) {
//	buffer := new(bytes.Buffer)
//	encoder := xml.NewEncoder(buffer)
//
//	encoder.
//}

const (
	empty = ""
	tab   = "\t"
)

func PrettyXml(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := xml.NewEncoder(buffer)
	encoder.Indent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func TestEtree(t *testing.T) {
	adi := buildAdi()
	fmt.Println(adi)
}

type Labels map[string]string

func TestWrap(t *testing.T) {
	document := etree.NewDocument()
	document.CreateProcInst(Target, DefaultProcInst)
	document.CreateDirective(DefaultDocktype)
	root := document.CreateElement(ADI)

	title := NewTitleAssetClass("InDemand", "indemand.com", "The_Titanic", "The Titanic asset package")
	movie := NewMovieAssetClass("InDemand", "indemand.com", "The_Titanic", "The Titanic asset package")

	compose(
		applyMultipleMetadata(movie),
		applyElement(title),
	)(root)

	document.Indent(2)
	buffer := new(bytes.Buffer)
	_, err := document.WriteTo(buffer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buffer.String())
}

func TestDecorateSection(t *testing.T) {

	document := etree.NewDocument()
	document.CreateProcInst(Target, DefaultProcInst)
	document.CreateDirective(DefaultDocktype)
	root := document.CreateElement(ADI)

	//var fn AssetElementFunc
	fn := AssetElementFunc(
		func(element *etree.Element) *etree.Element {
			return element.CreateElement(ASSET)
		},
	)

	title := NewTitleAssetClass("InDemand", "indemand.com", "The_Titanic", "The Titanic asset package")

	fn = DecorateSection(fn)

	element := fn(root)

	title.Construct(element)
	//fn(title.Construct(root))

	document.Indent(2)
	buffer := new(bytes.Buffer)
	_, err := document.WriteTo(buffer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buffer.String())

}

func TestNewDis(t *testing.T) {

	class := NewPackageAssetClass("InDemand", "indemand.com", "The_Titanic", "The Titanic asset package")
	movie := NewMovieAssetClass("InDemand", "indemand.com", "The_Titanic", "The Titanic asset package")
	title := NewTitleAssetClass("InDemand", "indemand.com", "The_Titanic", "The Titanic asset package")

	node := NewADI(class, title, movie, movie, movie)
	document := node.BuildDoc()

	document.Indent(2)
	buffer := new(bytes.Buffer)
	n, err := document.WriteTo(buffer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	fmt.Println(buffer.String())
}

type ClassType interface {
	Name() string
}

type movie struct{}
type preview struct{}

func (m *movie) Name() string   { return "movie" }
func (p *preview) Name() string { return "preview" }

type registry map[string]func() ClassType

func TestNewClassType(t *testing.T) {
	registries := make(registry)
	registries["movie"] = func() ClassType {
		return &movie{}
	}

	registries["preview"] = func() ClassType {
		return &preview{}
	}

	for k := range registries {

		name := registries[k]().Name()

		fmt.Println(name)
	}
}

func TestEtreeNamesapces(t *testing.T) {
	document := etree.NewDocument()
	document.CreateProcInst(Target, DefaultProcInst)

	element := &etree.Element{
		Space: "sig",
		Tag:   "NPTpoint",
		Attr: []etree.Attr{
			{
				Key:   "nptPoint",
				Value: "10.00",
			},
		},
	}
	document.SetRoot(element)

	document.Indent(2)
	buffer := new(bytes.Buffer)
	n, err := document.WriteTo(buffer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	fmt.Println(buffer.String())

}

//type Node struct {
//	E *etree.Element
//	Content string
//	Nodes []Node
//}
//
//func TestFooBar(t *testing.T) {
//
//	document := etree.NewDocument()
//	document.CreateProcInst(Target, DefaultProcInst)
//	root := document.CreateElement(ADI)
//
//
//	var n Node
//
//
//
//}

//func walk(nodes []Node, f func(Node) bool) {
//	for _, n := range nodes {
//		if f(n) {
//			walk(n.Nodes, f)
//		}
//	}
//}
