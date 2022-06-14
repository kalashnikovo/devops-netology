# Домашнее задание к занятию "08.02 Работа с Playbook"

Данный плейбук предназначен для установки `Clickhouse` и `Vector` на хосты, указанные в `inventory` файле.

## group_vars

| Переменная  | Назначение  |
|:---|:---|
| `clickhouse_version` | версия `Clickhouse` |
| `clickhouse_packages` | `RPM` пакеты `Clickhouse`, которые необходимо скачать |
| `vector_url` | URL адрес для скачивания `RPM` пакетов `Vector` |
| `vector_version` | версия `Vector` |
| `vector_home:` | каталог для скачивания `RPM` пакетов `Vector` |

## Inventory файл

Группа "Clickhouse" состоит из 1 хоста `clickhouse-01` c типом подключения `docker` (используется контейнеризированный образ `Centos 8`)

Группа "vector" состоит из 1 хоста `vector-01` c типом подключения `docker` (используется контейнеризированный образ `Centos 8`)


## Playbook

Playbook состоит из 2 `play`.

Play "Install Clickhouse" применяется на группу хостов "Clickhouse" и предназначен для установки и запуска `Clickhouse`

Объявляем `handler` для запуска `clickhouse-server`.
```yaml
 handlers:
    - name: Start clickhouse service
      ansible.builtin.service:
        name: clickhouse-server
        state: restarted
```

| Имя таска | Описание |
|--------------|---------|
| `Clickhouse \| Get clickhouse distrib` | Скачивание `RPM` пакетов. Используется цикл с перменными `clickhouse_packages`. Так как не у всех пакетов есть `noarch` версии, используем перехват ошибки `rescue` |
| `Clickhouse \| Install clickhouse packages` | Установка `RPM` пакетов. В моем контейнере `Centos` не указаны репозитории по умолчанию, поэтому используем параметры `disablerepo: "*"` и `disable_gpg_check: true` - исключаем все репозитории при установке / обновлении и не проверяем GPG подписи пакетов. В `notify` указываем, что данный таск требует запуск handler `Start clickhouse service` |
| `Clickhouse \| Flush handlers` | Форсируем применение handler `Start clickhouse service`. Это необходимо для того, чтобы handler выполнился на текущем этапе, а не по завершению тасок. Если его не запустить сейчас, то сервис не будет запущен и следующий таск завершится с ошибкой |
| `Clickhouse \| Create database` | Создаем в `Clickhouse` БД с названием "logs". Также прописываем условия, при которых таск будет иметь состояние `failed` и `changed` |

Play "Install Vector" применяется на группу хостов "Vector" и предназначен для установки и запуска `Vector`

Объявляем `handler` для запуска `Start Vector service`.
```yaml
  handlers:
    - name: Start Vector service
      ansible.builtin.service:
        name: vector
        state: restarted
```

| Имя таска | Описание |
|--------------|---------|
| `Vector \| Create dir` | Создаем директорию для скачивания `RPM`. Выставляем на нее права 0644. Путь указан в переменной `vector_home` |
| `Vector \| Download packages` | Скачивание `RPM` в созданную на предудыщем шаге директорию. |
| `Vector \| Install packages` | Установка `RPM` пакетов. В моем контейнере `Centos` не указаны репозитории по умолчанию, поэтому используем параметры `disablerepo: "*"` и `disable_gpg_check: true` - исключаем все репозитории при установке / обновлении и не проверяем GPG подписи пакетов. В `notify` указываем, что данный таск требует запуск handler `Start Vector service` |