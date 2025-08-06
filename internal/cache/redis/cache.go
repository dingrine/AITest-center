package redis

import (
	"context"
	"errors"
	"examCenter/internal/encoding"
	"examCenter/pkg/json"
	"examCenter/pkg/redis"

	"github.com/CrazyThursdayV50/pkgo/builtin"
	"github.com/CrazyThursdayV50/pkgo/builtin/wrap"
)

type cache[T encoding.ValidData] struct {
	client *redis.Client
}

func NewCache[T encoding.ValidData](client *redis.Client) *cache[T] {
	return &cache[T]{
		client: client,
	}
}

func (r *cache[D]) set(ctx context.Context, key string, data D) error {
	cmd := r.client.Client.Set(ctx, key, data, -1)
	return cmd.Err()
}

func (r *cache[D]) Set(ctx context.Context, key string, data D) error {
	return r.set(ctx, key, data)
}

func (r *cache[D]) get(ctx context.Context, key string) (builtin.UnWrapper[D], error) {
	cmd := r.client.Client.Get(ctx, key)
	data, err := cmd.Result()
	if errors.Is(err, redis.Nil) {
		return wrap.Nil[D](), nil
	}

	if err != nil {
		return wrap.Nil[D](), err
	}

	var model D
	err = json.JSON().UnmarshalFromString(data, &model)
	if err != nil {
		return nil, err
	}
	return wrap.Wrap(model), nil
}

func (r *cache[D]) Get(ctx context.Context, key string) (builtin.UnWrapper[D], error) {
	return r.get(ctx, key)
}

func (r *cache[D]) has(ctx context.Context, key string) (bool, error) {
	cmd := r.client.Client.Keys(ctx, key)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	return len(cmd.Val()) > 0, nil
}

func (r *cache[D]) Has(ctx context.Context, key string) (bool, error) {
	return r.has(ctx, key)
}

func (r *cache[D]) del(ctx context.Context, key string) error {
	cmd := r.client.Client.Del(ctx, key)
	return cmd.Err()
}

func (r *cache[D]) Del(ctx context.Context, key string) error {
	return r.del(ctx, key)
}
