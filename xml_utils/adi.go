package xml_utils

import (
	"github.com/beevik/etree"
	"sort"
)

// asset class types
const (
	pckAssetClassType     AssetClass = "package"
	titleAssetClassType   AssetClass = "title"
	movieAssetClassType   AssetClass = "movie"
	previewAssetClassType AssetClass = "preview"
	captionAssetClassType AssetClass = "closed caption"
	posterAssetClassType  AssetClass = "poster"
	assetsGroupClassType  AssetClass = "assetsGroup"
	rootClassType         AssetClass = "root"
)

type Rank int
type AssetClass string

const (
	CaptionRank Rank = iota
	PreviewRank
	PosterRank
	MovieRank
	TitleRank
	PckRank
)

// xml tags used to generate ADI
const (
	Target          = "xml"
	DefaultProcInst = `version="1.0" encoding="UTF-8"`
	DefaultDocktype = `DOCTYPE ADI SYSTEM "ADI.DTD"`
	ADI             = "ADI"
	METADATA        = "Metadata"
	AMS             = "AMS"
	ASSET           = "Asset"
)

const (
	PROVICER      = "Provider"
	PRODUCT       = "Product"
	ASSET_NAME    = "Asset_Name"
	VersionMajor  = "Version_Major"
	VersionMinor  = "Version_Minor"
	DESCRIPTION   = "Description"
	CREATION_DATE = "Creation_Date"
	ProvierId     = "Provider_ID"
	ASSET_ID      = "Asset_ID"
	ASSET_CLASS   = "Asset_Class"
	VERT          = "Verb"
	TITLE         = "Title"
)

//
//type ADI struct {
//	XMLName  xml.Name `xml:"ADI"`
//	Text     string   `xml:",chardata"`
//	Metadata struct {
//		Text string `xml:",chardata"`
//		AMS  struct {
//			Text         string `xml:",chardata"`
//			Provider     string `xml:"Provider,attr"`
//			Product      string `xml:"Product,attr"`
//			AssetName    string `xml:"Asset_Name,attr"`
//			VersionMajor string `xml:"Version_Major,attr"`
//			VersionMinor string `xml:"Version_Minor,attr"`
//			Description  string `xml:"Description,attr"`
//			CreationDate string `xml:"Creation_Date,attr"`
//			ProviderID   string `xml:"Provider_ID,attr"`
//			AssetID      string `xml:"Asset_ID,attr"`
//			AssetClass   string `xml:"Asset_Class,attr"`
//		} `xml:"AMS"`
//		AppData struct {
//			Text  string `xml:",chardata"`
//			App   string `xml:"App,attr"`
//			Name  string `xml:"Name,attr"`
//			Value string `xml:"Value,attr"`
//		} `xml:"App_Data"`
//	} `xml:"Metadata"`
//	Asset struct {
//		Text     string `xml:",chardata"`
//		Metadata struct {
//			Text string `xml:",chardata"`
//			AMS  struct {
//				Text         string `xml:",chardata"`
//				Provider     string `xml:"Provider,attr"`
//				Product      string `xml:"Product,attr"`
//				AssetName    string `xml:"Asset_Name,attr"`
//				VersionMajor string `xml:"Version_Major,attr"`
//				VersionMinor string `xml:"Version_Minor,attr"`
//				Description  string `xml:"Description,attr"`
//				CreationDate string `xml:"Creation_Date,attr"`
//				ProviderID   string `xml:"Provider_ID,attr"`
//				AssetID      string `xml:"Asset_ID,attr"`
//				AssetClass   string `xml:"Asset_Class,attr"`
//			} `xml:"AMS"`
//			AppData []struct {
//				Text  string `xml:",chardata"`
//				App   string `xml:"App,attr"`
//				Name  string `xml:"Name,attr"`
//				Value string `xml:"Value,attr"`
//			} `xml:"App_Data"`
//		} `xml:"Metadata"`
//		Asset []struct {
//			Text     string `xml:",chardata"`
//			Metadata struct {
//				Text string `xml:",chardata"`
//				AMS  struct {
//					Text         string `xml:",chardata"`
//					Provider     string `xml:"Provider,attr"`
//					Product      string `xml:"Product,attr"`
//					AssetName    string `xml:"Asset_Name,attr"`
//					VersionMajor string `xml:"Version_Major,attr"`
//					VersionMinor string `xml:"Version_Minor,attr"`
//					Description  string `xml:"Description,attr"`
//					CreationDate string `xml:"Creation_Date,attr"`
//					ProviderID   string `xml:"Provider_ID,attr"`
//					AssetID      string `xml:"Asset_ID,attr"`
//					AssetClass   string `xml:"Asset_Class,attr"`
//				} `xml:"AMS"`
//				AppData []struct {
//					Text  string `xml:",chardata"`
//					App   string `xml:"App,attr"`
//					Name  string `xml:"Name,attr"`
//					Value string `xml:"Value,attr"`
//				} `xml:"App_Data"`
//			} `xml:"Metadata"`
//			Content struct {
//				Text  string `xml:",chardata"`
//				Value string `xml:"Value,attr"`
//			} `xml:"Content"`
//		} `xml:"Asset"`
//	} `xml:"Asset"`
//}

//type ADI2 struct {
//	XMLName  xml.Name `xml:"ADI"`
//	//Text     string   `xml:",chardata"`
//	Metadata Metadata `xml:"Metadata"`
//}
//
//type Metadata struct {
//	//Text    string `xml:",chardata"`
//	AMS AMS     `xml:"AMS"`
//	AppData AppData `xml:"App_Data,omitempty"`
//}
//
//type AMS struct {
//	Text         string `xml:",chardata"`
//	Provider     string `xml:"Provider,attr"`
//	Product      string `xml:"Product,attr"`
//	AssetName    string `xml:"Asset_Name,attr"`
//	VersionMajor string `xml:"Version_Major,attr"`
//	VersionMinor string `xml:"Version_Minor,attr"`
//	Description  string `xml:"Description,attr"`
//	CreationDate string `xml:"Creation_Date,attr"`
//	ProviderID   string `xml:"Provider_ID,attr"`
//	AssetID      string `xml:"Asset_ID,attr"`
//	AssetClass   string `xml:"Asset_Class,attr"`
//}

type AppData struct {
	App   string `xml:"App,attr"`
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

//func (a AppData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	name := start.Name
//	attrs := start.Attr
//
//	fmt.Println(name)
//	fmt.Println(attrs)
//
//	starts := start.Copy()
//	end := start.End()
//	fmt.Println(starts)
//	fmt.Println(end)
//
//	//e.
//
//	return e.Encode(a)
//}

type AssetDistributionInterface interface {
	Construct(*etree.Element)
	Class() AssetClass
	Rank() Rank
}

type CommonFields struct {
	Provider     string
	Product      string
	AssetName    string
	VersionMajor string
	VersionMinor string
	Description  string
	ProviderID   string
	CreationDate string
	AssetID      string
}

type PackageAssetClass struct {
	*CommonFields
	Verb      string
	ClassType AssetClass
	Order     Rank
}

type TitleAssetClass struct {
	*CommonFields
	TitleBrief           string
	Title                string
	SummarySort          string
	Rating               string
	ClosedCaption        string
	RunTime              string
	DisplayRunTime       string
	Year                 string
	Category             string
	Genre                string
	ShowType             string
	LicensingWindowStart string
	LicensingWindowEnd   string
	ClassType            AssetClass
	Order                Rank
}

type MovieAssetClass struct {
	*CommonFields
	ClassType             AssetClass
	Order                 Rank
	Encryption            string
	Type                  string
	AudioType             string
	HDContent             string
	CopyProtection        string
	BitRate               string
	ContentFileSize       string
	ContentChecksum       string
	CopyProtectionVerbose string
}

type Elements []AssetDistributionInterface

func (e Elements) Len() int           { return len(e) }
func (e Elements) Less(i, j int) bool { return e[i].Rank() < e[j].Rank() }
func (e Elements) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

type Node struct {
	Elements
	Doc *etree.Document
}

func NewADI(adi ...AssetDistributionInterface) *Node {

	elements := make(Elements, 0, len(adi))
	for _, v := range adi {
		elements = append(elements, v)
	}

	node := &Node{
		Elements: elements,
		Doc:      etree.NewDocument(),
	}
	return node
}

func (n *Node) BuildDoc() *etree.Document {
	//sort.Sort(n.Elements)
	reverse := sort.Reverse(n.Elements)
	sort.Sort(reverse)

	n.Doc.CreateProcInst(Target, DefaultProcInst)
	n.Doc.CreateDirective(DefaultDocktype)
	root := n.Doc.CreateElement(ADI)

	wraps := make([]wrap, 0, len(n.Elements))

	for _, v := range n.Elements {
		if v.Class() == pckAssetClassType {
			wraps = append(wraps, applyElement(v))
		} else if v.Class() == titleAssetClassType {
			wraps = append(wraps, noop(v))
		} else {
			wraps = append(wraps, applyMultipleMetadata(v))
		}
	}

	// reverse the order
	//for i, j := 0, len(wraps)-1; i < j; i, j = i+1, j-1 {
	//	wraps[i], wraps[j] = wraps[j], wraps[i]
	//}

	compose(wraps...)(root)

	return n.Doc
}

func NewPackageAssetClass(provider, providerID, assetName, description string) *PackageAssetClass {
	class := &PackageAssetClass{
		CommonFields: new(CommonFields),
		ClassType:    pckAssetClassType,
		Order:        PckRank,
	}

	class.ProviderID = providerID
	class.Provider = provider
	class.AssetName = assetName
	class.Description = description
	return class
}

func NewTitleAssetClass(provider, providerID, assetName, description string) *TitleAssetClass {
	class := &TitleAssetClass{
		CommonFields: new(CommonFields),
		ClassType:    titleAssetClassType,
		Order:        TitleRank,
	}

	class.ProviderID = providerID
	class.Provider = provider
	class.AssetName = assetName
	class.Description = description
	return class
}

func NewMovieAssetClass(provider, providerID, assetName, description string) *MovieAssetClass {
	class := &MovieAssetClass{
		CommonFields: new(CommonFields),
		ClassType:    movieAssetClassType,
		Order:        MovieRank,
	}

	class.ProviderID = providerID
	class.Provider = provider
	class.AssetName = assetName
	class.Description = description
	return class
}

func (p *PackageAssetClass) Construct(e *etree.Element) {

	metadata := e.CreateElement(METADATA)
	amsSpec := metadata.CreateElement(AMS)

	amsSpec.CreateAttr(VersionMajor, "1")
	amsSpec.CreateAttr(VersionMinor, "0")
	amsSpec.CreateAttr(PROVICER, p.Provider)
	amsSpec.CreateAttr(ProvierId, p.ProviderID)
	amsSpec.CreateAttr(DESCRIPTION, p.Description)
	amsSpec.CreateAttr(ASSET_CLASS, string(p.ClassType))
}

func (p PackageAssetClass) Class() AssetClass {
	return p.ClassType
}

func (p PackageAssetClass) Rank() Rank {
	return p.Order
}

//type AssetElementFunc func(*etree.Element)

type AssetElementFunc func(*etree.Element) *etree.Element

func DecorateSection(fn AssetElementFunc) AssetElementFunc {
	return func(element *etree.Element) *etree.Element {
		return fn(element)
	}
}

func (m *MovieAssetClass) Construct(e *etree.Element) {

	//asset := e.CreateElement(ASSET)
	//metadata := asset.CreateElement(METADATA)
	//amsSpec := metadata.CreateElement(AMS)

	e.CreateAttr(VersionMajor, "1")
	e.CreateAttr(VersionMinor, "0")
	e.CreateAttr(PROVICER, m.Provider)
	e.CreateAttr(ProvierId, m.ProviderID)
	e.CreateAttr(DESCRIPTION, m.Description)
	e.CreateAttr(ASSET_CLASS, string(m.ClassType))
}

func (m MovieAssetClass) Class() AssetClass {
	return m.ClassType
}

func (m MovieAssetClass) Rank() Rank {
	return m.Order
}

func (t *TitleAssetClass) Construct(e *etree.Element) {

	metadata := e.CreateElement(METADATA)
	amsSpec := metadata.CreateElement(AMS)

	amsSpec.CreateAttr(VersionMajor, "1")
	amsSpec.CreateAttr(VersionMinor, "0")
	amsSpec.CreateAttr(PROVICER, t.Provider)
	amsSpec.CreateAttr(ProvierId, t.ProviderID)
	amsSpec.CreateAttr(DESCRIPTION, t.Description)
	amsSpec.CreateAttr(ASSET_CLASS, string(t.ClassType))
}

func (t TitleAssetClass) Class() AssetClass {
	return t.ClassType
}

func (t TitleAssetClass) Rank() Rank {
	return t.Order
}

type wrap func(*etree.Element) *etree.Element

func noop(a AssetDistributionInterface) wrap {
	return func(element *etree.Element) *etree.Element {
		a.Construct(element)
		return element
	}
}

func applyElement(a AssetDistributionInterface) wrap {
	return func(element *etree.Element) *etree.Element {
		a.Construct(element)
		createElement := element.CreateElement(ASSET)
		return createElement
	}
}

func applyMultipleMetadata(a AssetDistributionInterface) wrap {
	return func(element *etree.Element) *etree.Element {

		createElement := element.CreateElement(ASSET)
		metadata := createElement.CreateElement(METADATA)
		amsSpec := metadata.CreateElement(AMS)

		a.Construct(amsSpec)
		return element
	}
}

func compose(fns ...wrap) wrap {
	return func(element *etree.Element) *etree.Element {

		f := fns[0]

		fs := fns[1:]

		if len(fns) == 1 {
			return f(element)
		}
		return f(compose(fs...)(element))
	}
}
