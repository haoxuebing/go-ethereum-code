// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quoter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// QuoterMetaData contains all meta data concerning the Quoter contract.
var QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__AddressHasNoToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Fiat24CryptoDeposit__EthRefundFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__InputTokenOutputTokenSame\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__NoPoolAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__NotOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__NotRateUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Fiat24CryptoDeposit__NotTokensWalletProvider\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__NotValidInputToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Fiat24CryptoDeposit__NotValidOutputToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Fiat24CryptoDeposit__Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Fiat24CryptoDeposit__SwapOutputAmountZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"name\":\"Fiat24CryptoDeposit__UsdcAmountHigherMaxDepositAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"Fiat24CryptoDeposit__UsdcAmountLowerMinDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Fiat24CryptoDeposit__ValueZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"walletId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedByWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedTokenViaEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedTokenViaUsd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdeur\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdchf\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgbp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcnh\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"marketClosed\",\"type\":\"bool\"}],\"name\":\"ExchangeRatesUpdatedByOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdeur\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdchf\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgbp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcnh\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"marketClosed\",\"type\":\"bool\"}],\"name\":\"ExchangeRatesUpdatedByRobot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"MoneyExchangedExactIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"}],\"name\":\"MoneyExchangedExactOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UsdcDepositAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CRYPTO_DESK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_DESK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_HUNDRET_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DIGITS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATES_UPDATER_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATES_UPDATER_ROBOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREASURY_DESK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_PERIPHERY_PAYMENTS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_QUOTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XXX24_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cnh24\",\"type\":\"address\"}],\"name\":\"addCNH24\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exchangeSpread\",\"type\":\"uint256\"}],\"name\":\"changeExchangeSpread\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketClosedSpread\",\"type\":\"uint256\"}],\"name\":\"changeMarketClosedSpread\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxUsdcDepositAmount\",\"type\":\"uint256\"}],\"name\":\"changeMaxUsdcDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minUsdcDepositAmount\",\"type\":\"uint256\"}],\"name\":\"changeMinUsdcDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slippage\",\"type\":\"uint256\"}],\"name\":\"changeSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_standardFee\",\"type\":\"uint256\"}],\"name\":\"changeStandardFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"}],\"name\":\"changeUsdcAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdcDepositAddress\",\"type\":\"address\"}],\"name\":\"changeUsdcDepositAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chf24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cnh24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdcAmount\",\"type\":\"uint256\"}],\"name\":\"depositByWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositTokenViaEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositTokenViaUsdc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eur24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exchangeRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeSpread\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"f24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"f24AirdropPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"f24AirdropStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"f24DeskAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"f24PerUSDC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"f24timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fiat24account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gbp24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"}],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usdcAmount\",\"type\":\"uint256\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeInUSDC\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"}],\"name\":\"getPoolFeeOfMostLiquidPool\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exactOut\",\"type\":\"bool\"}],\"name\":\"getSpread\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fiat24account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usd24\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eur24\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chf24\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gbp24\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_f24\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_f24timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_f24DeskAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcDepositAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketClosedSpread\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxUsdcDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minUsdcDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_inputAmount\",\"type\":\"uint256\"}],\"name\":\"moneyExchangeExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inputToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputAmount\",\"type\":\"uint256\"}],\"name\":\"moneyExchangeExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slippage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usd_eur\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usd_chf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usd_gbp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usd_cnh\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isMarketClosed\",\"type\":\"bool\"}],\"name\":\"updateExchangeRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usdc_usd24\",\"type\":\"uint256\"}],\"name\":\"updateUsdcUsd24ExchangeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usd24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validXXX24Tokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterMetaData.ABI instead.
var QuoterABI = QuoterMetaData.ABI

// Quoter is an auto generated Go binding around an Ethereum contract.
type Quoter struct {
	QuoterCaller     // Read-only binding to the contract
	QuoterTransactor // Write-only binding to the contract
	QuoterFilterer   // Log filterer for contract events
}

// QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterSession struct {
	Contract     *Quoter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterCallerSession struct {
	Contract *QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterTransactorSession struct {
	Contract     *QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterRaw struct {
	Contract *Quoter // Generic contract binding to access the raw methods on
}

// QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterCallerRaw struct {
	Contract *QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterTransactorRaw struct {
	Contract *QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoter creates a new instance of Quoter, bound to a specific deployed contract.
func NewQuoter(address common.Address, backend bind.ContractBackend) (*Quoter, error) {
	contract, err := bindQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quoter{QuoterCaller: QuoterCaller{contract: contract}, QuoterTransactor: QuoterTransactor{contract: contract}, QuoterFilterer: QuoterFilterer{contract: contract}}, nil
}

// NewQuoterCaller creates a new read-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterCaller(address common.Address, caller bind.ContractCaller) (*QuoterCaller, error) {
	contract, err := bindQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterCaller{contract: contract}, nil
}

// NewQuoterTransactor creates a new write-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoterTransactor, error) {
	contract, err := bindQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterTransactor{contract: contract}, nil
}

// NewQuoterFilterer creates a new log filterer instance of Quoter, bound to a specific deployed contract.
func NewQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoterFilterer, error) {
	contract, err := bindQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterFilterer{contract: contract}, nil
}

// bindQuoter binds a generic wrapper to an already deployed contract.
func bindQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transact(opts, method, params...)
}

// CRYPTODESK is a free data retrieval call binding the contract method 0xd0e1ec2d.
//
// Solidity: function CRYPTO_DESK() view returns(uint256)
func (_Quoter *QuoterCaller) CRYPTODESK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "CRYPTO_DESK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CRYPTODESK is a free data retrieval call binding the contract method 0xd0e1ec2d.
//
// Solidity: function CRYPTO_DESK() view returns(uint256)
func (_Quoter *QuoterSession) CRYPTODESK() (*big.Int, error) {
	return _Quoter.Contract.CRYPTODESK(&_Quoter.CallOpts)
}

// CRYPTODESK is a free data retrieval call binding the contract method 0xd0e1ec2d.
//
// Solidity: function CRYPTO_DESK() view returns(uint256)
func (_Quoter *QuoterCallerSession) CRYPTODESK() (*big.Int, error) {
	return _Quoter.Contract.CRYPTODESK(&_Quoter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Quoter *QuoterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Quoter *QuoterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Quoter.Contract.DEFAULTADMINROLE(&_Quoter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Quoter *QuoterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Quoter.Contract.DEFAULTADMINROLE(&_Quoter.CallOpts)
}

// FEEDESK is a free data retrieval call binding the contract method 0xbc1b2bcc.
//
// Solidity: function FEE_DESK() view returns(uint256)
func (_Quoter *QuoterCaller) FEEDESK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "FEE_DESK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDESK is a free data retrieval call binding the contract method 0xbc1b2bcc.
//
// Solidity: function FEE_DESK() view returns(uint256)
func (_Quoter *QuoterSession) FEEDESK() (*big.Int, error) {
	return _Quoter.Contract.FEEDESK(&_Quoter.CallOpts)
}

// FEEDESK is a free data retrieval call binding the contract method 0xbc1b2bcc.
//
// Solidity: function FEE_DESK() view returns(uint256)
func (_Quoter *QuoterCallerSession) FEEDESK() (*big.Int, error) {
	return _Quoter.Contract.FEEDESK(&_Quoter.CallOpts)
}

// FEEHUNDRETPERCENT is a free data retrieval call binding the contract method 0xc5881e21.
//
// Solidity: function FEE_HUNDRET_PERCENT() view returns(uint256)
func (_Quoter *QuoterCaller) FEEHUNDRETPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "FEE_HUNDRET_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEHUNDRETPERCENT is a free data retrieval call binding the contract method 0xc5881e21.
//
// Solidity: function FEE_HUNDRET_PERCENT() view returns(uint256)
func (_Quoter *QuoterSession) FEEHUNDRETPERCENT() (*big.Int, error) {
	return _Quoter.Contract.FEEHUNDRETPERCENT(&_Quoter.CallOpts)
}

// FEEHUNDRETPERCENT is a free data retrieval call binding the contract method 0xc5881e21.
//
// Solidity: function FEE_HUNDRET_PERCENT() view returns(uint256)
func (_Quoter *QuoterCallerSession) FEEHUNDRETPERCENT() (*big.Int, error) {
	return _Quoter.Contract.FEEHUNDRETPERCENT(&_Quoter.CallOpts)
}

// MAXDIGITS is a free data retrieval call binding the contract method 0x24fb9edf.
//
// Solidity: function MAX_DIGITS() view returns(uint256)
func (_Quoter *QuoterCaller) MAXDIGITS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "MAX_DIGITS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDIGITS is a free data retrieval call binding the contract method 0x24fb9edf.
//
// Solidity: function MAX_DIGITS() view returns(uint256)
func (_Quoter *QuoterSession) MAXDIGITS() (*big.Int, error) {
	return _Quoter.Contract.MAXDIGITS(&_Quoter.CallOpts)
}

// MAXDIGITS is a free data retrieval call binding the contract method 0x24fb9edf.
//
// Solidity: function MAX_DIGITS() view returns(uint256)
func (_Quoter *QuoterCallerSession) MAXDIGITS() (*big.Int, error) {
	return _Quoter.Contract.MAXDIGITS(&_Quoter.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Quoter *QuoterCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Quoter *QuoterSession) OPERATORROLE() ([32]byte, error) {
	return _Quoter.Contract.OPERATORROLE(&_Quoter.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Quoter *QuoterCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Quoter.Contract.OPERATORROLE(&_Quoter.CallOpts)
}

// RATESUPDATEROPERATORROLE is a free data retrieval call binding the contract method 0x55275af5.
//
// Solidity: function RATES_UPDATER_OPERATOR_ROLE() view returns(bytes32)
func (_Quoter *QuoterCaller) RATESUPDATEROPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "RATES_UPDATER_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATESUPDATEROPERATORROLE is a free data retrieval call binding the contract method 0x55275af5.
//
// Solidity: function RATES_UPDATER_OPERATOR_ROLE() view returns(bytes32)
func (_Quoter *QuoterSession) RATESUPDATEROPERATORROLE() ([32]byte, error) {
	return _Quoter.Contract.RATESUPDATEROPERATORROLE(&_Quoter.CallOpts)
}

// RATESUPDATEROPERATORROLE is a free data retrieval call binding the contract method 0x55275af5.
//
// Solidity: function RATES_UPDATER_OPERATOR_ROLE() view returns(bytes32)
func (_Quoter *QuoterCallerSession) RATESUPDATEROPERATORROLE() ([32]byte, error) {
	return _Quoter.Contract.RATESUPDATEROPERATORROLE(&_Quoter.CallOpts)
}

// RATESUPDATERROBOTROLE is a free data retrieval call binding the contract method 0xb4a08b6b.
//
// Solidity: function RATES_UPDATER_ROBOT_ROLE() view returns(bytes32)
func (_Quoter *QuoterCaller) RATESUPDATERROBOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "RATES_UPDATER_ROBOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATESUPDATERROBOTROLE is a free data retrieval call binding the contract method 0xb4a08b6b.
//
// Solidity: function RATES_UPDATER_ROBOT_ROLE() view returns(bytes32)
func (_Quoter *QuoterSession) RATESUPDATERROBOTROLE() ([32]byte, error) {
	return _Quoter.Contract.RATESUPDATERROBOTROLE(&_Quoter.CallOpts)
}

// RATESUPDATERROBOTROLE is a free data retrieval call binding the contract method 0xb4a08b6b.
//
// Solidity: function RATES_UPDATER_ROBOT_ROLE() view returns(bytes32)
func (_Quoter *QuoterCallerSession) RATESUPDATERROBOTROLE() ([32]byte, error) {
	return _Quoter.Contract.RATESUPDATERROBOTROLE(&_Quoter.CallOpts)
}

// TREASURYDESK is a free data retrieval call binding the contract method 0x0f2f45d2.
//
// Solidity: function TREASURY_DESK() view returns(uint256)
func (_Quoter *QuoterCaller) TREASURYDESK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "TREASURY_DESK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TREASURYDESK is a free data retrieval call binding the contract method 0x0f2f45d2.
//
// Solidity: function TREASURY_DESK() view returns(uint256)
func (_Quoter *QuoterSession) TREASURYDESK() (*big.Int, error) {
	return _Quoter.Contract.TREASURYDESK(&_Quoter.CallOpts)
}

// TREASURYDESK is a free data retrieval call binding the contract method 0x0f2f45d2.
//
// Solidity: function TREASURY_DESK() view returns(uint256)
func (_Quoter *QuoterCallerSession) TREASURYDESK() (*big.Int, error) {
	return _Quoter.Contract.TREASURYDESK(&_Quoter.CallOpts)
}

// UNISWAPFACTORY is a free data retrieval call binding the contract method 0xc74c0fac.
//
// Solidity: function UNISWAP_FACTORY() view returns(address)
func (_Quoter *QuoterCaller) UNISWAPFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "UNISWAP_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPFACTORY is a free data retrieval call binding the contract method 0xc74c0fac.
//
// Solidity: function UNISWAP_FACTORY() view returns(address)
func (_Quoter *QuoterSession) UNISWAPFACTORY() (common.Address, error) {
	return _Quoter.Contract.UNISWAPFACTORY(&_Quoter.CallOpts)
}

// UNISWAPFACTORY is a free data retrieval call binding the contract method 0xc74c0fac.
//
// Solidity: function UNISWAP_FACTORY() view returns(address)
func (_Quoter *QuoterCallerSession) UNISWAPFACTORY() (common.Address, error) {
	return _Quoter.Contract.UNISWAPFACTORY(&_Quoter.CallOpts)
}

// UNISWAPPERIPHERYPAYMENTS is a free data retrieval call binding the contract method 0x063286cc.
//
// Solidity: function UNISWAP_PERIPHERY_PAYMENTS() view returns(address)
func (_Quoter *QuoterCaller) UNISWAPPERIPHERYPAYMENTS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "UNISWAP_PERIPHERY_PAYMENTS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPPERIPHERYPAYMENTS is a free data retrieval call binding the contract method 0x063286cc.
//
// Solidity: function UNISWAP_PERIPHERY_PAYMENTS() view returns(address)
func (_Quoter *QuoterSession) UNISWAPPERIPHERYPAYMENTS() (common.Address, error) {
	return _Quoter.Contract.UNISWAPPERIPHERYPAYMENTS(&_Quoter.CallOpts)
}

// UNISWAPPERIPHERYPAYMENTS is a free data retrieval call binding the contract method 0x063286cc.
//
// Solidity: function UNISWAP_PERIPHERY_PAYMENTS() view returns(address)
func (_Quoter *QuoterCallerSession) UNISWAPPERIPHERYPAYMENTS() (common.Address, error) {
	return _Quoter.Contract.UNISWAPPERIPHERYPAYMENTS(&_Quoter.CallOpts)
}

// UNISWAPQUOTER is a free data retrieval call binding the contract method 0xb268630b.
//
// Solidity: function UNISWAP_QUOTER() view returns(address)
func (_Quoter *QuoterCaller) UNISWAPQUOTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "UNISWAP_QUOTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPQUOTER is a free data retrieval call binding the contract method 0xb268630b.
//
// Solidity: function UNISWAP_QUOTER() view returns(address)
func (_Quoter *QuoterSession) UNISWAPQUOTER() (common.Address, error) {
	return _Quoter.Contract.UNISWAPQUOTER(&_Quoter.CallOpts)
}

// UNISWAPQUOTER is a free data retrieval call binding the contract method 0xb268630b.
//
// Solidity: function UNISWAP_QUOTER() view returns(address)
func (_Quoter *QuoterCallerSession) UNISWAPQUOTER() (common.Address, error) {
	return _Quoter.Contract.UNISWAPQUOTER(&_Quoter.CallOpts)
}

// UNISWAPROUTER is a free data retrieval call binding the contract method 0xd8264920.
//
// Solidity: function UNISWAP_ROUTER() view returns(address)
func (_Quoter *QuoterCaller) UNISWAPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "UNISWAP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPROUTER is a free data retrieval call binding the contract method 0xd8264920.
//
// Solidity: function UNISWAP_ROUTER() view returns(address)
func (_Quoter *QuoterSession) UNISWAPROUTER() (common.Address, error) {
	return _Quoter.Contract.UNISWAPROUTER(&_Quoter.CallOpts)
}

// UNISWAPROUTER is a free data retrieval call binding the contract method 0xd8264920.
//
// Solidity: function UNISWAP_ROUTER() view returns(address)
func (_Quoter *QuoterCallerSession) UNISWAPROUTER() (common.Address, error) {
	return _Quoter.Contract.UNISWAPROUTER(&_Quoter.CallOpts)
}

// USDCDIVISOR is a free data retrieval call binding the contract method 0x8c9b6c28.
//
// Solidity: function USDC_DIVISOR() view returns(uint256)
func (_Quoter *QuoterCaller) USDCDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "USDC_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDCDIVISOR is a free data retrieval call binding the contract method 0x8c9b6c28.
//
// Solidity: function USDC_DIVISOR() view returns(uint256)
func (_Quoter *QuoterSession) USDCDIVISOR() (*big.Int, error) {
	return _Quoter.Contract.USDCDIVISOR(&_Quoter.CallOpts)
}

// USDCDIVISOR is a free data retrieval call binding the contract method 0x8c9b6c28.
//
// Solidity: function USDC_DIVISOR() view returns(uint256)
func (_Quoter *QuoterCallerSession) USDCDIVISOR() (*big.Int, error) {
	return _Quoter.Contract.USDCDIVISOR(&_Quoter.CallOpts)
}

// XXX24DIVISOR is a free data retrieval call binding the contract method 0x1557c3b4.
//
// Solidity: function XXX24_DIVISOR() view returns(uint256)
func (_Quoter *QuoterCaller) XXX24DIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "XXX24_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XXX24DIVISOR is a free data retrieval call binding the contract method 0x1557c3b4.
//
// Solidity: function XXX24_DIVISOR() view returns(uint256)
func (_Quoter *QuoterSession) XXX24DIVISOR() (*big.Int, error) {
	return _Quoter.Contract.XXX24DIVISOR(&_Quoter.CallOpts)
}

// XXX24DIVISOR is a free data retrieval call binding the contract method 0x1557c3b4.
//
// Solidity: function XXX24_DIVISOR() view returns(uint256)
func (_Quoter *QuoterCallerSession) XXX24DIVISOR() (*big.Int, error) {
	return _Quoter.Contract.XXX24DIVISOR(&_Quoter.CallOpts)
}

// Chf24 is a free data retrieval call binding the contract method 0xaee6b443.
//
// Solidity: function chf24() view returns(address)
func (_Quoter *QuoterCaller) Chf24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "chf24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chf24 is a free data retrieval call binding the contract method 0xaee6b443.
//
// Solidity: function chf24() view returns(address)
func (_Quoter *QuoterSession) Chf24() (common.Address, error) {
	return _Quoter.Contract.Chf24(&_Quoter.CallOpts)
}

// Chf24 is a free data retrieval call binding the contract method 0xaee6b443.
//
// Solidity: function chf24() view returns(address)
func (_Quoter *QuoterCallerSession) Chf24() (common.Address, error) {
	return _Quoter.Contract.Chf24(&_Quoter.CallOpts)
}

// Cnh24 is a free data retrieval call binding the contract method 0xc0154ccb.
//
// Solidity: function cnh24() view returns(address)
func (_Quoter *QuoterCaller) Cnh24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "cnh24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cnh24 is a free data retrieval call binding the contract method 0xc0154ccb.
//
// Solidity: function cnh24() view returns(address)
func (_Quoter *QuoterSession) Cnh24() (common.Address, error) {
	return _Quoter.Contract.Cnh24(&_Quoter.CallOpts)
}

// Cnh24 is a free data retrieval call binding the contract method 0xc0154ccb.
//
// Solidity: function cnh24() view returns(address)
func (_Quoter *QuoterCallerSession) Cnh24() (common.Address, error) {
	return _Quoter.Contract.Cnh24(&_Quoter.CallOpts)
}

// Eur24 is a free data retrieval call binding the contract method 0x55e64444.
//
// Solidity: function eur24() view returns(address)
func (_Quoter *QuoterCaller) Eur24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "eur24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eur24 is a free data retrieval call binding the contract method 0x55e64444.
//
// Solidity: function eur24() view returns(address)
func (_Quoter *QuoterSession) Eur24() (common.Address, error) {
	return _Quoter.Contract.Eur24(&_Quoter.CallOpts)
}

// Eur24 is a free data retrieval call binding the contract method 0x55e64444.
//
// Solidity: function eur24() view returns(address)
func (_Quoter *QuoterCallerSession) Eur24() (common.Address, error) {
	return _Quoter.Contract.Eur24(&_Quoter.CallOpts)
}

// ExchangeRates is a free data retrieval call binding the contract method 0x503fa44f.
//
// Solidity: function exchangeRates(address , address ) view returns(uint256)
func (_Quoter *QuoterCaller) ExchangeRates(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "exchangeRates", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRates is a free data retrieval call binding the contract method 0x503fa44f.
//
// Solidity: function exchangeRates(address , address ) view returns(uint256)
func (_Quoter *QuoterSession) ExchangeRates(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Quoter.Contract.ExchangeRates(&_Quoter.CallOpts, arg0, arg1)
}

// ExchangeRates is a free data retrieval call binding the contract method 0x503fa44f.
//
// Solidity: function exchangeRates(address , address ) view returns(uint256)
func (_Quoter *QuoterCallerSession) ExchangeRates(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Quoter.Contract.ExchangeRates(&_Quoter.CallOpts, arg0, arg1)
}

// ExchangeSpread is a free data retrieval call binding the contract method 0xacd0032d.
//
// Solidity: function exchangeSpread() view returns(uint256)
func (_Quoter *QuoterCaller) ExchangeSpread(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "exchangeSpread")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeSpread is a free data retrieval call binding the contract method 0xacd0032d.
//
// Solidity: function exchangeSpread() view returns(uint256)
func (_Quoter *QuoterSession) ExchangeSpread() (*big.Int, error) {
	return _Quoter.Contract.ExchangeSpread(&_Quoter.CallOpts)
}

// ExchangeSpread is a free data retrieval call binding the contract method 0xacd0032d.
//
// Solidity: function exchangeSpread() view returns(uint256)
func (_Quoter *QuoterCallerSession) ExchangeSpread() (*big.Int, error) {
	return _Quoter.Contract.ExchangeSpread(&_Quoter.CallOpts)
}

// F24 is a free data retrieval call binding the contract method 0xcbf99d38.
//
// Solidity: function f24() view returns(address)
func (_Quoter *QuoterCaller) F24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "f24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// F24 is a free data retrieval call binding the contract method 0xcbf99d38.
//
// Solidity: function f24() view returns(address)
func (_Quoter *QuoterSession) F24() (common.Address, error) {
	return _Quoter.Contract.F24(&_Quoter.CallOpts)
}

// F24 is a free data retrieval call binding the contract method 0xcbf99d38.
//
// Solidity: function f24() view returns(address)
func (_Quoter *QuoterCallerSession) F24() (common.Address, error) {
	return _Quoter.Contract.F24(&_Quoter.CallOpts)
}

// F24AirdropPaused is a free data retrieval call binding the contract method 0xf7f15445.
//
// Solidity: function f24AirdropPaused() view returns(bool)
func (_Quoter *QuoterCaller) F24AirdropPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "f24AirdropPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// F24AirdropPaused is a free data retrieval call binding the contract method 0xf7f15445.
//
// Solidity: function f24AirdropPaused() view returns(bool)
func (_Quoter *QuoterSession) F24AirdropPaused() (bool, error) {
	return _Quoter.Contract.F24AirdropPaused(&_Quoter.CallOpts)
}

// F24AirdropPaused is a free data retrieval call binding the contract method 0xf7f15445.
//
// Solidity: function f24AirdropPaused() view returns(bool)
func (_Quoter *QuoterCallerSession) F24AirdropPaused() (bool, error) {
	return _Quoter.Contract.F24AirdropPaused(&_Quoter.CallOpts)
}

// F24AirdropStart is a free data retrieval call binding the contract method 0x8f6ac994.
//
// Solidity: function f24AirdropStart() view returns(uint256)
func (_Quoter *QuoterCaller) F24AirdropStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "f24AirdropStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// F24AirdropStart is a free data retrieval call binding the contract method 0x8f6ac994.
//
// Solidity: function f24AirdropStart() view returns(uint256)
func (_Quoter *QuoterSession) F24AirdropStart() (*big.Int, error) {
	return _Quoter.Contract.F24AirdropStart(&_Quoter.CallOpts)
}

// F24AirdropStart is a free data retrieval call binding the contract method 0x8f6ac994.
//
// Solidity: function f24AirdropStart() view returns(uint256)
func (_Quoter *QuoterCallerSession) F24AirdropStart() (*big.Int, error) {
	return _Quoter.Contract.F24AirdropStart(&_Quoter.CallOpts)
}

// F24DeskAddress is a free data retrieval call binding the contract method 0x0fbf0944.
//
// Solidity: function f24DeskAddress() view returns(address)
func (_Quoter *QuoterCaller) F24DeskAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "f24DeskAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// F24DeskAddress is a free data retrieval call binding the contract method 0x0fbf0944.
//
// Solidity: function f24DeskAddress() view returns(address)
func (_Quoter *QuoterSession) F24DeskAddress() (common.Address, error) {
	return _Quoter.Contract.F24DeskAddress(&_Quoter.CallOpts)
}

// F24DeskAddress is a free data retrieval call binding the contract method 0x0fbf0944.
//
// Solidity: function f24DeskAddress() view returns(address)
func (_Quoter *QuoterCallerSession) F24DeskAddress() (common.Address, error) {
	return _Quoter.Contract.F24DeskAddress(&_Quoter.CallOpts)
}

// F24PerUSDC is a free data retrieval call binding the contract method 0x04f26978.
//
// Solidity: function f24PerUSDC() view returns(uint256)
func (_Quoter *QuoterCaller) F24PerUSDC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "f24PerUSDC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// F24PerUSDC is a free data retrieval call binding the contract method 0x04f26978.
//
// Solidity: function f24PerUSDC() view returns(uint256)
func (_Quoter *QuoterSession) F24PerUSDC() (*big.Int, error) {
	return _Quoter.Contract.F24PerUSDC(&_Quoter.CallOpts)
}

// F24PerUSDC is a free data retrieval call binding the contract method 0x04f26978.
//
// Solidity: function f24PerUSDC() view returns(uint256)
func (_Quoter *QuoterCallerSession) F24PerUSDC() (*big.Int, error) {
	return _Quoter.Contract.F24PerUSDC(&_Quoter.CallOpts)
}

// F24timelock is a free data retrieval call binding the contract method 0x20af28c5.
//
// Solidity: function f24timelock() view returns(address)
func (_Quoter *QuoterCaller) F24timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "f24timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// F24timelock is a free data retrieval call binding the contract method 0x20af28c5.
//
// Solidity: function f24timelock() view returns(address)
func (_Quoter *QuoterSession) F24timelock() (common.Address, error) {
	return _Quoter.Contract.F24timelock(&_Quoter.CallOpts)
}

// F24timelock is a free data retrieval call binding the contract method 0x20af28c5.
//
// Solidity: function f24timelock() view returns(address)
func (_Quoter *QuoterCallerSession) F24timelock() (common.Address, error) {
	return _Quoter.Contract.F24timelock(&_Quoter.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_Quoter *QuoterCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "fees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_Quoter *QuoterSession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _Quoter.Contract.Fees(&_Quoter.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_Quoter *QuoterCallerSession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _Quoter.Contract.Fees(&_Quoter.CallOpts, arg0)
}

// Fiat24account is a free data retrieval call binding the contract method 0xb6a09e25.
//
// Solidity: function fiat24account() view returns(address)
func (_Quoter *QuoterCaller) Fiat24account(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "fiat24account")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fiat24account is a free data retrieval call binding the contract method 0xb6a09e25.
//
// Solidity: function fiat24account() view returns(address)
func (_Quoter *QuoterSession) Fiat24account() (common.Address, error) {
	return _Quoter.Contract.Fiat24account(&_Quoter.CallOpts)
}

// Fiat24account is a free data retrieval call binding the contract method 0xb6a09e25.
//
// Solidity: function fiat24account() view returns(address)
func (_Quoter *QuoterCallerSession) Fiat24account() (common.Address, error) {
	return _Quoter.Contract.Fiat24account(&_Quoter.CallOpts)
}

// Gbp24 is a free data retrieval call binding the contract method 0x1fb0292f.
//
// Solidity: function gbp24() view returns(address)
func (_Quoter *QuoterCaller) Gbp24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "gbp24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gbp24 is a free data retrieval call binding the contract method 0x1fb0292f.
//
// Solidity: function gbp24() view returns(address)
func (_Quoter *QuoterSession) Gbp24() (common.Address, error) {
	return _Quoter.Contract.Gbp24(&_Quoter.CallOpts)
}

// Gbp24 is a free data retrieval call binding the contract method 0x1fb0292f.
//
// Solidity: function gbp24() view returns(address)
func (_Quoter *QuoterCallerSession) Gbp24() (common.Address, error) {
	return _Quoter.Contract.Gbp24(&_Quoter.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xbaaa61be.
//
// Solidity: function getExchangeRate(address _inputToken, address _outputToken) view returns(uint256)
func (_Quoter *QuoterCaller) GetExchangeRate(opts *bind.CallOpts, _inputToken common.Address, _outputToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "getExchangeRate", _inputToken, _outputToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xbaaa61be.
//
// Solidity: function getExchangeRate(address _inputToken, address _outputToken) view returns(uint256)
func (_Quoter *QuoterSession) GetExchangeRate(_inputToken common.Address, _outputToken common.Address) (*big.Int, error) {
	return _Quoter.Contract.GetExchangeRate(&_Quoter.CallOpts, _inputToken, _outputToken)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xbaaa61be.
//
// Solidity: function getExchangeRate(address _inputToken, address _outputToken) view returns(uint256)
func (_Quoter *QuoterCallerSession) GetExchangeRate(_inputToken common.Address, _outputToken common.Address) (*big.Int, error) {
	return _Quoter.Contract.GetExchangeRate(&_Quoter.CallOpts, _inputToken, _outputToken)
}

// GetFee is a free data retrieval call binding the contract method 0xd250185c.
//
// Solidity: function getFee(uint256 _tokenId, uint256 _usdcAmount) view returns(uint256 feeInUSDC)
func (_Quoter *QuoterCaller) GetFee(opts *bind.CallOpts, _tokenId *big.Int, _usdcAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "getFee", _tokenId, _usdcAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xd250185c.
//
// Solidity: function getFee(uint256 _tokenId, uint256 _usdcAmount) view returns(uint256 feeInUSDC)
func (_Quoter *QuoterSession) GetFee(_tokenId *big.Int, _usdcAmount *big.Int) (*big.Int, error) {
	return _Quoter.Contract.GetFee(&_Quoter.CallOpts, _tokenId, _usdcAmount)
}

// GetFee is a free data retrieval call binding the contract method 0xd250185c.
//
// Solidity: function getFee(uint256 _tokenId, uint256 _usdcAmount) view returns(uint256 feeInUSDC)
func (_Quoter *QuoterCallerSession) GetFee(_tokenId *big.Int, _usdcAmount *big.Int) (*big.Int, error) {
	return _Quoter.Contract.GetFee(&_Quoter.CallOpts, _tokenId, _usdcAmount)
}

// GetPoolFeeOfMostLiquidPool is a free data retrieval call binding the contract method 0x2c7692bd.
//
// Solidity: function getPoolFeeOfMostLiquidPool(address _inputToken, address _outputToken) view returns(uint24)
func (_Quoter *QuoterCaller) GetPoolFeeOfMostLiquidPool(opts *bind.CallOpts, _inputToken common.Address, _outputToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "getPoolFeeOfMostLiquidPool", _inputToken, _outputToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolFeeOfMostLiquidPool is a free data retrieval call binding the contract method 0x2c7692bd.
//
// Solidity: function getPoolFeeOfMostLiquidPool(address _inputToken, address _outputToken) view returns(uint24)
func (_Quoter *QuoterSession) GetPoolFeeOfMostLiquidPool(_inputToken common.Address, _outputToken common.Address) (*big.Int, error) {
	return _Quoter.Contract.GetPoolFeeOfMostLiquidPool(&_Quoter.CallOpts, _inputToken, _outputToken)
}

// GetPoolFeeOfMostLiquidPool is a free data retrieval call binding the contract method 0x2c7692bd.
//
// Solidity: function getPoolFeeOfMostLiquidPool(address _inputToken, address _outputToken) view returns(uint24)
func (_Quoter *QuoterCallerSession) GetPoolFeeOfMostLiquidPool(_inputToken common.Address, _outputToken common.Address) (*big.Int, error) {
	return _Quoter.Contract.GetPoolFeeOfMostLiquidPool(&_Quoter.CallOpts, _inputToken, _outputToken)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Quoter *QuoterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Quoter *QuoterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Quoter.Contract.GetRoleAdmin(&_Quoter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Quoter *QuoterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Quoter.Contract.GetRoleAdmin(&_Quoter.CallOpts, role)
}

// GetSpread is a free data retrieval call binding the contract method 0x8f1a3879.
//
// Solidity: function getSpread(address _inputToken, address _outputToken, bool exactOut) view returns(uint256)
func (_Quoter *QuoterCaller) GetSpread(opts *bind.CallOpts, _inputToken common.Address, _outputToken common.Address, exactOut bool) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "getSpread", _inputToken, _outputToken, exactOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpread is a free data retrieval call binding the contract method 0x8f1a3879.
//
// Solidity: function getSpread(address _inputToken, address _outputToken, bool exactOut) view returns(uint256)
func (_Quoter *QuoterSession) GetSpread(_inputToken common.Address, _outputToken common.Address, exactOut bool) (*big.Int, error) {
	return _Quoter.Contract.GetSpread(&_Quoter.CallOpts, _inputToken, _outputToken, exactOut)
}

// GetSpread is a free data retrieval call binding the contract method 0x8f1a3879.
//
// Solidity: function getSpread(address _inputToken, address _outputToken, bool exactOut) view returns(uint256)
func (_Quoter *QuoterCallerSession) GetSpread(_inputToken common.Address, _outputToken common.Address, exactOut bool) (*big.Int, error) {
	return _Quoter.Contract.GetSpread(&_Quoter.CallOpts, _inputToken, _outputToken, exactOut)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Quoter *QuoterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Quoter *QuoterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Quoter.Contract.HasRole(&_Quoter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Quoter *QuoterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Quoter.Contract.HasRole(&_Quoter.CallOpts, role, account)
}

// MarketClosed is a free data retrieval call binding the contract method 0x48a5dec6.
//
// Solidity: function marketClosed() view returns(bool)
func (_Quoter *QuoterCaller) MarketClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "marketClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketClosed is a free data retrieval call binding the contract method 0x48a5dec6.
//
// Solidity: function marketClosed() view returns(bool)
func (_Quoter *QuoterSession) MarketClosed() (bool, error) {
	return _Quoter.Contract.MarketClosed(&_Quoter.CallOpts)
}

// MarketClosed is a free data retrieval call binding the contract method 0x48a5dec6.
//
// Solidity: function marketClosed() view returns(bool)
func (_Quoter *QuoterCallerSession) MarketClosed() (bool, error) {
	return _Quoter.Contract.MarketClosed(&_Quoter.CallOpts)
}

// MarketClosedSpread is a free data retrieval call binding the contract method 0x41f33e19.
//
// Solidity: function marketClosedSpread() view returns(uint256)
func (_Quoter *QuoterCaller) MarketClosedSpread(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "marketClosedSpread")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketClosedSpread is a free data retrieval call binding the contract method 0x41f33e19.
//
// Solidity: function marketClosedSpread() view returns(uint256)
func (_Quoter *QuoterSession) MarketClosedSpread() (*big.Int, error) {
	return _Quoter.Contract.MarketClosedSpread(&_Quoter.CallOpts)
}

// MarketClosedSpread is a free data retrieval call binding the contract method 0x41f33e19.
//
// Solidity: function marketClosedSpread() view returns(uint256)
func (_Quoter *QuoterCallerSession) MarketClosedSpread() (*big.Int, error) {
	return _Quoter.Contract.MarketClosedSpread(&_Quoter.CallOpts)
}

// MaxUsdcDepositAmount is a free data retrieval call binding the contract method 0x972d0bba.
//
// Solidity: function maxUsdcDepositAmount() view returns(uint256)
func (_Quoter *QuoterCaller) MaxUsdcDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "maxUsdcDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdcDepositAmount is a free data retrieval call binding the contract method 0x972d0bba.
//
// Solidity: function maxUsdcDepositAmount() view returns(uint256)
func (_Quoter *QuoterSession) MaxUsdcDepositAmount() (*big.Int, error) {
	return _Quoter.Contract.MaxUsdcDepositAmount(&_Quoter.CallOpts)
}

// MaxUsdcDepositAmount is a free data retrieval call binding the contract method 0x972d0bba.
//
// Solidity: function maxUsdcDepositAmount() view returns(uint256)
func (_Quoter *QuoterCallerSession) MaxUsdcDepositAmount() (*big.Int, error) {
	return _Quoter.Contract.MaxUsdcDepositAmount(&_Quoter.CallOpts)
}

// MinUsdcDepositAmount is a free data retrieval call binding the contract method 0x0caa0540.
//
// Solidity: function minUsdcDepositAmount() view returns(uint256)
func (_Quoter *QuoterCaller) MinUsdcDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "minUsdcDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinUsdcDepositAmount is a free data retrieval call binding the contract method 0x0caa0540.
//
// Solidity: function minUsdcDepositAmount() view returns(uint256)
func (_Quoter *QuoterSession) MinUsdcDepositAmount() (*big.Int, error) {
	return _Quoter.Contract.MinUsdcDepositAmount(&_Quoter.CallOpts)
}

// MinUsdcDepositAmount is a free data retrieval call binding the contract method 0x0caa0540.
//
// Solidity: function minUsdcDepositAmount() view returns(uint256)
func (_Quoter *QuoterCallerSession) MinUsdcDepositAmount() (*big.Int, error) {
	return _Quoter.Contract.MinUsdcDepositAmount(&_Quoter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Quoter *QuoterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Quoter *QuoterSession) Paused() (bool, error) {
	return _Quoter.Contract.Paused(&_Quoter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Quoter *QuoterCallerSession) Paused() (bool, error) {
	return _Quoter.Contract.Paused(&_Quoter.CallOpts)
}

// Slippage is a free data retrieval call binding the contract method 0x3e032a3b.
//
// Solidity: function slippage() view returns(uint256)
func (_Quoter *QuoterCaller) Slippage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "slippage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Slippage is a free data retrieval call binding the contract method 0x3e032a3b.
//
// Solidity: function slippage() view returns(uint256)
func (_Quoter *QuoterSession) Slippage() (*big.Int, error) {
	return _Quoter.Contract.Slippage(&_Quoter.CallOpts)
}

// Slippage is a free data retrieval call binding the contract method 0x3e032a3b.
//
// Solidity: function slippage() view returns(uint256)
func (_Quoter *QuoterCallerSession) Slippage() (*big.Int, error) {
	return _Quoter.Contract.Slippage(&_Quoter.CallOpts)
}

// StandardFee is a free data retrieval call binding the contract method 0xf8140a7e.
//
// Solidity: function standardFee() view returns(uint256)
func (_Quoter *QuoterCaller) StandardFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "standardFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StandardFee is a free data retrieval call binding the contract method 0xf8140a7e.
//
// Solidity: function standardFee() view returns(uint256)
func (_Quoter *QuoterSession) StandardFee() (*big.Int, error) {
	return _Quoter.Contract.StandardFee(&_Quoter.CallOpts)
}

// StandardFee is a free data retrieval call binding the contract method 0xf8140a7e.
//
// Solidity: function standardFee() view returns(uint256)
func (_Quoter *QuoterCallerSession) StandardFee() (*big.Int, error) {
	return _Quoter.Contract.StandardFee(&_Quoter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Quoter *QuoterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Quoter *QuoterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Quoter.Contract.SupportsInterface(&_Quoter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Quoter *QuoterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Quoter.Contract.SupportsInterface(&_Quoter.CallOpts, interfaceId)
}

// Usd24 is a free data retrieval call binding the contract method 0x44c3b29f.
//
// Solidity: function usd24() view returns(address)
func (_Quoter *QuoterCaller) Usd24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "usd24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usd24 is a free data retrieval call binding the contract method 0x44c3b29f.
//
// Solidity: function usd24() view returns(address)
func (_Quoter *QuoterSession) Usd24() (common.Address, error) {
	return _Quoter.Contract.Usd24(&_Quoter.CallOpts)
}

// Usd24 is a free data retrieval call binding the contract method 0x44c3b29f.
//
// Solidity: function usd24() view returns(address)
func (_Quoter *QuoterCallerSession) Usd24() (common.Address, error) {
	return _Quoter.Contract.Usd24(&_Quoter.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Quoter *QuoterCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Quoter *QuoterSession) Usdc() (common.Address, error) {
	return _Quoter.Contract.Usdc(&_Quoter.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Quoter *QuoterCallerSession) Usdc() (common.Address, error) {
	return _Quoter.Contract.Usdc(&_Quoter.CallOpts)
}

// UsdcDepositAddress is a free data retrieval call binding the contract method 0x983318d9.
//
// Solidity: function usdcDepositAddress() view returns(address)
func (_Quoter *QuoterCaller) UsdcDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "usdcDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcDepositAddress is a free data retrieval call binding the contract method 0x983318d9.
//
// Solidity: function usdcDepositAddress() view returns(address)
func (_Quoter *QuoterSession) UsdcDepositAddress() (common.Address, error) {
	return _Quoter.Contract.UsdcDepositAddress(&_Quoter.CallOpts)
}

// UsdcDepositAddress is a free data retrieval call binding the contract method 0x983318d9.
//
// Solidity: function usdcDepositAddress() view returns(address)
func (_Quoter *QuoterCallerSession) UsdcDepositAddress() (common.Address, error) {
	return _Quoter.Contract.UsdcDepositAddress(&_Quoter.CallOpts)
}

// ValidXXX24Tokens is a free data retrieval call binding the contract method 0x1b723877.
//
// Solidity: function validXXX24Tokens(address ) view returns(bool)
func (_Quoter *QuoterCaller) ValidXXX24Tokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "validXXX24Tokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidXXX24Tokens is a free data retrieval call binding the contract method 0x1b723877.
//
// Solidity: function validXXX24Tokens(address ) view returns(bool)
func (_Quoter *QuoterSession) ValidXXX24Tokens(arg0 common.Address) (bool, error) {
	return _Quoter.Contract.ValidXXX24Tokens(&_Quoter.CallOpts, arg0)
}

// ValidXXX24Tokens is a free data retrieval call binding the contract method 0x1b723877.
//
// Solidity: function validXXX24Tokens(address ) view returns(bool)
func (_Quoter *QuoterCallerSession) ValidXXX24Tokens(arg0 common.Address) (bool, error) {
	return _Quoter.Contract.ValidXXX24Tokens(&_Quoter.CallOpts, arg0)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Quoter *QuoterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Quoter *QuoterSession) Weth() (common.Address, error) {
	return _Quoter.Contract.Weth(&_Quoter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Quoter *QuoterCallerSession) Weth() (common.Address, error) {
	return _Quoter.Contract.Weth(&_Quoter.CallOpts)
}

// AddCNH24 is a paid mutator transaction binding the contract method 0x3c3ef2fb.
//
// Solidity: function addCNH24(address _cnh24) returns()
func (_Quoter *QuoterTransactor) AddCNH24(opts *bind.TransactOpts, _cnh24 common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "addCNH24", _cnh24)
}

// AddCNH24 is a paid mutator transaction binding the contract method 0x3c3ef2fb.
//
// Solidity: function addCNH24(address _cnh24) returns()
func (_Quoter *QuoterSession) AddCNH24(_cnh24 common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.AddCNH24(&_Quoter.TransactOpts, _cnh24)
}

// AddCNH24 is a paid mutator transaction binding the contract method 0x3c3ef2fb.
//
// Solidity: function addCNH24(address _cnh24) returns()
func (_Quoter *QuoterTransactorSession) AddCNH24(_cnh24 common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.AddCNH24(&_Quoter.TransactOpts, _cnh24)
}

// ChangeExchangeSpread is a paid mutator transaction binding the contract method 0x1118bb3b.
//
// Solidity: function changeExchangeSpread(uint256 _exchangeSpread) returns()
func (_Quoter *QuoterTransactor) ChangeExchangeSpread(opts *bind.TransactOpts, _exchangeSpread *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeExchangeSpread", _exchangeSpread)
}

// ChangeExchangeSpread is a paid mutator transaction binding the contract method 0x1118bb3b.
//
// Solidity: function changeExchangeSpread(uint256 _exchangeSpread) returns()
func (_Quoter *QuoterSession) ChangeExchangeSpread(_exchangeSpread *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeExchangeSpread(&_Quoter.TransactOpts, _exchangeSpread)
}

// ChangeExchangeSpread is a paid mutator transaction binding the contract method 0x1118bb3b.
//
// Solidity: function changeExchangeSpread(uint256 _exchangeSpread) returns()
func (_Quoter *QuoterTransactorSession) ChangeExchangeSpread(_exchangeSpread *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeExchangeSpread(&_Quoter.TransactOpts, _exchangeSpread)
}

// ChangeMarketClosedSpread is a paid mutator transaction binding the contract method 0xa75f4eb1.
//
// Solidity: function changeMarketClosedSpread(uint256 _marketClosedSpread) returns()
func (_Quoter *QuoterTransactor) ChangeMarketClosedSpread(opts *bind.TransactOpts, _marketClosedSpread *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeMarketClosedSpread", _marketClosedSpread)
}

// ChangeMarketClosedSpread is a paid mutator transaction binding the contract method 0xa75f4eb1.
//
// Solidity: function changeMarketClosedSpread(uint256 _marketClosedSpread) returns()
func (_Quoter *QuoterSession) ChangeMarketClosedSpread(_marketClosedSpread *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeMarketClosedSpread(&_Quoter.TransactOpts, _marketClosedSpread)
}

// ChangeMarketClosedSpread is a paid mutator transaction binding the contract method 0xa75f4eb1.
//
// Solidity: function changeMarketClosedSpread(uint256 _marketClosedSpread) returns()
func (_Quoter *QuoterTransactorSession) ChangeMarketClosedSpread(_marketClosedSpread *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeMarketClosedSpread(&_Quoter.TransactOpts, _marketClosedSpread)
}

// ChangeMaxUsdcDepositAmount is a paid mutator transaction binding the contract method 0xf80b0de9.
//
// Solidity: function changeMaxUsdcDepositAmount(uint256 _maxUsdcDepositAmount) returns()
func (_Quoter *QuoterTransactor) ChangeMaxUsdcDepositAmount(opts *bind.TransactOpts, _maxUsdcDepositAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeMaxUsdcDepositAmount", _maxUsdcDepositAmount)
}

// ChangeMaxUsdcDepositAmount is a paid mutator transaction binding the contract method 0xf80b0de9.
//
// Solidity: function changeMaxUsdcDepositAmount(uint256 _maxUsdcDepositAmount) returns()
func (_Quoter *QuoterSession) ChangeMaxUsdcDepositAmount(_maxUsdcDepositAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeMaxUsdcDepositAmount(&_Quoter.TransactOpts, _maxUsdcDepositAmount)
}

// ChangeMaxUsdcDepositAmount is a paid mutator transaction binding the contract method 0xf80b0de9.
//
// Solidity: function changeMaxUsdcDepositAmount(uint256 _maxUsdcDepositAmount) returns()
func (_Quoter *QuoterTransactorSession) ChangeMaxUsdcDepositAmount(_maxUsdcDepositAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeMaxUsdcDepositAmount(&_Quoter.TransactOpts, _maxUsdcDepositAmount)
}

// ChangeMinUsdcDepositAmount is a paid mutator transaction binding the contract method 0xeeaa315e.
//
// Solidity: function changeMinUsdcDepositAmount(uint256 _minUsdcDepositAmount) returns()
func (_Quoter *QuoterTransactor) ChangeMinUsdcDepositAmount(opts *bind.TransactOpts, _minUsdcDepositAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeMinUsdcDepositAmount", _minUsdcDepositAmount)
}

// ChangeMinUsdcDepositAmount is a paid mutator transaction binding the contract method 0xeeaa315e.
//
// Solidity: function changeMinUsdcDepositAmount(uint256 _minUsdcDepositAmount) returns()
func (_Quoter *QuoterSession) ChangeMinUsdcDepositAmount(_minUsdcDepositAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeMinUsdcDepositAmount(&_Quoter.TransactOpts, _minUsdcDepositAmount)
}

// ChangeMinUsdcDepositAmount is a paid mutator transaction binding the contract method 0xeeaa315e.
//
// Solidity: function changeMinUsdcDepositAmount(uint256 _minUsdcDepositAmount) returns()
func (_Quoter *QuoterTransactorSession) ChangeMinUsdcDepositAmount(_minUsdcDepositAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeMinUsdcDepositAmount(&_Quoter.TransactOpts, _minUsdcDepositAmount)
}

// ChangeSlippage is a paid mutator transaction binding the contract method 0xbec872b0.
//
// Solidity: function changeSlippage(uint256 _slippage) returns()
func (_Quoter *QuoterTransactor) ChangeSlippage(opts *bind.TransactOpts, _slippage *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeSlippage", _slippage)
}

// ChangeSlippage is a paid mutator transaction binding the contract method 0xbec872b0.
//
// Solidity: function changeSlippage(uint256 _slippage) returns()
func (_Quoter *QuoterSession) ChangeSlippage(_slippage *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeSlippage(&_Quoter.TransactOpts, _slippage)
}

// ChangeSlippage is a paid mutator transaction binding the contract method 0xbec872b0.
//
// Solidity: function changeSlippage(uint256 _slippage) returns()
func (_Quoter *QuoterTransactorSession) ChangeSlippage(_slippage *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeSlippage(&_Quoter.TransactOpts, _slippage)
}

// ChangeStandardFee is a paid mutator transaction binding the contract method 0x81d892dd.
//
// Solidity: function changeStandardFee(uint256 _standardFee) returns()
func (_Quoter *QuoterTransactor) ChangeStandardFee(opts *bind.TransactOpts, _standardFee *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeStandardFee", _standardFee)
}

// ChangeStandardFee is a paid mutator transaction binding the contract method 0x81d892dd.
//
// Solidity: function changeStandardFee(uint256 _standardFee) returns()
func (_Quoter *QuoterSession) ChangeStandardFee(_standardFee *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeStandardFee(&_Quoter.TransactOpts, _standardFee)
}

// ChangeStandardFee is a paid mutator transaction binding the contract method 0x81d892dd.
//
// Solidity: function changeStandardFee(uint256 _standardFee) returns()
func (_Quoter *QuoterTransactorSession) ChangeStandardFee(_standardFee *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeStandardFee(&_Quoter.TransactOpts, _standardFee)
}

// ChangeUsdcAddress is a paid mutator transaction binding the contract method 0xe9b39e2a.
//
// Solidity: function changeUsdcAddress(address _usdcAddress) returns()
func (_Quoter *QuoterTransactor) ChangeUsdcAddress(opts *bind.TransactOpts, _usdcAddress common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeUsdcAddress", _usdcAddress)
}

// ChangeUsdcAddress is a paid mutator transaction binding the contract method 0xe9b39e2a.
//
// Solidity: function changeUsdcAddress(address _usdcAddress) returns()
func (_Quoter *QuoterSession) ChangeUsdcAddress(_usdcAddress common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeUsdcAddress(&_Quoter.TransactOpts, _usdcAddress)
}

// ChangeUsdcAddress is a paid mutator transaction binding the contract method 0xe9b39e2a.
//
// Solidity: function changeUsdcAddress(address _usdcAddress) returns()
func (_Quoter *QuoterTransactorSession) ChangeUsdcAddress(_usdcAddress common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeUsdcAddress(&_Quoter.TransactOpts, _usdcAddress)
}

// ChangeUsdcDepositAddress is a paid mutator transaction binding the contract method 0x2f5008f9.
//
// Solidity: function changeUsdcDepositAddress(address _usdcDepositAddress) returns()
func (_Quoter *QuoterTransactor) ChangeUsdcDepositAddress(opts *bind.TransactOpts, _usdcDepositAddress common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "changeUsdcDepositAddress", _usdcDepositAddress)
}

// ChangeUsdcDepositAddress is a paid mutator transaction binding the contract method 0x2f5008f9.
//
// Solidity: function changeUsdcDepositAddress(address _usdcDepositAddress) returns()
func (_Quoter *QuoterSession) ChangeUsdcDepositAddress(_usdcDepositAddress common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeUsdcDepositAddress(&_Quoter.TransactOpts, _usdcDepositAddress)
}

// ChangeUsdcDepositAddress is a paid mutator transaction binding the contract method 0x2f5008f9.
//
// Solidity: function changeUsdcDepositAddress(address _usdcDepositAddress) returns()
func (_Quoter *QuoterTransactorSession) ChangeUsdcDepositAddress(_usdcDepositAddress common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.ChangeUsdcDepositAddress(&_Quoter.TransactOpts, _usdcDepositAddress)
}

// DepositByWallet is a paid mutator transaction binding the contract method 0x746459f9.
//
// Solidity: function depositByWallet(address _client, address _outputToken, uint256 _usdcAmount) returns(uint256)
func (_Quoter *QuoterTransactor) DepositByWallet(opts *bind.TransactOpts, _client common.Address, _outputToken common.Address, _usdcAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "depositByWallet", _client, _outputToken, _usdcAmount)
}

// DepositByWallet is a paid mutator transaction binding the contract method 0x746459f9.
//
// Solidity: function depositByWallet(address _client, address _outputToken, uint256 _usdcAmount) returns(uint256)
func (_Quoter *QuoterSession) DepositByWallet(_client common.Address, _outputToken common.Address, _usdcAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.DepositByWallet(&_Quoter.TransactOpts, _client, _outputToken, _usdcAmount)
}

// DepositByWallet is a paid mutator transaction binding the contract method 0x746459f9.
//
// Solidity: function depositByWallet(address _client, address _outputToken, uint256 _usdcAmount) returns(uint256)
func (_Quoter *QuoterTransactorSession) DepositByWallet(_client common.Address, _outputToken common.Address, _usdcAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.DepositByWallet(&_Quoter.TransactOpts, _client, _outputToken, _usdcAmount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _outputToken) payable returns(uint256)
func (_Quoter *QuoterTransactor) DepositETH(opts *bind.TransactOpts, _outputToken common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "depositETH", _outputToken)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _outputToken) payable returns(uint256)
func (_Quoter *QuoterSession) DepositETH(_outputToken common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.DepositETH(&_Quoter.TransactOpts, _outputToken)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address _outputToken) payable returns(uint256)
func (_Quoter *QuoterTransactorSession) DepositETH(_outputToken common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.DepositETH(&_Quoter.TransactOpts, _outputToken)
}

// DepositTokenViaEth is a paid mutator transaction binding the contract method 0xa01c3e2b.
//
// Solidity: function depositTokenViaEth(address _inputToken, address _outputToken, uint256 _amount) returns(uint256)
func (_Quoter *QuoterTransactor) DepositTokenViaEth(opts *bind.TransactOpts, _inputToken common.Address, _outputToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "depositTokenViaEth", _inputToken, _outputToken, _amount)
}

// DepositTokenViaEth is a paid mutator transaction binding the contract method 0xa01c3e2b.
//
// Solidity: function depositTokenViaEth(address _inputToken, address _outputToken, uint256 _amount) returns(uint256)
func (_Quoter *QuoterSession) DepositTokenViaEth(_inputToken common.Address, _outputToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.DepositTokenViaEth(&_Quoter.TransactOpts, _inputToken, _outputToken, _amount)
}

// DepositTokenViaEth is a paid mutator transaction binding the contract method 0xa01c3e2b.
//
// Solidity: function depositTokenViaEth(address _inputToken, address _outputToken, uint256 _amount) returns(uint256)
func (_Quoter *QuoterTransactorSession) DepositTokenViaEth(_inputToken common.Address, _outputToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.DepositTokenViaEth(&_Quoter.TransactOpts, _inputToken, _outputToken, _amount)
}

// DepositTokenViaUsdc is a paid mutator transaction binding the contract method 0xb461e299.
//
// Solidity: function depositTokenViaUsdc(address _inputToken, address _outputToken, uint256 _amount) returns(uint256)
func (_Quoter *QuoterTransactor) DepositTokenViaUsdc(opts *bind.TransactOpts, _inputToken common.Address, _outputToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "depositTokenViaUsdc", _inputToken, _outputToken, _amount)
}

// DepositTokenViaUsdc is a paid mutator transaction binding the contract method 0xb461e299.
//
// Solidity: function depositTokenViaUsdc(address _inputToken, address _outputToken, uint256 _amount) returns(uint256)
func (_Quoter *QuoterSession) DepositTokenViaUsdc(_inputToken common.Address, _outputToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.DepositTokenViaUsdc(&_Quoter.TransactOpts, _inputToken, _outputToken, _amount)
}

// DepositTokenViaUsdc is a paid mutator transaction binding the contract method 0xb461e299.
//
// Solidity: function depositTokenViaUsdc(address _inputToken, address _outputToken, uint256 _amount) returns(uint256)
func (_Quoter *QuoterTransactorSession) DepositTokenViaUsdc(_inputToken common.Address, _outputToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.DepositTokenViaUsdc(&_Quoter.TransactOpts, _inputToken, _outputToken, _amount)
}

// GetQuote is a paid mutator transaction binding the contract method 0x6b77e75a.
//
// Solidity: function getQuote(address _inputToken, address _outputToken, uint24 _fee, uint256 _amount) returns(uint256)
func (_Quoter *QuoterTransactor) GetQuote(opts *bind.TransactOpts, _inputToken common.Address, _outputToken common.Address, _fee *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "getQuote", _inputToken, _outputToken, _fee, _amount)
}

// GetQuote is a paid mutator transaction binding the contract method 0x6b77e75a.
//
// Solidity: function getQuote(address _inputToken, address _outputToken, uint24 _fee, uint256 _amount) returns(uint256)
func (_Quoter *QuoterSession) GetQuote(_inputToken common.Address, _outputToken common.Address, _fee *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.GetQuote(&_Quoter.TransactOpts, _inputToken, _outputToken, _fee, _amount)
}

// GetQuote is a paid mutator transaction binding the contract method 0x6b77e75a.
//
// Solidity: function getQuote(address _inputToken, address _outputToken, uint24 _fee, uint256 _amount) returns(uint256)
func (_Quoter *QuoterTransactorSession) GetQuote(_inputToken common.Address, _outputToken common.Address, _fee *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.GetQuote(&_Quoter.TransactOpts, _inputToken, _outputToken, _fee, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Quoter *QuoterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Quoter *QuoterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.GrantRole(&_Quoter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Quoter *QuoterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.GrantRole(&_Quoter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x76902b83.
//
// Solidity: function initialize(address _fiat24account, address _usd24, address _eur24, address _chf24, address _gbp24, address _usdc, address _weth, address _f24, address _f24timelock, address _f24DeskAddress, address _usdcDepositAddress) returns()
func (_Quoter *QuoterTransactor) Initialize(opts *bind.TransactOpts, _fiat24account common.Address, _usd24 common.Address, _eur24 common.Address, _chf24 common.Address, _gbp24 common.Address, _usdc common.Address, _weth common.Address, _f24 common.Address, _f24timelock common.Address, _f24DeskAddress common.Address, _usdcDepositAddress common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "initialize", _fiat24account, _usd24, _eur24, _chf24, _gbp24, _usdc, _weth, _f24, _f24timelock, _f24DeskAddress, _usdcDepositAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x76902b83.
//
// Solidity: function initialize(address _fiat24account, address _usd24, address _eur24, address _chf24, address _gbp24, address _usdc, address _weth, address _f24, address _f24timelock, address _f24DeskAddress, address _usdcDepositAddress) returns()
func (_Quoter *QuoterSession) Initialize(_fiat24account common.Address, _usd24 common.Address, _eur24 common.Address, _chf24 common.Address, _gbp24 common.Address, _usdc common.Address, _weth common.Address, _f24 common.Address, _f24timelock common.Address, _f24DeskAddress common.Address, _usdcDepositAddress common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.Initialize(&_Quoter.TransactOpts, _fiat24account, _usd24, _eur24, _chf24, _gbp24, _usdc, _weth, _f24, _f24timelock, _f24DeskAddress, _usdcDepositAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x76902b83.
//
// Solidity: function initialize(address _fiat24account, address _usd24, address _eur24, address _chf24, address _gbp24, address _usdc, address _weth, address _f24, address _f24timelock, address _f24DeskAddress, address _usdcDepositAddress) returns()
func (_Quoter *QuoterTransactorSession) Initialize(_fiat24account common.Address, _usd24 common.Address, _eur24 common.Address, _chf24 common.Address, _gbp24 common.Address, _usdc common.Address, _weth common.Address, _f24 common.Address, _f24timelock common.Address, _f24DeskAddress common.Address, _usdcDepositAddress common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.Initialize(&_Quoter.TransactOpts, _fiat24account, _usd24, _eur24, _chf24, _gbp24, _usdc, _weth, _f24, _f24timelock, _f24DeskAddress, _usdcDepositAddress)
}

// MoneyExchangeExactIn is a paid mutator transaction binding the contract method 0x72a023fa.
//
// Solidity: function moneyExchangeExactIn(address _inputToken, address _outputToken, uint256 _inputAmount) returns(uint256)
func (_Quoter *QuoterTransactor) MoneyExchangeExactIn(opts *bind.TransactOpts, _inputToken common.Address, _outputToken common.Address, _inputAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "moneyExchangeExactIn", _inputToken, _outputToken, _inputAmount)
}

// MoneyExchangeExactIn is a paid mutator transaction binding the contract method 0x72a023fa.
//
// Solidity: function moneyExchangeExactIn(address _inputToken, address _outputToken, uint256 _inputAmount) returns(uint256)
func (_Quoter *QuoterSession) MoneyExchangeExactIn(_inputToken common.Address, _outputToken common.Address, _inputAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.MoneyExchangeExactIn(&_Quoter.TransactOpts, _inputToken, _outputToken, _inputAmount)
}

// MoneyExchangeExactIn is a paid mutator transaction binding the contract method 0x72a023fa.
//
// Solidity: function moneyExchangeExactIn(address _inputToken, address _outputToken, uint256 _inputAmount) returns(uint256)
func (_Quoter *QuoterTransactorSession) MoneyExchangeExactIn(_inputToken common.Address, _outputToken common.Address, _inputAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.MoneyExchangeExactIn(&_Quoter.TransactOpts, _inputToken, _outputToken, _inputAmount)
}

// MoneyExchangeExactOut is a paid mutator transaction binding the contract method 0x8f28f584.
//
// Solidity: function moneyExchangeExactOut(address _inputToken, address _outputToken, uint256 _outputAmount) returns(uint256)
func (_Quoter *QuoterTransactor) MoneyExchangeExactOut(opts *bind.TransactOpts, _inputToken common.Address, _outputToken common.Address, _outputAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "moneyExchangeExactOut", _inputToken, _outputToken, _outputAmount)
}

// MoneyExchangeExactOut is a paid mutator transaction binding the contract method 0x8f28f584.
//
// Solidity: function moneyExchangeExactOut(address _inputToken, address _outputToken, uint256 _outputAmount) returns(uint256)
func (_Quoter *QuoterSession) MoneyExchangeExactOut(_inputToken common.Address, _outputToken common.Address, _outputAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.MoneyExchangeExactOut(&_Quoter.TransactOpts, _inputToken, _outputToken, _outputAmount)
}

// MoneyExchangeExactOut is a paid mutator transaction binding the contract method 0x8f28f584.
//
// Solidity: function moneyExchangeExactOut(address _inputToken, address _outputToken, uint256 _outputAmount) returns(uint256)
func (_Quoter *QuoterTransactorSession) MoneyExchangeExactOut(_inputToken common.Address, _outputToken common.Address, _outputAmount *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.MoneyExchangeExactOut(&_Quoter.TransactOpts, _inputToken, _outputToken, _outputAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Quoter *QuoterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Quoter *QuoterSession) Pause() (*types.Transaction, error) {
	return _Quoter.Contract.Pause(&_Quoter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Quoter *QuoterTransactorSession) Pause() (*types.Transaction, error) {
	return _Quoter.Contract.Pause(&_Quoter.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Quoter *QuoterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Quoter *QuoterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.RenounceRole(&_Quoter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Quoter *QuoterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.RenounceRole(&_Quoter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Quoter *QuoterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Quoter *QuoterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.RevokeRole(&_Quoter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Quoter *QuoterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.RevokeRole(&_Quoter.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Quoter *QuoterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Quoter *QuoterSession) Unpause() (*types.Transaction, error) {
	return _Quoter.Contract.Unpause(&_Quoter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Quoter *QuoterTransactorSession) Unpause() (*types.Transaction, error) {
	return _Quoter.Contract.Unpause(&_Quoter.TransactOpts)
}

// UpdateExchangeRates is a paid mutator transaction binding the contract method 0x488da0dc.
//
// Solidity: function updateExchangeRates(uint256 _usd_eur, uint256 _usd_chf, uint256 _usd_gbp, uint256 _usd_cnh, bool _isMarketClosed) returns()
func (_Quoter *QuoterTransactor) UpdateExchangeRates(opts *bind.TransactOpts, _usd_eur *big.Int, _usd_chf *big.Int, _usd_gbp *big.Int, _usd_cnh *big.Int, _isMarketClosed bool) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "updateExchangeRates", _usd_eur, _usd_chf, _usd_gbp, _usd_cnh, _isMarketClosed)
}

// UpdateExchangeRates is a paid mutator transaction binding the contract method 0x488da0dc.
//
// Solidity: function updateExchangeRates(uint256 _usd_eur, uint256 _usd_chf, uint256 _usd_gbp, uint256 _usd_cnh, bool _isMarketClosed) returns()
func (_Quoter *QuoterSession) UpdateExchangeRates(_usd_eur *big.Int, _usd_chf *big.Int, _usd_gbp *big.Int, _usd_cnh *big.Int, _isMarketClosed bool) (*types.Transaction, error) {
	return _Quoter.Contract.UpdateExchangeRates(&_Quoter.TransactOpts, _usd_eur, _usd_chf, _usd_gbp, _usd_cnh, _isMarketClosed)
}

// UpdateExchangeRates is a paid mutator transaction binding the contract method 0x488da0dc.
//
// Solidity: function updateExchangeRates(uint256 _usd_eur, uint256 _usd_chf, uint256 _usd_gbp, uint256 _usd_cnh, bool _isMarketClosed) returns()
func (_Quoter *QuoterTransactorSession) UpdateExchangeRates(_usd_eur *big.Int, _usd_chf *big.Int, _usd_gbp *big.Int, _usd_cnh *big.Int, _isMarketClosed bool) (*types.Transaction, error) {
	return _Quoter.Contract.UpdateExchangeRates(&_Quoter.TransactOpts, _usd_eur, _usd_chf, _usd_gbp, _usd_cnh, _isMarketClosed)
}

// UpdateUsdcUsd24ExchangeRate is a paid mutator transaction binding the contract method 0x310a0325.
//
// Solidity: function updateUsdcUsd24ExchangeRate(uint256 _usdc_usd24) returns()
func (_Quoter *QuoterTransactor) UpdateUsdcUsd24ExchangeRate(opts *bind.TransactOpts, _usdc_usd24 *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "updateUsdcUsd24ExchangeRate", _usdc_usd24)
}

// UpdateUsdcUsd24ExchangeRate is a paid mutator transaction binding the contract method 0x310a0325.
//
// Solidity: function updateUsdcUsd24ExchangeRate(uint256 _usdc_usd24) returns()
func (_Quoter *QuoterSession) UpdateUsdcUsd24ExchangeRate(_usdc_usd24 *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.UpdateUsdcUsd24ExchangeRate(&_Quoter.TransactOpts, _usdc_usd24)
}

// UpdateUsdcUsd24ExchangeRate is a paid mutator transaction binding the contract method 0x310a0325.
//
// Solidity: function updateUsdcUsd24ExchangeRate(uint256 _usdc_usd24) returns()
func (_Quoter *QuoterTransactorSession) UpdateUsdcUsd24ExchangeRate(_usdc_usd24 *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.UpdateUsdcUsd24ExchangeRate(&_Quoter.TransactOpts, _usdc_usd24)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Quoter *QuoterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Quoter *QuoterSession) Receive() (*types.Transaction, error) {
	return _Quoter.Contract.Receive(&_Quoter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Quoter *QuoterTransactorSession) Receive() (*types.Transaction, error) {
	return _Quoter.Contract.Receive(&_Quoter.TransactOpts)
}

// QuoterDepositedByWalletIterator is returned from FilterDepositedByWallet and is used to iterate over the raw logs and unpacked data for DepositedByWallet events raised by the Quoter contract.
type QuoterDepositedByWalletIterator struct {
	Event *QuoterDepositedByWallet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterDepositedByWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterDepositedByWallet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterDepositedByWallet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterDepositedByWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterDepositedByWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterDepositedByWallet represents a DepositedByWallet event raised by the Quoter contract.
type QuoterDepositedByWallet struct {
	TokenId       *big.Int
	ClientAddress common.Address
	WalletId      *big.Int
	WalletAddress common.Address
	OutputToken   common.Address
	UsdcAmount    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositedByWallet is a free log retrieval operation binding the contract event 0x3f65d89a3e87576d9d84da7409131cdbf5f3b8247415b3313e707cc91d5bb4a2.
//
// Solidity: event DepositedByWallet(uint256 indexed tokenId, address indexed clientAddress, uint256 indexed walletId, address walletAddress, address outputToken, uint256 usdcAmount)
func (_Quoter *QuoterFilterer) FilterDepositedByWallet(opts *bind.FilterOpts, tokenId []*big.Int, clientAddress []common.Address, walletId []*big.Int) (*QuoterDepositedByWalletIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "DepositedByWallet", tokenIdRule, clientAddressRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &QuoterDepositedByWalletIterator{contract: _Quoter.contract, event: "DepositedByWallet", logs: logs, sub: sub}, nil
}

// WatchDepositedByWallet is a free log subscription operation binding the contract event 0x3f65d89a3e87576d9d84da7409131cdbf5f3b8247415b3313e707cc91d5bb4a2.
//
// Solidity: event DepositedByWallet(uint256 indexed tokenId, address indexed clientAddress, uint256 indexed walletId, address walletAddress, address outputToken, uint256 usdcAmount)
func (_Quoter *QuoterFilterer) WatchDepositedByWallet(opts *bind.WatchOpts, sink chan<- *QuoterDepositedByWallet, tokenId []*big.Int, clientAddress []common.Address, walletId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "DepositedByWallet", tokenIdRule, clientAddressRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterDepositedByWallet)
				if err := _Quoter.contract.UnpackLog(event, "DepositedByWallet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedByWallet is a log parse operation binding the contract event 0x3f65d89a3e87576d9d84da7409131cdbf5f3b8247415b3313e707cc91d5bb4a2.
//
// Solidity: event DepositedByWallet(uint256 indexed tokenId, address indexed clientAddress, uint256 indexed walletId, address walletAddress, address outputToken, uint256 usdcAmount)
func (_Quoter *QuoterFilterer) ParseDepositedByWallet(log types.Log) (*QuoterDepositedByWallet, error) {
	event := new(QuoterDepositedByWallet)
	if err := _Quoter.contract.UnpackLog(event, "DepositedByWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterDepositedEthIterator is returned from FilterDepositedEth and is used to iterate over the raw logs and unpacked data for DepositedEth events raised by the Quoter contract.
type QuoterDepositedEthIterator struct {
	Event *QuoterDepositedEth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterDepositedEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterDepositedEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterDepositedEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterDepositedEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterDepositedEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterDepositedEth represents a DepositedEth event raised by the Quoter contract.
type QuoterDepositedEth struct {
	Sender       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositedEth is a free log retrieval operation binding the contract event 0x34d7d1806267c56e6c814b66febca4a2a46f1dd9d8af4146c62b585f6b84c2c9.
//
// Solidity: event DepositedEth(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) FilterDepositedEth(opts *bind.FilterOpts, sender []common.Address) (*QuoterDepositedEthIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "DepositedEth", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterDepositedEthIterator{contract: _Quoter.contract, event: "DepositedEth", logs: logs, sub: sub}, nil
}

// WatchDepositedEth is a free log subscription operation binding the contract event 0x34d7d1806267c56e6c814b66febca4a2a46f1dd9d8af4146c62b585f6b84c2c9.
//
// Solidity: event DepositedEth(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) WatchDepositedEth(opts *bind.WatchOpts, sink chan<- *QuoterDepositedEth, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "DepositedEth", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterDepositedEth)
				if err := _Quoter.contract.UnpackLog(event, "DepositedEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedEth is a log parse operation binding the contract event 0x34d7d1806267c56e6c814b66febca4a2a46f1dd9d8af4146c62b585f6b84c2c9.
//
// Solidity: event DepositedEth(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) ParseDepositedEth(log types.Log) (*QuoterDepositedEth, error) {
	event := new(QuoterDepositedEth)
	if err := _Quoter.contract.UnpackLog(event, "DepositedEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterDepositedTokenViaEthIterator is returned from FilterDepositedTokenViaEth and is used to iterate over the raw logs and unpacked data for DepositedTokenViaEth events raised by the Quoter contract.
type QuoterDepositedTokenViaEthIterator struct {
	Event *QuoterDepositedTokenViaEth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterDepositedTokenViaEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterDepositedTokenViaEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterDepositedTokenViaEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterDepositedTokenViaEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterDepositedTokenViaEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterDepositedTokenViaEth represents a DepositedTokenViaEth event raised by the Quoter contract.
type QuoterDepositedTokenViaEth struct {
	Sender       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositedTokenViaEth is a free log retrieval operation binding the contract event 0xd1469a782115ba96fa7fabe422d9723867f349cfbaf9b999e327e3db4f8baa3e.
//
// Solidity: event DepositedTokenViaEth(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) FilterDepositedTokenViaEth(opts *bind.FilterOpts, sender []common.Address) (*QuoterDepositedTokenViaEthIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "DepositedTokenViaEth", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterDepositedTokenViaEthIterator{contract: _Quoter.contract, event: "DepositedTokenViaEth", logs: logs, sub: sub}, nil
}

// WatchDepositedTokenViaEth is a free log subscription operation binding the contract event 0xd1469a782115ba96fa7fabe422d9723867f349cfbaf9b999e327e3db4f8baa3e.
//
// Solidity: event DepositedTokenViaEth(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) WatchDepositedTokenViaEth(opts *bind.WatchOpts, sink chan<- *QuoterDepositedTokenViaEth, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "DepositedTokenViaEth", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterDepositedTokenViaEth)
				if err := _Quoter.contract.UnpackLog(event, "DepositedTokenViaEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedTokenViaEth is a log parse operation binding the contract event 0xd1469a782115ba96fa7fabe422d9723867f349cfbaf9b999e327e3db4f8baa3e.
//
// Solidity: event DepositedTokenViaEth(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) ParseDepositedTokenViaEth(log types.Log) (*QuoterDepositedTokenViaEth, error) {
	event := new(QuoterDepositedTokenViaEth)
	if err := _Quoter.contract.UnpackLog(event, "DepositedTokenViaEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterDepositedTokenViaUsdIterator is returned from FilterDepositedTokenViaUsd and is used to iterate over the raw logs and unpacked data for DepositedTokenViaUsd events raised by the Quoter contract.
type QuoterDepositedTokenViaUsdIterator struct {
	Event *QuoterDepositedTokenViaUsd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterDepositedTokenViaUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterDepositedTokenViaUsd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterDepositedTokenViaUsd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterDepositedTokenViaUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterDepositedTokenViaUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterDepositedTokenViaUsd represents a DepositedTokenViaUsd event raised by the Quoter contract.
type QuoterDepositedTokenViaUsd struct {
	Sender       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositedTokenViaUsd is a free log retrieval operation binding the contract event 0xb9e470208de2a9058ebfcea09f875f4b4fab9492dd164962f84d744002e6465d.
//
// Solidity: event DepositedTokenViaUsd(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) FilterDepositedTokenViaUsd(opts *bind.FilterOpts, sender []common.Address) (*QuoterDepositedTokenViaUsdIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "DepositedTokenViaUsd", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterDepositedTokenViaUsdIterator{contract: _Quoter.contract, event: "DepositedTokenViaUsd", logs: logs, sub: sub}, nil
}

// WatchDepositedTokenViaUsd is a free log subscription operation binding the contract event 0xb9e470208de2a9058ebfcea09f875f4b4fab9492dd164962f84d744002e6465d.
//
// Solidity: event DepositedTokenViaUsd(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) WatchDepositedTokenViaUsd(opts *bind.WatchOpts, sink chan<- *QuoterDepositedTokenViaUsd, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "DepositedTokenViaUsd", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterDepositedTokenViaUsd)
				if err := _Quoter.contract.UnpackLog(event, "DepositedTokenViaUsd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedTokenViaUsd is a log parse operation binding the contract event 0xb9e470208de2a9058ebfcea09f875f4b4fab9492dd164962f84d744002e6465d.
//
// Solidity: event DepositedTokenViaUsd(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) ParseDepositedTokenViaUsd(log types.Log) (*QuoterDepositedTokenViaUsd, error) {
	event := new(QuoterDepositedTokenViaUsd)
	if err := _Quoter.contract.UnpackLog(event, "DepositedTokenViaUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterExchangeRatesUpdatedByOperatorIterator is returned from FilterExchangeRatesUpdatedByOperator and is used to iterate over the raw logs and unpacked data for ExchangeRatesUpdatedByOperator events raised by the Quoter contract.
type QuoterExchangeRatesUpdatedByOperatorIterator struct {
	Event *QuoterExchangeRatesUpdatedByOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterExchangeRatesUpdatedByOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterExchangeRatesUpdatedByOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterExchangeRatesUpdatedByOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterExchangeRatesUpdatedByOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterExchangeRatesUpdatedByOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterExchangeRatesUpdatedByOperator represents a ExchangeRatesUpdatedByOperator event raised by the Quoter contract.
type QuoterExchangeRatesUpdatedByOperator struct {
	Sender       common.Address
	Usdeur       *big.Int
	Usdchf       *big.Int
	Usdgbp       *big.Int
	Usdcnh       *big.Int
	MarketClosed bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExchangeRatesUpdatedByOperator is a free log retrieval operation binding the contract event 0x9f360fb4784547b0e5ca1a9940b6c40efc68cd7558feb8c2d2ad9b9669ed83be.
//
// Solidity: event ExchangeRatesUpdatedByOperator(address indexed sender, uint256 usdeur, uint256 usdchf, uint256 usdgbp, uint256 usdcnh, bool marketClosed)
func (_Quoter *QuoterFilterer) FilterExchangeRatesUpdatedByOperator(opts *bind.FilterOpts, sender []common.Address) (*QuoterExchangeRatesUpdatedByOperatorIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "ExchangeRatesUpdatedByOperator", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterExchangeRatesUpdatedByOperatorIterator{contract: _Quoter.contract, event: "ExchangeRatesUpdatedByOperator", logs: logs, sub: sub}, nil
}

// WatchExchangeRatesUpdatedByOperator is a free log subscription operation binding the contract event 0x9f360fb4784547b0e5ca1a9940b6c40efc68cd7558feb8c2d2ad9b9669ed83be.
//
// Solidity: event ExchangeRatesUpdatedByOperator(address indexed sender, uint256 usdeur, uint256 usdchf, uint256 usdgbp, uint256 usdcnh, bool marketClosed)
func (_Quoter *QuoterFilterer) WatchExchangeRatesUpdatedByOperator(opts *bind.WatchOpts, sink chan<- *QuoterExchangeRatesUpdatedByOperator, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "ExchangeRatesUpdatedByOperator", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterExchangeRatesUpdatedByOperator)
				if err := _Quoter.contract.UnpackLog(event, "ExchangeRatesUpdatedByOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExchangeRatesUpdatedByOperator is a log parse operation binding the contract event 0x9f360fb4784547b0e5ca1a9940b6c40efc68cd7558feb8c2d2ad9b9669ed83be.
//
// Solidity: event ExchangeRatesUpdatedByOperator(address indexed sender, uint256 usdeur, uint256 usdchf, uint256 usdgbp, uint256 usdcnh, bool marketClosed)
func (_Quoter *QuoterFilterer) ParseExchangeRatesUpdatedByOperator(log types.Log) (*QuoterExchangeRatesUpdatedByOperator, error) {
	event := new(QuoterExchangeRatesUpdatedByOperator)
	if err := _Quoter.contract.UnpackLog(event, "ExchangeRatesUpdatedByOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterExchangeRatesUpdatedByRobotIterator is returned from FilterExchangeRatesUpdatedByRobot and is used to iterate over the raw logs and unpacked data for ExchangeRatesUpdatedByRobot events raised by the Quoter contract.
type QuoterExchangeRatesUpdatedByRobotIterator struct {
	Event *QuoterExchangeRatesUpdatedByRobot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterExchangeRatesUpdatedByRobotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterExchangeRatesUpdatedByRobot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterExchangeRatesUpdatedByRobot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterExchangeRatesUpdatedByRobotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterExchangeRatesUpdatedByRobotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterExchangeRatesUpdatedByRobot represents a ExchangeRatesUpdatedByRobot event raised by the Quoter contract.
type QuoterExchangeRatesUpdatedByRobot struct {
	Sender       common.Address
	Usdeur       *big.Int
	Usdchf       *big.Int
	Usdgbp       *big.Int
	Usdcnh       *big.Int
	MarketClosed bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExchangeRatesUpdatedByRobot is a free log retrieval operation binding the contract event 0x8d901df24d3fcb3c29b3949e685435c71fbb313dffa57e49da96fc7abd47b537.
//
// Solidity: event ExchangeRatesUpdatedByRobot(address indexed sender, uint256 usdeur, uint256 usdchf, uint256 usdgbp, uint256 usdcnh, bool marketClosed)
func (_Quoter *QuoterFilterer) FilterExchangeRatesUpdatedByRobot(opts *bind.FilterOpts, sender []common.Address) (*QuoterExchangeRatesUpdatedByRobotIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "ExchangeRatesUpdatedByRobot", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterExchangeRatesUpdatedByRobotIterator{contract: _Quoter.contract, event: "ExchangeRatesUpdatedByRobot", logs: logs, sub: sub}, nil
}

// WatchExchangeRatesUpdatedByRobot is a free log subscription operation binding the contract event 0x8d901df24d3fcb3c29b3949e685435c71fbb313dffa57e49da96fc7abd47b537.
//
// Solidity: event ExchangeRatesUpdatedByRobot(address indexed sender, uint256 usdeur, uint256 usdchf, uint256 usdgbp, uint256 usdcnh, bool marketClosed)
func (_Quoter *QuoterFilterer) WatchExchangeRatesUpdatedByRobot(opts *bind.WatchOpts, sink chan<- *QuoterExchangeRatesUpdatedByRobot, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "ExchangeRatesUpdatedByRobot", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterExchangeRatesUpdatedByRobot)
				if err := _Quoter.contract.UnpackLog(event, "ExchangeRatesUpdatedByRobot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExchangeRatesUpdatedByRobot is a log parse operation binding the contract event 0x8d901df24d3fcb3c29b3949e685435c71fbb313dffa57e49da96fc7abd47b537.
//
// Solidity: event ExchangeRatesUpdatedByRobot(address indexed sender, uint256 usdeur, uint256 usdchf, uint256 usdgbp, uint256 usdcnh, bool marketClosed)
func (_Quoter *QuoterFilterer) ParseExchangeRatesUpdatedByRobot(log types.Log) (*QuoterExchangeRatesUpdatedByRobot, error) {
	event := new(QuoterExchangeRatesUpdatedByRobot)
	if err := _Quoter.contract.UnpackLog(event, "ExchangeRatesUpdatedByRobot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterMoneyExchangedExactInIterator is returned from FilterMoneyExchangedExactIn and is used to iterate over the raw logs and unpacked data for MoneyExchangedExactIn events raised by the Quoter contract.
type QuoterMoneyExchangedExactInIterator struct {
	Event *QuoterMoneyExchangedExactIn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterMoneyExchangedExactInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterMoneyExchangedExactIn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterMoneyExchangedExactIn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterMoneyExchangedExactInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterMoneyExchangedExactInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterMoneyExchangedExactIn represents a MoneyExchangedExactIn event raised by the Quoter contract.
type QuoterMoneyExchangedExactIn struct {
	Sender       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMoneyExchangedExactIn is a free log retrieval operation binding the contract event 0xa6ad14cbf43fd219354877f33abdfa78542d2d066c833c03a195d83b1ccbba39.
//
// Solidity: event MoneyExchangedExactIn(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) FilterMoneyExchangedExactIn(opts *bind.FilterOpts, sender []common.Address) (*QuoterMoneyExchangedExactInIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "MoneyExchangedExactIn", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterMoneyExchangedExactInIterator{contract: _Quoter.contract, event: "MoneyExchangedExactIn", logs: logs, sub: sub}, nil
}

// WatchMoneyExchangedExactIn is a free log subscription operation binding the contract event 0xa6ad14cbf43fd219354877f33abdfa78542d2d066c833c03a195d83b1ccbba39.
//
// Solidity: event MoneyExchangedExactIn(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) WatchMoneyExchangedExactIn(opts *bind.WatchOpts, sink chan<- *QuoterMoneyExchangedExactIn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "MoneyExchangedExactIn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterMoneyExchangedExactIn)
				if err := _Quoter.contract.UnpackLog(event, "MoneyExchangedExactIn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoneyExchangedExactIn is a log parse operation binding the contract event 0xa6ad14cbf43fd219354877f33abdfa78542d2d066c833c03a195d83b1ccbba39.
//
// Solidity: event MoneyExchangedExactIn(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) ParseMoneyExchangedExactIn(log types.Log) (*QuoterMoneyExchangedExactIn, error) {
	event := new(QuoterMoneyExchangedExactIn)
	if err := _Quoter.contract.UnpackLog(event, "MoneyExchangedExactIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterMoneyExchangedExactOutIterator is returned from FilterMoneyExchangedExactOut and is used to iterate over the raw logs and unpacked data for MoneyExchangedExactOut events raised by the Quoter contract.
type QuoterMoneyExchangedExactOutIterator struct {
	Event *QuoterMoneyExchangedExactOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterMoneyExchangedExactOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterMoneyExchangedExactOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterMoneyExchangedExactOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterMoneyExchangedExactOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterMoneyExchangedExactOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterMoneyExchangedExactOut represents a MoneyExchangedExactOut event raised by the Quoter contract.
type QuoterMoneyExchangedExactOut struct {
	Sender       common.Address
	InputToken   common.Address
	OutputToken  common.Address
	InputAmount  *big.Int
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMoneyExchangedExactOut is a free log retrieval operation binding the contract event 0x0817d793c7fe339a7cd6f6cf8e2b6deadf3e9a29c74b5e0b87966710b03827f1.
//
// Solidity: event MoneyExchangedExactOut(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) FilterMoneyExchangedExactOut(opts *bind.FilterOpts, sender []common.Address) (*QuoterMoneyExchangedExactOutIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "MoneyExchangedExactOut", senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterMoneyExchangedExactOutIterator{contract: _Quoter.contract, event: "MoneyExchangedExactOut", logs: logs, sub: sub}, nil
}

// WatchMoneyExchangedExactOut is a free log subscription operation binding the contract event 0x0817d793c7fe339a7cd6f6cf8e2b6deadf3e9a29c74b5e0b87966710b03827f1.
//
// Solidity: event MoneyExchangedExactOut(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) WatchMoneyExchangedExactOut(opts *bind.WatchOpts, sink chan<- *QuoterMoneyExchangedExactOut, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "MoneyExchangedExactOut", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterMoneyExchangedExactOut)
				if err := _Quoter.contract.UnpackLog(event, "MoneyExchangedExactOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoneyExchangedExactOut is a log parse operation binding the contract event 0x0817d793c7fe339a7cd6f6cf8e2b6deadf3e9a29c74b5e0b87966710b03827f1.
//
// Solidity: event MoneyExchangedExactOut(address indexed sender, address inputToken, address outputToken, uint256 inputAmount, uint256 outputAmount)
func (_Quoter *QuoterFilterer) ParseMoneyExchangedExactOut(log types.Log) (*QuoterMoneyExchangedExactOut, error) {
	event := new(QuoterMoneyExchangedExactOut)
	if err := _Quoter.contract.UnpackLog(event, "MoneyExchangedExactOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Quoter contract.
type QuoterPausedIterator struct {
	Event *QuoterPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterPaused represents a Paused event raised by the Quoter contract.
type QuoterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Quoter *QuoterFilterer) FilterPaused(opts *bind.FilterOpts) (*QuoterPausedIterator, error) {

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &QuoterPausedIterator{contract: _Quoter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Quoter *QuoterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *QuoterPaused) (event.Subscription, error) {

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterPaused)
				if err := _Quoter.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Quoter *QuoterFilterer) ParsePaused(log types.Log) (*QuoterPaused, error) {
	event := new(QuoterPaused)
	if err := _Quoter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Quoter contract.
type QuoterRoleAdminChangedIterator struct {
	Event *QuoterRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterRoleAdminChanged represents a RoleAdminChanged event raised by the Quoter contract.
type QuoterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Quoter *QuoterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*QuoterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &QuoterRoleAdminChangedIterator{contract: _Quoter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Quoter *QuoterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *QuoterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterRoleAdminChanged)
				if err := _Quoter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Quoter *QuoterFilterer) ParseRoleAdminChanged(log types.Log) (*QuoterRoleAdminChanged, error) {
	event := new(QuoterRoleAdminChanged)
	if err := _Quoter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Quoter contract.
type QuoterRoleGrantedIterator struct {
	Event *QuoterRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterRoleGranted represents a RoleGranted event raised by the Quoter contract.
type QuoterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Quoter *QuoterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*QuoterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterRoleGrantedIterator{contract: _Quoter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Quoter *QuoterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *QuoterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterRoleGranted)
				if err := _Quoter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Quoter *QuoterFilterer) ParseRoleGranted(log types.Log) (*QuoterRoleGranted, error) {
	event := new(QuoterRoleGranted)
	if err := _Quoter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Quoter contract.
type QuoterRoleRevokedIterator struct {
	Event *QuoterRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterRoleRevoked represents a RoleRevoked event raised by the Quoter contract.
type QuoterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Quoter *QuoterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*QuoterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &QuoterRoleRevokedIterator{contract: _Quoter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Quoter *QuoterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *QuoterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterRoleRevoked)
				if err := _Quoter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Quoter *QuoterFilterer) ParseRoleRevoked(log types.Log) (*QuoterRoleRevoked, error) {
	event := new(QuoterRoleRevoked)
	if err := _Quoter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Quoter contract.
type QuoterUnpausedIterator struct {
	Event *QuoterUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterUnpaused represents a Unpaused event raised by the Quoter contract.
type QuoterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Quoter *QuoterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*QuoterUnpausedIterator, error) {

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &QuoterUnpausedIterator{contract: _Quoter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Quoter *QuoterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *QuoterUnpaused) (event.Subscription, error) {

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterUnpaused)
				if err := _Quoter.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Quoter *QuoterFilterer) ParseUnpaused(log types.Log) (*QuoterUnpaused, error) {
	event := new(QuoterUnpaused)
	if err := _Quoter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterUsdcDepositAddressChangedIterator is returned from FilterUsdcDepositAddressChanged and is used to iterate over the raw logs and unpacked data for UsdcDepositAddressChanged events raised by the Quoter contract.
type QuoterUsdcDepositAddressChangedIterator struct {
	Event *QuoterUsdcDepositAddressChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuoterUsdcDepositAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterUsdcDepositAddressChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuoterUsdcDepositAddressChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuoterUsdcDepositAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterUsdcDepositAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterUsdcDepositAddressChanged represents a UsdcDepositAddressChanged event raised by the Quoter contract.
type QuoterUsdcDepositAddressChanged struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUsdcDepositAddressChanged is a free log retrieval operation binding the contract event 0x333d753eb88dfde031e8da4385e0440b26920614852643bc6c80e79718dac36d.
//
// Solidity: event UsdcDepositAddressChanged(address oldAddress, address newAddress)
func (_Quoter *QuoterFilterer) FilterUsdcDepositAddressChanged(opts *bind.FilterOpts) (*QuoterUsdcDepositAddressChangedIterator, error) {

	logs, sub, err := _Quoter.contract.FilterLogs(opts, "UsdcDepositAddressChanged")
	if err != nil {
		return nil, err
	}
	return &QuoterUsdcDepositAddressChangedIterator{contract: _Quoter.contract, event: "UsdcDepositAddressChanged", logs: logs, sub: sub}, nil
}

// WatchUsdcDepositAddressChanged is a free log subscription operation binding the contract event 0x333d753eb88dfde031e8da4385e0440b26920614852643bc6c80e79718dac36d.
//
// Solidity: event UsdcDepositAddressChanged(address oldAddress, address newAddress)
func (_Quoter *QuoterFilterer) WatchUsdcDepositAddressChanged(opts *bind.WatchOpts, sink chan<- *QuoterUsdcDepositAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Quoter.contract.WatchLogs(opts, "UsdcDepositAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterUsdcDepositAddressChanged)
				if err := _Quoter.contract.UnpackLog(event, "UsdcDepositAddressChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUsdcDepositAddressChanged is a log parse operation binding the contract event 0x333d753eb88dfde031e8da4385e0440b26920614852643bc6c80e79718dac36d.
//
// Solidity: event UsdcDepositAddressChanged(address oldAddress, address newAddress)
func (_Quoter *QuoterFilterer) ParseUsdcDepositAddressChanged(log types.Log) (*QuoterUsdcDepositAddressChanged, error) {
	event := new(QuoterUsdcDepositAddressChanged)
	if err := _Quoter.contract.UnpackLog(event, "UsdcDepositAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
