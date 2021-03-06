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
# Generated by Django 1.11.5 on 2017-12-22 03:45
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('configuration', '0012_add_id_to_ports'),
    ]

    operations = [
        migrations.AlterField(
            model_name='application',
            name='name',
            field=models.CharField(default='', max_length=64, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='configmap',
            name='name',
            field=models.CharField(default='', max_length=64, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='deplpyment',
            name='name',
            field=models.CharField(max_length=64, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='secret',
            name='name',
            field=models.CharField(default='', max_length=64, verbose_name='名称'),
        ),
        migrations.AlterField(
            model_name='service',
            name='name',
            field=models.CharField(default='', max_length=64, verbose_name='名称'),
        ),
    ]
