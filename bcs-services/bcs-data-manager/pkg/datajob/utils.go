/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 *  Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *  Licensed under the MIT License (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *  http://opensource.org/licenses/MIT
 *  Unless required by applicable law or agreed to in writing, software distributed under
 *  the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 *  either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package datajob

import (
	"github.com/Tencent/bk-bcs/bcs-services/bcs-data-manager/pkg/common"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-data-manager/pkg/types"
)

func getClusterWorkloadCount(opts *types.JobCommonOpts, clients *types.Clients) int64 {
	workloadList := make([]*types.WorkloadMeta, 0)
	switch opts.ClusterType {
	case types.Kubernetes:
		namespaceList := common.GetK8sNamespaceList(opts.ClusterID, opts.ProjectID, clients.K8sStorageCli)
		for _, namespace := range namespaceList {
			workloads := common.GetK8sWorkloadList(opts.ClusterID, opts.ProjectID, namespace.Name,
				clients.K8sStorageCli)
			workloadList = append(workloadList, workloads...)
		}
	case types.Mesos:
		workloads := common.GetMesosWorkloadList(opts.ClusterID, opts.ProjectID,
			clients.K8sStorageCli)
		workloadList = append(workloadList, workloads...)
	}
	return int64(len(workloadList))
}
