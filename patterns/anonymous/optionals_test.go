package anonymous

import (
	"fmt"
	"testing"
)

func TestTransformationTypeFunc_Apply(t *testing.T) {
	create := Create("controller://10.208.127.11", WithScheme(schemeChange))
	execute := create.Execute()
	fmt.Println(execute)
}
