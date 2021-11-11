// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("apm-server", "XPackFields", asset.ModuleFieldsPri, AssetXPackFields); err != nil {
		panic(err)
	}
}

// AssetXPackFields returns asset data.
// This is the base64 encoded zlib format compressed contents of x-pack/apm-server.
func AssetXPackFields() string {
	return "eJyck0Fv2zAMhe/+FUTPsX+ADwOKnTsU2O6FKj87RG1JI6l2+feDHLtxFqNbl1so6vGJ33NNLzi15NJUm7igzhvHUE8wYa/1r+T8S0VkbCNaurt/fKAflz56OPfdVUQd1AunUm7pS0VEVLo3qrSoHsiFbu+g1gTPPXtKEhPEGHqYlQQ/MwuHgUb2CIqOejjLAiXN/khOyY6gI6vFQdxEPWPsyE4JTUWkxyj25GPoeWjJJKOic4u284CagpvQbl3NdZolWhok5rRUulNwE/uWejcqluIqtvxd9bosbiO2J3dtZHv5/TXvJ+v9vZMdAOvvUVC7YRAMztBtthT7Kw6rW22qapMMTe7jSHxP7l+yUHRW1uQElGeQUQqft0JXnKGYKryhpnNSRmcI/kTPsDcgEAc1yRNCeYtCXtlDP0V58QBrEoRjtywrB7aWJr0iP8YwVDsb/ppFEIw6Z458HEecd3hWnF9lR1bCa+niQBOPIyt8DN3s9mKn7OX/0vZH2KDGYSbYLHv5MHe70rdxvAwQaIpB8WQ8ofExB7vK2c3Gdvb2LU/PkAJ5k8iYbYhzABbwzV+na56arFdzFn7sJeqnfd1f3KyfQTF54+xwZllmrCx/BwAA//9RTKg4"
}
