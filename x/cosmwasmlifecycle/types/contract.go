package types

// Can Execute returns true if the contract can be executed
// based on the module parameters, strikes and available deposits
func (c *Contract) CanExecute(p Params, executionType ExecutionType) bool {
	hasTooManyStrikes := c.Strikes >= p.StrikesToDisableExecution
	hasEnoughDeposits := c.Deposit.IsGTE(p.MinDeposit)
	hasExecutionType := c.ExecutionType == ExecutionType_BEGIN_AND_END_BLOCK || c.ExecutionType == executionType

	return hasTooManyStrikes && hasEnoughDeposits && hasExecutionType
}

// Check if the contract has enough strikes to disable execution
// based on the params of the module
func (c *Contract) HasEnoughStrikesToDisableExecution(strikesToDisabeExecution int64) bool {
	return c.Strikes >= strikesToDisabeExecution
}
