/*
Copyright 2021 The Fluid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

// Runtime for Juice
const (
	JuiceFSRuntime = "juicefs"

	JuiceFSMountType = "JuiceFS"

	JuiceFSNamespace = "juicefs-system"

	JuiceFSChart = JuiceFSRuntime

	JuiceFSFuseImageEnv = "JUICEFS_FUSE_IMAGE_ENV"

	DefaultJuiceFSFuseImage = "juicedata/juicefs-fuse:v1.0.4-4.9.2"

	DefaultJuiceFSRuntimeImage = "juicedata/juicefs-fuse:v1.0.4-4.9.2"

	DefaultJuiceFSMigrateImage = "juicedata/juicefs-fuse:nightly"

	NightlyTag = "nightly"

	JuiceFSCeMountPath = "/bin/mount.juicefs"
	JuiceFSMountPath   = "/sbin/mount.juicefs"
	JuiceCeCliPath     = "/usr/local/bin/juicefs"
	JuiceCliPath       = "/usr/bin/juicefs"

	JuiceFSFuseContainer = "juicefs-fuse"

	JuiceFSWorkerContainer = "juicefs-worker"

	JuiceFSDefaultCacheDir = "/var/jfsCache"
)
