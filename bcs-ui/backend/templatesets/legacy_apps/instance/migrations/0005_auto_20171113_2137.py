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
# Generated by Django 1.11.5 on 2017-11-13 13:37
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('instance', '0004_auto_20171102_2102'),
    ]

    operations = [
        migrations.AddField(
            model_name='versioninstance',
            name='history',
            field=models.TextField(default=[], help_text='以json格式存储', verbose_name='历史变更数据'),
            preserve_default=False,
        ),
        migrations.AddField(
            model_name='versioninstance',
            name='ns_id',
            field=models.IntegerField(default=0, verbose_name='命名空间ID'),
            preserve_default=False,
        ),
        migrations.AddField(
            model_name='versioninstance',
            name='template_id',
            field=models.IntegerField(default=0, help_text='该字段只在db中查看使用', verbose_name='关联的模板 ID'),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name='instanceconfig',
            name='category',
            field=models.CharField(choices=[('application', 'Application'), ('deployment', 'Deplpyment'), ('service', 'Service'), ('configmap', 'ConfigMap'), ('secret', 'Secret')], max_length=32, verbose_name='资源类型'),
        ),
        migrations.AlterField(
            model_name='instanceconfig',
            name='namespace',
            field=models.CharField(max_length=32, verbose_name='命名空间ID'),
        ),
        migrations.AlterField(
            model_name='versioninstance',
            name='namespaces',
            field=models.TextField(help_text='该字段已经废弃', verbose_name='命名空间ID'),
        ),
    ]
