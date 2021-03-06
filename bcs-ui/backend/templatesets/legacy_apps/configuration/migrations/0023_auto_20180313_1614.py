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
# Generated by Django 1.11.5 on 2018-03-13 08:14
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('configuration', '0022_auto_20180206_1206'),
    ]

    operations = [
        migrations.AlterField(
            model_name='application',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='configmap',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='deplpyment',
            name='name',
            field=models.CharField(max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8sconfigmap',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8sdaemonset',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8sdeployment',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8sjob',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8ssecret',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8sservice',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='k8sstatefulset',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='secret',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='service',
            name='name',
            field=models.CharField(
                default='', max_length=255, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='template',
            name='name',
            field=models.CharField(
                max_length=255, unique=True, verbose_name='名称'),
        ),
    ]
