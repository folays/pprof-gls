package pprof_gls

import (
	"context"
	"reflect"
	"runtime/pprof"
	"unsafe"
)

var (
	LabelContextKey_reflect_value reflect.Value // runtime.pprof.labelContextKey : struct{}
	LabelContextKey_any           any           // runtime.pprof.labelContextKey : struct{}

	labelmap_reflect_type reflect.Type // *runtime.pprof.labelMap : *map[string]string
)

func init() {
	ctx := pprof.WithLabels(context.Background(), pprof.Labels())

	LabelContextKey_reflect_value = reflect.ValueOf(ctx).Elem().FieldByName("key")
	LabelContextKey_any = ReflectGetAnyFromUnexported(LabelContextKey_reflect_value)

	labelmap_reflect_type = reflect.ValueOf(ctx).Elem().FieldByName("val").Elem().Type()
}

func ReflectGetAnyFromUnexported(v reflect.Value) (any any) {
	//return v.Interface() // Go angrily won't let you do it, when it could have

	return reflect.NewAt(reflect.PointerTo(v.Type()).Elem(), v.Addr().UnsafePointer()).Elem().Interface()
}

//go:linkname runtime_getProfLabel runtime/pprof.runtime_getProfLabel
func runtime_getProfLabel() (labels unsafe.Pointer)

//go:linkname runtime_setProfLabel runtime/pprof.runtime_setProfLabel
func runtime_setProfLabel(labels unsafe.Pointer)

func Do(ctx context.Context, labels pprof.LabelSet, fn func(context.Context)) {
	if parentLabels := ctx.Value(LabelContextKey_any); parentLabels == nil {
		if glsLabels := runtime_getProfLabel(); glsLabels != nil {
			ctx = context.WithValue(ctx, LabelContextKey_any, reflect.NewAt(labelmap_reflect_type.Elem(), glsLabels).Interface())
		}

	}

	pprof.Do(ctx, labels, fn)
}
