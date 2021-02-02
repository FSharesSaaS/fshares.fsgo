package rex

import (
	eos "github.com/FSharesSaaS/fshares.fsgo"
)

func NewConsolidate(owner eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner eos.AccountName
}
