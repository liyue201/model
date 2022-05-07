package model

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Contract struct {
	mgm.IDField `json:",inline" bson:",inline"`
	// mandatory: true. Contract name
	Name string `json:"name" bson:"name"`
	// mandatory: false. git repo if the contract is in OVR git
	GitRepo string `json:"gitRepo" bson:"gitRepo"`
	// mandatory: false. git repo release tag if the contract is in OVR git
	GitRelease string 	`json:"gitRelease" bson:"gitRelease"`
	// mandatory: false. git repo commit id if the contract is in OVR git
	GitCommitId string 	`json:"gitCommitId" bson:"gitCommitId"`
	// mandatory: true. contract network id.
	NetworkId primitive.ObjectID `json:"networkId,omitempty" bson:"networkId"`
	// mandatory: true. address of the contract.
	ContractAddress string `json:"contractAddress" bson:"contractAddress"`
	// mandatory: false. key used to deploy contract
	DeployKey string 	`json:"deployKey" bson:"deployKey"`
	// mandatory: false. contract owner public key
	OwnerKey string 	`json:"ownerKey" bson:"ownerKey"`
	// mandatory: false. Contract description
	Description string `json:"description,omitempty" bson:"description"`
	// mandatory: false. user id who created the contract.
	CreatedBy *primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy"`
	// mandatory: false. time the contract is first deployed.
	CreatedAt *time.Time 	`json:"createdAt,omitempty" bson:"createdAt"`
	// mandatory: false. time the contract is last updated.
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty" bson:"LastUpdatedAt"`
	// mandatory: false. contract abi json string.
	ABI string	`json:"abi,omitempty" bson:"abi"`
}

func ListContracts(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*Contract, error) {
	var cs []*Contract
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&Token{}).SimpleFindWithCtx(ctx, &cs, filter, opt)
	return cs, err
}

func CountContracts(filter interface{}) (count int64, err error) {
	return mgm.Coll(&Contract{}).CountDocuments(mgm.Ctx(), filter)
}

func (m *Contract) CollectionName() string {
	return "contracts"
}