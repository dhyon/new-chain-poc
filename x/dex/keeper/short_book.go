package keeper

import (
	"encoding/binary"

	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetShortBookCount get the total number of shortBook
func (k Keeper) GetShortBookCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ShortBookCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetShortBookCount set the total number of shortBook
func (k Keeper) SetShortBookCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ShortBookCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendShortBook appends a shortBook in the store with a new id and update the count
func (k Keeper) AppendShortBook(
	ctx sdk.Context,
	shortBook types.ShortBook,
) uint64 {
	// Create the shortBook
	count := k.GetShortBookCount(ctx)

	// Set the ID of the appended value
	shortBook.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	appendedValue := k.cdc.MustMarshal(&shortBook)
	store.Set(GetKeyForShortBook(shortBook), appendedValue)

	// Update shortBook count
	k.SetShortBookCount(ctx, count+1)

	return count
}

// SetShortBook set a specific shortBook in the store
func (k Keeper) SetShortBook(ctx sdk.Context, shortBook types.ShortBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	b := k.cdc.MustMarshal(&shortBook)
	store.Set(GetKeyForShortBook(shortBook), b)
}

// GetShortBook returns a shortBook from its id
// DO NOT USE
func (k Keeper) GetShortBook(ctx sdk.Context, id uint64) (val types.ShortBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	b := store.Get(GetShortBookIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetShortBookByPrice(ctx sdk.Context, price uint32) (val types.ShortBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	b := store.Get(GetKeyForPrice(price))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShortBook removes a shortBook from the store
// DO NOT USE
func (k Keeper) RemoveShortBook(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	store.Delete(GetShortBookIDBytes(id))
}

func (k Keeper) RemoveShortBookByPrice(ctx sdk.Context, price uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	store.Delete(GetKeyForPrice(price))
}

// GetAllShortBook returns all shortBook
func (k Keeper) GetAllShortBook(ctx sdk.Context) (list []types.ShortBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShortBookKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ShortBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetShortBookIDBytes returns the byte representation of the ID
func GetShortBookIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetShortBookIDFromBytes returns ID in uint64 format from a byte array
func GetShortBookIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func GetKeyForShortBook(shortBook types.ShortBook) []byte {
	return GetKeyForPrice(uint32(shortBook.Entry.Price))
}
