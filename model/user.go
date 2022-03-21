package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	// ID
	mgm.IDField `json:",inline" bson:",inline"`
	Roles []primitive.ObjectID `json:"roles,omitempty" bson:"roles"`
	//Name                   string               `json:"name,omitempty" bson:"name"`
	Email                  string               `json:"email,omitempty" bson:"email"`
	Password               string               `json:"password,omitempty" bson:"password"`
	PublicKey              string               `json:"publicKey,omitempty" bson:"publicKey"`
	SignedMessage          string               `json:"signedMessage,omitempty" bson:"signedMessage"`
	InvitationCode         string               `json:"invitationCode,omitempty" bson:"invitationCode"`
	ReferenceCode          string               `json:"referenceCode" bson:"referenceCode"`
	ReferenceCodeUsedCount uint64               `json:"referenceCodeUsedCount,omitempty" bson:"referenceCodeUsedCount"`
	ReferencedBy           *primitive.ObjectID  `json:"referencedBy,omitempty" bson:"referencedBy"`
	ReferencedRank         uint64               `json:"referencedRank,omitempty" bson:"referencedRank"`
	FollowCommunities      []primitive.ObjectID `json:"followCommunities,omitempty" bson:"followCommunities"`
	FavoriteCommunities    []primitive.ObjectID `json:"favoriteCommunities,omitempty" bson:"favoriteCommunities"`
	LikePosts              []primitive.ObjectID `json:"likePosts,omitempty" bson:"likePosts"`
	Status                 uint64               `json:"status,omitempty" bson:"status"`
	VerificationCode       string               `json:"verificationCode,omitempty" bson:"verificationCode"`
	PasswordResetCode      string               `json:"passwordResetCode,omitempty" bson:"passwordResetCode"`
	Nfts                   []OldNft             `json:"nfts,omitempty" bson:"nfts"`
	Balances               []Balance            `json:"balances,omitempty" bson:"balances"`
	Guesses                []Guess              `json:"guesses,omitempty" bson:"guesses"`
	LastLoginTime          int64                `json:"lastLoginTime" bson:"lastLoginTime"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	// 最近一次更新时间
	UpdatedAt *time.Time `json:"lastUpdatedAt,omitempty" bson:"lastUpdatedAt"`
}

type Guess struct {
	CommunityId  primitive.ObjectID `json:"communityId,omitempty" bson:"communityId"`
	GuessContent string             `json:"guessContent,omitempty" bson:"guessContent"`
	GuessAt      *time.Time         `json:"guessAt,omitempty" bson:"guessAt"`
	GuessResult  GuessResult        `json:"guessResult,omitempty" bson:"guessResult"`
}

type GuessResult struct {
	IsCorrect bool   `json:"isCorrect,omitempty" bson:"isCorrect"`
	GuessRank uint64 `json:"guessRank,omitempty" bson:"guessRank"`
}

type OldNft struct {
	Uri         string `json:"uri,omitempty" bson:"uri"`
	Title       string `json:"title,omitempty" bson:"title"`
	SubTitle    string `json:"subTitle,omitempty" bson:"subTitle"`
	Description string `json:"description,omitempty" bson:"description"`
}

// Balance Add TokenAmount Decimal type
type Balance struct {
	TokenType    string               `json:"tokenType,omitempty" bson:"tokenType"`
	Amount       float64              `json:"amount,omitempty" bson:"amount"`
	TokenBalance primitive.Decimal128 `json:"tokenBalance,omitempty" bson:"tokenBalance"`
	Currency     string               `json:"currency,omitempty" bson:"currency"`
}
