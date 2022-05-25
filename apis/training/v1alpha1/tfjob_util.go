// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	commonv1 "github.com/alibaba/kubedl/pkg/job_controller/api/v1"
)

// IsTFJobChieforMaster returns true if the type is Master or Chief.
func IsTFJobChieforMaster(typ commonv1.ReplicaType) bool {
	return typ == TFReplicaTypeChief || typ == TFReplicaTypeMaster || typ == commonv1.JobReplicaTypeAIMaster
}

// IsTFJobWorker returns true if the type is Worker.
func IsTFJobWorker(typ commonv1.ReplicaType) bool {
	return typ == TFReplicaTypeWorker
}

// IsTFJobEvaluator returns true if the type is Evaluator.
func IsTFJobEvaluator(typ commonv1.ReplicaType) bool {
	return typ == TFReplicaTypeEval
}
