package prepare

import (
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"

	"github.com/razvanmarinn/tinyman-go-sdk/utils"
)

// AppOptOutTransactions prepares a transaction group to opt-out of Tinyman
func AppOptOutTransactions(validatorAppID uint64, senderAddress string, sp types.SuggestedParams) (*utils.TransactionGroup, error) {
	addr, err := types.DecodeAddress(senderAddress)
	if err != nil {
		return nil, err
	}

	tx, err := transaction.MakeApplicationClearStateTx(
		validatorAppID,
		nil,
		nil,
		nil,
		nil,
		sp,
		addr,
		nil,
		types.Digest{},
		[32]byte{},
		types.Address{},
	)
	if err != nil {
		return nil, err
	}

	txGroup, err := utils.NewTransactionGroup([]types.Transaction{tx})
	if err != nil {
		return nil, err
	}

	return txGroup, nil
}
