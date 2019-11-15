package xsd

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"time"
)

// An ADI container element for holding anything.
type ADIContainerType struct {
	Asset []AssetType `xml:"urn:cablelabs:md:xsd:core:3.0 Asset,omitempty"`
	Ext   ExtType     `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
}

// Must match the pattern (M2T|MP4|ASF|3GP|AVI|MSSmoothStreaming|LiveStreaming|F4M|DASH|private :.+)
type AVContainerType string

type AcquiredSignal struct {
	UTCPoint                 UTCPointDescriptorType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCPoint"`
	NPTPoint                 NPTPointDescriptorType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 NPTPoint"`
	SCTE35PointDescriptor    SCTE35PointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 SCTE35PointDescriptor"`
	BinaryData               BinarySignalType          `xml:"urn:cablelabs:md:xsd:signaling:3.0 BinaryData"`
	StreamTimes              StreamTimesType           `xml:"urn:cablelabs:md:xsd:signaling:3.0 StreamTimes,omitempty"`
	Ext                      ExtType                   `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType        `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string                    `xml:"acquisitionSignalID,attr"`
	AcquisitionTime          time.Time                 `xml:"acquisitionTime,attr,omitempty"`
	SignalPointID            string                    `xml:"signalPointID,attr,omitempty"`
}

func (t *AcquiredSignal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AcquiredSignal
	var layout struct {
		*T
		AcquisitionTime *xsdDateTime `xml:"acquisitionTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AcquisitionTime = (*xsdDateTime)(&layout.T.AcquisitionTime)
	return e.EncodeElement(layout, start)
}
func (t *AcquiredSignal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AcquiredSignal
	var overlay struct {
		*T
		AcquisitionTime *xsdDateTime `xml:"acquisitionTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AcquisitionTime = (*xsdDateTime)(&overlay.T.AcquisitionTime)
	return d.DecodeElement(&overlay, &start)
}

// Acquisiton Point Info Type - information about a specific acquisiton point
type AcquisitionPointInfoType struct {
	UTCPoint                 UTCPointDescriptorType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCPoint"`
	NPTPoint                 NPTPointDescriptorType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 NPTPoint"`
	SCTE35PointDescriptor    SCTE35PointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 SCTE35PointDescriptor"`
	BinaryData               BinarySignalType          `xml:"urn:cablelabs:md:xsd:signaling:3.0 BinaryData"`
	StreamTimes              StreamTimesType           `xml:"urn:cablelabs:md:xsd:signaling:3.0 StreamTimes,omitempty"`
	Ext                      ExtType                   `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType        `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string                    `xml:"acquisitionSignalID,attr"`
	AcquisitionTime          time.Time                 `xml:"acquisitionTime,attr,omitempty"`
	SignalPointID            string                    `xml:"signalPointID,attr,omitempty"`
}

func (t *AcquisitionPointInfoType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AcquisitionPointInfoType
	var layout struct {
		*T
		AcquisitionTime *xsdDateTime `xml:"acquisitionTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AcquisitionTime = (*xsdDateTime)(&layout.T.AcquisitionTime)
	return e.EncodeElement(layout, start)
}
func (t *AcquisitionPointInfoType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AcquisitionPointInfoType
	var overlay struct {
		*T
		AcquisitionTime *xsdDateTime `xml:"acquisitionTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AcquisitionTime = (*xsdDateTime)(&overlay.T.AcquisitionTime)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern create|replace|delete|noop|private:.+
type Action string

// Must match the pattern (V|MV|GV|AL|GL|AC|SC|N|BN|RP|private:.+)
type AdvisoryType string

// Alternate Content Type - support for switching a stream to an alternate stream
type AlternateContentType struct {
	AltContentIdentity string `xml:"altContentIdentity,attr,omitempty"`
	ZoneIdentity       string `xml:"zoneIdentity,attr,omitempty"`
}

// An alternate identifier for an asset
type AlternateIdType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	IdentifierSystem   string             `xml:"identifierSystem,attr"`
}

// May be one of 0, 1, 2, 3
type AnalogProtectionSystemType byte

type ArgRefType struct {
	Variable string `xml:"variable,attr"`
}

// Name of an asset as supplied by the provider; typically a descriptive, preferably unique name that identifies and describes the asset.
type AssetNameType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	Deprecated         bool               `xml:"deprecated,attr"`
}

// An abstract base type that defines a reference to an asset.
type AssetRefBaseType struct {
	Ext   ExtType `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	UriId string  `xml:"uriId,attr"`
}

// A type that defines a reference to an asset.
type AssetRefType struct {
	Ext   ExtType `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	UriId string  `xml:"uriId,attr"`
}

// An abstract base type from which all other MD3.0 assets are derived.
type AssetType struct {
	AlternateId          []AlternateIdType  `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string             `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType      `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType        `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType    `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType            `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType       `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string             `xml:"uriId,attr"`
	ProviderVersionNum   int                `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time          `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time          `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time          `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string             `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time          `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string             `xml:"eTag,attr,omitempty"`
	State                StateType          `xml:"state,attr,omitempty"`
	StateDetail          string             `xml:"stateDetail,attr,omitempty"`
}

func (t *AssetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssetType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *AssetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssetType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern (Adult|Mature|General|Family|Teen|Children|private:.+)
type AudienceType string

// Must match the pattern (Dolby ProLogic|Dolby Digital|Stereo|Mono|Dolby 5.1|private:.+)
type AudioTypeType string

type AudioVideoAssetType struct {
	AudioType                  []AudioTypeType            `xml:"urn:cablelabs:md:xsd:content:3.0 AudioType,omitempty"`
	ScreenFormat               ScreenFormatType           `xml:"urn:cablelabs:md:xsd:content:3.0 ScreenFormat,omitempty"`
	Resolution                 ResolutionType             `xml:"urn:cablelabs:md:xsd:content:3.0 Resolution,omitempty"`
	FrameRate                  FrameRateType              `xml:"urn:cablelabs:md:xsd:content:3.0 FrameRate,omitempty"`
	Codec                      CodecType                  `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	AVContainer                AVContainerType            `xml:"urn:cablelabs:md:xsd:content:3.0 AVContainer,omitempty"`
	BitRate                    int                        `xml:"urn:cablelabs:md:xsd:content:3.0 BitRate,omitempty"`
	AlternateBitRateResolution []BitRateResolutionType    `xml:"urn:cablelabs:md:xsd:content:3.0 AlternateBitRateResolution,omitempty"`
	Duration                   string                     `xml:"urn:cablelabs:md:xsd:content:3.0 Duration,omitempty"`
	Language                   []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	SubtitleLanguage           []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 SubtitleLanguage,omitempty"`
	DubbedLanguage             []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 DubbedLanguage,omitempty"`
	Rating                     []RatingType               `xml:"urn:cablelabs:md:xsd:content:3.0 Rating,omitempty"`
	Audience                   []AudienceType             `xml:"urn:cablelabs:md:xsd:content:3.0 Audience,omitempty"`
	EncryptionInfo             EncryptionInfoType         `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionInfo,omitempty"`
	CopyControlInfo            CopyControlInfoType        `xml:"urn:cablelabs:md:xsd:content:3.0 CopyControlInfo,omitempty"`
	IsResumeEnabled            bool                       `xml:"urn:cablelabs:md:xsd:content:3.0 IsResumeEnabled,omitempty"`
	TrickModesRestricted       []TrickModeRestrictionType `xml:"urn:cablelabs:md:xsd:content:3.0 TrickModesRestricted,omitempty"`
	TrickRef                   []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 TrickRef,omitempty"`
	ThreeDMode                 int                        `xml:"urn:cablelabs:md:xsd:content:3.0 ThreeDMode,omitempty"`
	POGroupRef                 []EffectiveAssetRefType    `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef             []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl                  string                     `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize            uint64                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum            ChecksumType               `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority        int                        `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef                 string                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType                  NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId                []AlternateIdType          `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact          string                     `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName                  AssetNameType              `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product                    ProductType                `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider                   NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description                DescriptionType            `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                        ExtType                    `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef            AssetRefType               `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                      string                     `xml:"uriId,attr"`
	ProviderVersionNum         int                        `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum         int                        `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime           time.Time                  `xml:"creationDateTime,attr,omitempty"`
	StartDateTime              time.Time                  `xml:"startDateTime,attr,omitempty"`
	EndDateTime                time.Time                  `xml:"endDateTime,attr,omitempty"`
	NotifyURI                  string                     `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime       time.Time                  `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                       string                     `xml:"eTag,attr,omitempty"`
	State                      StateType                  `xml:"state,attr,omitempty"`
	StateDetail                string                     `xml:"stateDetail,attr,omitempty"`
}

func (t *AudioVideoAssetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AudioVideoAssetType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *AudioVideoAssetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AudioVideoAssetType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// See Section 10.3.1 - avail_descriptor()
type AvailDescriptorType struct {
	Ext             Ext  `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	ProviderAvailId uint `xml:"providerAvailId,attr"`
}

// See Section 9.3.5 - bandwidth_reservation()
type BandwidthReservationType struct {
	Ext Ext `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
}

type BarkerType struct {
	AudioType                  []AudioTypeType            `xml:"urn:cablelabs:md:xsd:content:3.0 AudioType,omitempty"`
	ScreenFormat               ScreenFormatType           `xml:"urn:cablelabs:md:xsd:content:3.0 ScreenFormat,omitempty"`
	Resolution                 ResolutionType             `xml:"urn:cablelabs:md:xsd:content:3.0 Resolution,omitempty"`
	FrameRate                  FrameRateType              `xml:"urn:cablelabs:md:xsd:content:3.0 FrameRate,omitempty"`
	Codec                      CodecType                  `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	AVContainer                AVContainerType            `xml:"urn:cablelabs:md:xsd:content:3.0 AVContainer,omitempty"`
	BitRate                    int                        `xml:"urn:cablelabs:md:xsd:content:3.0 BitRate,omitempty"`
	AlternateBitRateResolution []BitRateResolutionType    `xml:"urn:cablelabs:md:xsd:content:3.0 AlternateBitRateResolution,omitempty"`
	Duration                   string                     `xml:"urn:cablelabs:md:xsd:content:3.0 Duration,omitempty"`
	Language                   []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	SubtitleLanguage           []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 SubtitleLanguage,omitempty"`
	DubbedLanguage             []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 DubbedLanguage,omitempty"`
	Rating                     []RatingType               `xml:"urn:cablelabs:md:xsd:content:3.0 Rating,omitempty"`
	Audience                   []AudienceType             `xml:"urn:cablelabs:md:xsd:content:3.0 Audience,omitempty"`
	EncryptionInfo             EncryptionInfoType         `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionInfo,omitempty"`
	CopyControlInfo            CopyControlInfoType        `xml:"urn:cablelabs:md:xsd:content:3.0 CopyControlInfo,omitempty"`
	IsResumeEnabled            bool                       `xml:"urn:cablelabs:md:xsd:content:3.0 IsResumeEnabled,omitempty"`
	TrickModesRestricted       []TrickModeRestrictionType `xml:"urn:cablelabs:md:xsd:content:3.0 TrickModesRestricted,omitempty"`
	TrickRef                   []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 TrickRef,omitempty"`
	ThreeDMode                 int                        `xml:"urn:cablelabs:md:xsd:content:3.0 ThreeDMode,omitempty"`
	POGroupRef                 []EffectiveAssetRefType    `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef             []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl                  string                     `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize            uint64                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum            ChecksumType               `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority        int                        `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef                 string                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType                  NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId                []AlternateIdType          `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact          string                     `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName                  AssetNameType              `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product                    ProductType                `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider                   NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description                DescriptionType            `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                        ExtType                    `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef            AssetRefType               `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                      string                     `xml:"uriId,attr"`
	ProviderVersionNum         int                        `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum         int                        `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime           time.Time                  `xml:"creationDateTime,attr,omitempty"`
	StartDateTime              time.Time                  `xml:"startDateTime,attr,omitempty"`
	EndDateTime                time.Time                  `xml:"endDateTime,attr,omitempty"`
	NotifyURI                  string                     `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime       time.Time                  `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                       string                     `xml:"eTag,attr,omitempty"`
	State                      StateType                  `xml:"state,attr,omitempty"`
	StateDetail                string                     `xml:"stateDetail,attr,omitempty"`
}

func (t *BarkerType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BarkerType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *BarkerType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BarkerType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

type BatchInfoType struct {
	Source      MovieType   `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Source,omitempty"`
	Destination []MovieType `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Destination,omitempty"`
	Ext         ExtType     `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Ext,omitempty"`
	BatchId     string      `xml:"batchId,attr"`
}

type BinarySignalType struct {
	Value      []byte     `xml:",chardata"`
	SignalType SignalType `xml:"signalType,attr"`
}

func (t *BinarySignalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BinarySignalType
	var layout struct {
		*T
		Value *xsdBase64Binary `xml:",chardata"`
	}
	layout.T = (*T)(t)
	layout.Value = (*xsdBase64Binary)(&layout.T.Value)
	return e.EncodeElement(layout, start)
}
func (t *BinarySignalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BinarySignalType
	var overlay struct {
		*T
		Value *xsdBase64Binary `xml:",chardata"`
	}
	overlay.T = (*T)(t)
	overlay.Value = (*xsdBase64Binary)(&overlay.T.Value)
	return d.DecodeElement(&overlay, &start)
}

type BinaryType struct {
	Value      []byte     `xml:",chardata"`
	SignalType SignalType `xml:"signalType,attr,omitempty"`
}

func (t *BinaryType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BinaryType
	var layout struct {
		*T
		Value *xsdBase64Binary `xml:",chardata"`
	}
	layout.T = (*T)(t)
	layout.Value = (*xsdBase64Binary)(&layout.T.Value)
	return e.EncodeElement(layout, start)
}
func (t *BinaryType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BinaryType
	var overlay struct {
		*T
		Value      *xsdBase64Binary `xml:",chardata"`
		SignalType *SignalType      `xml:"signalType,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Value = (*xsdBase64Binary)(&overlay.T.Value)
	overlay.SignalType = (*SignalType)(&overlay.T.SignalType)
	return d.DecodeElement(&overlay, &start)
}

// Pair of bitrate/resolution values (adaptive streaming).
type BitRateResolutionType struct {
	BitRate    int            `xml:"urn:cablelabs:md:xsd:content:3.0 BitRate,omitempty"`
	Resolution ResolutionType `xml:"urn:cablelabs:md:xsd:content:3.0 Resolution,omitempty"`
}

type BoxCoverType struct {
	X_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 X_Resolution,omitempty"`
	Y_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 Y_Resolution,omitempty"`
	Language             string                  `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	Codec                ImageCodecType          `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	POGroupRef           []EffectiveAssetRefType `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef       []AssetRefType          `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl            string                  `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize      uint64                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum      ChecksumType            `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority  int                     `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef           string                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId          []AlternateIdType       `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string                  `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType           `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType             `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType         `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType                 `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType            `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string                  `xml:"uriId,attr"`
	ProviderVersionNum   int                     `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                     `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time               `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time               `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time               `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string                  `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time               `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string                  `xml:"eTag,attr,omitempty"`
	State                StateType               `xml:"state,attr,omitempty"`
	StateDetail          string                  `xml:"stateDetail,attr,omitempty"`
}

func (t *BoxCoverType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BoxCoverType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *BoxCoverType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BoxCoverType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// See Section 9.4.2 - break_duration()
type BreakDurationType struct {
	Ext        Ext    `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	AutoReturn bool   `xml:"autoReturn,attr"`
	Duration   uint64 `xml:"duration,attr"`
}

// A sequence of Bucket elements
type BucketListType struct {
	Bucket []BucketType `xml:"urn:cablelabs:md:xsd:core:3.0 Bucket,omitempty"`
}

// Bucket are logical Asset containers associated with the {ProviderId} portion of the asset uriId.
type BucketType struct {
	TotalStorage             uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 TotalStorage,omitempty"`
	AvailableStorage         uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 AvailableStorage,omitempty"`
	InputProtocols           string `xml:"urn:cablelabs:md:xsd:core:3.0 InputProtocols,omitempty"`
	TotalInputBandwidth      uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 TotalInputBandwidth,omitempty"`
	AvailableInputBandwidth  uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 AvailableInputBandwidth,omitempty"`
	OutputProtocols          string `xml:"urn:cablelabs:md:xsd:core:3.0 OutputProtocols,omitempty"`
	TotalOutputBandwidth     uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 TotalOutputBandwidth,omitempty"`
	AvailableOutputBandwidth uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 AvailableOutputBandwidth,omitempty"`
	NumAssets                uint64 `xml:"urn:cablelabs:md:xsd:core:3.0 NumAssets,omitempty"`
	ProviderId               string `xml:"providerId,attr"`
}

// May be one of 0, 1, 2, 3
type CGMSAType byte

// Must match the pattern [0-9#\*]+
type Chars string

// Must match the pattern [0-9A-Fa-f]{32}
type ChecksumType string

// Must match the pattern (MPEG2|AVC MP@L30| AVC MP@L40|AVC MP@L42|AVC HP@L30|AVC HP@L40|AVC HP@L42|MPEG4-MVC|private :.+)
type CodecType string

type Component struct {
	Ext           Ext       `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	ComponentTag  byte      `xml:"componentTag,attr"`
	UtcSpliceTime time.Time `xml:"utcSpliceTime,attr"`
}

func (t *Component) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Component
	var layout struct {
		*T
		UtcSpliceTime *xsdDateTime `xml:"utcSpliceTime,attr"`
	}
	layout.T = (*T)(t)
	layout.UtcSpliceTime = (*xsdDateTime)(&layout.T.UtcSpliceTime)
	return e.EncodeElement(layout, start)
}
func (t *Component) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Component
	var overlay struct {
		*T
		UtcSpliceTime *xsdDateTime `xml:"utcSpliceTime,attr"`
	}
	overlay.T = (*T)(t)
	overlay.UtcSpliceTime = (*xsdDateTime)(&overlay.T.UtcSpliceTime)
	return d.DecodeElement(&overlay, &start)
}

// Conditioning Into Type - Conditioning information communicated to an acquisition point. For example, communicate ABR (Adaptive Bit Rate) information.
type ConditioningInfoType struct {
	Segment                []string `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 Segment,omitempty"`
	Ext                    ExtType  `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 Ext,omitempty"`
	Duration               string   `xml:"duration,attr,omitempty"`
	AcquisitionSignalIDRef string   `xml:"acquisitionSignalIDRef,attr,omitempty"`
	StartOffset            string   `xml:"startOffset,attr,omitempty"`
}

// May be one of 0, 1
type ConstrainedImageTriggerType byte

type ContentAssetType struct {
	POGroupRef           []EffectiveAssetRefType `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef       []AssetRefType          `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl            string                  `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize      uint64                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum      ChecksumType            `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority  int                     `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef           string                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId          []AlternateIdType       `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string                  `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType           `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType             `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType         `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType                 `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType            `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string                  `xml:"uriId,attr"`
	ProviderVersionNum   int                     `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                     `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time               `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time               `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time               `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string                  `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time               `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string                  `xml:"eTag,attr,omitempty"`
	State                StateType               `xml:"state,attr,omitempty"`
	StateDetail          string                  `xml:"stateDetail,attr,omitempty"`
}

func (t *ContentAssetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ContentAssetType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *ContentAssetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ContentAssetType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// List of URI to Content on the AMS.
type ContentRefListType struct {
	ContentRef []string `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
}

type CopyControlInfoType struct {
	IsCopyProtection        bool                        `xml:"urn:cablelabs:md:xsd:content:3.0 IsCopyProtection,omitempty"`
	IsCopyProtectionVerbose bool                        `xml:"urn:cablelabs:md:xsd:content:3.0 IsCopyProtectionVerbose,omitempty"`
	AnalogProtectionSystem  AnalogProtectionSystemType  `xml:"urn:cablelabs:md:xsd:content:3.0 AnalogProtectionSystem,omitempty"`
	EncryptionModeIndicator EncryptionModeIndicatorType `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionModeIndicator,omitempty"`
	ConstrainedImageTrigger ConstrainedImageTriggerType `xml:"urn:cablelabs:md:xsd:content:3.0 ConstrainedImageTrigger,omitempty"`
	CGMS_A                  CGMSAType                   `xml:"urn:cablelabs:md:xsd:content:3.0 CGMS_A,omitempty"`
	RequiresOutputControl   bool                        `xml:"urn:cablelabs:md:xsd:content:3.0 RequiresOutputControl,omitempty"`
	Ext                     ExtType                     `xml:"urn:cablelabs:md:xsd:content:3.0 Ext,omitempty"`
}

// Must match the pattern [a-zA-Z]{2}
type CountryType string

// Must match the pattern [a-zA-Z]{3}
type CurrencyType string

// See Section 10.3.2 - DTMF_descriptor()
type DTMFDescriptorType struct {
	Ext     Ext   `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	Preroll byte  `xml:"preroll,attr,omitempty"`
	Chars   Chars `xml:"chars,attr,omitempty"`
}

// A pair of dateTime values that together represent a start and an end.
type DateTimeRangeType struct {
	StartDateTime time.Time `xml:"startDateTime,attr"`
	EndDateTime   time.Time `xml:"endDateTime,attr"`
}

func (t *DateTimeRangeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DateTimeRangeType
	var layout struct {
		*T
		StartDateTime *xsdDateTime `xml:"startDateTime,attr"`
		EndDateTime   *xsdDateTime `xml:"endDateTime,attr"`
	}
	layout.T = (*T)(t)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	return e.EncodeElement(layout, start)
}
func (t *DateTimeRangeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DateTimeRangeType
	var overlay struct {
		*T
		StartDateTime *xsdDateTime `xml:"startDateTime,attr"`
		EndDateTime   *xsdDateTime `xml:"endDateTime,attr"`
	}
	overlay.T = (*T)(t)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	return d.DecodeElement(&overlay, &start)
}

type DeliveryRestrictions struct {
	Ext                    Ext  `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	WebDeliveryAllowedFlag bool `xml:"webDeliveryAllowedFlag,attr"`
	NoRegionalBlackoutFlag bool `xml:"noRegionalBlackoutFlag,attr"`
	ArchiveAllowedFlag     bool `xml:"archiveAllowedFlag,attr"`
	DeviceRestrictions     byte `xml:"deviceRestrictions,attr"`
}

type DeprecatedBooleanType struct {
	Value      bool `xml:",chardata"`
	Deprecated bool `xml:"deprecated,attr"`
}

type DeprecatedStringType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	Deprecated         bool               `xml:"deprecated,attr"`
}

// Description of an asset.
type DescriptionType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	Deprecated         bool               `xml:"deprecated,attr"`
}

type DescriptorData []byte

func (t *DescriptorData) UnmarshalText(text []byte) error {
	return (*xsdHexBinary)(t).UnmarshalText(text)
}
func (t DescriptorData) MarshalText() ([]byte, error) {
	return xsdHexBinary(t).MarshalText()
}

// Must match the pattern ([0-9]{1,2}):[0-5][0-9](:[0-5][0-9])?
type DisplayRunTimeType string

// Identifies an Asset based on a asset reference which includes optional start/end range limits on being active, use to set the days of the week and hours within the day
type EffectiveAssetRefType struct {
	Ext           ExtType   `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	Order         uint      `xml:"order,attr,omitempty"`
	StartDateTime time.Time `xml:"startDateTime,attr,omitempty"`
	EndDateTime   time.Time `xml:"endDateTime,attr,omitempty"`
	Mon           bool      `xml:"mon,attr,omitempty"`
	Tue           bool      `xml:"tue,attr,omitempty"`
	Wed           bool      `xml:"wed,attr,omitempty"`
	Thu           bool      `xml:"thu,attr,omitempty"`
	Fri           bool      `xml:"fri,attr,omitempty"`
	Sat           bool      `xml:"sat,attr,omitempty"`
	Sun           bool      `xml:"sun,attr,omitempty"`
	StartTime     time.Time `xml:"startTime,attr,omitempty"`
	Duration      string    `xml:"duration,attr,omitempty"`
	UriId         string    `xml:"uriId,attr"`
}

func (t *EffectiveAssetRefType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EffectiveAssetRefType
	var layout struct {
		*T
		StartDateTime *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime   *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		StartTime     *xsdTime     `xml:"startTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.StartTime = (*xsdTime)(&layout.T.StartTime)
	return e.EncodeElement(layout, start)
}
func (t *EffectiveAssetRefType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EffectiveAssetRefType
	var overlay struct {
		*T
		StartDateTime *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime   *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		Mon           *bool        `xml:"mon,attr,omitempty"`
		Tue           *bool        `xml:"tue,attr,omitempty"`
		Wed           *bool        `xml:"wed,attr,omitempty"`
		Thu           *bool        `xml:"thu,attr,omitempty"`
		Fri           *bool        `xml:"fri,attr,omitempty"`
		Sat           *bool        `xml:"sat,attr,omitempty"`
		Sun           *bool        `xml:"sun,attr,omitempty"`
		StartTime     *xsdTime     `xml:"startTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.Mon = (*bool)(&overlay.T.Mon)
	overlay.Tue = (*bool)(&overlay.T.Tue)
	overlay.Wed = (*bool)(&overlay.T.Wed)
	overlay.Thu = (*bool)(&overlay.T.Thu)
	overlay.Fri = (*bool)(&overlay.T.Fri)
	overlay.Sat = (*bool)(&overlay.T.Sat)
	overlay.Sun = (*bool)(&overlay.T.Sun)
	overlay.StartTime = (*xsdTime)(&overlay.T.StartTime)
	return d.DecodeElement(&overlay, &start)
}

type EncryptedPacket struct {
	Ext                 Ext  `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	EncryptionAlgorithm byte `xml:"encryptionAlgorithm,attr"`
	CwIndex             byte `xml:"cwIndex,attr"`
}

// Describes an encrypted asset.
type EncryptionInfoType struct {
	VendorName           NonEmptyStringType `xml:"urn:cablelabs:md:xsd:content:3.0 VendorName,omitempty"`
	ReceiverType         NonEmptyStringType `xml:"urn:cablelabs:md:xsd:content:3.0 ReceiverType"`
	ReceiverVersion      uint               `xml:"urn:cablelabs:md:xsd:content:3.0 ReceiverVersion,omitempty"`
	Encryption           NonEmptyStringType `xml:"urn:cablelabs:md:xsd:content:3.0 Encryption,omitempty"`
	EncryptionAlgorithm  NonEmptyStringType `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionAlgorithm,omitempty"`
	EncryptionDateTime   time.Time          `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionDateTime,omitempty"`
	EncryptionSystemInfo NonEmptyStringType `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionSystemInfo,omitempty"`
	EncryptionKeyBlock   NonEmptyStringType `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionKeyBlock,omitempty"`
	Ext                  ExtType            `xml:"urn:cablelabs:md:xsd:content:3.0 Ext,omitempty"`
}

func (t *EncryptionInfoType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EncryptionInfoType
	var layout struct {
		*T
		EncryptionDateTime *xsdDateTime `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionDateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EncryptionDateTime = (*xsdDateTime)(&layout.T.EncryptionDateTime)
	return e.EncodeElement(layout, start)
}
func (t *EncryptionInfoType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EncryptionInfoType
	var overlay struct {
		*T
		EncryptionDateTime *xsdDateTime `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionDateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EncryptionDateTime = (*xsdDateTime)(&overlay.T.EncryptionDateTime)
	return d.DecodeElement(&overlay, &start)
}

// May be one of 0, 1, 2, 3
type EncryptionModeIndicatorType byte

// A list of errors that occured.
type ErrorListType struct {
	Error []ErrorType `xml:"urn:cablelabs:md:xsd:core:3.0 Error,omitempty"`
}

// An error code and error message.
type ErrorType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	Code               uint64             `xml:"code,attr,omitempty"`
	Lang               Lang               `xml:"lang,attr,omitempty"`
}

type Event struct {
	Ext                        Ext               `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	Program                    Program           `xml:"http://www.scte.org/schemas/35/2016 Program"`
	Component                  []Component       `xml:"http://www.scte.org/schemas/35/2016 Component"`
	BreakDuration              BreakDurationType `xml:"http://www.scte.org/schemas/35/2016 BreakDuration,omitempty"`
	SpliceEventId              uint              `xml:"spliceEventId,attr,omitempty"`
	SpliceEventCancelIndicator bool              `xml:"spliceEventCancelIndicator,attr,omitempty"`
	OutOfNetworkIndicator      bool              `xml:"outOfNetworkIndicator,attr,omitempty"`
	UniqueProgramId            uint              `xml:"uniqueProgramId,attr,omitempty"`
	AvailNum                   byte              `xml:"availNum,attr,omitempty"`
	AvailsExpected             byte              `xml:"availsExpected,attr,omitempty"`
}

func (t *Event) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Event
	var overlay struct {
		*T
		SpliceEventCancelIndicator *bool `xml:"spliceEventCancelIndicator,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SpliceEventCancelIndicator = (*bool)(&overlay.T.SpliceEventCancelIndicator)
	return d.DecodeElement(&overlay, &start)
}

// Event Schedule Type - support insertion of a signal on a repetitive and/or scheduled basis.
type EventScheduleType struct {
	StartUTC    UTCPointDescriptorType `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 StartUTC"`
	StopUTC     UTCPointDescriptorType `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 StopUTC,omitempty"`
	StartOffset string                 `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 StartOffset"`
	StopOffset  string                 `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 StopOffset,omitempty"`
	Ext         ExtType                `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 Ext,omitempty"`
	Interval    string                 `xml:"interval,attr,omitempty"`
}

type Ext struct {
	Items []string `xml:",any"`
}

// This type may contain elements or attributes from any namespace and is provided for future extensibility.
type ExtType struct {
	Items []string `xml:",any"`
}

// Must match the pattern (24|30|60|private:\d+)
type FrameRateType string

// Must match the pattern (JPG|BMP|GIF|TIF|PNG|private:.+)
type ImageCodecType string

type Lang string

// A string which can be specified in multiple lanaguages for localization.
type LocalizableStringType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	Lang               Lang               `xml:"lang,attr,omitempty"`
}

// Must match the pattern gt|gteq|private:.+
type LowerTestType string

type MovieType struct {
	AudioType                  []AudioTypeType            `xml:"urn:cablelabs:md:xsd:content:3.0 AudioType,omitempty"`
	ScreenFormat               ScreenFormatType           `xml:"urn:cablelabs:md:xsd:content:3.0 ScreenFormat,omitempty"`
	Resolution                 ResolutionType             `xml:"urn:cablelabs:md:xsd:content:3.0 Resolution,omitempty"`
	FrameRate                  FrameRateType              `xml:"urn:cablelabs:md:xsd:content:3.0 FrameRate,omitempty"`
	Codec                      CodecType                  `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	AVContainer                AVContainerType            `xml:"urn:cablelabs:md:xsd:content:3.0 AVContainer,omitempty"`
	BitRate                    int                        `xml:"urn:cablelabs:md:xsd:content:3.0 BitRate,omitempty"`
	AlternateBitRateResolution []BitRateResolutionType    `xml:"urn:cablelabs:md:xsd:content:3.0 AlternateBitRateResolution,omitempty"`
	Duration                   string                     `xml:"urn:cablelabs:md:xsd:content:3.0 Duration,omitempty"`
	Language                   []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	SubtitleLanguage           []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 SubtitleLanguage,omitempty"`
	DubbedLanguage             []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 DubbedLanguage,omitempty"`
	Rating                     []RatingType               `xml:"urn:cablelabs:md:xsd:content:3.0 Rating,omitempty"`
	Audience                   []AudienceType             `xml:"urn:cablelabs:md:xsd:content:3.0 Audience,omitempty"`
	EncryptionInfo             EncryptionInfoType         `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionInfo,omitempty"`
	CopyControlInfo            CopyControlInfoType        `xml:"urn:cablelabs:md:xsd:content:3.0 CopyControlInfo,omitempty"`
	IsResumeEnabled            bool                       `xml:"urn:cablelabs:md:xsd:content:3.0 IsResumeEnabled,omitempty"`
	TrickModesRestricted       []TrickModeRestrictionType `xml:"urn:cablelabs:md:xsd:content:3.0 TrickModesRestricted,omitempty"`
	TrickRef                   []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 TrickRef,omitempty"`
	ThreeDMode                 int                        `xml:"urn:cablelabs:md:xsd:content:3.0 ThreeDMode,omitempty"`
	POGroupRef                 []EffectiveAssetRefType    `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef             []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl                  string                     `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize            uint64                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum            ChecksumType               `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority        int                        `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef                 string                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType                  NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId                []AlternateIdType          `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact          string                     `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName                  AssetNameType              `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product                    ProductType                `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider                   NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description                DescriptionType            `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                        ExtType                    `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef            AssetRefType               `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                      string                     `xml:"uriId,attr"`
	ProviderVersionNum         int                        `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum         int                        `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime           time.Time                  `xml:"creationDateTime,attr,omitempty"`
	StartDateTime              time.Time                  `xml:"startDateTime,attr,omitempty"`
	EndDateTime                time.Time                  `xml:"endDateTime,attr,omitempty"`
	NotifyURI                  string                     `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime       time.Time                  `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                       string                     `xml:"eTag,attr,omitempty"`
	State                      StateType                  `xml:"state,attr,omitempty"`
	StateDetail                string                     `xml:"stateDetail,attr,omitempty"`
}

func (t *MovieType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MovieType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *MovieType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MovieType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// The date portion of NPT postition expressed as a date
type NPTDateType time.Time

func (t *NPTDateType) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t NPTDateType) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

// The NPT for a point that specifies a point of interest or the start
// or end point of a region. NPT always has an nptPoint but may also include a date
// constraint on the NPT. The date is useful when defining NPT offsets into what was
// previously a live stream
type NPTPointDescriptorType struct {
	Ext      ExtType      `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	NptDate  NPTDateType  `xml:"nptDate,attr,omitempty"`
	NptPoint NPTPointType `xml:"nptPoint,attr"`
}

// Must match the pattern [0-9]*\.[0-9]{3}|BOS|EOS
type NPTPointType string

// Must be at least 1 items long
type NonEmptyStringType string

// Base type describing a person, useful for actors, singers, producers, directors, etc.
type PersonType struct {
	FirstName    string `xml:"firstName,attr"`
	LastName     string `xml:"lastName,attr,omitempty"`
	SortableName string `xml:"sortableName,attr,omitempty"`
	FullName     string `xml:"fullName,attr"`
}

type PosterType struct {
	X_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 X_Resolution,omitempty"`
	Y_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 Y_Resolution,omitempty"`
	Language             string                  `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	Codec                ImageCodecType          `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	POGroupRef           []EffectiveAssetRefType `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef       []AssetRefType          `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl            string                  `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize      uint64                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum      ChecksumType            `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority  int                     `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef           string                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId          []AlternateIdType       `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string                  `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType           `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType             `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType         `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType                 `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType            `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string                  `xml:"uriId,attr"`
	ProviderVersionNum   int                     `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                     `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time               `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time               `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time               `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string                  `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time               `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string                  `xml:"eTag,attr,omitempty"`
	State                StateType               `xml:"state,attr,omitempty"`
	StateDetail          string                  `xml:"stateDetail,attr,omitempty"`
}

func (t *PosterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PosterType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *PosterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PosterType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

type PreviewType struct {
	AudioType                  []AudioTypeType            `xml:"urn:cablelabs:md:xsd:content:3.0 AudioType,omitempty"`
	ScreenFormat               ScreenFormatType           `xml:"urn:cablelabs:md:xsd:content:3.0 ScreenFormat,omitempty"`
	Resolution                 ResolutionType             `xml:"urn:cablelabs:md:xsd:content:3.0 Resolution,omitempty"`
	FrameRate                  FrameRateType              `xml:"urn:cablelabs:md:xsd:content:3.0 FrameRate,omitempty"`
	Codec                      CodecType                  `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	AVContainer                AVContainerType            `xml:"urn:cablelabs:md:xsd:content:3.0 AVContainer,omitempty"`
	BitRate                    int                        `xml:"urn:cablelabs:md:xsd:content:3.0 BitRate,omitempty"`
	AlternateBitRateResolution []BitRateResolutionType    `xml:"urn:cablelabs:md:xsd:content:3.0 AlternateBitRateResolution,omitempty"`
	Duration                   string                     `xml:"urn:cablelabs:md:xsd:content:3.0 Duration,omitempty"`
	Language                   []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	SubtitleLanguage           []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 SubtitleLanguage,omitempty"`
	DubbedLanguage             []string                   `xml:"urn:cablelabs:md:xsd:content:3.0 DubbedLanguage,omitempty"`
	Rating                     []RatingType               `xml:"urn:cablelabs:md:xsd:content:3.0 Rating,omitempty"`
	Audience                   []AudienceType             `xml:"urn:cablelabs:md:xsd:content:3.0 Audience,omitempty"`
	EncryptionInfo             EncryptionInfoType         `xml:"urn:cablelabs:md:xsd:content:3.0 EncryptionInfo,omitempty"`
	CopyControlInfo            CopyControlInfoType        `xml:"urn:cablelabs:md:xsd:content:3.0 CopyControlInfo,omitempty"`
	IsResumeEnabled            bool                       `xml:"urn:cablelabs:md:xsd:content:3.0 IsResumeEnabled,omitempty"`
	TrickModesRestricted       []TrickModeRestrictionType `xml:"urn:cablelabs:md:xsd:content:3.0 TrickModesRestricted,omitempty"`
	TrickRef                   []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 TrickRef,omitempty"`
	ThreeDMode                 int                        `xml:"urn:cablelabs:md:xsd:content:3.0 ThreeDMode,omitempty"`
	POGroupRef                 []EffectiveAssetRefType    `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef             []AssetRefType             `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl                  string                     `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize            uint64                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum            ChecksumType               `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority        int                        `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef                 string                     `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType                  NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId                []AlternateIdType          `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact          string                     `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName                  AssetNameType              `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product                    ProductType                `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider                   NonEmptyStringType         `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description                DescriptionType            `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                        ExtType                    `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef            AssetRefType               `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                      string                     `xml:"uriId,attr"`
	ProviderVersionNum         int                        `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum         int                        `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime           time.Time                  `xml:"creationDateTime,attr,omitempty"`
	StartDateTime              time.Time                  `xml:"startDateTime,attr,omitempty"`
	EndDateTime                time.Time                  `xml:"endDateTime,attr,omitempty"`
	NotifyURI                  string                     `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime       time.Time                  `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                       string                     `xml:"eTag,attr,omitempty"`
	State                      StateType                  `xml:"state,attr,omitempty"`
	StateDetail                string                     `xml:"stateDetail,attr,omitempty"`
}

func (t *PreviewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PreviewType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *PreviewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PreviewType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// A decimal value for price combined with the corresponding currency attribute.
type PriceType struct {
	Value    float64      `xml:",chardata"`
	Currency CurrencyType `xml:"currency,attr,omitempty"`
	Retail   float64      `xml:"retail,attr,omitempty"`
}

// See Section 9.3.6 - private_command()
type PrivateCommandType struct {
	PrivateBytes []byte `xml:"http://www.scte.org/schemas/35/2016 PrivateBytes,omitempty"`
	Ext          Ext    `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	Identifier   uint   `xml:"identifier,attr"`
}

func (t *PrivateCommandType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PrivateCommandType
	var layout struct {
		*T
		PrivateBytes *xsdHexBinary `xml:"http://www.scte.org/schemas/35/2016 PrivateBytes,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PrivateBytes = (*xsdHexBinary)(&layout.T.PrivateBytes)
	return e.EncodeElement(layout, start)
}
func (t *PrivateCommandType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PrivateCommandType
	var overlay struct {
		*T
		PrivateBytes *xsdHexBinary `xml:"http://www.scte.org/schemas/35/2016 PrivateBytes,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PrivateBytes = (*xsdHexBinary)(&overlay.T.PrivateBytes)
	return d.DecodeElement(&overlay, &start)
}

type ProcessRuleSelectorType struct {
	Arg            []ArgRefType `xml:"urn:cablelabs:md:xsd:core:3.0 Arg,omitempty"`
	Ext            []ExtType    `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	RuleId         string       `xml:"ruleId,attr"`
	RuleRepository string       `xml:"ruleRepository,attr"`
}

type ProcessStatusAcknowledgementType struct {
	StatusCode               StatusCodeType     `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 StatusCode,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string             `xml:"acquisitionSignalID,attr,omitempty"`
	BatchId                  string             `xml:"batchId,attr,omitempty"`
}

type ProcessStatusNotificationType struct {
	StatusCode               StatusCodeType     `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 StatusCode,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string             `xml:"acquisitionSignalID,attr,omitempty"`
	BatchId                  string             `xml:"batchId,attr,omitempty"`
}

type ProcessStatusRequestType struct {
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string             `xml:"acquisitionSignalID,attr,omitempty"`
	BatchId                  string             `xml:"batchId,attr,omitempty"`
}

type ProcessStatusResponseType struct {
	StatusCode               StatusCodeType     `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 StatusCode,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string             `xml:"acquisitionSignalID,attr,omitempty"`
	BatchId                  string             `xml:"batchId,attr"`
}

// Processing Notification Type - message to acquisition point to direct processing
type ProcessingNotificationType struct {
	BatchInfo                BatchInfoType      `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 BatchInfo,omitempty"`
	StatusCode               StatusCodeType     `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 StatusCode,omitempty"`
	Ext                      ExtType            `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Ext,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr,omitempty"`
}

// The types of product that is delivered through these terms.
type ProductType struct {
	NonEmptyStringType NonEmptyStringType `xml:",chardata"`
	Deprecated         bool               `xml:"deprecated,attr"`
}

type Program struct {
	Ext           Ext       `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	UtcSpliceTime time.Time `xml:"utcSpliceTime,attr"`
}

func (t *Program) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Program
	var layout struct {
		*T
		UtcSpliceTime *xsdDateTime `xml:"utcSpliceTime,attr"`
	}
	layout.T = (*T)(t)
	layout.UtcSpliceTime = (*xsdDateTime)(&layout.T.UtcSpliceTime)
	return e.EncodeElement(layout, start)
}
func (t *Program) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Program
	var overlay struct {
		*T
		UtcSpliceTime *xsdDateTime `xml:"utcSpliceTime,attr"`
	}
	overlay.T = (*T)(t)
	overlay.UtcSpliceTime = (*xsdDateTime)(&overlay.T.UtcSpliceTime)
	return d.DecodeElement(&overlay, &start)
}

// Content rating combined with the corresponding rating system attribute.
type RatingType struct {
	Value        string `xml:",chardata"`
	RatingSystem string `xml:"ratingSystem,attr,omitempty"`
}

// Must match the pattern (480i|720p|1080i|1080p|private:.+)
type ResolutionType string

// Response Signal Type - extension of AcquisitionPointInfoType from the signaling schema to support actions to take
type ResponseSignalType struct {
	EventSchedule            EventScheduleType         `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 EventSchedule,omitempty"`
	AlternateContent         []AlternateContentType    `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 AlternateContent,omitempty"`
	UTCPoint                 UTCPointDescriptorType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCPoint"`
	NPTPoint                 NPTPointDescriptorType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 NPTPoint"`
	SCTE35PointDescriptor    SCTE35PointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 SCTE35PointDescriptor"`
	BinaryData               BinarySignalType          `xml:"urn:cablelabs:md:xsd:signaling:3.0 BinaryData"`
	StreamTimes              StreamTimesType           `xml:"urn:cablelabs:md:xsd:signaling:3.0 StreamTimes,omitempty"`
	Ext                      ExtType                   `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	Action                   Action                    `xml:"action,attr,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType        `xml:"acquisitionPointIdentity,attr"`
	AcquisitionSignalID      string                    `xml:"acquisitionSignalID,attr"`
	AcquisitionTime          time.Time                 `xml:"acquisitionTime,attr,omitempty"`
	SignalPointID            string                    `xml:"signalPointID,attr,omitempty"`
}

func (t *ResponseSignalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ResponseSignalType
	var layout struct {
		*T
		AcquisitionTime *xsdDateTime `xml:"acquisitionTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AcquisitionTime = (*xsdDateTime)(&layout.T.AcquisitionTime)
	return e.EncodeElement(layout, start)
}
func (t *ResponseSignalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ResponseSignalType
	var overlay struct {
		*T
		AcquisitionTime *xsdDateTime `xml:"acquisitionTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AcquisitionTime = (*xsdDateTime)(&overlay.T.AcquisitionTime)
	return d.DecodeElement(&overlay, &start)
}

// Type definition for an SCTE 35 splice info section.
type SCTE35PointDescriptorType struct {
	SpliceInsert               SpliceInsertType             `xml:"urn:cablelabs:md:xsd:signaling:3.0 SpliceInsert,omitempty"`
	AvailDescriptorInfo        []AvailDescriptorType        `xml:"urn:cablelabs:md:xsd:signaling:3.0 AvailDescriptorInfo,omitempty"`
	DTMFDescriptorInfo         []DTMFDescriptorType         `xml:"urn:cablelabs:md:xsd:signaling:3.0 DTMFDescriptorInfo,omitempty"`
	SegmentationDescriptorInfo []SegmentationDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 SegmentationDescriptorInfo,omitempty"`
	UniqueDescriptorInfo       []UniqueDescriptorType       `xml:"urn:cablelabs:md:xsd:signaling:3.0 UniqueDescriptorInfo,omitempty"`
	Ext                        []ExtType                    `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	SpliceCommandType          uint                         `xml:"spliceCommandType,attr"`
}

// Must match the pattern [0-9]*
type SCTE35PointType string

// Must match the pattern [0-9]{1,2}:[0-5][0-9]:[0-5][0-9]:[0-9]{1,2}
type SMPTETimeType string

// Must match the pattern (Standard|Widescreen|Letterbox|OAR|private:.+)
type ScreenFormatType string

// See Section 10.3.3 - segmentation_descriptor()
type SegmentationDescriptorType struct {
	DeliveryRestrictions             DeliveryRestrictions   `xml:"http://www.scte.org/schemas/35/2016 DeliveryRestrictions,omitempty"`
	SegmentationUpid                 []SegmentationUpidType `xml:"http://www.scte.org/schemas/35/2016 SegmentationUpid,omitempty"`
	Component                        []_anon3               `xml:"http://www.scte.org/schemas/35/2016 Component,omitempty"`
	Ext                              Ext                    `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	SegmentationEventId              uint                   `xml:"segmentationEventId,attr,omitempty"`
	SegmentationEventCancelIndicator bool                   `xml:"segmentationEventCancelIndicator,attr,omitempty"`
	SegmentationDuration             uint64                 `xml:"segmentationDuration,attr,omitempty"`
	SegmentationTypeId               byte                   `xml:"segmentationTypeId,attr,omitempty"`
	SegmentNum                       byte                   `xml:"segmentNum,attr,omitempty"`
	SegmentsExpected                 byte                   `xml:"segmentsExpected,attr,omitempty"`
	SubSegmentNum                    byte                   `xml:"subSegmentNum,attr,omitempty"`
	SubSegmentsExpected              byte                   `xml:"subSegmentsExpected,attr,omitempty"`
}

func (t *SegmentationDescriptorType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SegmentationDescriptorType
	var overlay struct {
		*T
		SegmentationEventCancelIndicator *bool `xml:"segmentationEventCancelIndicator,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SegmentationEventCancelIndicator = (*bool)(&overlay.T.SegmentationEventCancelIndicator)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern text|hexbinary|base-64|private:.+
type SegmentationUpidFormat string

// See Section 11.4 - segmentation_upid()
type SegmentationUpidType struct {
	Value                  string                 `xml:",chardata"`
	SegmentationUpidType   byte                   `xml:"segmentationUpidType,attr,omitempty"`
	FormatIdentifier       uint                   `xml:"formatIdentifier,attr,omitempty"`
	SegmentationUpidFormat SegmentationUpidFormat `xml:"segmentationUpidFormat,attr,omitempty"`
}

// Must match the pattern (Series|Sports|Music|Ad|Miniseries|Movie|Kids|Events|Lifestyle|Other|Paid Programming|Barker|private:.+)
type ShowTypeType string

// SignalGroupAssetType is an extension of the CableLabs 3 AssetType.  Implementers which do not deal with CableLabs Assets may safely ignore this type.
type SignalGroupAssetType struct {
	SignalPoint          []SignalPointType  `xml:"urn:cablelabs:md:xsd:signaling:3.0 SignalPoint,omitempty"`
	SignalRegion         []SignalRegionType `xml:"urn:cablelabs:md:xsd:signaling:3.0 SignalRegion,omitempty"`
	AlternateId          []AlternateIdType  `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string             `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType      `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType        `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType    `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType            `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType       `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string             `xml:"uriId,attr"`
	ProviderVersionNum   int                `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time          `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time          `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time          `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string             `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time          `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string             `xml:"eTag,attr,omitempty"`
	State                StateType          `xml:"state,attr,omitempty"`
	StateDetail          string             `xml:"stateDetail,attr,omitempty"`
}

func (t *SignalGroupAssetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SignalGroupAssetType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *SignalGroupAssetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SignalGroupAssetType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Specify a group of signal points and/or regions.
type SignalGroupType struct {
	SignalPoint  []SignalPointType  `xml:"urn:cablelabs:md:xsd:signaling:3.0 SignalPoint,omitempty"`
	SignalRegion []SignalRegionType `xml:"urn:cablelabs:md:xsd:signaling:3.0 SignalRegion,omitempty"`
	Ext          ExtType            `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
}

// Based type for signals
type SignalPointType struct {
	AlternateId        []string               `xml:"urn:cablelabs:md:xsd:signaling:3.0 AlternateId,omitempty"`
	NPTPointDescriptor NPTPointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 NPTPointDescriptor,omitempty"`
	SignaledPointInfo  SignaledPointInfoType  `xml:"urn:cablelabs:md:xsd:signaling:3.0 SignaledPointInfo,omitempty"`
	Ext                ExtType                `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	SignalPointID      string                 `xml:"signalPointID,attr,omitempty"`
}

// Signal Processing Event Type - Type to carry one or more acquired signals across an interface
type SignalProcessingEventType struct {
	AcquiredSignal []AcquiredSignal `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 AcquiredSignal"`
	Ext            ExtType          `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 Ext,omitempty"`
}

// Signal Processing Notification Type - message to acquisition point to direct processing of signals
type SignalProcessingNotificationType struct {
	ResponseSignal           []ResponseSignalType   `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 ResponseSignal,omitempty"`
	ConditioningInfo         []ConditioningInfoType `xml:"urn:cablelabs:iptvservices:esam:xsd:signal:1 ConditioningInfo,omitempty"`
	BatchInfo                BatchInfoType          `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 BatchInfo,omitempty"`
	StatusCode               StatusCodeType         `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 StatusCode,omitempty"`
	Ext                      ExtType                `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Ext,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType     `xml:"acquisitionPointIdentity,attr,omitempty"`
}

// Type definition for a region of interest. The region can be defined by reference or by fully describing each point that bounds the region. The End Point is optional since not all End Points are signaled (ex. SCTE 35 out point signals only)
type SignalRegionType struct {
	AlternateId    []string        `xml:"urn:cablelabs:md:xsd:signaling:3.0 AlternateId,omitempty"`
	StartPoint     SignalPointType `xml:"urn:cablelabs:md:xsd:signaling:3.0 StartPoint"`
	EndPoint       SignalPointType `xml:"urn:cablelabs:md:xsd:signaling:3.0 EndPoint,omitempty"`
	StartPointRef  AssetRefType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 StartPointRef"`
	EndPointRef    AssetRefType    `xml:"urn:cablelabs:md:xsd:signaling:3.0 EndPointRef,omitempty"`
	Ext            ExtType         `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	SignalRegionID string          `xml:"signalRegionID,attr,omitempty"`
	Duration       string          `xml:"duration,attr,omitempty"`
}

type SignalStateRequest struct {
	Ext                      ExtType            `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Ext,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr"`
	UriId                    string             `xml:"uriId,attr"`
}

// Must match the pattern SpliceInfoSection|private:.+
type SignalType string

type SignalValidityTimeRange struct {
	UTCStart UTCPointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCStart"`
	UTCEnd   UTCPointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCEnd"`
	Ext      ExtType                `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	Order    uint                   `xml:"order,attr,omitempty"`
}

// Specifies a bounded interval of time when a signal should be considered valid. If a signal arrives outside the valid time range it shall not be considered valid.
type SignalValidityTimeRangeType struct {
	UTCStart UTCPointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCStart"`
	UTCEnd   UTCPointDescriptorType `xml:"urn:cablelabs:md:xsd:signaling:3.0 UTCEnd"`
	Ext      ExtType                `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	Order    uint                   `xml:"order,attr,omitempty"`
}

// Specify information about a signaled point in a stream. The information may specfy information about an anticipated signal or be populated after the signal arrives.
type SignaledPointInfoType struct {
	SignalValidityTimeRange SignalValidityTimeRange    `xml:"urn:cablelabs:md:xsd:signaling:3.0 SignalValidityTimeRange,omitempty"`
	SCTE35PointDescriptor   SCTE35PointDescriptorType  `xml:"urn:cablelabs:md:xsd:signaling:3.0 SCTE35PointDescriptor,omitempty"`
	StreamTimes             StreamTimesType            `xml:"urn:cablelabs:md:xsd:signaling:3.0 StreamTimes,omitempty"`
	AcquisitionPointInfo    []AcquisitionPointInfoType `xml:"urn:cablelabs:md:xsd:signaling:3.0 AcquisitionPointInfo,omitempty"`
	Ext                     ExtType                    `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	AcquisitionSignalID     string                     `xml:"acquisitionSignalID,attr,omitempty"`
}

// Must match the pattern [+-]?(100|([0-9]{1,2})(\.[0-9]{1,2}))(,[+-]?(100|([0-9]{1,2})(\.[0-9]{1,2})))?
type SpeedScaleType string

// See Section 9.3 - Splice Commands
type SpliceCommandType struct {
	Ext Ext `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
}

type SpliceDescriptorType struct {
	Ext Ext `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
}

// See Section 9.2 - Splice Info Section
type SpliceInfoSectionType struct {
	Ext                    Ext                        `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	EncryptedPacket        EncryptedPacket            `xml:"http://www.scte.org/schemas/35/2016 EncryptedPacket,omitempty"`
	SpliceNull             SpliceNullType             `xml:"http://www.scte.org/schemas/35/2016 SpliceNull"`
	SpliceSchedule         SpliceScheduleType         `xml:"http://www.scte.org/schemas/35/2016 SpliceSchedule"`
	SpliceInsert           SpliceInsertType           `xml:"http://www.scte.org/schemas/35/2016 SpliceInsert"`
	TimeSignal             TimeSignalType             `xml:"http://www.scte.org/schemas/35/2016 TimeSignal"`
	BandwidthReservation   BandwidthReservationType   `xml:"http://www.scte.org/schemas/35/2016 BandwidthReservation"`
	PrivateCommand         PrivateCommandType         `xml:"http://www.scte.org/schemas/35/2016 PrivateCommand"`
	AvailDescriptor        AvailDescriptorType        `xml:"http://www.scte.org/schemas/35/2016 AvailDescriptor"`
	DTMFDescriptor         DTMFDescriptorType         `xml:"http://www.scte.org/schemas/35/2016 DTMFDescriptor"`
	SegmentationDescriptor SegmentationDescriptorType `xml:"http://www.scte.org/schemas/35/2016 SegmentationDescriptor"`
	TimeDescriptor         TimeDescriptorType         `xml:"http://www.scte.org/schemas/35/2016 TimeDescriptor"`
	PtsAdjustment          uint64                     `xml:"ptsAdjustment,attr,omitempty"`
	ProtocolVersion        byte                       `xml:"protocolVersion,attr,omitempty"`
	Tier                   uint                       `xml:"tier,attr,omitempty"`
}

func (t *SpliceInfoSectionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SpliceInfoSectionType
	var overlay struct {
		*T
		PtsAdjustment *uint64 `xml:"ptsAdjustment,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PtsAdjustment = (*uint64)(&overlay.T.PtsAdjustment)
	return d.DecodeElement(&overlay, &start)
}

// See Section 9.3.3 - splice_insert().
type SpliceInsertType struct {
	Program                    _anon1            `xml:"http://www.scte.org/schemas/35/2016 Program"`
	Component                  []_anon2          `xml:"http://www.scte.org/schemas/35/2016 Component"`
	BreakDuration              BreakDurationType `xml:"http://www.scte.org/schemas/35/2016 BreakDuration,omitempty"`
	Ext                        Ext               `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	SpliceEventId              uint              `xml:"spliceEventId,attr,omitempty"`
	SpliceEventCancelIndicator bool              `xml:"spliceEventCancelIndicator,attr,omitempty"`
	OutOfNetworkIndicator      bool              `xml:"outOfNetworkIndicator,attr,omitempty"`
	SpliceImmediateFlag        bool              `xml:"spliceImmediateFlag,attr,omitempty"`
	UniqueProgramId            uint              `xml:"uniqueProgramId,attr,omitempty"`
	AvailNum                   byte              `xml:"availNum,attr,omitempty"`
	AvailsExpected             byte              `xml:"availsExpected,attr,omitempty"`
}

func (t *SpliceInsertType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SpliceInsertType
	var overlay struct {
		*T
		SpliceEventCancelIndicator *bool `xml:"spliceEventCancelIndicator,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SpliceEventCancelIndicator = (*bool)(&overlay.T.SpliceEventCancelIndicator)
	return d.DecodeElement(&overlay, &start)
}

// See Section 9.3.1 - splice_null()
type SpliceNullType struct {
	Ext Ext `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
}

// See Section 9.3.2 - splice_schedule()
type SpliceScheduleType struct {
	Event []Event `xml:"http://www.scte.org/schemas/35/2016 Event,omitempty"`
	Ext   Ext     `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
}

// See Section 9.4.1 - splice_time()
type SpliceTimeType struct {
	Ext     Ext    `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	PtsTime uint64 `xml:"ptsTime,attr,omitempty"`
}

// Must match the pattern (Provisioned|Processing|Failed|Verified|Available|Deleting|Deleted|private:.+)
type StateType string

type StatusCodeType struct {
	Note       []NonEmptyStringType `xml:"urn:cablelabs:md:xsd:core:3.0 Note,omitempty"`
	ClassCode  int                  `xml:"classCode,attr"`
	DetailCode int                  `xml:"detailCode,attr,omitempty"`
}

// A base type for still image assets.
type StillImageAssetType struct {
	X_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 X_Resolution,omitempty"`
	Y_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 Y_Resolution,omitempty"`
	Language             string                  `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	Codec                ImageCodecType          `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	POGroupRef           []EffectiveAssetRefType `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef       []AssetRefType          `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl            string                  `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize      uint64                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum      ChecksumType            `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority  int                     `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef           string                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId          []AlternateIdType       `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string                  `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType           `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType             `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType         `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType                 `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType            `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string                  `xml:"uriId,attr"`
	ProviderVersionNum   int                     `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                     `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time               `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time               `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time               `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string                  `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time               `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string                  `xml:"eTag,attr,omitempty"`
	State                StateType               `xml:"state,attr,omitempty"`
	StateDetail          string                  `xml:"stateDetail,attr,omitempty"`
}

func (t *StillImageAssetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StillImageAssetType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *StillImageAssetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StillImageAssetType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// A single platform specific time comprised of a type and value pair.
type StreamTimeType struct {
	Ext       ExtType            `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	TimeType  NonEmptyStringType `xml:"timeType,attr,omitempty"`
	TimeValue NonEmptyStringType `xml:"timeValue,attr,omitempty"`
}

// StreamTimes contains one or  more time values that are generated and interpreted by an underlying subsystem. While implementation specific, this data is intended to be passed as is through intervening subsystems unaltered.
type StreamTimesType struct {
	StreamTime []StreamTimeType `xml:"urn:cablelabs:md:xsd:signaling:3.0 StreamTime"`
	Ext        ExtType          `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
}

type ThumbnailType struct {
	X_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 X_Resolution,omitempty"`
	Y_Resolution         uint                    `xml:"urn:cablelabs:md:xsd:content:3.0 Y_Resolution,omitempty"`
	Language             string                  `xml:"urn:cablelabs:md:xsd:content:3.0 Language,omitempty"`
	Codec                ImageCodecType          `xml:"urn:cablelabs:md:xsd:content:3.0 Codec,omitempty"`
	POGroupRef           []EffectiveAssetRefType `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef       []AssetRefType          `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl            string                  `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize      uint64                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum      ChecksumType            `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority  int                     `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef           string                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId          []AlternateIdType       `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string                  `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType           `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType             `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType         `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType                 `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType            `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string                  `xml:"uriId,attr"`
	ProviderVersionNum   int                     `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                     `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time               `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time               `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time               `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string                  `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time               `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string                  `xml:"eTag,attr,omitempty"`
	State                StateType               `xml:"state,attr,omitempty"`
	StateDetail          string                  `xml:"stateDetail,attr,omitempty"`
}

func (t *ThumbnailType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ThumbnailType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *ThumbnailType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ThumbnailType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// See Section 10.3.4 - time_descriptor()
type TimeDescriptorType struct {
	Ext        Ext    `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	TaiSeconds uint64 `xml:"taiSeconds,attr,omitempty"`
	TaiNs      uint   `xml:"taiNs,attr,omitempty"`
	UtcOffset  uint   `xml:"utcOffset,attr,omitempty"`
}

// See Section 9.3.4 - time_signal()
type TimeSignalType struct {
	SpliceTime SpliceTimeType `xml:"http://www.scte.org/schemas/35/2016 SpliceTime"`
	Ext        Ext            `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
}

type TrickModeExclusionType struct {
	Type      TrickModeType  `xml:"type,attr"`
	Scale     SpeedScaleType `xml:"scale,attr,omitempty"`
	Lowertest LowerTestType  `xml:"lowertest,attr,omitempty"`
	Uppertest UpperTestType  `xml:"uppertest,attr,omitempty"`
}

type TrickModeRestrictionType struct {
	TrickModeExclusion       []TrickModeExclusionType `xml:"urn:cablelabs:md:xsd:core:3.0 TrickModeExclusion"`
	TrickModeRestrictionRule ProcessRuleSelectorType  `xml:"urn:cablelabs:md:xsd:core:3.0 TrickModeRestrictionRule,omitempty"`
	Ext                      ExtType                  `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	StartDateTime            time.Time                `xml:"startDateTime,attr,omitempty"`
	EndDateTime              time.Time                `xml:"endDateTime,attr,omitempty"`
	Mon                      bool                     `xml:"mon,attr,omitempty"`
	Tue                      bool                     `xml:"tue,attr,omitempty"`
	Wed                      bool                     `xml:"wed,attr,omitempty"`
	Thu                      bool                     `xml:"thu,attr,omitempty"`
	Fri                      bool                     `xml:"fri,attr,omitempty"`
	Sat                      bool                     `xml:"sat,attr,omitempty"`
	Sun                      bool                     `xml:"sun,attr,omitempty"`
	StartTime                time.Time                `xml:"startTime,attr,omitempty"`
	Duration                 string                   `xml:"duration,attr,omitempty"`
}

func (t *TrickModeRestrictionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TrickModeRestrictionType
	var layout struct {
		*T
		StartDateTime *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime   *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		StartTime     *xsdTime     `xml:"startTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.StartTime = (*xsdTime)(&layout.T.StartTime)
	return e.EncodeElement(layout, start)
}
func (t *TrickModeRestrictionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TrickModeRestrictionType
	var overlay struct {
		*T
		StartDateTime *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime   *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		Mon           *bool        `xml:"mon,attr,omitempty"`
		Tue           *bool        `xml:"tue,attr,omitempty"`
		Wed           *bool        `xml:"wed,attr,omitempty"`
		Thu           *bool        `xml:"thu,attr,omitempty"`
		Fri           *bool        `xml:"fri,attr,omitempty"`
		Sat           *bool        `xml:"sat,attr,omitempty"`
		Sun           *bool        `xml:"sun,attr,omitempty"`
		StartTime     *xsdTime     `xml:"startTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.Mon = (*bool)(&overlay.T.Mon)
	overlay.Tue = (*bool)(&overlay.T.Tue)
	overlay.Wed = (*bool)(&overlay.T.Wed)
	overlay.Thu = (*bool)(&overlay.T.Thu)
	overlay.Fri = (*bool)(&overlay.T.Fri)
	overlay.Sat = (*bool)(&overlay.T.Sat)
	overlay.Sun = (*bool)(&overlay.T.Sun)
	overlay.StartTime = (*xsdTime)(&overlay.T.StartTime)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern trick|jump|all|pause|private:.+
type TrickModeType string

// Describes a trick-play asset.
type TrickType struct {
	BitRate              int                     `xml:"urn:cablelabs:md:xsd:content:3.0 BitRate"`
	VendorName           NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 VendorName,omitempty"`
	VendorProduct        NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 VendorProduct"`
	ForVersion           NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 ForVersion"`
	TrickMode            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 TrickMode"`
	POGroupRef           []EffectiveAssetRefType `xml:"urn:cablelabs:md:xsd:content:3.0 POGroupRef,omitempty"`
	SignalGroupRef       []AssetRefType          `xml:"urn:cablelabs:md:xsd:content:3.0 SignalGroupRef,omitempty"`
	SourceUrl            string                  `xml:"urn:cablelabs:md:xsd:content:3.0 SourceUrl,omitempty"`
	ContentFileSize      uint64                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentFileSize,omitempty"`
	ContentCheckSum      ChecksumType            `xml:"urn:cablelabs:md:xsd:content:3.0 ContentCheckSum,omitempty"`
	PropagationPriority  int                     `xml:"urn:cablelabs:md:xsd:content:3.0 PropagationPriority,omitempty"`
	ContentRef           string                  `xml:"urn:cablelabs:md:xsd:content:3.0 ContentRef,omitempty"`
	MediaType            NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:content:3.0 MediaType,omitempty"`
	AlternateId          []AlternateIdType       `xml:"urn:cablelabs:md:xsd:core:3.0 AlternateId,omitempty"`
	ProviderQAContact    string                  `xml:"urn:cablelabs:md:xsd:core:3.0 ProviderQAContact,omitempty"`
	AssetName            AssetNameType           `xml:"urn:cablelabs:md:xsd:core:3.0 AssetName,omitempty"`
	Product              ProductType             `xml:"urn:cablelabs:md:xsd:core:3.0 Product,omitempty"`
	Provider             NonEmptyStringType      `xml:"urn:cablelabs:md:xsd:core:3.0 Provider,omitempty"`
	Description          DescriptionType         `xml:"urn:cablelabs:md:xsd:core:3.0 Description,omitempty"`
	Ext                  ExtType                 `xml:"urn:cablelabs:md:xsd:core:3.0 Ext,omitempty"`
	MasterSourceRef      AssetRefType            `xml:"urn:cablelabs:md:xsd:core:3.0 MasterSourceRef,omitempty"`
	UriId                string                  `xml:"uriId,attr"`
	ProviderVersionNum   int                     `xml:"providerVersionNum,attr,omitempty"`
	InternalVersionNum   int                     `xml:"internalVersionNum,attr,omitempty"`
	CreationDateTime     time.Time               `xml:"creationDateTime,attr,omitempty"`
	StartDateTime        time.Time               `xml:"startDateTime,attr,omitempty"`
	EndDateTime          time.Time               `xml:"endDateTime,attr,omitempty"`
	NotifyURI            string                  `xml:"notifyURI,attr,omitempty"`
	LastModifiedDateTime time.Time               `xml:"lastModifiedDateTime,attr,omitempty"`
	ETag                 string                  `xml:"eTag,attr,omitempty"`
	State                StateType               `xml:"state,attr,omitempty"`
	StateDetail          string                  `xml:"stateDetail,attr,omitempty"`
}

func (t *TrickType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TrickType
	var layout struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDateTime = (*xsdDateTime)(&layout.T.CreationDateTime)
	layout.StartDateTime = (*xsdDateTime)(&layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(&layout.T.EndDateTime)
	layout.LastModifiedDateTime = (*xsdDateTime)(&layout.T.LastModifiedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *TrickType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TrickType
	var overlay struct {
		*T
		CreationDateTime     *xsdDateTime `xml:"creationDateTime,attr,omitempty"`
		StartDateTime        *xsdDateTime `xml:"startDateTime,attr,omitempty"`
		EndDateTime          *xsdDateTime `xml:"endDateTime,attr,omitempty"`
		LastModifiedDateTime *xsdDateTime `xml:"lastModifiedDateTime,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDateTime = (*xsdDateTime)(&overlay.T.CreationDateTime)
	overlay.StartDateTime = (*xsdDateTime)(&overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(&overlay.T.EndDateTime)
	overlay.LastModifiedDateTime = (*xsdDateTime)(&overlay.T.LastModifiedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Type to hold the UTC time of a point.
type UTCPointDescriptorType struct {
	Ext      ExtType      `xml:"urn:cablelabs:md:xsd:signaling:3.0 Ext,omitempty"`
	UtcPoint UTCPointType `xml:"utcPoint,attr"`
}

// UTC Point  Type expressed as XSD:dateTime
type UTCPointType time.Time

func (t *UTCPointType) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t UTCPointType) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// SCTE 35 supports a specific list of descriptor types but recognizes other descriptors may be sent. The Unique descriptor type allows an implementaiton to handle such descriptors in an implementation specific manner.
type UniqueDescriptorType struct {
	DescriptorTag  byte           `xml:"descriptorTag,attr"`
	DescriptorData DescriptorData `xml:"descriptorData,attr,omitempty"`
}

// Must match the pattern lt|lteq|private:.+
type UpperTestType string

type UriProcessingRequest struct {
	Ext                      ExtType            `xml:"urn:cablelabs:iptvservices:esam:xsd:common:1 Ext,omitempty"`
	AcquisitionPointIdentity NonEmptyStringType `xml:"acquisitionPointIdentity,attr"`
	BatchId                  string             `xml:"batchId,attr"`
	UriId                    string             `xml:"uriId,attr"`
}

type _anon1 struct {
	Ext        Ext            `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	SpliceTime SpliceTimeType `xml:"http://www.scte.org/schemas/35/2016 SpliceTime,omitempty"`
}

type _anon2 struct {
	Ext          Ext            `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	SpliceTime   SpliceTimeType `xml:"http://www.scte.org/schemas/35/2016 SpliceTime,omitempty"`
	ComponentTag byte           `xml:"componentTag,attr"`
}

type _anon3 struct {
	Ext          Ext    `xml:"http://www.scte.org/schemas/35/2016 Ext,omitempty"`
	ComponentTag byte   `xml:"componentTag,attr"`
	PtsOffset    uint64 `xml:"ptsOffset,attr"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdHexBinary []byte

func (b *xsdHexBinary) UnmarshalText(text []byte) (err error) {
	*b, err = hex.DecodeString(string(text))
	return
}
func (b xsdHexBinary) MarshalText() ([]byte, error) {
	n := hex.EncodedLen(len(b))
	buf := make([]byte, n)
	hex.Encode(buf, []byte(b))
	return buf, nil
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("15:04:05.999999999")), nil
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
