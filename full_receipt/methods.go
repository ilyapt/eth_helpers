package full_receipt

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetTransactionReceipt(ctx context.Context, cli *ethclient.Client, txHash common.Hash) (*Receipt, error) {
	var receipt Receipt
	if err := cli.Client().CallContext(ctx, &receipt, "eth_getTransactionReceipt", txHash); err != nil {
		return nil, err
	}
	return &receipt, nil
}

func GetBlockReceipts(ctx context.Context, cli *ethclient.Client, block *big.Int) ([]*Receipt, error) {
	var receipts []*Receipt
	if err := cli.Client().CallContext(ctx, &receipts, "eth_getBlockReceipts", (*hexutil.Big)(block)); err != nil {
		return nil, err
	}
	return receipts, nil
}
