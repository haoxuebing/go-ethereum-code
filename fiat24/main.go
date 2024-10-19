package main

import (
	"fmt"
	"math/big"

	"go-ethereum/fiat24/quoter"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var UniV3QuoterV1Addr = common.HexToAddress("0x4582f67698843Dfb6A9F195C0dDee05B0A8C973F")

func main() {
	endpoint_url := "wss://arbitrum-one-rpc.publicnode.com"
	client, err := ethclient.Dial(endpoint_url)
	if err != nil {
		panic(err)
	}
	selfQuoter, err := quoter.NewQuoter(UniV3QuoterV1Addr, client)
	if err != nil {
		panic(err)
	}

	selfQuoterRaw := quoter.QuoterRaw{
		Contract: selfQuoter,
	}

	quoteTokenAddr := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")
	baseTokenAddr := common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")
	fee := big.NewInt(500)
	amountIn := big.NewInt(1000)

	var rlt []interface{}
	err = selfQuoterRaw.Call(&bind.CallOpts{}, &rlt, "getQuote", quoteTokenAddr, baseTokenAddr, fee, amountIn)
	if err != nil {
		panic(err)
	}
	out := abi.ConvertType(rlt[0], new(big.Int)).(*big.Int)
	fmt.Println(out)
}
