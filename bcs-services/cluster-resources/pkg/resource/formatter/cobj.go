/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package formatter

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util/mapx"
)

// FormatCRD ...
func FormatCRD(manifest map[string]interface{}) map[string]interface{} {
	ret := CommonFormatRes(manifest)
	ret["name"] = mapx.Get(manifest, "metadata.name", "--")
	ret["scope"] = mapx.Get(manifest, "spec.scope", "--")
	ret["kind"] = mapx.Get(manifest, "spec.names.kind", "--")
	ret["apiVersion"] = parseCObjAPIVersion(manifest)
	return ret
}

// FormatCObj ...
func FormatCObj(manifest map[string]interface{}) map[string]interface{} {
	return CommonFormatRes(manifest)
}

// 根据 CRD 配置解析 cobj ApiVersion
// ref: https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/#specify-multiple-versions
func parseCObjAPIVersion(manifest map[string]interface{}) string {
	group, _ := mapx.GetItems(manifest, "spec.group")
	versions, _ := mapx.GetItems(manifest, "spec.versions")

	if versions != nil && len(versions.([]interface{})) != 0 {
		for _, ver := range versions.([]interface{}) {
			v, _ := ver.(map[string]interface{})
			if v["served"].(bool) {
				return fmt.Sprintf("%s/%s", group, v["name"])
			}
		}
		return fmt.Sprintf("%s/%s", group, versions.([]interface{})[0].(map[string]interface{})["name"])
	}

	version, _ := mapx.GetItems(manifest, "spec.version")
	if version != nil && version != "" {
		return fmt.Sprintf("%s/%s", group, version)
	}
	return fmt.Sprintf("%s/v1alpha1", group)
}
