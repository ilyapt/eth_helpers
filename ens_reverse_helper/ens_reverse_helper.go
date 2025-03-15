package ens_reverse_helper

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

var helper abi.ABI
var helperAddress = common.HexToAddress("0x959c0233c4f754fD10E39b1F6937A4B1CCBcC525")

func init() {
	var err error
	helper, err = abi.JSON(strings.NewReader(`[{
      "inputs": [
        {
          "internalType": "address[]",
          "name": "addr",
          "type": "address[]"
        }
      ],
      "name": "names",
      "outputs": [
        {
          "internalType": "string[]",
          "name": "",
          "type": "string[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }]`))
	if err != nil {
		panic(err)
	}
}

const method = "names"

func ReverseLookup(ctx context.Context, cli *ethclient.Client, addresses []common.Address) ([]string, error) {
	data, err := helper.Pack(method, addresses)
	if err != nil {
		return nil, err
	}
	response, err := cli.CallContract(ctx, ethereum.CallMsg{
		To:   &helperAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var result []string
	if err := helper.UnpackIntoInterface(&result, method, response); err != nil {
		return nil, err
	}
	return result, nil
}
