package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

 func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
 	 count := k.GetPostCount(ctx)
	 post.Id = count

	 store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	 byteKey := make([]byte, 8)

	 binary.BigEndian.PutUint64(byteKey,post.Id)

	 appendedValue := k.cdc.MustMarshal(&post)

 	 store.Set(byteKey, appendedValue)
 	 k.SetPostCount(ctx, count+1)
 	 return count
 }

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
	  return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
  }

  func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.PostCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Post-count- to count
	store.Set(byteKey, bz)
  }
  