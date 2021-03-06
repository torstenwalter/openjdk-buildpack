/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import (
	"os"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/libcfbuildpack/buildpack"
)

// Version returns the selected version of Java using the following precedence:
//
// 1. $BP_JAVA_VERSION
// 2. Build Plan Version
// 3. Buildpack Metadata "default_version"
// 4. Empty string
func Version(buildpack buildpack.Buildpack, dependency buildplan.Dependency) string {
	if version, ok := os.LookupEnv("BP_JAVA_VERSION"); ok {
		return version
	}

	if dependency.Version != "" {
		return dependency.Version
	}

	if version, ok := buildpack.Metadata["default_version"].(string); ok {
		return version
	}

	return ""
}
