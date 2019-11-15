package xml_utils

import (
	"bytes"
	"fmt"
	"github.com/beevik/etree"
)

func buildAdi() string {
	document := etree.NewDocument()
	document.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	element := document.CreateElement("ADI")
	createElement := element.CreateElement("Metadata")

	ams := createElement.CreateElement("AMS")
	ams.CreateAttr("Version_Major", "1")
	ams.CreateAttr("Version_Minor", "0")
	ams.CreateAttr("Verb", "")
	ams.CreateAttr("Asset_Class", "title")
	ams.CreateAttr("Provider", "InDemand")
	ams.CreateAttr("Product", "First-Run")
	ams.CreateAttr("Asset_Name", "The_Titanic")
	ams.CreateAttr("Description", "The Titanic asset package")
	ams.CreateAttr("Creation_Date", "2002-01-11")
	ams.CreateAttr("Provider_ID", "indemand.com")
	ams.CreateAttr("Provider", "InDemand")
	ams.CreateAttr("Asset_ID", "UNVA2001081701004001")

	topAsset := element.CreateElement("Asset")
	titleMetadata := topAsset.CreateElement("Metadata")

	titleAms := titleMetadata.CreateElement("AMS")
	titleAms.CreateAttr("Version_Major", "1")
	titleAms.CreateAttr("Version_Minor", "0")
	titleAms.CreateAttr("Verb", "")
	titleAms.CreateAttr("Asset_Class", "title")
	titleAms.CreateAttr("Provider", "InDemand")
	titleAms.CreateAttr("Product", "First-Run")
	titleAms.CreateAttr("Asset_Name", "The_Titanic")
	titleAms.CreateAttr("Description", "The Titanic asset package")
	titleAms.CreateAttr("Creation_Date", "2002-01-11")
	titleAms.CreateAttr("Provider_ID", "indemand.com")
	titleAms.CreateAttr("Provider", "InDemand")
	titleAms.CreateAttr("Asset_ID", "UNVA2001081701004001")

	appData := titleMetadata.CreateElement("App_Data")
	appData.CreateAttr("App", "SVOD")
	appData.CreateAttr("Name", "Type")
	appData.CreateAttr("Value", "title")

	document.Indent(2)
	buffer := new(bytes.Buffer)
	n, err := document.WriteTo(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	return buffer.String()
}
