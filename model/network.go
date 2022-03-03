package model

import "github.com/kamva/mgm/v3"

type Network struct {
	mgm.IDField  `json:",inline" bson:",inline"`
	// mandatory: true. Network name.
	Name 	string `json:"name,omitempty" bson:"name"`
	// mandatory: false. Network rpc url.
	RpcUrl  string `json:"rpcUrl,omitempty" bson:"rpcUrl"`
}
