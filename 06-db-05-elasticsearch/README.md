# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

Dockerfile
```docker
FROM centos:7
RUN cd /opt && \
    groupadd elasticsearch && \
    useradd -c "elasticsearch" -g elasticsearch elasticsearch &&\
    yum update -y && yum -y install wget perl-Digest-SHA && \
    wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.0-linux-x86_64.tar.gz && \
    wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.0-linux-x86_64.tar.gz.sha512 && \
    shasum -a 512 -c elasticsearch-8.2.0-linux-x86_64.tar.gz.sha512 && \
    tar -xzf elasticsearch-8.2.0-linux-x86_64.tar.gz && \
	rm elasticsearch-8.2.0-linux-x86_64.tar.gz elasticsearch-8.2.0-linux-x86_64.tar.gz.sha512 && \ 
	mkdir /var/lib/data && chmod -R 777 /var/lib/data && \
	chown -R elasticsearch:elasticsearch /opt/elasticsearch-8.2.0 && \
	yum -y remove wget perl-Digest-SHA && \
	yum clean all
USER elasticsearch
WORKDIR /opt/elasticsearch-8.2.0/
COPY elasticsearch.yml  config/
EXPOSE 9200 9300
ENTRYPOINT ["bin/elasticsearch"]
```

Но так как URL `elastic` у меня недоступны, то пришлось собирать образ с копированием этих файлов с локальной машины
```docker
FROM centos:7
COPY elasticsearch-8.2.0-linux-x86_64.tar.gz  /opt
COPY elasticsearch-8.2.0-linux-x86_64.tar.gz.sha512  /opt	
RUN cd /opt && \
    groupadd elasticsearch && \
    useradd -c "elasticsearch" -g elasticsearch elasticsearch &&\
    yum update -y && yum -y install perl-Digest-SHA && \
    shasum -a 512 -c elasticsearch-8.2.0-linux-x86_64.tar.gz.sha512 && \
    tar -xzf elasticsearch-8.2.0-linux-x86_64.tar.gz && \
    rm elasticsearch-8.2.0-linux-x86_64.tar.gz elasticsearch-8.2.0-linux-x86_64.tar.gz.sha512 && \ 
    mkdir /var/lib/data && chmod -R 777 /var/lib/data && \
    chown -R elasticsearch:elasticsearch /opt/elasticsearch-8.2.0 && \
    yum clean all
USER elasticsearch
WORKDIR /opt/elasticsearch-8.2.0/
COPY elasticsearch.yml  config/
EXPOSE 9200 9300
ENTRYPOINT ["bin/elasticsearch"]
```

Файл elasticsearch.yml
```yml
node:
  name: netology_test
path:
  data: /var/lib/data
xpack.ml.enabled: false
```
[Ссылка на образ](https://hub.docker.com/r/nkiselyov/elasticsearch)

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl --insecure -u elastic https://localhost:9200
Enter host password for user 'elastic':
{
  "name" : "netology_test",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "JjKYfo_yRCadwob52-kEwQ",
  "version" : {
    "number" : "8.2.0",
    "build_flavor" : "default",
    "build_type" : "tar",
    "build_hash" : "b174af62e8dd9f4ac4d25875e9381ffe2b9282c5",
    "build_date" : "2022-04-20T10:35:10.180408517Z",
    "build_snapshot" : false,
    "lucene_version" : "9.1.0",
    "minimum_wire_compatibility_version" : "7.17.0",
    "minimum_index_compatibility_version" : "7.0.0"
  },
  "tagline" : "You Know, for Search"
}
```

## Задача 2

Ознакомтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html) 
и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1| 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

```bash
curl -X PUT --insecure -u elastic "https://localhost:9200/ind-1?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 1,  
      "number_of_replicas": 0 
    }
  }
}
'

curl -X PUT --insecure -u elastic "https://localhost:9200/ind-2?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 2,  
      "number_of_replicas": 1 
    }
  }
}
'

curl -X PUT --insecure -u elastic "https://localhost:9200/ind-3?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 4,  
      "number_of_replicas": 2 
    }
  }
}
'
```

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.
```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X GET --insecure -u elastic:elastic "https://localhost:9200/_cat/indices?v=true"
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   ind-1 z8a9XNImQdqAysDUI2QTcQ   1   0          0            0       225b           225b
yellow open   ind-3 8MksEhg3Tgi-2qJ--tBSrA   4   2          0            0       450b           450b
yellow open   ind-2 GP0S7lRLQ-KTkfZe2rfVVg   2   1          0            0       450b           450b
```
Получите состояние кластера `elasticsearch`, используя API.
```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X GET --insecure -u elastic:elastic "https://localhost:9200/_cluster/health?pretty"
{
  "cluster_name" : "elasticsearch",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 7,
  "active_shards" : 7,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 41.17647058823529
}
```
Полагаю, индексы и кластер находятся в yellow, так как при создании индексов мы указали количество реплик больше 1. В кластере у нас 1 нода, поэтому реплицировать индексы некуда.

Удалите все индексы.
```bash
curl -X DELETE --insecure -u elastic:elastic "https://localhost:9200/ind-1?pretty"
curl -X DELETE --insecure -u elastic:elastic "https://localhost:9200/ind-2?pretty"
curl -X DELETE --insecure -u elastic:elastic "https://localhost:9200/ind-3?pretty"
```

## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создали директорию `/opt/elasticsearch-8.2.0/snapshots`. В конфигурационный файл `elasticsearch.yml` внесли параметр `path.repo: ["/opt/elasticsearch-8.2.0/snapshots"]`

Выполнили запрос на регистрацию `snapshot` репозитория

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X PUT --insecure -u elastic:elastic "https://localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
",
  "settings": {
    "location": "/opt/elasticsearch-8.2.0/snapshots"
  }
}
'
> {
>   "type": "fs",
>   "settings": {
>     "location": "/opt/elasticsearch-8.2.0/snapshots"
>   }
> }
> '

{
  "acknowledged" : true
}
```

Создали индекс `test` с 0 реплик и 1 шардом

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X GET --insecure -u elastic:elastic "https://localhost:9200/_cat/indices?v=true"
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test  Vljz_EYpR5e1E7_ogrOu_w   1   0          0            0       225b           225b
```
Сделали snapshot кластера

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X PUT --insecure -u elastic:elastic "https://localhost:9200/_snapshot/netology_backup/%3Cmy_snapshot_%7Bnow%2Fd%7D%3E?pretty"
{
  "accepted" : true
}
```
Список файлов в директории со `snapshot`ами

```bash
[elasticsearch@fe18b5f27b80 snapshots]$ ll
total 36
-rw-r--r-- 1 elasticsearch elasticsearch  1107 May 23 14:45 index-0
-rw-r--r-- 1 elasticsearch elasticsearch     8 May 23 14:45 index.latest
drwxr-xr-x 5 elasticsearch elasticsearch  4096 May 23 14:45 indices
-rw-r--r-- 1 elasticsearch elasticsearch 16595 May 23 14:45 meta-y1W7p4kFTO2S7PM8rZ75AQ.dat
-rw-r--r-- 1 elasticsearch elasticsearch   401 May 23 14:45 snap-y1W7p4kFTO2S7PM8rZ75AQ.dat
```

Удалили индекс `test` и создали индекс `test-2`.

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X GET --insecure -u elastic:elastic "https://localhost:9200/_cat/indices?v=true"
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test2 6mmyeUKPSiKsXyPrDVrdGA   1   0          0            0       225b           225b
```

Получили список доступных `snapshot`ов

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X GET --insecure -u elastic:elastic "https://localhost:9200/_snapshot/netology_backup/*?verbose=false&pretty"
{
  "snapshots" : [
    {
      "snapshot" : "my_snapshot_2022.05.23",
      "uuid" : "y1W7p4kFTO2S7PM8rZ75AQ",
      "repository" : "netology_backup",
      "indices" : [
        ".geoip_databases",
        ".security-7",
        "test"
      ],
      "data_streams" : [ ],
      "state" : "SUCCESS"
    }
  ],
  "total" : 1,
  "remaining" : 0
}
```

Восстановили состояние кластера `elasticsearch` из `snapshot`, созданного ранее.

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X POST --insecure -u elastic:elastic "https://localhost:9200/_snapshot/netology_backup/my_snapshot_2022.05.23/_restore?pretty" -H 'Content-Type: application/json' -d'
> {
>   "indices": "*",
>   "include_global_state": true
> }
> '

{
  "accepted" : true
}
```
Итоговый список индексов

```bash
vagrant@ubuntu-20:~/docker/elasticsearch$ curl -X GET --insecure -u elastic:elastic "https://localhost:9200/_cat/indices?v=true"
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test2 6mmyeUKPSiKsXyPrDVrdGA   1   0          0            0       225b           225b
green  open   test  d1DglgrzRHaUcZSzggitjw   1   0          0            0       225b           225b
```
---
