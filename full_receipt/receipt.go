package full_receipt

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// Receipt based on types.Receipt but with From and To fields
type Receipt struct {
	From common.Address  `json:"from"`
	To   *common.Address `json:"to"`

	// Consensus fields: These fields are defined by the Yellow Paper
	Type              uint8       `json:"type,omitempty"`
	PostState         []byte      `json:"root"`
	Status            uint64      `json:"status"`
	CumulativeGasUsed uint64      `json:"cumulativeGasUsed"`
	Bloom             types.Bloom `json:"logsBloom"`
	Logs              []*Log      `json:"logs"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	TxHash            common.Hash    `json:"transactionHash"`
	ContractAddress   common.Address `json:"contractAddress"`
	GasUsed           uint64         `json:"gasUsed"`
	EffectiveGasPrice *big.Int       `json:"effectiveGasPrice"` // required, but tag omitted for backwards compatibility
	BlobGasUsed       uint64         `json:"blobGasUsed,omitempty"`
	BlobGasPrice      *big.Int       `json:"blobGasPrice,omitempty"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint        `json:"transactionIndex"`
}

func (r *Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		BlobGasPrice      *hexutil.Big    `json:"blobGasPrice,omitempty"`
		BlobGasUsed       hexutil.Uint64  `json:"blobGasUsed,omitempty"`
		BlockHash         common.Hash     `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big    `json:"blockNumber,omitempty"`
		ContractAddress   *common.Address `json:"contractAddress"`
		CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
		EffectiveGasPrice *hexutil.Big    `json:"effectiveGasPrice"`
		From              common.Address  `json:"from"`
		GasUsed           hexutil.Uint64  `json:"gasUsed"`
		Logs              []*Log          `json:"logs"`
		Bloom             types.Bloom     `json:"logsBloom"`
		PostState         *hexutil.Bytes  `json:"root,omitempty"`
		Status            hexutil.Uint64  `json:"status"`
		To                *common.Address `json:"to"`
		TxHash            common.Hash     `json:"transactionHash"`
		TransactionIndex  hexutil.Uint    `json:"transactionIndex"`
		Type              hexutil.Uint64  `json:"type"`
	}
	var enc Receipt
	enc.From = r.From
	enc.To = r.To
	enc.Type = hexutil.Uint64(r.Type)
	if len(r.PostState) != 0 {
		enc.PostState = (*hexutil.Bytes)(&r.PostState)
	}
	enc.Status = hexutil.Uint64(r.Status)
	enc.CumulativeGasUsed = hexutil.Uint64(r.CumulativeGasUsed)
	enc.Bloom = r.Bloom
	enc.Logs = r.Logs
	enc.TxHash = r.TxHash
	if r.ContractAddress != (common.Address{}) {
		enc.ContractAddress = &r.ContractAddress
	}
	enc.GasUsed = hexutil.Uint64(r.GasUsed)
	enc.EffectiveGasPrice = (*hexutil.Big)(r.EffectiveGasPrice)
	enc.BlobGasUsed = hexutil.Uint64(r.BlobGasUsed)
	enc.BlobGasPrice = (*hexutil.Big)(r.BlobGasPrice)
	enc.BlockHash = r.BlockHash
	enc.BlockNumber = (*hexutil.Big)(r.BlockNumber)
	enc.TransactionIndex = hexutil.Uint(r.TransactionIndex)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		From              common.Address  `json:"from"`
		To                *common.Address `json:"to"`
		Type              *hexutil.Uint64 `json:"type,omitempty"`
		PostState         *hexutil.Bytes  `json:"root"`
		Status            *hexutil.Uint64 `json:"status"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulativeGasUsed"`
		Bloom             *types.Bloom    `json:"logsBloom"`
		Logs              []*Log          `json:"logs"`
		TxHash            *common.Hash    `json:"transactionHash"`
		ContractAddress   *common.Address `json:"contractAddress"`
		GasUsed           *hexutil.Uint64 `json:"gasUsed"`
		EffectiveGasPrice *hexutil.Big    `json:"effectiveGasPrice"`
		BlobGasUsed       *hexutil.Uint64 `json:"blobGasUsed,omitempty"`
		BlobGasPrice      *hexutil.Big    `json:"blobGasPrice,omitempty"`
		BlockHash         *common.Hash    `json:"blockHash,omitempty"`
		BlockNumber       *hexutil.Big    `json:"blockNumber,omitempty"`
		TransactionIndex  *hexutil.Uint   `json:"transactionIndex"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	r.From = dec.From
	r.To = dec.To
	if dec.Type != nil {
		r.Type = uint8(*dec.Type)
	}
	if dec.PostState != nil {
		r.PostState = *dec.PostState
	}
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	if dec.CumulativeGasUsed == nil {
		return errors.New("missing required field 'cumulativeGasUsed' for Receipt")
	}
	r.CumulativeGasUsed = uint64(*dec.CumulativeGasUsed)
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom = *dec.Bloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = dec.Logs
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash = *dec.TxHash
	if dec.ContractAddress != nil {
		r.ContractAddress = *dec.ContractAddress
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed = uint64(*dec.GasUsed)
	if dec.EffectiveGasPrice != nil {
		r.EffectiveGasPrice = (*big.Int)(dec.EffectiveGasPrice)
	}
	if dec.BlobGasUsed != nil {
		r.BlobGasUsed = uint64(*dec.BlobGasUsed)
	}
	if dec.BlobGasPrice != nil {
		r.BlobGasPrice = (*big.Int)(dec.BlobGasPrice)
	}
	if dec.BlockHash != nil {
		r.BlockHash = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		r.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.TransactionIndex != nil {
		r.TransactionIndex = uint(*dec.TransactionIndex)
	}
	return nil
}
