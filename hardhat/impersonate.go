package hardhat

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Impersonate(ctx context.Context, cli *ethclient.Client, addr common.Address) error {
	var ok bool
	if err := cli.Client().CallContext(ctx, &ok, "hardhat_impersonateAccount", addr); err != nil {
		return err
	}
	if !ok {
		return errors.New("nok")
	}
	return nil
}
