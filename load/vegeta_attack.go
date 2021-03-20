package load

import (
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"net/http"
)

func attack() {
	vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodPost,
		URL:    "http://localhost:7010/",
	})
	attacker := vegeta.NewAttacker()
	attacker.Stop()
}
