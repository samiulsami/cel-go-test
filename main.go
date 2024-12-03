package main

import (
	"encoding/json"
	"log"

	"github.com/google/cel-go/cel"
	"k8s.io/klog/v2"
)

func main() {
	expression := `conditions.filter(c, c.status == "True")`

	env, err := cel.NewEnv(
		cel.Variable("conditions", cel.ListType(cel.MapType(cel.StringType, cel.DynType))),
	)
	if err != nil {
		klog.Fatalf("failed to create CEL environment: %v", err)
	}

	ast, iss := env.Compile(expression)
	if iss != nil && iss.Err() != nil {
		klog.Fatalf("failed to compile CEL expression: %v", iss.Err())
	}

	program, err := env.Program(ast)
	if err != nil {
		klog.Fatalf("failed to create CEL program: %v", err)
	}

	out, details, err := program.Eval(getsampleObj())
	if err != nil {
		klog.Fatalf("failed to evaluate CEL expression: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(out.Value(), "", "  ")
	if err != nil {
		klog.Fatalf("failed to marshal CEL expression result: %v", err)
	}
	log.Println(string(prettyJSON))

	klog.Infof("CEL expression result: %s", out.Value())

	klog.Infof("CEL expression details: %v", details)
}

func getsampleObj() map[string]interface{} {
	return map[string]interface{}{
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
