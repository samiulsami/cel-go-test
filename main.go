package main

import (
	"fmt"
	"reflect"

	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/klog/v2"
)

func main() {
	expression := `status.phase == "Running"`

	env, err := cel.NewEnv(
		// cel.Variable("conditions", cel.ListType(cel.MapType(cel.StringType, cel.DynType))),
		cel.Variable("status", cel.MapType(cel.StringType, cel.DynType)),
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

func getsampleObj2() map[string]interface{} {
	return map[string]interface{}{
		"components": map[string]interface{}{
			"test1": map[string]interface{}{
				"status": map[string]interface{}{
					"phase": "Running",
				},
			},
			"test2": map[string]interface{}{
				"status": map[string]interface{}{
					"phase": "Running",
				},
			},
			"test3": map[string]interface{}{
				"status": map[string]interface{}{
					"phase": "Stopped",
				},
			},
			"test4": map[string]interface{}{
				"status": map[string]interface{}{
					"phase": "Stopped",
				},
			},
		},
	}
}

func getsampleObj() map[string]interface{} {
	return map[string]interface{}{
		"status": map[string]interface{}{
			"phase": "Running",
		},
		"conditions": []map[string]interface{}{
			{
				"lastTransitionTime": "2024-10-25T09:22:51Z",
				"message":            "The KubeDB operator has started the provisioning of MongoDB: monitoring/mongodb",
				"reason":             "DatabaseProvisioningStartedSuccessfully",
				"status":             "True",
				"type":               "ProvisioningStarted",
			},
			{
				"lastTransitionTime": "2024-11-07T11:20:14Z",
				"message":            "All desired replicas are ready.",
				"reason":             "AllReplicasReady",
				"status":             "True",
				"type":               "ReplicaReady",
			},
			{
				"lastTransitionTime": "2024-11-07T11:20:45Z",
				"message":            "The MongoDB: monitoring/mongodb is accepting client requests.",
				"observedGeneration": 2,
				"reason":             "DatabaseAcceptingConnectionRequest",
				"status":             "True",
				"type":               "AcceptingConnection",
			},
			{
				"lastTransitionTime": "2024-11-07T11:20:45Z",
				"message":            "The MongoDB: monitoring/mongodb is ready.",
				"observedGeneration": 2,
				"reason":             "ReadinessCheckSucceeded",
				"status":             "True",
				"type":               "Ready",
			},
			{
				"lastTransitionTime": "2024-10-25T09:23:23Z",
				"message":            "The MongoDB: monitoring/mongodb is successfully provisioned.",
				"observedGeneration": 2,
				"reason":             "DatabaseSuccessfullyProvisioned",
				"status":             "True",
				"type":               "Provisioned",
			},
		},
		"observedGeneration": 2,
		"phase":              "Ready",
	}
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
