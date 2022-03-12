package keeper

import (
	"encoding/binary"

	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetLongBookCount get the total number of longBook
func (k Keeper) GetLongBookCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LongBookCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLongBookCount set the total number of longBook
func (k Keeper) SetLongBookCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LongBookCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLongBook appends a longBook in the store with a new id and update the count
func (k Keeper) AppendLongBook(
	ctx sdk.Context,
	longBook types.LongBook,
) uint64 {
	// Create the longBook
	count := k.GetLongBookCount(ctx)

	// Set the ID of the appended value
	longBook.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	appendedValue := k.cdc.MustMarshal(&longBook)
	store.Set(GetKeyForLongBook(longBook), appendedValue)

	// Update longBook count
	k.SetLongBookCount(ctx, count+1)

	return count
}

// SetLongBook set a specific longBook in the store
func (k Keeper) SetLongBook(ctx sdk.Context, longBook types.LongBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	b := k.cdc.MustMarshal(&longBook)
	store.Set(GetKeyForLongBook(longBook), b)
}

// GetLongBook returns a longBook from its id
// DO NOT USE
func (k Keeper) GetLongBook(ctx sdk.Context, id uint64) (val types.LongBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	b := store.Get(GetLongBookIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetLongBookByPrice(ctx sdk.Context, price uint32) (val types.LongBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	b := store.Get(GetKeyForPrice(price))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLongBook removes a longBook from the store
// DO NOT USE
func (k Keeper) RemoveLongBook(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	store.Delete(GetLongBookIDBytes(id))
}

func (k Keeper) RemoveLongBookByPrice(ctx sdk.Context, price uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	store.Delete(GetKeyForPrice(price))
}

// GetAllLongBook returns all longBook
func (k Keeper) GetAllLongBook(ctx sdk.Context) (list []types.LongBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LongBookKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LongBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLongBookIDBytes returns the byte representation of the ID
func GetLongBookIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLongBookIDFromBytes returns ID in uint64 format from a byte array
func GetLongBookIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func GetKeyForLongBook(longBook types.LongBook) []byte {
	return GetKeyForPrice(uint32(longBook.Entry.Price))
}

func GetKeyForPrice(price uint32) []byte {
	key := make([]byte, 4)
	binary.BigEndian.PutUint32(key, uint32(price))
	return key
}
