<h1 align="center">CosmWasm Lifecycle Module</h1>

![IMG](./docs/logo.jpg)

### Abstract 

CosmWasm Lifecycle blockchain module leverages [CosmoSDK's lifecycle](https://docs.cosmos.network/main/build/building-modules/beginblock-endblock) to facilitate the execution of smart contracts at the initiation and conclusion of each block. Given the necessity for swift and resource-light execution in both stages, this module mandates [Gov](https://docs.cosmos.network/main/build/modules/gov) voting and demands a collateral deposit for each smart contract on an individual basis. This collateral deposit will be burned if the smart contract fails to execute multiple times.


### Usecases

Automatic execution of smart contracts is very useful because enable multiple usecases that involves consensus models like oracle data voting, automatic LP rebalancing, automatic disputes, automatic rewards claiming, restaking... 

### Drawbacks

But automatic smart contract executions in consensus is a double edge sword because too many smart contract executions or too many operations in the smart contracts can slowdown the block production. 

### Solution to the drawbacks

This module addresses the earlier issue by empowering the chain's governance to collectively vote on enabling each smart contract's automatic execution.

Additionally, a secondary measure is implemented, requiring a collateral deposit for each smart contract if it wants to be executed. This deposit will be burned if the smart contract fails the execution multiple times