/**
 * Created by g7tianyi on Apr 24, 2022
 */

package model

// BinanceTransactionReceipt Binance transaction receipt
type BinanceTransactionReceipt struct {
	// Error 出错时，其他字段均为空
	Error string `json:"error,omitempty" bson:"error"`
	// TRUE if the transaction was successful, FALSE if the EVM reverted the transaction.
	Status bool `bson:"status" json:"status,omitempty"`
	// Hash of the block where this transaction was in.
	BlockHash string `json:"blockHash,omitempty" bson:"blockHash"`
	// Block number where this transaction was in.
	BlockNumber int64 `json:"blockNumber,omitempty" bson:"blockNumber"`
	// Hash of the transaction
	TransactionHash string `json:"transactionHash,omitempty" bson:"transactionHash"`
	// Integer of the transactions index position in the block.
	TransactionIndex int64 `json:"transactionIndex,omitempty" bson:"transactionIndex"`
	// Address of the sender.
	From string `json:"from,omitempty" bson:"from"`
	// Address of the receiver. null when it’s a contract creation transaction.
	To string `json:"to,omitempty" bson:"to"`
	// The contract address created, if the transaction was a contract creation, otherwise null.
	ContractAddress string `json:"contractAddress,omitempty" bson:"contractAddress"`
	// The amount of gas used by this specific transaction alone.
	GasUsed int64 `json:"gasUsed,omitempty" bson:"gasUsed"`
	// The bloom filter for the logs of the block. null if a pending block.
	LogsBloom string `json:"logsBloom,omitempty" bson:"logsBloom"`
	// Not sure what it is as of now
	Type string `json:"type,omitempty" bson:"type"`
}
