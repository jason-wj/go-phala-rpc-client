package phatypes

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type Result struct {
	StakePool StakePool
}

type StakePool struct {
	LockAccount        types.Address
	OwnerRewardAccount types.Address
}
