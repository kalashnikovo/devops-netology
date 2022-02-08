# Домашнее задание к занятию "4.2. Использование Python для решения типовых DevOps задач"

## Обязательная задача 1

Есть скрипт:
```python
#!/usr/bin/env python3
a = 1
b = '2'
c = a + b
```

### Вопросы:
| Вопрос  | Ответ                                                                                             |
| ------------- |---------------------------------------------------------------------------------------------------|
| Какое значение будет присвоено переменной `c`?  | `unsupported operand type(s) for +: 'int' and 'str'`, так как переменные `a` и `b` имеют разные типы |
| Как получить для переменной `c` значение 12?  | `c = str(a) + b`                                                                                                 |
| Как получить для переменной `c` значение 3?  | `c = a + int(b)`                                                                                             |

## Обязательная задача 2
Мы устроились на работу в компанию, где раньше уже был DevOps Engineer. Он написал скрипт, позволяющий узнать, какие файлы модифицированы в репозитории, относительно локальных изменений. Этим скриптом недовольно начальство, потому что в его выводе есть не все изменённые файлы, а также непонятен полный путь к директории, где они находятся. Как можно доработать скрипт ниже, чтобы он исполнял требования вашего руководителя?

```python
#!/usr/bin/env python3

import os

bash_command = ["cd ~/netology/sysadm-homeworks", "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
is_change = False
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(prepare_result)
        break
```

### Ваш скрипт:
```python
#!/usr/bin/env python3

import os

bash_command = ["cd ./", "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
# Неиспользуемая переменная
#is_change = False
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(prepare_result)
# break прерывает выполнение скрипта после первого успешного нахождения, поэтому он не нужен
#        break

```

### Вывод скрипта при запуске при тестировании:
```
04-script-01-bash/README.md
```

## Обязательная задача 3
1. Доработать скрипт выше так, чтобы он мог проверять не только локальный репозиторий в текущей директории, а также умел воспринимать путь к репозиторию, который мы передаём как входной параметр. Мы точно знаем, что начальство коварное и будет проверять работу этого скрипта в директориях, которые не являются локальными репозиториями.

### Ваш скрипт:
```python
#!/usr/bin/env python3

import os
import sys

path = "./"
# Проверка аргумента пути к репозиторию
# Если аргументы введены, то проверяем директорию
if len(sys.argv) >= 2:
    path = sys.argv[1]
    # Если директории не существует - завершаем выполнение скрипта с ошибкой
    if not os.path.isdir(path):
        sys.exit("Directory doesn't exist: " + path)

bash_command = ["cd " + path, "git status 2>&1"]
git_command = ["git rev-parse --show-toplevel"]

# Если запрос git status завершился с ошибкой (не является репозиторием), то прерываем работу скрипта
result_os = os.popen(' && '.join(bash_command)).read()
if result_os.find('not a git') != -1:
    sys.exit("not a git repository: " + path)

# Получаем топ-левел каталог репозитория
git_top_level = (os.popen(' && '.join(git_command)).read()).replace('\n', '/')

for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(git_top_level + prepare_result)

```

### Вывод скрипта при запуске при тестировании:
```
Скрипт с параметром (репозиторий есть)

PS D:\Git\Homework2\devops-netology> python.exe .\test.py "D:\Git\Homework2\devops-netology\"
D:/Git/Homework2/devops-netology/04-script-02-py/README.md
D:/Git/Homework2/devops-netology/README.md

Скрипт с параметром (репозитория нет)

PS D:\Git\Homework2\devops-netology> python.exe .\test.py "D:\Git\Homework2\"
not a git repository: D:\Git\Homework2\

Скрипт без параметра (берется текущий каталог)

PS D:\Git\Homework2\devops-netology> python.exe .\test.py
D:/Git/Homework2/devops-netology/04-script-02-py/README.md
D:/Git/Homework2/devops-netology/README.md
```

## Обязательная задача 4
1. Наша команда разрабатывает несколько веб-сервисов, доступных по http. Мы точно знаем, что на их стенде нет никакой балансировки, кластеризации, за DNS прячется конкретный IP сервера, где установлен сервис. Проблема в том, что отдел, занимающийся нашей инфраструктурой очень часто меняет нам сервера, поэтому IP меняются примерно раз в неделю, при этом сервисы сохраняют за собой DNS имена. Это бы совсем никого не беспокоило, если бы несколько раз сервера не уезжали в такой сегмент сети нашей компании, который недоступен для разработчиков. Мы хотим написать скрипт, который опрашивает веб-сервисы, получает их IP, выводит информацию в стандартный вывод в виде: <URL сервиса> - <его IP>. Также, должна быть реализована возможность проверки текущего IP сервиса c его IP из предыдущей проверки. Если проверка будет провалена - оповестить об этом в стандартный вывод сообщением: [ERROR] <URL сервиса> IP mismatch: <старый IP> <Новый IP>. Будем считать, что наша разработка реализовала сервисы: `drive.google.com`, `mail.google.com`, `google.com`.

### Ваш скрипт:
```python
import socket
import time
# Нулевые значения
service_addr = {
    'drive.google.com': '0',
    'mail.google.com': '0',
    'google.com': '0'
}

# Получаем текущие на момент запуска скрипта значения (необходимо для последующего сравнения).
for item in service_addr:
    initial_addr = socket.gethostbyname(item)
    service_addr[item] = initial_addr


while True:
    # Перебираем все ключи в словаре
    for item in service_addr:
        old_addr = service_addr[item]
        new_addr = socket.gethostbyname(item)
        # Если старое и новое не совпадают - адрес изменился. Перезаписываем значение в словаре и выводим ошибку
        if new_addr != old_addr:
            service_addr[item] = new_addr
            print("[ERROR] "+item+" IP mismatch: old IP "+old_addr+", new IP "+new_addr)
        print(item + " - " + service_addr[item])
    print("######################################")
    time.sleep(10)

```

### Вывод скрипта при запуске при тестировании:
```
PS D:\Git\Homework2\devops-netology> python.exe .\test.py
drive.google.com - 74.125.205.194
mail.google.com - 74.125.131.83
google.com - 142.251.1.138
######################################
drive.google.com - 74.125.205.194
mail.google.com - 74.125.131.83
google.com - 142.251.1.138
######################################
drive.google.com - 74.125.205.194
mail.google.com - 74.125.131.83
google.com - 142.251.1.138
######################################
drive.google.com - 74.125.205.194
mail.google.com - 74.125.131.83
[ERROR] google.com IP mismatch: old IP 142.251.1.138, new IP 142.251.1.113
google.com - 142.251.1.113
######################################
drive.google.com - 74.125.205.194
mail.google.com - 74.125.131.83
google.com - 142.251.1.113
######################################
drive.google.com - 74.125.205.194
mail.google.com - 74.125.131.83
google.com - 142.251.1.113
######################################
drive.google.com - 74.125.205.194
[ERROR] mail.google.com IP mismatch: old IP 74.125.131.83, new IP 74.125.131.18
mail.google.com - 74.125.131.18
[ERROR] google.com IP mismatch: old IP 142.251.1.113, new IP 142.251.1.139
google.com - 142.251.1.139
######################################
```