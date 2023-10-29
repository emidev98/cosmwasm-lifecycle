<h1 align="center">Lifecycle Hooks Module</h1>

![IMG](./docs/logo.jpg)

Lifecycle Hooks is a CosmosSDK blockchain module that uses the [modules lifecycle](https://docs.cosmos.network/main/build/building-modules/beginblock-endblock) to facilitate the execution of smart contracts at the beign and end of each block. 


### Usecases

Automatic execution of smart contracts is very useful because it enables multiple usecases like for example:
- laverage consensus of multiple parties to do oracle data voting, 
- decentralized LP rebalancing by a protocol itself, 
- automatic disputes resolution, 
- rewards claiming and restaking,

As you can see there are multiple use cases that will be enabled with this module.

### Drawbacks

Automatic smart contract executions in consensus is a double edge sword because too many smart contract executions or too many operations in the smart contracts can slowdown the block production significantly. 

### Solution to the drawbacks

<table>
    <tr>
        <td align="center"><img src="./docs/icons/governance.jpg" height="150px"></td>
        <td align="center"><img src="./docs/icons/collateral.jpg" height="150px"></td>
        <td align="center"><img src="./docs/icons/execution.jpg" height="150px"></td>
    </tr>
    <tr>
        <th><h3 align="center">Governance</h3></th>
        <th><h3 align="center">Collateral</h3></th>
        <th><h3 align="center">Execution</h3></th>
    </tr>
    <tr>
        <td>
            The first messure to the issues is to involve `chain governance` to collectively vote on enabling each smart contract execution at begin or end block. That way the overall community can decide if the usecase and optimization of the smart contract is good enough for the required computation.
        </td>
        <td>
            A secondary measure is to `require a collateral deposit` for each smart contract that will be executed on block lifecycle. This deposit will burned if the smart contract fails the execution many times (which is defined in the module params). 
        </td>
        <td>
            The latest measure is to allow smart contracts execution each `n number of blocks`. Which means that execution at end and begin block decreases computation complexity because it does not have to load all the wasm environment and try an execution each time.
        </td>
    </tr>
</table>

