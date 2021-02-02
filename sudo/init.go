package sudo

import eos "github.com/FSharesSaaS/fshares.fsgo"

func init() {
	eos.RegisterAction(AN("eosio.wrap"), ActN("exec"), Exec{})
}

var AN = eos.AN
var ActN = eos.ActN
