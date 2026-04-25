// Copyright 2020 FairwindsOps Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsr

import (
	"os"
	"strings"
)

// DefaultAZsEnvVar is the env var Gemini reads to discover the cluster-wide
// fallback list of FSR target AZs. SnapshotGroups that omit
// spec.fastSnapshotRestore.availabilityZones use this value.
const DefaultAZsEnvVar = "GEMINI_DEFAULT_FSR_AZS"

// EnabledEnvVar is the cluster-wide FSR opt-in. Default behaviour (unset or
// any value not listed below) is DISABLED — the controller will not initialize
// the AWS client and ReconcileFSR short-circuits for every SnapshotGroup.
// Set to "true", "1", "yes", or "on" (case-insensitive) to enable. When
// disabled, this is a pure skip: it does not clean up snapshots already
// FSR-enabled in AWS.
const EnabledEnvVar = "GEMINI_FSR_ENABLED"

// EnabledFromEnv reports the cluster-wide FSR opt-in state. Returns false
// unless the env var is explicitly set to a truthy value.
func EnabledFromEnv() bool {
	raw := strings.ToLower(strings.TrimSpace(os.Getenv(EnabledEnvVar)))
	switch raw {
	case "true", "1", "yes", "on":
		return true
	default:
		return false
	}
}

// DefaultAZsFromEnv parses GEMINI_DEFAULT_FSR_AZS into a slice. Accepts a
// comma-separated list with optional whitespace ("ap-northeast-1a, ap-northeast-1c").
// Returns nil if the env var is unset or empty.
func DefaultAZsFromEnv() []string {
	raw := os.Getenv(DefaultAZsEnvVar)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
