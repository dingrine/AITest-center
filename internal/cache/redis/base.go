package redis

import (
	"context"
	"examCenter/internal/encoding"
	"examCenter/pkg/redis"
	"fmt"

	"github.com/CrazyThursdayV50/pkgo/builtin"
)

type tokenCache[D encoding.ValidData] struct {
	cache    *cache[D]
	tokenKey func(int64) string
}

func newTokenCache[D encoding.ValidData](client *redis.Client, tokenKey func(int64) string) *tokenCache[D] {
	return &tokenCache[D]{
		cache:    NewCache[D](client),
		tokenKey: tokenKey,
	}
}

func (r *tokenCache[D]) set(ctx context.Context, tokenId int64, data D) error {
	key := r.tokenKey(tokenId)
	return r.cache.set(ctx, key, data)
}

func (r *tokenCache[D]) Set(ctx context.Context, tokenId int64, data D) error {
	fmt.Println("redis set", tokenId)
	return r.set(ctx, tokenId, data)
}

func (r *tokenCache[D]) get(ctx context.Context, tokenId int64) (builtin.UnWrapper[D], error) {
	key := r.tokenKey(tokenId)
	return r.cache.get(ctx, key)
}

func (r *tokenCache[D]) Get(ctx context.Context, tokenId int64) (builtin.UnWrapper[D], error) {
	return r.get(ctx, tokenId)
}

func (r *tokenCache[D]) has(ctx context.Context, tokenId int64) (bool, error) {
	key := r.tokenKey(tokenId)
	return r.cache.has(ctx, key)
}

func (r *tokenCache[D]) Has(ctx context.Context, tokenId int64) (bool, error) {
	return r.has(ctx, tokenId)
}

func (r *tokenCache[D]) del(ctx context.Context, tokenId int64) error {
	key := r.tokenKey(tokenId)
	return r.cache.del(ctx, key)
}

func (r *tokenCache[D]) Del(ctx context.Context, tokenId int64) error {
	return r.del(ctx, tokenId)
}
