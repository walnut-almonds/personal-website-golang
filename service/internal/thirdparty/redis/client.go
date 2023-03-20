package redis

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

type IRedisClient interface {
	CheckIfKeyExists(ctx context.Context, key string) (bool, error)
	GetString(ctx context.Context, key string) (string, error)
	GetJSON(ctx context.Context, key string, val interface{}) error
	GetJSONWithMultiKeys(ctx context.Context, val interface{}, keys ...string) error
	SubscribeChannel(ctx context.Context, channels ...string) <-chan *redis.Message
	AddToSet(ctx context.Context, key string, values ...interface{}) error
	AddToSortedSet(ctx context.Context, key string, score float64, value interface{}) error
	GetAllInSet(ctx context.Context, key string) ([]string, error)
	GetAllInSortedSet(ctx context.Context, key string) ([]string, error)
	RemoveFromSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	RemoveFromSortedSet(ctx context.Context, key string, value interface{}) error
	RemoveAllFromSet(ctx context.Context, key string) (int64, error)
	RemoveAllFromSortedSet(ctx context.Context, key string) (int64, error)
	CheckInSet(ctx context.Context, key string, values interface{}) (bool, error)
	GetSetCount(ctx context.Context, key string) (int64, error)
	Exists(ctx context.Context, key ...string) (bool, error)
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	Del(ctx context.Context, key string) (int64, error)
	MGet(ctx context.Context, key string) ([]interface{}, error)
	HSet(ctx context.Context, key string, field string, val interface{}) error
	HMSet(ctx context.Context, key string, val ...interface{}) error
	HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error)
	HGet(ctx context.Context, key string, field string, val interface{}) error
	HLen(ctx context.Context, key string) (int64, error)
	HDel(ctx context.Context, key string, field string) (int64, error)
	ZAdd(ctx context.Context, key string, member string, score float64) (int64, error)
	ZRem(ctx context.Context, key string, member string) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	Publish(ctx context.Context, channel string, message interface{}) (int64, error)
	Subscribe(ctx context.Context, channel string) <-chan *redis.Message
	HMGetByKey(ctx context.Context, key string, values []string) (map[string]string, error)
	HMSetByKey(ctx context.Context, key string, values map[string]interface{}) (bool, error)
	IncrByFloat(ctx context.Context, key string, value float64) (float64, error)
	HIncrByFloat(ctx context.Context, key string, field string, value float64) (float64, error)

	Pipeline() redis.Pipeliner
}

type redisClient struct {
	client *redis.Client
}

func (r *redisClient) CheckIfKeyExists(ctx context.Context, key string) (bool, error) {
	val, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	return val > 0, nil
}

func (r *redisClient) GetString(ctx context.Context, key string) (string, error) {
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) GetJSON(ctx context.Context, key string, val interface{}) error {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	err = json.Unmarshal(data, val)
	if err != nil {
		return xerrors.Errorf("無法解析 JSON 資料 %s: %w", key, err)
	}

	return nil
}

func (r *redisClient) GetJSONWithMultiKeys(ctx context.Context, val interface{}, keys ...string) error {
	data, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 一次取得多筆資料 %v: %w", keys, err)
	}

	rawStrings := convertInterfaceArrayToStringArray(data)
	rawString := "[" + strings.Join(rawStrings, ",") + "]"

	if err := json.Unmarshal([]byte(rawString), val); err != nil {
		return xerrors.Errorf("無法解析 JSON 資料 %s: %w", keys, err)
	}

	return nil
}

func (r *redisClient) SubscribeChannel(ctx context.Context, channels ...string) <-chan *redis.Message {
	return r.client.Subscribe(ctx, channels...).Channel()
}

func (r *redisClient) AddToSet(ctx context.Context, key string, values ...interface{}) error {
	_, err := r.client.SAdd(ctx, key, values...).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 新增值 %v 進 Set %s: %w", values, key, err)
	}
	return err
}

func (r *redisClient) AddToSortedSet(ctx context.Context, key string, score float64, value interface{}) error {
	_, err := r.client.ZAdd(ctx, key, &redis.Z{Score: score, Member: value}).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 新增值 %v 進 SortedSet %s: %w", value, key, err)
	}
	return err
}

func (r *redisClient) GetAllInSet(ctx context.Context, key string) ([]string, error) {
	values, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return values, nil
}

func (r *redisClient) GetAllInSortedSet(ctx context.Context, key string) ([]string, error) {
	values, err := r.client.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return values, nil
}

func (r *redisClient) RemoveFromSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	count, err := r.client.SRem(ctx, key, values...).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 的 Set %s 刪除 %v: %w", key, values, err)
	}
	return count, err
}

func (r *redisClient) RemoveFromSortedSet(ctx context.Context, key string, value interface{}) error {
	_, err := r.client.ZRem(ctx, key, value).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 的 SortedSet %s 刪除 %v: %w", key, value, err)
	}
	return err
}

func (r *redisClient) RemoveAllFromSet(ctx context.Context, key string) (int64, error) {
	count, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 刪除 Set %s 的全部值: %w", key, err)
	}
	return count, nil
}

func (r *redisClient) RemoveAllFromSortedSet(ctx context.Context, key string) (int64, error) {
	count, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 刪除 Sorted Set %s 的全部值: %w", key, err)
	}
	return count, nil
}

func (r *redisClient) CheckInSet(ctx context.Context, key string, values interface{}) (bool, error) {
	existed, err := r.client.SIsMember(ctx, key, values).Result()
	if err != nil {
		return false, xerrors.Errorf("無法從 Redis 確認 %s 是否存在: %w", key, err)
	}
	return existed, nil
}

func (r *redisClient) GetSetCount(ctx context.Context, key string) (int64, error) {
	count, err := r.client.SCard(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return count, nil
}

func (r *redisClient) Exists(ctx context.Context, key ...string) (bool, error) {
	data, err := r.client.Exists(ctx, key...).Result()
	if err != nil {
		return false, xerrors.Errorf("無法使用 Redis Exists %s", err)
	}
	return data > 0, nil
}

func (r *redisClient) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	if err := r.client.Set(ctx, key, val, expiration).Err(); err != nil {
		return xerrors.Errorf("無法 Set Redis 值 %s: %w", key, err)
	}
	return nil
}

func (r *redisClient) Get(ctx context.Context, key string, val interface{}) error {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	err = json.Unmarshal(data, val)
	if err != nil {
		return xerrors.Errorf("無法解決 JSON 資料 %s: %w", key, err)
	}

	return nil
}

func (r *redisClient) Del(ctx context.Context, key string) (int64, error) {
	data, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return data, xerrors.Errorf("無法從 Redis Del %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) MGet(ctx context.Context, key string) ([]interface{}, error) {
	data, err := r.client.MGet(ctx, key).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return data, nil
}

func (r *redisClient) HMSet(ctx context.Context, key string, val ...interface{}) error {
	if err := r.client.HMSet(ctx, key, val...).Err(); err != nil {
		return xerrors.Errorf("無法 HMSet Redis 值 %s: %w", key, err)
	}
	return nil
}

func (r *redisClient) HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error) {
	data, err := r.client.HMGet(ctx, key, field...).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis HMGet %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) HSet(ctx context.Context, key string, field string, val interface{}) error {
	if err := r.client.HSet(ctx, key, field, val).Err(); err != nil {
		return xerrors.Errorf("無法 HSet Redis 值 %s: %w", key, err)
	}
	return nil
}

func (r *redisClient) HGet(ctx context.Context, key string, field string, val interface{}) error {
	data, err := r.client.HGet(ctx, key, field).Bytes()
	if err != nil {
		return xerrors.Errorf("無法從 Redis HGet %s-%s: %w", key, field, err)
	}
	if strVal, ok := val.(*string); ok {
		*strVal = string(data)
		return nil
	}
	if err = json.Unmarshal(data, val); err != nil {
		return xerrors.Errorf("無法解決 JSON 資料 %s-%s: %w", key, field, err)
	}

	return nil
}

func (r *redisClient) HLen(ctx context.Context, key string) (int64, error) {
	data, err := r.client.HLen(ctx, key).Result()
	if err != nil {
		return data, xerrors.Errorf("無法從 Redis HLen %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) HDel(ctx context.Context, key string, field string) (int64, error) {
	data, err := r.client.HDel(ctx, key, field).Result()
	if err != nil {
		return data, xerrors.Errorf("無法從 Redis HDel %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) ZAdd(ctx context.Context, key string, member string, score float64) (int64, error) {
	memberObj := &redis.Z{
		Score:  score,
		Member: member,
	}
	data, err := r.client.ZAdd(ctx, key, memberObj).Result()
	if err != nil {
		return data, xerrors.Errorf("無法從 Redis ZAdd %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	data, err := r.client.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		return data, xerrors.Errorf("無法從 Redis ZRange %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) ZRem(ctx context.Context, key string, member string) (int64, error) {
	data, err := r.client.ZRem(ctx, key, member).Result()
	if err != nil {
		return data, xerrors.Errorf("無法從 Redis ZRem %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	data, err := r.client.Publish(ctx, channel, message).Result()
	if err != nil {
		return data, xerrors.Errorf("無法 Publish 資料到 Redis HDel %s: %w", channel, err)
	}
	return data, nil
}

func (r *redisClient) Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	return r.client.Subscribe(ctx, channel).Channel()
}

func (r *redisClient) HMGetByKey(ctx context.Context, key string, fields []string) (map[string]string, error) {
	data, err := r.client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從Redis 取得 %s: %w", key, err)
	}

	outputData := map[string]string{}
	for i, v := range data {
		if s, ok := v.(string); ok {
			outputData[fields[i]] = s
		}
	}

	return outputData, nil
}

func (r *redisClient) HMSetByKey(ctx context.Context, key string, values map[string]interface{}) (bool, error) {
	result, err := r.client.HMSet(ctx, key, convertToMap(values)).Result()
	if err != nil {
		return false, xerrors.Errorf("無法從Redis 新增 Key %s 值 %w", key, err)
	}

	return result, nil
}

func (r *redisClient) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	data, err := r.client.IncrByFloat(ctx, key, value).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從Redis 新增 %s: %w", key, err)
	}

	return data, nil
}

func (r *redisClient) HIncrByFloat(ctx context.Context, key string, field string, value float64) (float64, error) {
	data, err := r.client.HIncrByFloat(ctx, key, field, value).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從Redis 新增 %s: %w", key, err)
	}

	return data, nil
}

func convertToMap(value interface{}) map[string]interface{} {
	valueMap := map[string]interface{}{}

	matchByte, err := json.Marshal(value)
	if err != nil {
		return valueMap
	}

	if err := json.Unmarshal(matchByte, &valueMap); err != nil {
		return valueMap
	}

	return valueMap
}

func (r *redisClient) Pipeline() redis.Pipeliner {
	return r.client.Pipeline()
}
