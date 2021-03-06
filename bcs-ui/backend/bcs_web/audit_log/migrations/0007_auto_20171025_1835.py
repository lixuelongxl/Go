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
# Generated by Django 1.11.5 on 2017-10-25 10:35
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('activity_log', '0006_auto_20171013_1456'),
    ]

    operations = [
        migrations.AlterField(
            model_name='useractivitylog',
            name='resource_type',
            field=models.CharField(blank=True, choices=[('项目', 'project'), ('角色', 'role'), ('权限', 'permission'), ('用户', 'user'), ('集群', 'cluster'), ('节点', 'node'), ('模板', 'template'), ('应用', 'application')], help_text='操作对象类型', max_length=32, null=True),
        ),
    ]
