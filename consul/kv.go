package consul

import (
	"encoding/json"
	consulapi "github.com/hashicorp/consul/api"
)

func SetKV[T any](kv *consulapi.KV, q *consulapi.WriteOptions, key string, value T) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = kv.Put(&consulapi.KVPair{Key: key, Value: v}, q)
	return err
}

func GetKV[v any](kv *consulapi.KV, key string, q *consulapi.QueryOptions, value v) (bool, error) {
	pair, _, err := kv.Get(key, q)
	if err != nil {
		return false, err
	}
	if pair == nil {
		return false, nil
	}

	if err = json.Unmarshal(pair.Value, &value); err != nil {
		return false, err
	}
	return true, nil
}

func ListKVs[v any](kv *consulapi.KV, prefix string, q *consulapi.QueryOptions) ([]v, error) {
	pairs, _, err := kv.List(prefix, q)
	if err != nil {
		return nil, err
	}

	elements := make([]v, 0, len(pairs))
	for _, kvPair := range pairs {
		var e v
		if err := json.Unmarshal(kvPair.Value, &e); err != nil {
			return nil, err
		}
		elements = append(elements, e)
	}

	return elements, nil
}

func DeleteKV(kv *consulapi.KV, key string, w *consulapi.WriteOptions) error {
	_, err := kv.Delete(key, w)
	return err
}

func DeleteKVTree(kv *consulapi.KV, prefix string, w *consulapi.WriteOptions) error {
	_, err := kv.DeleteTree(prefix, w)
	return err
}
