# Generated by Django 3.2.2 on 2022-02-09 03:38

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('activity_log', '0010_auto_20210813_1216'),
    ]

    operations = [
        migrations.AlterField(
            model_name='useractivitylog',
            name='activity_type',
            field=models.CharField(choices=[('add', '创建'), ('modify', '更新'), ('rollback', '回滚'), ('delete', '删除'), ('begin', '开始'), ('end', '结束'), ('start', '启动'), ('pause', '暂停'), ('carryon', '继续'), ('stop', '停止'), ('restart', '重启'), ('retrieve', '查询'), ('reschedule', '重新调度')], default='', help_text='操作类型', max_length=32),
        ),
    ]
