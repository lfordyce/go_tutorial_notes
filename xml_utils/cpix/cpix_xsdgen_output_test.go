package cpix

import (
	"aqwari.net/xml/xsdgen"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	b, err := ioutil.ReadFile("/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/cpix/schema/cpix.xsd")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestCpixSchemaGen(t *testing.T) {
	var cfg xsdgen.Config
	cfg.Option(xsdgen.PackageName("cpix"))
	//cfg.Option(xsdgen.LogOutput(log.New(os.Stderr, "", 0)), xsdgen.LogLevel(2))
	//cfg.Option(xsdgen.DefaultOptions...)
	if err := cfg.GenCLI(
		"/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/cpix/schema/cpix.xsd",
		"/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/cpix/schema/pskc.xsd",
		"/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/cpix/schema/xenc-schema.xsd",
		"/Users/LFordyc1/Go/Projects/generalNotes/xml_utils/cpix/schema/xmldsig-core-schema.xsd",
	); err != nil {
		t.Fatal(err)
	}
}

type testLogger testing.T

func (t *testLogger) Printf(format string, v ...interface{}) {
	t.Logf(format, v...)
}

func TestFullCpix(t *testing.T) {
	gen := testGen(t, "schema/cpix.xsd", "schema/pskc.xsd", "schema/xenc-schema.xsd", "schema/xmldsig-core-schema.xsd")
	fmt.Println(gen)
}

func testGen(t *testing.T, files ...string) string {
	file, err := ioutil.TempFile("", "xsdgen")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	var cfg xsdgen.Config
	cfg.Option(xsdgen.DefaultOptions...)
	cfg.Option(xsdgen.LogOutput((*testLogger)(t)))

	args := []string{"-v", "-o", file.Name()}
	err = cfg.GenCLI(append(args, files...)...)
	if err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatal(err)
	}
	return string(data)

}
