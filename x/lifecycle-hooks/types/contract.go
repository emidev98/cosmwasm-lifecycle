package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewCleanContract(executionType ExecutionType, executionFrequency int64, deposit sdk.Coin) Contract {
	return Contract{
		ExecutionType:        executionType,
		ExecutionFrequency:   executionFrequency,
		Deposit:              deposit,
		Strikes:              0,
		LatestBlockExecution: 0,
	}
}

// Compare contract strikes to the max allowed strikes
func (c *Contract) HaveMaxStrikesToPenalize(maxStrikes int64) bool {
	return c.Strikes >= maxStrikes
}

// Check if the contract has less stake than min deposit
func (c *Contract) HaveLessThanMinDeposit(minDeposit sdk.Coin) bool {
	return c.Deposit.IsLT(minDeposit)
}

// Can Execute returns true if the contract can be executed
func (c *Contract) CanExecute(currentBlockHeight int64, currentLifecycle ExecutionType) bool {
	matchExecutionType := c.ExecutionType == ExecutionType_BEGIN_AND_END_BLOCK || c.ExecutionType == currentLifecycle
	isFirstExecution := c.LatestBlockExecution == 0
	if matchExecutionType && isFirstExecution {
		return true
	}

	executionFrequencyReached := (c.LatestBlockExecution + c.ExecutionFrequency) <= currentBlockHeight
	if matchExecutionType && executionFrequencyReached {
		return true
	}

	return false
}
