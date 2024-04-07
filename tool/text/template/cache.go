package template

import (
	"github.com/patrickmn/go-cache"
	"os"
	"reflect"
)

// export THRIFTGO_TPL_OPTIMIZE=1
var useCache = os.Getenv("THRIFTGO_TPL_OPTIMIZE") == "1"
var fieldCache *cache.Cache

func init() {
	fieldCache = cache.New(cache.NoExpiration, 0)
	methodCache = cache.New(cache.NoExpiration, 0)
}

func CacheFieldByName(t reflect.Type, name string) (reflect.StructField, bool) {
	if !useCache {
		return t.FieldByName(name)
	}

	key := t.String() + "_" + name

	if st, ok := fieldCache.Get(key); ok {
		return st.(reflect.StructField), true
	}

	st, ok := t.FieldByName(name)
	if ok {
		fieldCache.Set(key, st, cache.NoExpiration)
	}

	return st, ok
}

var methodCache *cache.Cache

func CacheMethodByName(val reflect.Value, name string) reflect.Value {
	if false {
		return val.MethodByName(name)
	}

	m, ok := CacheMethodTypeByName(val.Type(), name)
	if !ok {
		return val.MethodByName(name)
	}
	return val.Method(m.Index)
}

func CacheMethodTypeByName(t reflect.Type, name string) (reflect.Method, bool) {

	key := t.String() + "_" + name

	if st, ok := methodCache.Get(key); ok {
		return st.(reflect.Method), true
	}

	st, ok := t.MethodByName(name)
	if ok {
		methodCache.Set(key, st, cache.NoExpiration)
	}

	return st, ok
}
