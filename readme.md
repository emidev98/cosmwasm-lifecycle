<h1 align="center">CosmWasm Lifecycle Module</h1>

![IMG](./docs/logo.jpg)

CosmWasm Lifecycle blockchain module leverages [CosmoSDK's lifecycle](https://docs.cosmos.network/main/build/building-modules/beginblock-endblock) to facilitate the execution of smart contracts at the initiation and conclusion of each block. Given the necessity for swift and resource-light execution in both stages, this module mandates [Gov](https://docs.cosmos.network/main/build/modules/gov) voting and demands a collateral deposit for each smart contract on an individual basis. This collateral deposit will be burned if the smart contract fails to execute multiple times.






