package prepare

import (
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"

	"github.com/razvanmarinn/tinyman-go-sdk/utils"
	"github.com/razvanmarinn/tinyman-go-sdk/v2/constants"
	"github.com/razvanmarinn/tinyman-go-sdk/v2/contracts"
)

// RedeemTransactions prepares a transaction group to redeem a specified excess asset amount from a pool.
func RedeemTransactions(
	validatorAppID,
	asset1ID,
	asset2ID,
	liquidityAssetID,
	assetID,
	assetAmount uint64,
	senderAddress string,
	sp types.SuggestedParams,
) (*utils.TransactionGroup, error) {
	var err error
	var tx1 types.Transaction
	var tx2 types.Transaction
	var tx3 types.Transaction

	poolAccount, err := contracts.PoolLogicSigAccount(validatorAppID, asset1ID, asset2ID)
	if err != nil {
		return nil, err
	}

	poolAddress, err := poolAccount.Address()
	if err != nil {
		return nil, err
	}

	tx1, err = transaction.MakePaymentTxn(senderAddress, poolAddress.String(), constants.RedeemFee, []byte("fee"), "", sp)
	if err != nil {
		return nil, err
	}

	foreignAssets := []uint64{asset1ID, asset2ID, liquidityAssetID}
	if asset2ID == 0 {
		foreignAssets = []uint64{asset1ID, liquidityAssetID}
	}

	tx2, err = transaction.MakeApplicationNoOpTx(
		validatorAppID,
		[][]byte{[]byte("redeem")},
		[]string{senderAddress},
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

	if assetID != 0 {
		tx3, err = transaction.MakeAssetTransferTxn(poolAddress.String(), senderAddress, assetAmount, nil, sp, "", assetID)
		if err != nil {
			return nil, err
		}
	} else {
		tx3, err = transaction.MakePaymentTxn(poolAddress.String(), senderAddress, assetAmount, nil, "", sp)
		if err != nil {
			return nil, err
		}
	}

	txGroup, err := utils.NewTransactionGroup([]types.Transaction{tx1, tx2, tx3})
	if err != nil {
		return nil, err
	}

	if err := txGroup.SignWithLogicSig(poolAccount); err != nil {
		return nil, err
	}

	return txGroup, nil
}
