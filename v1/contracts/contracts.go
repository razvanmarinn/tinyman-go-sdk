package contracts

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/algorand/go-algorand-sdk/v2/crypto"
)

var POOL_LOGICSIG_TEMPLATE = ("BoAYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgQBbNQA0ADEYEkQxGYEBEkSBAUM=")

// PoolLogicSigAccount creates a logic signature account of the pool
func PoolLogicSigAccount(validatorAppID, asset1ID, asset2ID uint64) (*crypto.LogicSigAccount, error) {
	if asset2ID > asset1ID {
		asset1ID, asset2ID = asset2ID, asset1ID
	}
	program, err := base64.StdEncoding.DecodeString(POOL_LOGICSIG_TEMPLATE)
	if err != nil {
		return nil, err
	}
	binary.BigEndian.PutUint64(program[3:11], validatorAppID) 
	binary.BigEndian.PutUint64(program[11:19], asset1ID)     
	binary.BigEndian.PutUint64(program[19:27], asset2ID)      

	poolAccount, _ := crypto.MakeLogicSigAccountEscrowChecked(program, nil)

	return &poolAccount, nil
}
