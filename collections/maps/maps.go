package maps

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func mergeMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

type myIntMap map[int]string

func (m myIntMap) sort() (index []int) {
	for k := range m {
		index = append(index, k)
	}
	sort.Ints(index)
	return
}

type myStringMap map[string]string

func (m myStringMap) sort() (index []string) {
	for k := range m {
		index = append(index, k)
	}
	sort.Strings(index)
	return
}

type Package struct {
	Provider     string `json:"Provider"`
	Product      string `json:"Product"`
	AssetName    string `json:"Asset_Name"`
	VersionMajor string `json:"Version_Major"`
	VersionMinor string `json:"Version_Minor"`
	Description  string `json:"Description"`
	CreationDate string `json:"Creation_Date"`
	ProviderID   string `json:"Provider_ID"`
	Verb         string `json:"Verb"`

	//AssetID             string `json:"Asset_ID"`
	//AssetClass          string `json:"Asset_Class"`
	//MetadataSpecVersion string `json:"Metadata_Spec_Version" spec:"MOD"`
	//TitleSection        Title  `json:"Title_Section,omitempty"`
	//MovieSection        Movie  `json:"Movie_Section,omitempty"`
}

type PassObj func(*map[string]interface{}) map[string]string

func Init(m interface{}, fn PassObj) error {

	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("only accepts structs; got %T", v)
	}

	t := v.Type()

	out := make(map[string]interface{}, v.NumField())

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)

		if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			// check for possible comma as in "...,omitempty"
			var commaIdx int
			if commaIdx = strings.Index(jsonTag, ","); commaIdx < 0 {
				commaIdx = len(jsonTag)
			}

			s := jsonTag[:commaIdx]

			if v.Field(i).CanInterface() {
				out[s] = v.Field(i).Interface()
			} else {
				return fmt.Errorf("unexported field (lower case) not allowed: %T", v.Type().Field(i).Name)
			}
		}
	}
	fn(&out)
	return nil
}
