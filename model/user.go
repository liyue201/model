package model

import (
	"context"
	"time"

	"github.com/Overealityio/overeality-server-model/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User 用户信息
type User struct {
	// ID
	mgm.IDField `json:",inline" bson:",inline"`
	// 角色列表，目前应该已经没用了
	Roles []primitive.ObjectID `json:"roles,omitempty" bson:"roles"`
	// 用户名
	Name string `json:"name,omitempty" bson:"name"`
	// 简介
	Description string `json:"description,omitempty" bson:"description"`
	// 用户邮箱
	Email string `json:"email,omitempty" bson:"email"`
	// 用户密码
	Password string `json:"password,omitempty" bson:"password"`
	// metamask的Public Key
	PublicKey string `json:"publicKey,omitempty" bson:"publicKey"`
	// 修改PublicKey时的验证码
	PublicKeyUpdateVerificationCode string `json:"publicKeyUpdateVerificationCode,omitempty"  bson:"publicKeyUpdateVerificationCode"`
	SignedMessage                   string `json:"signedMessage,omitempty" bson:"signedMessage"`
	// 邀请码
	InvitationCode string `json:"invitationCode,omitempty" bson:"invitationCode"`
	// 推荐码
	ReferenceCode string `json:"referenceCode" bson:"referenceCode"`
	// 邀请用户总数
	ReferenceCodeUsedCount uint64 `json:"referenceCodeUsedCount,omitempty" bson:"referenceCodeUsedCount"`
	// 由谁邀请加入Overeality
	ReferencedBy *primitive.ObjectID `json:"referencedBy,omitempty" bson:"referencedBy"`
	// 邀请人数排名
	ReferencedRank uint64 `json:"referencedRank,omitempty" bson:"referencedRank"`
	// 关注的社区列表
	FollowCommunities []primitive.ObjectID `json:"followCommunities,omitempty" bson:"followCommunities"`
	// 收藏的社区列表
	FavoriteCommunities []primitive.ObjectID `json:"favoriteCommunities,omitempty" bson:"favoriteCommunities"`
	// 喜欢的文章集合
	LikePosts []primitive.ObjectID `json:"likePosts,omitempty" bson:"likePosts"`
	// 用户状态
	Status uint64 `json:"status,omitempty" bson:"status"`
	// 注册时的验证码
	VerificationCode string `json:"verificationCode,omitempty" bson:"verificationCode"`
	// 修改密码时的验证码，一般会发给用户邮箱
	PasswordResetCode string `json:"passwordResetCode,omitempty" bson:"passwordResetCode"`
	// 用户装备的NFT
	Nfts []OldNft `json:"nfts,omitempty" bson:"nfts"`
	// 用户的Token余额
	Balances []Balance `json:"balances,omitempty" bson:"balances"`
	// 暂时不清楚这个字段的意思，似乎已经被Deprecate
	Guesses []Guess `json:"guesses,omitempty" bson:"guesses"`
	// 是否为系统管理员
	IsAdmin bool `json:"isAdmin" bson:"isAdmin"`
	// 管理员角色
	AdminRoles []primitive.ObjectID `json:"adminRoles" bson:"adminRoles"`
	// 用户基于积分的总排名
	PointRank uint64 `json:"pointRank" bson:"pointRank"`
	// 用户总积分
	TotalPoints float64 `json:"totalPoints" bson:"totalPoints"`
	// 用户因邀请好友而获得的积分
	InvitationPoints float64 `json:"invitationPoints" bson:"invitationPoints"`
	// 用户积分等级
	PointLevel uint64 `json:"pointLevel" bson:"pointLevel"`
	// 最近一次登录时间
	LastLoginTime int64 `json:"lastLoginTime" bson:"lastLoginTime"`
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

type OldBalance struct {
	TokenType string  `json:"tokenType,omitempty" bson:"tokenType"`
	Amount    float64 `json:"amount,omitempty" bson:"amount"`
	Currency  string  `json:"currency,omitempty" bson:"currency"`
}

func (u *User) CollectionName() string {
	return "users"
}
func (u *User) Saving() error {
	u.UpdatedAt = util.TimeNow()
	return nil
}

func (u *User) Creating() error {
	u.CreatedAt = util.TimeNow()
	return nil
}

func (u *User) Validate() error {
	return nil
}

func (u *User) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(u).FindByIDWithCtx(ctx, id, u)
	return err
}

func (u *User) GetUserByEmail(ctx context.Context, email string) error {
	coll := mgm.Coll(u)
	err := coll.FirstWithCtx(ctx, bson.M{"email": email}, u)
	return err
}

func (u *User) GetUserByPublicKey(ctx context.Context, publicKey string) error {
	coll := mgm.Coll(u)
	err := coll.FirstWithCtx(ctx, bson.M{"publicKey": publicKey}, u)
	return err
}

func (u *User) UpdateUserLastLoginTime(ctx context.Context) error {
	u.LastLoginTime = time.Now().Unix()
	opt := &options.UpdateOptions{}
	opt.SetUpsert(false)
	err := mgm.Coll(u).UpdateWithCtx(ctx, u, opt)
	return err
}

func (u *User) PublicKeyLoginProcess(ctx context.Context) error {
	u.VerificationCode = ""
	u.LastLoginTime = time.Now().Unix()
	if u.ReferenceCode == "" {
		codeGenerator, _ := util.NewSonyflake()
		code, _ := codeGenerator.GenerateCode()
		u.ReferenceCode = code
	}
	opt := &options.UpdateOptions{}
	opt.SetUpsert(false)
	err := mgm.Coll(u).UpdateWithCtx(ctx, u, opt)
	return err
}

func (u *User) Update(ctx context.Context) error {
	opt := &options.UpdateOptions{}
	opt.SetUpsert(false)
	return mgm.Coll(u).UpdateWithCtx(ctx, u, opt)
}

func ListUsers(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*User, error) {
	var users []*User
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&User{}).SimpleFindWithCtx(ctx, &users, filter, opt)
	return users, err
}

func CountUsers(filter interface{}) (count int64, err error) {
	return mgm.Coll(&User{}).CountDocuments(mgm.Ctx(), filter)
}
