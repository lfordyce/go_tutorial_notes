package xml_utils

import (
	"encoding/xml"
	"reflect"
)

const (
	AcquisitionPointIdentity ResponseSignalAttr = "AcquisitionPointIdentity"
	AcquisitionSignalID      ResponseSignalAttr = "AcquisitionSignalID"
	SignalPointID            ResponseSignalAttr = "SignalPointID"
	Action                   ResponseSignalAttr = "Action"

	CableLabsSignalUrn         Namespaces = "urn:cablelabs:iptvservices:esam:xsd:signal:1"
	CableLabsSignalingUrn      XmlnsAttr  = "urn:cablelabs:md:xsd:signaling:3.0"
	CableLabsCommonUrn         XmlnsAttr  = "urn:cablelabs:iptvservices:esam:xsd:common:1"
	CableLabsSchemaInstanceXsi XmlnsAttr  = "http://www.w3.org/2001/XMLSchema-instance"
)

type ResponseSignalAttr string
type Namespaces string

type SignalProcessingNotification struct {
	XMLName xml.Name  `xml:"SignalProcessingNotification"`
	Text    string    `xml:",chardata"`
	Xmlns   string    `xml:"xmlns,attr"`
	Xsi     XmlnsAttr `xml:"xsi,attr"`
	Common  XmlnsAttr `xml:"common,attr"`
	Content XmlnsAttr `xml:"content,attr"`
	Core    XmlnsAttr `xml:"core,attr"`
	Sig     XmlnsAttr `xml:"sig,attr"`
	Batch   *BatchInfo
	ResponseSignalArray
	ConditioningInfoArray
}

type BatchInfo struct {
	XMLName xml.Name `xml:"BatchInfo"`
	Text    string   `xml:",chardata"`
	BatchId string   `xml:"batchId,attr"`
	Source  *Source
}

// prefix BatchInfo tag with common:
// <common:BatchInfo batchId="....">
//		<common:Source xsi:type="content:MovieType" uriId="...."/>
//	</common:BatchInfo>
func (b *BatchInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "common:" + start.Name.Local
	return e.EncodeElement(*b, start)
}

type Source struct {
	XMLName xml.Name `xml:"Source"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
	UriId   string   `xml:"uriId,attr"`
}

// prefect Source with common:
//	<common:Source xsi:type="content:MovieType" uriId="...."/>
func (s *Source) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "common:" + start.Name.Local
	return e.EncodeElement(*s, start)
}

type ResponseSignal struct {
	XMLName                  xml.Name `xml:"ResponseSignal"`
	Text                     string   `xml:",chardata"`
	AcquisitionPointIdentity string   `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string   `xml:"acquisitionSignalID,attr"`
	SignalPointID            string   `xml:"signalPointID,attr"`
	Action                   string   `xml:"action,attr"`
	NPT                      *NPTPoint
}

type NPTPoint struct {
	XMLName    xml.Name `xml:"NPTPoint"`
	Text       string   `xml:",chardata"`
	NptPoint   string   `xml:"nptPoint,attr"`
	Descriptor *SCTE35PointDescriptor
}

func (n *NPTPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "sig:" + start.Name.Local
	return e.EncodeElement(*n, start)
}

type SCTE35PointDescriptor struct {
	XMLName           xml.Name `xml:"SCTE35PointDescriptor"`
	Text              string   `xml:",chardata"`
	SpliceCommandType string   `xml:"spliceCommandType,attr"`
	Segmentation      *SegmentationDescriptorInfo
}

func (s *SCTE35PointDescriptor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "sig:" + start.Name.Local
	return e.EncodeElement(*s, start)
}

type SegmentationDescriptorInfo struct {
	XMLName          xml.Name `xml:"SegmentationDescriptorInfo"`
	Text             string   `xml:",chardata"`
	SegmentEventId   string   `xml:"segmentEventId,attr"`
	SegmentTypeId    string   `xml:"segmentTypeId,attr"`
	UpidType         string   `xml:"upidType,attr"`
	Upid             string   `xml:"upid,attr"`
	Duration         string   `xml:"duration,attr"`
	SegmentNum       string   `xml:"segmentNum,attr"`
	SegmentsExpected string   `xml:"segmentsExpected,attr"`
}

type ConditioningInfo struct {
	XMLName                xml.Name `xml:"ConditioningInfo"`
	StartOffset            string   `xml:"startOffset,attr"`
	AcquisitionSignalIDRef string   `xml:"acquisitionSignalIDRef,attr"`
	Duration               string   `xml:"duration,attr"`
	//Segment                string   `xml:"ConditioningInfo>Segment"`
	Segment *Segment
}

type Segment struct {
	XMLName  xml.Name `xml:"Segment"`
	Duration string   `xml:",chardata"`
}

func (s *SegmentationDescriptorInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "sig:" + start.Name.Local
	return e.EncodeElement(*s, start)
}

type ResponseSignalArray struct {
	ResponseSignalList []ResponseSignal
}

type ConditioningInfoArray struct {
	ConditioningInfoList []ConditioningInfo
}

func (r *ResponseSignalArray) AddResponseSignal(attr map[ResponseSignalAttr]string, nptPoint *NPTPoint) {
	var signal ResponseSignal
	matchFieldToKey(attr, &signal)

	//signal.XMLName.Local = "ResponseSignal"
	signal.NPT = nptPoint

	r.ResponseSignalList = append(r.ResponseSignalList, signal)
}

func (c *ConditioningInfoArray) AddConditioningInfo(startOffset, acquisitionSignalIDRef, duration string) {
	info := ConditioningInfo{
		StartOffset:            startOffset,
		AcquisitionSignalIDRef: acquisitionSignalIDRef,
		Duration:               duration,
		Segment: &Segment{
			Duration: duration,
		},
	}
	info.XMLName.Local = "ConditioningInfo"
	c.ConditioningInfoList = append(c.ConditioningInfoList, info)
}

type XmlnsAttr string

func (a XmlnsAttr) MarshalXMLAttr(n xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "xmlns:" + n.Local},
		Value: string(a),
	}, nil
}

func matchFieldToKey(data map[ResponseSignalAttr]string, elem interface{}) {
	value := reflect.ValueOf(elem).Elem()
	typeOfElem := value.Type()

	for i := 0; i < value.NumField(); i++ {

		f := value.Field(i)

		//fmt.Printf("%d: %s %s = %v\n", i, typeOfElem.Field(i).Name, f.Type().Kind(), f.Interface())

		if f.Type().Kind() == reflect.String && typeOfElem.Field(i).Type.Kind() == reflect.String {

			if s, ok := data[ResponseSignalAttr(typeOfElem.Field(i).Name)]; ok {
				f.SetString(s)
			}
		}
	}
}
