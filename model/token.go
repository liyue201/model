package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TokenStandardEnum string

const(
	TokenStandardERC20   TokenStandardEnum = "ERC20"
	TokenStandardERC721  TokenStandardEnum = "ERC721"
	TokenStandardERC1155 TokenStandardEnum = "ERC1155"
)

type Token struct {
	mgm.IDField  `json:",inline" bson:",inline"`
	// @dev Before transaction move to blockchain, this is not mandatory.
	// mandatory: true. contract address for this token
	ContractAddress      *string      `json:"contractAddress,omitempty" bson:"contractAddress"`
	// mandatory: true. the token belong to which network
	NetworkId       primitive.ObjectID     `json:"networkId,omitempty" bson:"networkId"`
	// mandatory: false. field for ERC721 and ERC1155.
	ContractTokenId      int64        `json:"contractTokenId,omitempty" bson:"contractTokenId"`
	// mandatory: true. Token name
	Name         string               `json:"name,omitempty" bson:"name"`
	// mandatory: false. Token currency. for ERC20/1155
	Currency  *string               `json:"currency,omitempty" bson:"currency"`
	// mandatory: true. Token Standard
	Standard     TokenStandardEnum    `json:"standard,omitempty" bson:"standard"`
	// mandatory: false. if the token is bind with a community.
	CommunityId *primitive.ObjectID `json:"communityId,omitempty" bson:"communityId"`
	// mandatory: true. Decimals exponent for token transaction. 0 for ERC721 and ERC1155
	Decimals     int                  `json:"decimals,omitempty" bson:"decimals"`
	// mandatory: true. total supply for this token. 1 for ERC721
	TotalSupply  primitive.Decimal128 `json:"totalSupply,omitempty" bson:"totalSupply"`
	// mandatory: false. frozen supply when doing token exchange. Only for ERC20
	FrozenSupply *primitive.Decimal128 `json:"frozenSupply,omitempty" bson:"frozenSupply"`
	// mandatory: false. field only used for ERC721 and ERC1155
	NftMetaData *NftMetaData `json:"nftMetaData,omitempty" bson:"nftMetaData"`
	// mandatory: true. create time of this token.
	CreatedAt   *time.Time         `json:"createdAt,omitempty" bson:"createdAt"`
	// mandatory: true. update time of this token.
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type NftMetaData struct {
	// mandatory: false. NFT title.
	Title       string             `json:"title,omitempty" bson:"title"`
	// mandatory: false. NFT subtitle.
	SubTitle    string             `json:"subTitle,omitempty" bson:"subTitle"`
	// mandatory: false. NFT description
	Description string             `json:"description,omitempty" bson:"description"`
	// mandatory: true. NFT URI
	Uri         string             `json:"uri,omitempty" bson:"uri"`
	// mandatory: false. field only used for ERC721
	NumId       int                `json:"numId,omitempty" bson:"numId"`
	// mandatory: true. user id who mint this NFT
	MintBy      primitive.ObjectID `json:"mintBy,omitempty" bson:"mintBy"`
	// mandatory: false. only used for ERC721
	Owner	 *primitive.ObjectID `json:"owner,omitempty" bson:"owner"`
	// mandatory: false. for extra metadata, store json in value if needed.
	ExtraData   map[string]string  `json:"extraData,omitempty" bson:"extraData"`
}