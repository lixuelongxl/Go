# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""


class ReleaseDataProcessor:
    create_fields = [
        "project_id",
        "template_id",
        "template_name",
        "version_id",
        "show_version_id",
        "show_version_name",
        "namespaces",
        "ns_list",
        "is_start",
        "instance_entity",
        "variable_info",
    ]

    update_fields = [
        "project_id",
        "template_id",
        "release_id",
        "namespace_id",
        "variable_info",
        "name",
        "resource_name",
    ]

    def __init__(self, raw_release_data):
        self.raw_release_data = raw_release_data

    def release_data(self, method):
        release_data = {}
        if method == "create":
            data_fields = self.create_fields
        elif method == "update":
            data_fields = self.update_fields
        else:
            data_fields = []

        for k in data_fields:
            release_data[k] = self.raw_release_data.get(k)
        return release_data
