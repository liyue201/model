package model

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type TransactionStatusEnum string
const (
	TransactionStatusSuccess TransactionStatusEnum = "success"
	TransactionStatusPending TransactionStatusEnum = "pending"
	TransactionStatusRevert  TransactionStatusEnum = "revert"
)

type Transaction struct {
	mgm.IDField    `json:",inline" bson:",inline"`
    // @dev mandatory: false before transaction move to blockchain
    // mandatory: true. blockchain used.
	NetworkId primitive.ObjectID `json:"networkId,omitempty" bson:"networkId"`
	// @dev mandatory: false before transaction move to blockchain
	// mandatory: true
	TxnHash   *string `json:"txnHash,omitempty" bson:"txnHash"`
	// mandatory: ture. a batch of tokens(erc20/721/1155) transfer in this transaction
	TokensTransferred []*TokenTransfer     `json:"tokensTransferred,omitempty" bson:"tokensTransferred"`
	// mandatory: true. OFuel token consumed for this transaction.
	GasFee   *primitive.Decimal128 `json:"gasFee,omitempty" bson:"gasFee"`
	// mandatory: false. description for this transaction
	Description   string   `json:"description,omitempty" bson:"description"`
	// mandatory: true. Create time of this transaction
	CreatedAt  *time.Time   `json:"createdAt,omitempty" bson:"createdAt"`
	// mandatory: true. transaction status
	Status     TransactionStatusEnum   `json:"status,omitempty" bson:"status"`
	// mandatory: false. If the transaction is made by operator uploading csv, batchId will be stored.
	BatchId    *primitive.ObjectID  `json:"batchId,omitempty" bson:"BatchId"`
	// mandatory: true. Related balance changes.
	States     []*BalanceState      `json:"states,omitempty" bson:"states"`
}

type TokenTransfer struct {
	// @dev Before transaction move to blockchain, this is mandatory.
	// mandatory: false. from user id if user is stored in mongo
	From     *primitive.ObjectID   `json:"from,omitempty" bson:"from"`
	// @dev Before transaction move to blockchain, this is not mandatory cause we cannot get users address
	// mandatory: true. need to transfer to uint256 on blockchain. Maybe a contract.
	FromAddress  *string     `json:"fromAddress,omitempty" bson:"fromAddress"`
	// @dev Before transaction move to blockchain, this is mandatory.
	// mandatory: false. to user id if user is stored in mongo
	To       *primitive.ObjectID   `json:"to,omitempty" bson:"to"`
	// @dev Before transaction move to blockchain, this is not mandatory cause we cannot get users address
	// mandatory: true for transactions on blockchain and need to transfer to uint256. Maybe a contract
	ToAddress    *string      `json:"toAddress,omitempty" bson:"toAddress"`
	// mandatory: true. Id for token stored in mongodb
	Token    primitive.ObjectID   `json:"token,omitempty" bson:"token"`
	// mandatory: true. valued transferred, 1 for ERC721, int for ERC1155, decimal for ERC20
	Value    primitive.Decimal128 `json:"value,omitempty" bson:"value"`
}

type BalanceState struct {
	// mandatory: false. user id if user is stored in mongo
	UserId      *primitive.ObjectID    `json:"userId,omitempty" bson:"userId"`
	// @dev before transaction move to blockchain, this is not mandatory cause we cannot get users address
	// mandatory: true for transactions on blockchain and need to transfer to uint256. Maybe a contract
	UserAddress  string  `json:"userAddress,omitempty" bson:"userAddress"`
	// mandatory: true. token value before transaction made
	Before      primitive.Decimal128  `json:"before,omitempty" bson:"before"`
	// mandatory: true. token value after transaction made
	After       primitive.Decimal128 `json:"after,omitempty" bson:"after"`
	// mandatory: true. Id for token stored in mongodb
	Token      primitive.ObjectID   `json:"token,omitempty" bson:"token"`
}

func ListTransactions(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*Transaction, error) {
	var ts []*Transaction
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&Transaction{}).SimpleFindWithCtx(ctx, &ts, filter, opt)
	return ts, err
}

func CountTransactions(filter interface{}) (count int64, err error) {
	return mgm.Coll(&Transaction{}).CountDocuments(mgm.Ctx(), filter)
}

func (t *Transaction) CollectionName() string {
	return "transactions"
}

func (t *Transaction) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(t).FindByIDWithCtx(ctx, id, t)
	return err
}