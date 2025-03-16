package hardhat

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ilyapt/eth_helpers/full_receipt"
	"math/big"
)

type TxArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to,omitempty"`
	Data     hexutil.Bytes   `json:"data,omitempty"`
	Gas      uint64          `json:"gas,omitempty"`
	GasPrice *big.Int        `json:"gasPrice,omitempty"`
	//GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	//GasTipCap *big.Int        // EIP-1559 tip per gas.
	Value *big.Int `json:"value,omitempty"`

	//AccessList types.AccessList // EIP-2930 access list.

	// For BlobTxType
	//BlobGasFeeCap *big.Int
	//BlobHashes    []common.Hash
}

func SendTransaction(ctx context.Context, cli *ethclient.Client, args TxArgs) (*full_receipt.Receipt, error) {
	var txId common.Hash
	if err := cli.Client().CallContext(ctx, &txId, "eth_sendTransaction", args); err != nil {
		return nil, err
	}
	return full_receipt.GetTransactionReceipt(ctx, cli, txId)
}
