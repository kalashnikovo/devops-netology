# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [centos:7](https://hub.docker.com/_/centos) как базовый и 
[документацию по установке и запуску Elastcisearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html):

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib`
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
- ссылку на образ в репозитории dockerhub
- ответ `elasticsearch` на запрос пути `/` в json виде

Dockerfile
```bash
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

Но так как URL elastic у меня недоступны, то пришлось собирать образ с копированием этих файлов с локальной машины
```yml
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

В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

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
**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository) 
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html) 
состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее. 

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

Подсказки:
- возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

---
