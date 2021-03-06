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
# Generated by Django 1.11.5 on 2018-10-10 02:12
from __future__ import unicode_literals

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('helm', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='App',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('creator', models.CharField(max_length=32, verbose_name='创建者')),
                ('updator', models.CharField(max_length=32, verbose_name='更新者')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('updated', models.DateTimeField(auto_now=True)),
                ('project_id', models.CharField(max_length=32, verbose_name='Project ID')),
                ('cluster_id', models.CharField(max_length=32, verbose_name='Cluster ID')),
                ('name', models.CharField(max_length=32)),
                ('namespace', models.CharField(max_length=32)),
                ('namespace_id', models.IntegerField(db_index=True, verbose_name='Namespace ID')),
                ('transitioning_result', models.BooleanField(default=True)),
                ('transitioning_message', models.TextField()),
                ('chart', models.ForeignKey(db_constraint=False, on_delete=django.db.models.deletion.CASCADE, to='helm.Chart')),
                ('release', models.ForeignKey(db_constraint=False, on_delete=django.db.models.deletion.CASCADE, to='helm.ChartRelease')),
            ],
            options={
                'db_table': 'bcs_k8s_app',
            },
        ),
        migrations.AlterUniqueTogether(
            name='app',
            unique_together=set([('namespace_id', 'chart')]),
        ),
        migrations.AlterIndexTogether(
            name='app',
            index_together=set([('namespace_id', 'chart')]),
        ),
    ]
