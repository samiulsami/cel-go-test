package main

func getsampleObj2() map[string]any {
	return map[string]any{
		"components": map[string]any{
			"test1": map[string]any{
				"status": map[string]any{
					"phase": "Running",
				},
				"otherkey": "hehe",
			},
			"test2": map[string]any{
				"status": map[string]any{
					"phase": "Running",
				},
				"otherkey": "othervalue",
			},
			"test3": map[string]any{
				"status": map[string]any{
					"phase": "Stopped",
				},
				"otherkey": "something else",
			},
			"test4": map[string]any{
				"status": map[string]any{
					"phase": "Stopped",
				},
				"otherkey": "what",
			},
		},
	}
}

func getsampleObj() map[string]any {
	return map[string]any{
		"status": map[string]any{
			"phase": "Running",
		},
		"conditions": []map[string]any{
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
