package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenStandardEnum string

const (
	TokenStandardERC20   TokenStandardEnum = "ERC20"
	TokenStandardERC721  TokenStandardEnum = "ERC721"
	TokenStandardERC1155 TokenStandardEnum = "ERC1155"
)

type SmartNftAbilityEnum uint8

const (
	SmartNftAbilityUserInvitation SmartNftAbilityEnum = iota
	SmartNftAbilityCelebritySupport
	SmartNftAbilityCelebrityShare
	SmartNftAbilityEventOFuel
	SmartNftAbilityRaffleReward
	SmartNftAbilityTokenInterests
)

var SmartNftAbilitiesMap = map[SmartNftAbilityEnum]string{
	// SmartNftAbilityUserInvitation receives an additional X% of invitation points for inviting new users
	SmartNftAbilityUserInvitation: "Invitation Point",
	// SmartNftAbilityCelebritySupport receives an additional X% of support points for a community when supporting for its celebrities.
	SmartNftAbilityCelebritySupport: "Support Point",
	// SmartNftAbilityCelebrityShare receives an additional X% of share points for a community when supporting for its celebrities.
	SmartNftAbilityCelebrityShare: "Share Point",
	// SmartNftAbilityEventOFuel receives an additional X% of OFuels from events hosted by Overeality
	SmartNftAbilityEventOFuel: "Event Reward",
	// SmartNftAbilityRaffleReward receives additional X% weekly raffle rewards when you become the winner of the raffle
	SmartNftAbilityRaffleReward: "Raffle Reward",
	// SmartNftAbilityTokenInterests receives an additional X% of all tokens that you currently own
	SmartNftAbilityTokenInterests: "Token Interests",
}

func (s SmartNftAbilityEnum) String() string {
	return SmartNftAbilitiesMap[s]
}

func (s SmartNftAbilityEnum) IntValue() uint8 {
	return uint8(s)
}

func NewSmartNftAbilityEnumFromString(strValue string) (SmartNftAbilityEnum, error) {
	switch strValue {
	default:
		return SmartNftAbilityUserInvitation, errors.New("Invalid SmartNft Ability Enum")
	case "Invitation Point":
		return SmartNftAbilityUserInvitation, nil
	case "Support Point":
		return SmartNftAbilityCelebritySupport, nil
	case "Share Point":
		return SmartNftAbilityCelebrityShare, nil
	case "Event Reward":
		return SmartNftAbilityEventOFuel, nil
	case "Raffle Reward":
		return SmartNftAbilityRaffleReward, nil
	case "Token Interests":
		return SmartNftAbilityTokenInterests, nil
	}
}

type SmartNftAbility struct {
	Ability     SmartNftAbilityEnum `json:"ability,omitempty" bson:"ability"`
	Percentage  int                 `json:"percentage,omitempty" bson:"percentage"`
	IconUri     string              `json:"iconUri,omitempty" bson:"iconUri"`
	Description string              `json:"description,omitempty" bson:"description"`
}

func (m SmartNftAbility) String() string {
	return fmt.Sprintf("%s(%d): %d", m.Ability.String(), m.Ability.IntValue(), m.Percentage)
}

func (m *SmartNftAbility) MarshalJSON() ([]byte, error) {
	type Alias SmartNftAbility
	return json.Marshal(&struct {
		Ability string `json:"ability"`
		*Alias
	}{
		Ability: m.Ability.String(),
		Alias:   (*Alias)(m),
	})
}

func (m *SmartNftAbility) UnmarshalJSON(data []byte) error {
	type Alias SmartNftAbility
	aux := &struct {
		Ability string `json:"ability"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ab, err := NewSmartNftAbilityEnumFromString(aux.Ability)
	if err != nil {
		return err
	}
	m.Ability = ab
	return nil
}

type Token struct {
	mgm.IDField `json:",inline" bson:",inline"`
	// @dev Before transaction move to blockchain, this is not mandatory.
	// mandatory: true. contract address for this token
	ContractAddress *string `json:"contractAddress,omitempty" bson:"contractAddress"`
	// mandatory: true. the token belong to which network
	NetworkId primitive.ObjectID `json:"networkId,omitempty" bson:"networkId"`
	// mandatory: false. field for ERC721 and ERC1155.
	ContractTokenId int64 `json:"contractTokenId,omitempty" bson:"contractTokenId"`
	// mandatory: true. Token name
	Name string `json:"name,omitempty" bson:"name"`
	// mandatory: false. Token currency. for ERC20/1155
	Currency *string `json:"currency,omitempty" bson:"currency"`
	// mandatory: true. Token Standard
	Standard TokenStandardEnum `json:"standard,omitempty" bson:"standard"`
	// mandatory: false. If the NFT can be equipped in DId Card.
	SmartNft bool `json:"smartNft,omitempty" bson:"smartNft"`
	// mandatory: false. if the token is bind with a community.
	CommunityId *primitive.ObjectID `json:"communityId,omitempty" bson:"communityId"`
	// mandatory: true. Decimals exponent for token transaction. 0 for ERC721 and ERC1155
	Decimals int `json:"decimals,omitempty" bson:"decimals"`
	// mandatory: true. total supply for this token. 1 for ERC721
	TotalSupply primitive.Decimal128 `json:"totalSupply,omitempty" bson:"totalSupply"`
	// mandatory: false. frozen supply when doing token exchange. Only for ERC20
	FrozenSupply *primitive.Decimal128 `json:"frozenSupply,omitempty" bson:"frozenSupply"`
	// mandatory: false. field only used for ERC721 and ERC1155
	NftMetaData *NftMetaData `json:"nftMetaData,omitempty" bson:"nftMetaData"`
	// mandatory: true. user id who created the token.
	CreatedBy *primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy"`
	// mandatory: true. create time of this token.
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	// mandatory: true. update time of this token.
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func (m *Token) Validate() error {
	return nil
}

type NftMetaData struct {
	// mandatory: false. NFT title.
	Title string `json:"title,omitempty" bson:"title"`
	// mandatory: false. NFT subtitle.
	SubTitle string `json:"subTitle,omitempty" bson:"subTitle"`
	// mandatory: false. NFT description
	Description string `json:"description,omitempty" bson:"description"`
	// mandatory: true. NFT URI
	Uri string `json:"uri,omitempty" bson:"uri"`
	// mandatory: false. field only used for ERC721
	NumId int `json:"numId,omitempty" bson:"numId"`
	// mandatory: true. user id who mint this NFT
	MintBy primitive.ObjectID `json:"mintBy,omitempty" bson:"mintBy"`
	// mandatory: false. only used for ERC721
	Owner *primitive.ObjectID `json:"owner,omitempty" bson:"owner"`
	// mandatory: false. only used for Did Card NFT
	LeastLevel int `json:"leastLevel,omitempty" bson:"leastLevel"`
	// mandatory: false. only used for Did Card NFT
	SmartNftAbilities []*SmartNftAbility `json:"smartNftAbilities,omitempty" bson:"smartNftAbilities"`
	// mandatory: false. for extra metadata, store json in value if needed.
	ExtraData map[string]string `json:"extraData,omitempty" bson:"extraData"`
}

func ListTokens(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*Token, error) {
	var ts []*Token
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&Token{}).SimpleFindWithCtx(ctx, &ts, filter, opt)
	return ts, err
}

func CountTokens(filter interface{}) (count int64, err error) {
	return mgm.Coll(&Token{}).CountDocuments(mgm.Ctx(), filter)
}

func (m *Token) CollectionName() string {
	return "tokens"
}
