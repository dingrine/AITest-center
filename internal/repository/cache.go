package repository

import (
	"context"
	"examCenter/internal/encoding"

	"github.com/CrazyThursdayV50/pkgo/builtin"
)

type TokenCache[T encoding.ValidData] interface {
	Set(ctx context.Context, tokenId int64, data T) error
	Get(ctx context.Context, tokenId int64) (builtin.UnWrapper[T], error)
	Has(ctx context.Context, tokenId int64) (bool, error)
	Del(ctx context.Context, tokenId int64) error
}

type Cache[T encoding.ValidData] interface {
	Set(ctx context.Context, key string, data T) error
	Get(ctx context.Context, key string) (builtin.UnWrapper[T], error)
	Has(ctx context.Context, key string) (bool, error)
	Del(ctx context.Context, key string) error
}
