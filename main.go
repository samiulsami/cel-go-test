package main

import (
	"fmt"
	"reflect"

	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/ext"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/klog/v2"
)

func main() {
	expression := `components.transformList(ca, c, c.status.phase=='Running', c)`

	env, err := cel.NewEnv(
		// cel.Variable("conditions", cel.ListType(cel.MapType(cel.StringType, cel.DynType))),
		// cel.Variable("status", cel.MapType(cel.StringType, cel.DynType)),
		ext.TwoVarComprehensions(),
	)
	if err != nil {
		klog.Fatalf("failed to create CEL environment: %v", err)
	}

	ast, iss := env.Parse(expression)
	if iss != nil && iss.Err() != nil {
		klog.Fatalf("failed to compile CEL expression: %v", iss.Err())
	}

	program, err := env.Program(ast)
	if err != nil {
		klog.Fatalf("failed to create CEL program: %v", err)
	}

	out, _, err := program.Eval(getsampleObj2())
	if err != nil {
		klog.Fatalf("failed to evaluate CEL program: %v", err)
	}
	fmt.Println(valueToJSON(out))
}

func valueToJSON(val ref.Val) string {
	v, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
	if err != nil {
		glog.Exit(err)
	}
	marshaller := protojson.MarshalOptions{Indent: "    "}
	bytes, err := marshaller.Marshal(v.(proto.Message))
	if err != nil {
		glog.Exit(err)
	}
	return string(bytes)
}
