package constants

const (
	TinyManURL       = "https://tinyman.org"
	AlgodTestnetHost = "https://testnet-api.algonode.cloud"

	// AlgodMainnetHost is the algorand main net url
	AlgodMainnetHost = "https://mainnet-api.algonode.cloud"

	// TestnetValidatorAppIdV1_1 is the Tinyman test net validator app id version 1.1
	TestnetValidatorAppIdV2 uint64 = 148607000

	// MainnetValidatorAppIdV1_1 is the Tinyman main net validator app id version 1.1
	MainnetValidatorAppIdV2 uint64 = 1002541853

	// TestnetValidatorAppId is an alias for the current Tinyman test net validator app id
	TestnetValidatorAppId = TestnetValidatorAppIdV2

	// MainnetValidatorAppId is an alias for the current Tinyman main net validator app id
	MainnetValidatorAppId  = MainnetValidatorAppIdV2
	POOL_LOGICSIG_TEMPLATE = "BoAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgQBbNQA0ADEYEkQxGYEBEkSBAUM="

	// App Arguments for different functions
	BOOTSTRAP_APP_ARGUMENT             = "bootstrap"
	ADD_LIQUIDITY_APP_ARGUMENT         = "add_liquidity"
	ADD_INITIAL_LIQUIDITY_APP_ARGUMENT = "add_initial_liquidity"
	REMOVE_LIQUIDITY_APP_ARGUMENT      = "remove_liquidity"
	SWAP_APP_ARGUMENT                  = "swap"
	FLASH_LOAN_APP_ARGUMENT            = "flash_loan"
	VERIFY_FLASH_LOAN_APP_ARGUMENT     = "verify_flash_loan"
	FLASH_SWAP_APP_ARGUMENT            = "flash_swap"
	VERIFY_FLASH_SWAP_APP_ARGUMENT     = "verify_flash_swap"
	CLAIM_FEES_APP_ARGUMENT            = "claim_fees"
	CLAIM_EXTRA_APP_ARGUMENT           = "claim_extra"
	SET_FEE_APP_ARGUMENT               = "set_fee"
	SET_FEE_COLLECTOR_APP_ARGUMENT     = "set_fee_collector"
	SET_FEE_SETTER_APP_ARGUMENT        = "set_fee_setter"
	SET_FEE_MANAGER_APP_ARGUMENT       = "set_fee_manager"
	AlgoTokenName                      = "Algo"
	TotalLiquidityTokens               = 0xFFFFFFFFFFFFFFFF

	// SwapFixedInput is a fixed-input swap type
	SwapFixedInput = "fixed-input"

	// SwapFixedOutput is a fixed-output swap type
	SwapFixedOutput        = "fixed-output"
	LiquidityTokenDecimals = 6

	LiquidityTokenUnitName = "TM1POOL"

	LiquidityAssetUnitName = "TMPOOL11"

	MinBalancePerAccount = 100000

	MinBalancePerAsset = 100000

	MinBalancePerApp = 100000

	MinBalancePerAppByteSlice = 50000

	MinBalancePerAppUint = 28500

	BootstrapTransactionAmountForAlgo = 859000

	BootstrapTransactionAmount = 960000
	// AlgoTokenUnitName is the algo token unit name
	AlgoTokenUnitName = "ALGO"

	// AlgoTokenDecimals is the algo token decimals
	AlgoTokenDecimals = 6

	// Flexible and Single modes for liquidity
	ADD_LIQUIDITY_FLEXIBLE_MODE_APP_ARGUMENT = "flexible"
	ADD_LIQUIDITY_SINGLE_MODE_APP_ARGUMENT   = "single"

	// Fixed input and output modes for liquidity
	FIXED_INPUT_APP_ARGUMENT  = "fixed-input"
	FIXED_OUTPUT_APP_ARGUMENT = "fixed-output"

	// Pool Token and Asset Minimum Total Balance
	LOCKED_POOL_TOKENS             = 1000
	ASSET_MIN_TOTAL                = 1000000
	MIN_POOL_BALANCE_ASA_ALGO_PAIR = 300000 + (100000 + (25000+3500)*12 + (25000+25000)*2)
	MIN_POOL_BALANCE_ASA_ASA_PAIR  = MIN_POOL_BALANCE_ASA_ALGO_PAIR + 100000

	// Application State Constants
	APP_LOCAL_INTS   = 12
	APP_LOCAL_BYTES  = 2
	APP_GLOBAL_INTS  = 0
	APP_GLOBAL_BYTES = 3
	// TestnetValidatorAppAddress = get_application_address(TESTNET_VALIDATOR_APP_ID)
	// BurnFee is a burn transaction fee
	BurnFee uint64 = 4000

	// MintFee is a mint transaction fee
	MintFee uint64 = 2000

	// RedeemFee is a redeem transaction fee
	RedeemFee uint64 = 2000

	// SwapFee is a swap transaction fee
	SwapFee uint64 = 2000

	TESTNET_SWAP_ROUTER_APP_ID_V1 = 184778019
	MAINNET_SWAP_ROUTER_APP_ID_V1 = 1083651166
)

var (
	// SwapTypeMapping is a mapping for swap type
	SwapTypeMapping = map[string]string{
		SwapFixedInput:  "fi",
		SwapFixedOutput: "fo",
	}
)
