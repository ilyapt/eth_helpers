package full_receipt

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Log based on types.Log but with BlockTimestamp field
type Log struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    []byte         `json:"data"`

	BlockNumber    uint64      `json:"blockNumber"`
	BlockTimestamp uint64      `json:"blockTimestamp"`
	TxHash         common.Hash `json:"transactionHash"`
	TxIndex        uint        `json:"transactionIndex"`
	BlockHash      common.Hash `json:"blockHash"`
	Index          uint        `json:"logIndex"`

	Removed bool `json:"removed"`
}

func (l *Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address        common.Address  `json:"address"`
		BlockHash      common.Hash     `json:"blockHash"`
		BlockNumber    hexutil.Uint64  `json:"blockNumber"`
		BlockTimestamp *hexutil.Uint64 `json:"blockTimestamp,omitempty"`
		Data           hexutil.Bytes   `json:"data"`
		Index          hexutil.Uint    `json:"logIndex"`
		Removed        bool            `json:"removed"`
		Topics         []common.Hash   `json:"topics"`
		TxHash         common.Hash     `json:"transactionHash"`
		TxIndex        hexutil.Uint    `json:"transactionIndex"`
	}
	var enc Log
	enc.Address = l.Address
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = hexutil.Uint64(l.BlockNumber)
	enc.TxHash = l.TxHash
	enc.TxIndex = hexutil.Uint(l.TxIndex)
	enc.BlockHash = l.BlockHash
	enc.Index = hexutil.Uint(l.Index)
	enc.Removed = l.Removed
	if l.BlockTimestamp > 0 {
		enc.BlockTimestamp = (*hexutil.Uint64)(&l.BlockTimestamp)
	}
	return json.Marshal(&enc)
}

func (l *Log) UnmarshalJSON(input []byte) error {
	type Log struct {
		Address        *common.Address `json:"address"`
		Topics         []common.Hash   `json:"topics"`
		Data           *hexutil.Bytes  `json:"data"`
		BlockNumber    *hexutil.Uint64 `json:"blockNumber"`
		BlockTimestamp *hexutil.Uint64 `json:"blockTimestamp,omitempty"`
		TxHash         *common.Hash    `json:"transactionHash"`
		TxIndex        *hexutil.Uint   `json:"transactionIndex"`
		BlockHash      *common.Hash    `json:"blockHash"`
		Index          *hexutil.Uint   `json:"logIndex"`
		Removed        *bool           `json:"removed"`
	}
	var dec Log
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for Log")
	}
	l.Address = *dec.Address
	if dec.Topics == nil {
		return errors.New("missing required field 'topics' for Log")
	}
	l.Topics = dec.Topics
	if dec.Data == nil {
		return errors.New("missing required field 'data' for Log")
	}
	l.Data = *dec.Data
	if dec.BlockNumber != nil {
		l.BlockNumber = uint64(*dec.BlockNumber)
	}
	if dec.BlockTimestamp != nil {
		l.BlockTimestamp = uint64(*dec.BlockTimestamp)
	}
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Log")
	}
	l.TxHash = *dec.TxHash
	if dec.TxIndex != nil {
		l.TxIndex = uint(*dec.TxIndex)
	}
	if dec.BlockHash != nil {
		l.BlockHash = *dec.BlockHash
	}
	if dec.Index != nil {
		l.Index = uint(*dec.Index)
	}
	if dec.Removed != nil {
		l.Removed = *dec.Removed
	}
	return nil
}
