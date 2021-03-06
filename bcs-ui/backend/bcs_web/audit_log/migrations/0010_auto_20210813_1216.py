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
# Generated by Django 3.2.2 on 2021-08-13 04:16

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('activity_log', '0009_auto_20190417_1213'),
    ]

    operations = [
        migrations.AlterField(
            model_name='useractivitylog',
            name='activity_status',
            field=models.CharField(choices=[('completed', '完成'), ('error', '错误'), ('succeed', '成功'), ('failed', '失败')], default='', help_text='操作状态', max_length=32),
        ),
        migrations.AlterField(
            model_name='useractivitylog',
            name='activity_type',
            field=models.CharField(choices=[('add', '创建'), ('modify', '更新'), ('rollback', '回滚'), ('delete', '删除'), ('begin', '开始'), ('end', '结束'), ('start', '启动'), ('pause', '暂停'), ('carryon', '继续'), ('stop', '停止'), ('restart', '重启'), ('retrieve', '查询')], default='', help_text='操作类型', max_length=32),
        ),
        migrations.AlterField(
            model_name='useractivitylog',
            name='resource_type',
            field=models.CharField(blank=True, choices=[('project', '项目'), ('cluster', '集群'), ('node', '节点'), ('namespace', '命名空间'), ('template', '模板集'), ('variable', '变量'), ('instance', '应用'), ('lb', 'LoadBalancer'), ('metric', 'Metric'), ('web_console', 'WebConsole'), ('helm_app', 'Helm'), ('hpa', 'HPA'), ('deployment', 'Deployment'), ('daemonset', 'DaemonSet'), ('statefulset', 'StatefulSet'), ('cronjob', 'CronJob'), ('job', 'Job'), ('pod', 'Pod'), ('ingress', 'Ingress'), ('service', 'Service'), ('endpoints', 'Endpoints'), ('configmap', 'Configmap'), ('secret', 'Secret'), ('persistentvolume', 'PersistentVolume'), ('persistentvolumeclaim', 'PersistentVolumeClaim'), ('storageclass', 'StorageClass'), ('serviceaccount', 'ServiceAccount'), ('crd', '自定义资源定义'), ('customobject', '自定义对象')], help_text='操作对象类型', max_length=32, null=True),
        ),
    ]
