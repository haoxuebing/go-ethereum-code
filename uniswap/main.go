package main

import (
	"fmt"
	"math/big"

	"go-ethereum/uniswap/factory"
	"go-ethereum/uniswap/pool"
	"go-ethereum/uniswap/quoter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	UniV3FactoryAddr  = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	UniV3QuoterV1Addr = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
	UniV3PoolAddr     = common.HexToAddress("0xbE3aD6a5669Dc0B8b12FeBC03608860C31E2eef6")
)

var (
	quoteTokenAddr = common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9")
	baseTokenAddr  = common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831")
)

func main() {
	endpoint_url := "wss://arbitrum-one-rpc.publicnode.com"
	client, err := ethclient.Dial(endpoint_url)
	if err != nil {
		panic(err)
	}

	selfFactory, err := factory.NewFactoryCaller(UniV3FactoryAddr, client)
	if err != nil {
		panic(err)
	}
	selfQuoter, err := quoter.NewQuoterCaller(UniV3QuoterV1Addr, client)
	if err != nil {
		panic(err)
	}

	fees := []int64{100, 500, 3000, 10000}
	amountIn := int64(100)

	for _, fee := range fees {
		poolAddr, err := selfFactory.GetPool(&bind.CallOpts{}, quoteTokenAddr, baseTokenAddr, big.NewInt(fee))
		if err != nil {
			panic(err)
		}
		fmt.Println(poolAddr)
		amountOut, err := selfQuoter.QuoteExactInputSingle(
			&bind.CallOpts{},
			quoteTokenAddr,
			baseTokenAddr,
			big.NewInt(fee),
			big.NewInt(amountIn),
			big.NewInt(0),
		)
		if err != nil {
			panic(err)
		}

		selfPool, err := pool.NewPoolCaller(poolAddr, client)
		if err != nil {
			panic(err)
		}

		fmt.Println(amountOut)

		poolFee, err := selfPool.Fee(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		fmt.Println(poolFee)
	}

	// factory, err := selfQuoter.Factory(&bind.CallOpts{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(factory)

}
