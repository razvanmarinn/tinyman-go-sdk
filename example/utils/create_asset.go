package utils

import (
	"context"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/razvanmarinn/tinyman-go-sdk/v1"
)

// CreateAsset create a new asset
func CreateAsset(
	assetName,
	unitName string,
	decimals uint32,
	totalIssuance uint64,
	userAddress string,
	account *crypto.Account,
	ac *algod.Client,
	tc *tinyman.Client,
) (uint64, error) {
	txParams, err := ac.SuggestedParams().Do(context.Background())
	if err != nil {
		return 0, err
	}

	var note []byte = nil
	addr := userAddress
	defaultFrozen := false
	reserve := addr
	freeze := addr
	clawback := addr
	manager := addr
	assetURL := "http://someurl"
	assetMetadataHash := "thisIsSomeLength32HashCommitment"

	txn, err := transaction.MakeAssetCreateTxn(addr, note, txParams,
		totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash,
	)
	if err != nil {
		return 0, err
	}

	txid, stx, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		return 0, err
	}

	if _, err := ac.SendRawTransaction(stx).Do(context.Background()); err != nil {
		return 0, err
	}

	confirmedTxn, err := transaction.WaitForConfirmation(ac, txid, 4, context.Background())
	if err != nil {
		return 0, err
	}

	ctx := context.Background()
	if err := OptInAssetIfNeeded(ctx, tc, account, confirmedTxn.AssetIndex); err != nil {
		return 0, err
	}

	return confirmedTxn.AssetIndex, nil
}
