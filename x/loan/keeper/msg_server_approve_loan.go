package keeper

import (
	"context"

	"simpleloan/x/loan/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	if msg.Creator == loan.Borrower {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot approve: borrow from yourself")
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)
	err := k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	loan.Lender = msg.Creator
	loan.State = "approved"
	k.SetLoan(ctx, loan)
	return &types.MsgApproveLoanResponse{}, nil
}
