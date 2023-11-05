package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/emidev98/lifecycle-hooks/x/lifecycle-hooks/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueryServer struct {
	Keeper
}

var _ types.QueryServer = QueryServer{}

func (k QueryServer) Contracts(c context.Context, req *types.QueryContractsRequest) (res *types.QueryContractsResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	contractsStore := prefix.NewStore(store, types.ContractsKey)

	pageRes, err := query.Paginate(contractsStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.Contract
		k.cdc.MustUnmarshal(value, &contract)
		res.Contracts = append(res.Contracts, contract)
		return nil
	})
	if err != nil {
		return nil, err
	}
	res.Pagination = pageRes
	return res, nil
}

func (k QueryServer) Contract(c context.Context, req *types.QueryContractRequest) (res *types.QueryContractResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	contractAddress, err := sdk.AccAddressFromBech32(req.ContractAddress)
	if err == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid contract address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	contract, _ := k.GetContract(ctx, contractAddress)
	res.Contract = contract
	return res, nil
}
