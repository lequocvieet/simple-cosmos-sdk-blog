package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	//get the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Post
	var post = types.Post{
		Creator:   msg.Creator,
		Title:     msg.Title,
		Body:      msg.Body,
		CreatedAt: ctx.BlockHeight(),
	}

	//Add a post to the store and get back ID
	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id: id}, nil
}
