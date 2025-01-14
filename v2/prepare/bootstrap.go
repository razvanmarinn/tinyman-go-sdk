package prepare

import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"

	"github.com/razvanmarinn/tinyman-go-sdk/utils"
	"github.com/razvanmarinn/tinyman-go-sdk/v2/constants"
	"github.com/razvanmarinn/tinyman-go-sdk/v2/contracts"
)

// BootstrapTransactions prepares a transaction group to bootstrap a new pool
func BootstrapTransactions(
	validatorAppID,
	asset1ID,
	asset2ID uint64,
	asset1UnitName,
	asset2UnitName,
	senderAddress string,
	sp types.SuggestedParams,
) (*utils.TransactionGroup, error) {
	var err error
	var tx1 types.Transaction
	var tx2 types.Transaction
	var tx3 types.Transaction
	var tx4 types.Transaction

	poolAccount, err := contracts.PoolLogicSigAccount(validatorAppID, asset1ID, asset2ID)
	if err != nil {
		return nil, err
	}

	poolAddress, err := poolAccount.Address()
	if err != nil {
		return nil, err
	}

	if asset1ID <= asset2ID {
		return nil, fmt.Errorf("prepare bootstraping error caused by asset1 id <= asset2 id")
	}

	bootstrapAmount := uint64(constants.BootstrapTransactionAmount)
	if asset2ID == 0 {
		asset2UnitName = constants.AlgoTokenUnitName
		bootstrapAmount = constants.BootstrapTransactionAmountForAlgo
	}

	tx1, err = transaction.MakePaymentTxn(senderAddress, poolAddress.String(), bootstrapAmount, nil, "", sp)
	if err != nil {
		return nil, err
	}

	asset1IDBytes, err := utils.IntToBytes(asset1ID)
	if err != nil {
		return nil, err
	}
	asset2IDBytes, err := utils.IntToBytes(asset2ID)
	if err != nil {
		return nil, err
	}

	appArgs := [][]byte{[]byte(constants.BOOTSTRAP_APP_ARGUMENT), asset1IDBytes, asset2IDBytes}
	foreignAssets := []uint64{asset1ID, asset2ID}
	if asset2ID == 0 {
		foreignAssets = []uint64{asset1ID}
	}
	tx2, err = transaction.MakeApplicationOptInTx(
		validatorAppID,
		appArgs,
		nil,
		nil,
		foreignAssets,
		sp,
		poolAddress,
		nil,
		types.Digest{},
		[32]byte{},
		types.Address{},
	)
	if err != nil {
		return nil, err
	}

	tx3, err = transaction.MakeAssetCreateTxn(
		poolAddress.String(),
		nil,
		sp,
		constants.TotalLiquidityTokens,
		constants.LiquidityTokenDecimals,
		false,
		"",
		"",
		"",
		"",
		constants.LiquidityAssetUnitName,
		fmt.Sprintf("TinymanPool1.1 %s-%s", asset1UnitName, asset2UnitName),
		constants.TinyManURL,
		"",
	)
	if err != nil {
		return nil, err
	}

	tx4, err = transaction.MakeAssetAcceptanceTxn(poolAddress.String(), nil, sp, asset1ID)
	if err != nil {
		return nil, err
	}

	txs := []types.Transaction{tx1, tx2, tx3, tx4}
	if asset2ID > 0 {
		tx, err := transaction.MakeAssetAcceptanceTxn(poolAddress.String(), nil, sp, asset2ID)
		if err != nil {
			return nil, err
		}

		txs = append(txs, tx)
	}

	txGroup, err := utils.NewTransactionGroup(txs)
	if err != nil {
		return nil, err
	}

	if err := txGroup.SignWithLogicSig(poolAccount); err != nil {
		return nil, err
	}

	return txGroup, nil
}
