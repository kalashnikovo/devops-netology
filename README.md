
# Домашнее задание к занятию "2.1. Системы контроля версий."

В каталоге terraform будут проигнорированы файлы:
1) все файлы и каталоги в каталоге  .terraform
2) файлы, содержащие .tfstate
3) файлы с именем crash.log
4) файлы с расширением .tfvars
5) файлы override.tf, override.tf.json, и любые файлы, которые содержат в названии _override.tf или _override.tf.json
6) файлы terraform.rc и .terraformrc


# Домашнее задание к занятию "3.1. Работа в терминале, лекция 1"
1) Ознакомьтесь с графическим интерфейсом VirtualBox, посмотрите как выглядит виртуальная машина, которую создал для вас Vagrant, какие аппаратные ресурсы ей выделены. Какие ресурсы выделены по-умолчанию?

По умолчанию выдано 2 CPU и 1024 MB RAM

2) Ознакомьтесь с возможностями конфигурации VirtualBox через Vagrantfile: документация. Как добавить оперативной памяти или ресурсов процессора виртуальной машине?

 
Необходимо изменить конфиг файл vagrantfile и перезапустить ВМ (vagrant reload). Например,

````
 Vagrant.configure("2") do |config|
 	config.vm.box = "bento/ubuntu-20.04"
	config.vm.provider "virtualbox" do |v|
		v.memory = 4128
		v.cpus = 4
	end
 end
````
3) Какой переменной можно задать длину журнала history, и на какой строчке manual это описывается?

HISTSIZE - количество команд для запоминания, примерно на 862 строке
HISTFILESIZE - максимальное количество строк в файле истории, примерно 846 строка

4) Что делает директива ignoreboth в bash?

Сокращение для ignorespace и ignoredups. В списке истории не будут сохраняться строки, начинающиеся на пробел, и строки, совпадающие с предыдущей

5) В каких сценариях использования применимы скобки {} и на какой строчке man bash это описано?

Когда список должен выполняться в текущем контексте интерпретатора - строка 257
6) С учётом ответа на предыдущий вопрос, как создать однократным вызовом touch 100000 файлов? Получится ли аналогичным образом создать 300000? Если нет, то почему?

``touch {1..100000}``

Создать 300000 не удалось

``bash: /usr/bin/touch: Argument list too long``

7) В man bash поищите по /\[\[. Что делает конструкция [[ -d /tmp ]]


``[[ -d /tmp ]]`` возвращает 0 или 1, в зависимости от значения выражения внутри. Каталог /tmp существует, поэтому возвращаться будет статус 0 (если я правильно разобрался)

8) Основываясь на знаниях о просмотре текущих (например, PATH) и установке новых переменных; командах, которые мы рассматривали, добейтесь в выводе type -a bash в виртуальной машине наличия первым пунктом в списке:

````
# Создаем каталог
mkdir /tmp/new_path_directory 

# Копируем в него файл bash
cp /bin/bash /tmp/new_path_directory/

# Добавляем каталог в переменную $PATH (в начало) 
PATH=/tmp/new_path_directory/bash:$PATH

# Выводим переменную, чтобы убедиться в корректности
echo $PATH

# Выполняем type -a
type -a bash
````

9) Чем отличается планирование команд с помощью batch и at?

At используется для планирования разовой задачи, которая должна выполниться в определенное время

Batch используется для выполнения разовой задачи, когда позволит средняя нагрузка (в man указано, когда средняя нагрузка упадет ниже 1.5)


# Домашнее задание к занятию «2.4. Инструменты Git»

1) Найдите полный хеш и комментарий коммита, хеш которого начинается на aefea.

Полный хэш ``aefead2207ef7e2aa5dc81a34aedf0cad4c32545``

Комментарий ``Update CHANGELOG.md``

Команда `git show aefea` или ``git show -s --format="%H %B" aefea``

2) Какому тегу соответствует коммит 85024d3

Тег `v0.12.23` \
Команда `git show -s --oneline 85024d3`

3) Сколько родителей у коммита b8d720? Напишите их хеши.

Два родителя `56cd7859e05c36c06b56d013b55a252d0bb7e158` и `9ea88f22fc6269854151c571162c5bcf958bee2b`

Команда `git show --pretty=%P b8d720`

4) Перечислите хеши и комментарии всех коммитов которые были сделаны между тегами v0.12.23 и v0.12.24.

````
33ff1c03b (tag: v0.12.24) v0.12.24
b14b74c49 [Website] vmc provider links
3f235065b Update CHANGELOG.md
6ae64e247 registry: Fix panic when server is unreachable
5c619ca1b website: Remove links to the getting started guide's old location
06275647e Update CHANGELOG.md
d5f9411f5 command: Fix bug when using terraform login on Windows
4b6d06cc5 Update CHANGELOG.md
dd01a3507 Update CHANGELOG.md
225466bc3 Cleanup after v0.12.23 release
````

Команда `git show -s --oneline v0.12.23..v0.12.24`

5) Найдите коммит, в котором была создана функция func providerSource, ее определение в коде выглядит так func providerSource(...) (вместо троеточего перечислены аргументы).

Коммит `8c928e835` \
Команда `git log -S"func providerSource(" --oneline`

6) Найдите все коммиты в которых была изменена функция globalPluginDirs.

````
78b122055 Remove config.go and update things using its aliases
52dbf9483 keep .terraform.d/plugins for discovery
41ab0aef7 Add missing OS_ARCH dir to global plugin paths
66ebff90c move some more plugin search path logic to command
8364383c3 Push plugin discovery down into command package
````

Команды

Ищем файл, где объявляется функция \
`git grep "func globalPluginDirs("`

Ищем коммиты с изменением этой функции в файле \
`git log -s -L :globalPluginDirs:plugins.go --oneline`

7) Кто автор функции synchronizedWriters

`Martin Atkins <mart@degeneration.co.uk>`

Ищем коммит, в котором была создана функция \
`git log -S"func synchronizedWriters(" --oneline`

Просматриваем самый ранний коммит из вывода предыщей команды. Убеждаемся, что в нем была создана эта функция и смотрим автора \
`git show 5ac311e2a`

Можно сразу посмотреть автора командой \
`git show 5ac311e2a -s --pretty=format:%an`


# Домашнее задание к занятию "3.2. Работа в терминале, лекция 2"

1) Какого типа команда cd? Попробуйте объяснить, почему она именно такого типа; опишите ход своих мыслей, если считаете что она могла бы быть другого типа.

cd - это shell builtin команда, то есть команда, которая вызывается напрямую в shell, а не как внешняя исполняемая. Если бы она была внешней,
то запускалась бы в отдельном процессе и меняла бы директорию для этого процесса (текущий каталог shell оставался бы неизменным).

2) Какая альтернатива без pipe команде grep <some_string> <some_file> | wc -l? 

`grep <some_string> <some_file> -c`

3) Какой процесс с PID 1 является родителем для всех процессов в вашей виртуальной машине Ubuntu 20.04?

systemd



```
vagrant@vagrant:/var/log$ ps -p 1
    PID TTY          TIME CMD
      1 ?        00:00:01 systemd

vagrant@vagrant:/var/log$ pstree -p
systemd(1)─┬─VBoxService(780)─┬─{VBoxService}(781)
           │                  ├─{VBoxService}(782)
           │                  ├─{VBoxService}(783)
           │                  ├─{VBoxService}(784)
           │                  ├─{VBoxService}(785)
           │                  ├─{VBoxService}(786)
           │                  ├─{VBoxService}(787)
           │                  └─{VBoxService}(788)
           ├─accounts-daemon(584)─┬─{accounts-daemon}(603)
           │                      └─{accounts-daemon}(633)
           ├─agetty(812)
           ├─atd(800)
           ├─cron(796)
           ├─dbus-daemon(585)
           ├─irqbalance(607)───{irqbalance}(618)
           ├─multipathd(533)─┬─{multipathd}(534)
           │                 ├─{multipathd}(535)
           │                 ├─{multipathd}(536)
           │                 ├─{multipathd}(537)
           │                 ├─{multipathd}(538)
           │                 └─{multipathd}(539)
           ├─networkd-dispat(613)
           ├─polkitd(644)─┬─{polkitd}(645)
           │              └─{polkitd}(647)
           ├─rpcbind(560)
           ├─rsyslogd(617)─┬─{rsyslogd}(626)
           │               ├─{rsyslogd}(627)
           │               └─{rsyslogd}(628)
           ├─sshd(804)───sshd(1078)───sshd(1115)───bash(1116)───pstree(1186)
           ├─systemd(1081)───(sd-pam)(1082)
           ├─systemd-journal(367)
           ├─systemd-logind(620)
           ├─systemd-network(401)
           ├─systemd-resolve(561)
           └─systemd-udevd(397)
```

4) Как будет выглядеть команда, которая перенаправит вывод stderr ls на другую сессию терминала?

``ls /root 2>/dev/pts/X``, \
где `/dev/pts/X` - псевдотерминал другой сессии

5) Получится ли одновременно передать команде файл на stdin и вывести ее stdout в другой файл? Приведите работающий пример.

``cat <test1 >test2``

````
vagrant@vagrant:~$ cat test1
test1 text
vagrant@vagrant:~$ cat test2
vagrant@vagrant:~$ cat <test1 >test2
vagrant@vagrant:~$ cat test2
test1 text
````

6) Получится ли находясь в графическом режиме, вывести данные из PTY в какой-либо из эмуляторов TTY? Сможете ли вы наблюдать выводимые данные?

Да, например, \
``echo "test" > /dev/tty1``

Находясь в tty1, я смог увидеть отправленные данные

7) Выполните команду bash 5>&1. К чему она приведет? Что будет, если вы выполните echo netology > /proc/$$/fd/5? Почему так происходит?

``bash 5>&1`` - создаем новый дескриптор 5 и перенаправляем его в STDOUT \
``echo netology > /proc/$$/fd/5`` - перенаправляем результат команды в дескриптор 5

8) Получится ли в качестве входного потока для pipe использовать только stderr команды, не потеряв при этом отображение stdout на pty? Напоминаем: по умолчанию через pipe передается только stdout команды слева от | на stdin команды справа. Это можно сделать, поменяв стандартные потоки местами через промежуточный новый дескриптор, который вы научились создавать в предыдущем вопросе.

Получится. Для этого нужно "поменять местами STDOUT и STDERR". Для этого применяется конструкция `N>&1 1>&2 2>&N` (где N - новый промежуточный дескриптор). \
Например, \
```
vagrant@vagrant:~$ vagrant@vagrant:~$ ls /var
backups  cache  crash  lib  local  lock  log  mail  opt  run  snap  spool  tmp
vagrant@vagrant:~$ ls
2  test0  test1  test2  tty1
vagrant@vagrant:~$ ls /root
ls: cannot open directory '/root': Permission denied
vagrant@vagrant:~$ (ls /var && ls && ls /root) 3>&1 1>&2 2>&3 | wc -l
backups  cache  crash  lib  local  lock  log  mail  opt  run  snap  spool  tmp
2  test0  test1  test2  tty1
1
```
В результате получили, что в pipe передается stdout с дескриптором 2 (stderr)

9) Что выведет команда cat /proc/$$/environ? Как еще можно получить аналогичный по содержанию вывод?

Выводится список переменных окружения для процесса, под которым выполняется текущая оболочка bash

Аналогичный вывод можно получить с помощью команды `printenv`

10) Используя man, опишите что доступно по адресам /proc/<PID>/cmdline, /proc/<PID>/exe

`/proc/[pid]/cmdline` - файл только на чтение, который содержит строку запуска процессов, кроме зомби-процессов (строка 236) \
`/proc/[pid]/exe` -  сожержит полное имя выполняемого файла для процесса (строка 279)

11) Узнайте, какую наиболее старшую версию набора инструкций SSE поддерживает ваш процессор с помощью /proc/cpuinfo

sse 4_2
```
vagrant@vagrant:~$ grep sse /proc/cpuinfo
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni pclmulqdq ssse3 cx16 pcid
sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx rdrand hypervisor lahf_lm abm 3dnowprefetch invpcid_single pti fsgsbase avx2 invpcid rdseed clflushopt md_clear flush_l1d
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni pclmulqdq ssse3 cx16 pcid
sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx rdrand hypervisor lahf_lm abm 3dnowprefetch invpcid_single pti fsgsbase avx2 invpcid rdseed clflushopt md_clear flush_l1d
```

12) При открытии нового окна терминала и vagrant ssh создается новая сессия и выделяется pty. Это можно подтвердить командой tty, которая упоминалась в лекции 3.2

По умолчанию при запуске команды через SSH не выделяется TTY. Если же не указывать команды, то TTY будет выдаваться, так как предполагается, что будет запущен сеанс оболочки.
Полагаю, что изменить поведение можно через `ssh localhost` с последующей авторизацией и выполнением `'tty'`. Либо через `ssh -t localhost 'tty'`

13) Бывает, что есть необходимость переместить запущенный процесс из одной сессии в другую. Попробуйте сделать это, воспользовавшись reptyr. Например, так можно перенести в screen процесс, который вы запустили по ошибке в обычной SSH-сессии.

Получилось после правки `/proc/sys/kernel/yama/ptrace_scope`

14) Узнайте что делает команда tee и почему в отличие от sudo echo команда с sudo tee будет работать.

`tee` - получает значения из stdin и записывает их в stdout и файл.
Так как `tee` запускается отдельным процессом из-под sudo, то получая в stdin через pipe данные от `echo` - у нее есть права записать в файл.

# Домашнее задание к занятию "3.3. Операционные системы, лекция 1"

1) Какой системный вызов делает команда cd?

`chdir("/tmp")                           = 0`

2) Используя strace выясните, где находится база данных file, на основании которой она делает свои догадки.

Команда `file` пыталась обратиться к файлам, но они не существуют
````
/home/vagrant/.magic.mgc
/home/vagrant/.magic
/etc/magic.mgc
````
Далее было обращение к
````
/etc/magic
/usr/share/misc/magic.mgc
````

3) Предположим, приложение пишет лог в текстовый файл. Этот файл оказался удален (deleted в lsof), однако возможности сигналом сказать приложению переоткрыть файлы или просто перезапустить приложение – нет. Так как приложение продолжает писать в удаленный файл, место на диске постепенно заканчивается. Основываясь на знаниях о перенаправлении потоков предложите способ обнуления открытого удаленного файла (чтобы освободить место на файловой системе).

Сделать Truncate \
``echo ""| sudo tee /proc/PID/fd/DESCRIPTOR``, \
где ``PID`` - это PID процесса, который записывает в удаленный файл, а ``DESRIPTOR`` - дескриптор, удаленного файла.

Например,
````
exec 5> output.file
ping localhost >&5
rm output.file

vagrant@vagrant:~$ sudo lsof | grep deleted

ping      1580                       vagrant    1w      REG              253,0   865859     131084 /home/vagrant/output.file (deleted)
ping      1580                       vagrant    5w      REG              253,0   865859     131084 /home/vagrant/output.file (deleted)

echo ""| tee /proc/1580/fd/5

````

4) Занимают ли зомби-процессы какие-то ресурсы в ОС (CPU, RAM, IO)?

Зомби-процессы не занимают какие-либо системные ресурсы, но сохраняют свой ID процесса (есть риск исчерпания доступных идентификаторов)

5) На какие файлы вы увидели вызовы группы open за первую секунду работы утилиты?
````
vagrant@vagrant:~$ sudo opensnoop-bpfcc
PID    COMM               FD ERR PATH
785    vminfo              4   0 /var/run/utmp
577    dbus-daemon        -1   2 /usr/local/share/dbus-1/system-services
577    dbus-daemon        18   0 /usr/share/dbus-1/system-services
577    dbus-daemon        -1   2 /lib/dbus-1/system-services
577    dbus-daemon        18   0 /var/lib/snapd/dbus-1/system-services/
````
6) Какой системный вызов использует uname -a? Приведите цитату из man по этому системному вызову, где описывается альтернативное местоположение в /proc, где можно узнать версию ядра и релиз ОС.


`Part of the utsname information is also accessible via /proc/sys/kernel/{ostype, hostname, osrelease, version, domainname}.`

7) Чем отличается последовательность команд через ; и через && в bash?

`;` - выполелнение команд последовательно \
`&&` - команда после && выполняется только если команда до && завершилась успешно (статус выхода 0)

`test -d /tmp/some_dir && echo Hi` - так как каталога /tmp/some_dir не существует, то статус выхода не равен 0 и echo Hi не будет выполняться

Есть ли смысл использовать в bash &&, если применить set -e?

`set -e` - останавливает выполнение скрипта при ошибке.
Я думаю, в скриптах имеет смысл применять `set -e` с `&&`, так как она прекращает действие скрипта (не игнорирует ошибку) при ошибке в команде после `&&`
Например, 

````
echo Hi && test -d /tmp/some_dir; echo Bye
````

Без  `set -e` скрипт выполнит `echo Bye`, а без этой команды - нет

8) Из каких опций состоит режим bash set -euxo pipefail и почему его хорошо было бы использовать в сценариях?
````
-e  Exit immediately if a command exits with a non-zero status.
-u  Treat unset variables as an error when substituting.
-x  Print commands and their arguments as they are executed.
-o pipefail     the return value of a pipeline is the status of
                           the last command to exit with a non-zero status,
                           or zero if no command exited with a non-zero status
````
Данный режим обеспечит прекращение выполнения скрипта в случае ошибок и выведет необходимую для траблшутинга информацию (по сути логирование выполнения)

9) Используя -o stat для ps, определите, какой наиболее часто встречающийся статус у процессов в системе. В man ps ознакомьтесь (/PROCESS STATE CODES) что значат дополнительные к основной заглавной буквы статуса процессов. Его можно не учитывать при расчете (считать S, Ss или Ssl равнозначными).
````shell
vagrant@vagrant:~$ ps -d -o stat | sort | uniq -c
7 I
40 I<
1 R+
25 S
2 S+
1 Sl
2 SN
````
```
               <   high-priority (not nice to other users)
               N    low-priority (nice to other users)
               L    has pages locked into memory (for real-time and
                    custom IO)
               s    is a session leader
               l    is multi-threaded (using CLONE_THREAD, like NPTL
                    pthreads do)
               +    is in the foreground process group

```

# Домашнее задание к занятию "3.4. Операционные системы, лекция 2"

1) Используя знания из лекции по systemd, создайте самостоятельно простой unit-файл для node_exporter

unit-файл (по аналогии с https://github.com/prometheus/node_exporter/tree/master/examples/systemd)
````shell
vagrant@vagrant:~$ cat /etc/systemd/system/node_exporter.service
[Unit]
Description=Node Exporter

[Service]
EnvironmentFile=/etc/default/node_exporter
ExecStart=/usr/bin/node_exporter $OPTIONS

[Install]
WantedBy=multi-user.target

vagrant@vagrant:~$ cat /etc/default/node_exporter
OPTIONS="--collector.textfile.directory /var/lib/node_exporter/textfile_collector"
````
Процесс корректно стартует, завершается, перезапускается (в том числе после ребута)

```shell
vagrant@vagrant:~$ sudo systemctl status node_exporter
● node_exporter.service - Node Exporter
     Loaded: loaded (/etc/systemd/system/node_exporter.service; enabled; vendor preset: enabled)
     Active: active (running) since Sun 2022-01-16 12:54:55 UTC; 4s ago
   Main PID: 900 (node_exporter)
      Tasks: 4 (limit: 1071)
     Memory: 2.3M
     CGroup: /system.slice/node_exporter.service
             └─900 /usr/bin/node_exporter --collector.textfile.directory /var/lib/node_exporter/textfile_collector

Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=thermal_zone
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=time
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=timex
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=udp_queues
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=uname
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=vmstat
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=xfs
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:115 level=info collector=zfs
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=node_exporter.go:199 level=info msg="Listening on" address=:9100
Jan 16 12:54:55 vagrant node_exporter[900]: ts=2022-01-16T12:54:55.026Z caller=tls_config.go:195 level=info msg="TLS is disabled." http2=false

vagrant@vagrant:~$ ps -e |grep node_exporter
    900 ?        00:00:00 node_exporter
    
vagrant@vagrant:~$ sudo cat /proc/900/environ
LANG=en_US.UTF-8LANGUAGE=en_US:PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/binINVOCATION_ID=468effef990c42ab846feabdd52e5734JOURNAL_STREAM=9:25903OPTIONS=--collector.textfile.directory /var/lib/node_exporter/textfile_collector

```

2) Ознакомьтесь с опциями node_exporter и выводом /metrics по-умолчанию. Приведите несколько опций, которые вы бы выбрали для базового мониторинга хоста по CPU, памяти, диску и сети

Для CPU (для каждого из возможных ядер)
```
node_cpu_seconds_total{cpu="0",mode="idle"}
node_cpu_seconds_total{cpu="0",mode="system"}
node_cpu_seconds_total{cpu="0",mode="user"}
process_cpu_seconds_total
```
Для ОЗУ
```
node_memory_MemAvailable_bytes
node_memory_MemFree_bytes
node_memory_Buffers_bytes
node_memory_Cached_bytes
```
По дискам (выбрать необходимые диски)
```
node_disk_io_time_seconds_total{device="sda"}
node_disk_read_time_seconds_total{device="sda"}
node_disk_write_time_seconds_total{device="sda"}
node_filesystem_avail_bytes
```
По сети
```
node_network_info
node_network_receive_bytes_total
node_network_receive_errs_total
node_network_transmit_bytes_total
node_network_transmit_errs_total
```

3) Установите в свою виртуальную машину Netdata

```shell
vagrant@vagrant:~$ sudo ss -tulpn | grep :19999
tcp    LISTEN   0        4096              0.0.0.0:19999          0.0.0.0:*      users:(("netdata",pid=2147,fd=4))
```
4) Можно ли по выводу dmesg понять, осознает ли ОС, что загружена не на настоящем оборудовании, а на системе виртуализации?

```shell
vagrant@vagrant:~$ dmesg | grep -i virtual
[    0.000000] DMI: innotek GmbH VirtualBox/VirtualBox, BIOS VirtualBox 12/01/2006
[    0.003991] CPU MTRRs all blank - virtualized system.
[    0.120325] Booting paravirtualized kernel on KVM
[    2.596500] systemd[1]: Detected virtualization oracle.
```

5) Как настроен sysctl fs.nr_open на системе по-умолчанию?

```shell
vagrant@vagrant:~$ /sbin/sysctl -n fs.nr_open
1048576
```

nr_open - означает максимальное число дескрипторов, которые может использовать процесс. Но этого значения не дает достичь другой лимит
```shell
vagrant@vagrant:~$ ulimit -n
1024
```
6) Запустите любой долгоживущий процесс (не ls, который отработает мгновенно, а, например, sleep 1h) в отдельном неймспейсе процессов; покажите, что ваш процесс работает под PID 1 через nsenter.

```shell
vagrant@vagrant:~$ sudo unshare -f --pid --mount-proc sleep 1h
vagrant@vagrant:~$ ps -aux | grep sleep
root        2517  0.0  0.4  11864  4624 pts/1    S+   14:02   0:00 sudo unshare -f --pid --mount-proc sleep 1h
root        2518  0.0  0.0   8080   592 pts/1    S+   14:02   0:00 unshare -f --pid --mount-proc sleep 1h
root        2519  0.0  0.0   8076   580 pts/1    S+   14:02   0:00 sleep 1h

vagrant@vagrant:~$ sudo nsenter --target 2519 --pid --mount
root@vagrant:/# ps aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0   8076   580 pts/1    S+   14:02   0:00 sleep 1h
root           2  0.0  0.4   9836  4144 pts/2    S    14:05   0:00 -bash
root          11  0.0  0.3  11492  3292 pts/2    R+   14:05   0:00 ps aux
```
7) Найдите информацию о том, что такое `:(){ :|:& };:`. 

Судя по информации, что я нашел в интернете - это fork bomb
Её можно переписать в виде
```shell
:(){
 :|:&
};:

или

bomb() { 
 bomb | bomb &
}; bomb
```
Данная функция рекурсивно вызывает себя и передается через пайп другому вызову этой же функции (в фоне).
По сути запускаются 2 фоновых процесса, каждый из которых вызывает еще два фоновых процесса и тд, пока не закончатся ресурсы. \

Вызов dmesg расскажет, какой механизм помог автоматической стабилизации. Как настроен этот механизм по-умолчанию, и как изменить число процессов, которое можно создать в сессии?


```shell
[ 5667.545996] cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-6.scope

vagrant@vagrant:~$ ulimit -u
3571
```
Данная команда выводит количество возможных одновременно запущенных процессов. Изменяя этот параметр, можно защититься от fork bomb (упремся в потолок количества процессов для пользователя, но ресурсы не будут утилизированы полностью)

```shell
ulimit -u 500
```
Данной командой можно уменьшить порог процессов до 500

# Домашнее задание к занятию "3.5. Файловые системы"

1) Узнайте о sparse (разряженных) файлах.

**Разреженные файлы** - это файлы, для которых выделяется пространство на диске только для участков с ненулевыми данными. 
Список всех "дыр" хранится в метаданных ФС и используется при операциях с файлами. 
В результате получается, что разреженный файл занимает меньше места на диске (более эффективное использование дискового пространства)

2) Могут ли файлы, являющиеся жесткой ссылкой на один объект, иметь разные права доступа и владельца? Почему?

Не могут, так жесткие ссылки имеют один и тот же inode (объект, который содержит метаданные файла).

3) Сделайте vagrant destroy на имеющийся инстанс Ubuntu. Замените содержимое Vagrantfile следующим. Данная конфигурация создаст новую виртуальную машину с двумя дополнительными неразмеченными дисками по 2.5 Гб. 
```bash
vagrant@vagrant:~$ lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                    8:0    0   64G  0 disk
├─sda1                 8:1    0  512M  0 part /boot/efi
├─sda2                 8:2    0    1K  0 part
└─sda5                 8:5    0 63.5G  0 part
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm  /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm  [SWAP]
sdb                    8:16   0  2.5G  0 disk
sdc                    8:32   0  2.5G  0 disk
```

4) Используя fdisk, разбейте первый диск на 2 раздела: 2 Гб, оставшееся пространство.
```bash
vagrant@vagrant:~$ sudo fdisk /dev/sdb

Welcome to fdisk (util-linux 2.34).
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.

Device does not contain a recognized partition table.
Created a new DOS disklabel with disk identifier 0x8e9c6d3a.

Command (m for help): n
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p):

Using default response p.
Partition number (1-4, default 1):
First sector (2048-5242879, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-5242879, default 5242879): +2G

Created a new partition 1 of type 'Linux' and of size 2 GiB.

Command (m for help): n
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p):

Using default response p.
Partition number (2-4, default 2):
First sector (4196352-5242879, default 4196352):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (4196352-5242879, default 5242879):

Created a new partition 2 of type 'Linux' and of size 511 MiB. 
Command (m for help): w
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

vagrant@vagrant:~$ lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                    8:0    0   64G  0 disk
├─sda1                 8:1    0  512M  0 part /boot/efi
├─sda2                 8:2    0    1K  0 part
└─sda5                 8:5    0 63.5G  0 part
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm  /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm  [SWAP]
sdb                    8:16   0  2.5G  0 disk
├─sdb1                 8:17   0    2G  0 part
└─sdb2                 8:18   0  511M  0 part
sdc                    8:32   0  2.5G  0 disk
```
5) Используя sfdisk, перенесите данную таблицу разделов на второй диск.

```bash
vagrant@vagrant:~$ sudo sfdisk -d /dev/sdb | sudo sfdisk /dev/sdc
Checking that no-one is using this disk right now ... OK

Disk /dev/sdc: 2.51 GiB, 2684354560 bytes, 5242880 sectors
Disk model: VBOX HARDDISK
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes

>>> Script header accepted.
>>> Script header accepted.
>>> Script header accepted.
>>> Script header accepted.
>>> Created a new DOS disklabel with disk identifier 0xb1cf5424.
/dev/sdc1: Created a new partition 1 of type 'Linux' and of size 2 GiB.
/dev/sdc2: Created a new partition 2 of type 'Linux' and of size 511 MiB.
/dev/sdc3: Done.

New situation:
Disklabel type: dos
Disk identifier: 0xb1cf5424

Device     Boot   Start     End Sectors  Size Id Type
/dev/sdc1          2048 4196351 4194304    2G 83 Linux
/dev/sdc2       4196352 5242879 1046528  511M 83 Linux

The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.
vagrant@vagrant:~$ lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                    8:0    0   64G  0 disk
├─sda1                 8:1    0  512M  0 part /boot/efi
├─sda2                 8:2    0    1K  0 part
└─sda5                 8:5    0 63.5G  0 part
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm  /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm  [SWAP]
sdb                    8:16   0  2.5G  0 disk
├─sdb1                 8:17   0    2G  0 part
└─sdb2                 8:18   0  511M  0 part
sdc                    8:32   0  2.5G  0 disk
├─sdc1                 8:33   0    2G  0 part
└─sdc2                 8:34   0  511M  0 part
```
6) Соберите mdadm RAID1 на паре разделов 2 Гб

```bash
vagrant@vagrant:~$ sudo mdadm --create /dev/md0 -l 1 -n 2 /dev/sdb1 /dev/sdc1
mdadm: Note: this array has metadata at the start and
    may not be suitable as a boot device.  If you plan to
    store '/boot' on this device please ensure that
    your boot-loader understands md/v1.x metadata, or use
    --metadata=0.90
Continue creating array? y
mdadm: Defaulting to version 1.2 metadata
mdadm: array /dev/md0 started.
```
7) Соберите mdadm RAID0 на второй паре маленьких разделов

```bash
vagrant@vagrant:~$ sudo mdadm --create /dev/md1 -l 0 -n 2 /dev/sdb2 /dev/sdc2
mdadm: Defaulting to version 1.2 metadata
mdadm: array /dev/md1 started.
vagrant@vagrant:~$ lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
sda                    8:0    0   64G  0 disk
├─sda1                 8:1    0  512M  0 part  /boot/efi
├─sda2                 8:2    0    1K  0 part
└─sda5                 8:5    0 63.5G  0 part
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
sdb                    8:16   0  2.5G  0 disk
├─sdb1                 8:17   0    2G  0 part
│ └─md0                9:0    0    2G  0 raid1
└─sdb2                 8:18   0  511M  0 part
  └─md1                9:1    0 1018M  0 raid0
sdc                    8:32   0  2.5G  0 disk
├─sdc1                 8:33   0    2G  0 part
│ └─md0                9:0    0    2G  0 raid1
└─sdc2                 8:34   0  511M  0 part
  └─md1                9:1    0 1018M  0 raid0
```
8) Создайте 2 независимых PV на получившихся md-устройствах.

```bash
vagrant@vagrant:~$ sudo pvcreate /dev/md1 /dev/md0
  Physical volume "/dev/md1" successfully created.
  Physical volume "/dev/md0" successfully created.
vagrant@vagrant:~$ sudo pvscan
  PV /dev/sda5   VG vgvagrant       lvm2 [<63.50 GiB / 0    free]
  PV /dev/md0                       lvm2 [<2.00 GiB]
  PV /dev/md1                       lvm2 [1018.00 MiB]
  Total: 3 [<66.49 GiB] / in use: 1 [<63.50 GiB] / in no VG: 2 [2.99 GiB] 
```

9) Создайте общую volume-group на этих двух PV

```bash
vagrant@vagrant:~$ sudo vgcreate VG1 /dev/md0 /dev/md1
  Volume group "VG1" successfully created
vagrant@vagrant:~$ sudo vgscan
  Found volume group "vgvagrant" using metadata type lvm2
  Found volume group "VG1" using metadata type lvm2
```
10) Создайте LV размером 100 Мб, указав его расположение на PV с RAID0.

```bash
vagrant@vagrant:~$ sudo lvcreate -L 100M -n LV1 VG1 /dev/md1
  Logical volume "LV1" created..
```

11) Создайте mkfs.ext4 ФС на получившемся LV.

```bash
vagrant@vagrant:~$ sudo mkfs.ext4 /dev/VG1/LV1
mke2fs 1.45.5 (07-Jan-2020)
Creating filesystem with 25600 4k blocks and 25600 inodes

Allocating group tables: done
Writing inode tables: done
Creating journal (1024 blocks): done
Writing superblocks and filesystem accounting information: done
```

12) Смонтируйте этот раздел в любую директорию, например, /tmp/new
```bash
vagrant@vagrant:~$ mkdir /tmp/new
vagrant@vagrant:~$ sudo mount /dev/VG1/LV1 /tmp/new 
```
13) Поместите туда тестовый файл, например wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz

```bash
vagrant@vagrant:~$ sudo wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz
--2022-01-17 16:47:25--  https://mirror.yandex.ru/ubuntu/ls-lR.gz
Resolving mirror.yandex.ru (mirror.yandex.ru)... 213.180.204.183, 2a02:6b8::183
Connecting to mirror.yandex.ru (mirror.yandex.ru)|213.180.204.183|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 21992768 (21M) [application/octet-stream]
Saving to: ‘/tmp/new/test.gz’

/tmp/new/test.gz              100%[=================================================>]  20.97M  2.75MB/s    in 9.5s

2022-01-17 16:47:35 (2.21 MB/s) - ‘/tmp/new/test.gz’ saved [21992768/21992768]
```
14) Прикрепите вывод lsblk

```bash
vagrant@vagrant:~$ lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
sda                    8:0    0   64G  0 disk
├─sda1                 8:1    0  512M  0 part  /boot/efi
├─sda2                 8:2    0    1K  0 part
└─sda5                 8:5    0 63.5G  0 part
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
sdb                    8:16   0  2.5G  0 disk
├─sdb1                 8:17   0    2G  0 part
│ └─md0                9:0    0    2G  0 raid1
└─sdb2                 8:18   0  511M  0 part
  └─md1                9:1    0 1018M  0 raid0
    └─VG1-LV1        253:2    0  100M  0 lvm   /tmp/new
sdc                    8:32   0  2.5G  0 disk
├─sdc1                 8:33   0    2G  0 part
│ └─md0                9:0    0    2G  0 raid1
└─sdc2                 8:34   0  511M  0 part
  └─md1                9:1    0 1018M  0 raid0
    └─VG1-LV1        253:2    0  100M  0 lvm   /tmp/new
```
15) Протестируйте целостность файла

```bash
vagrant@vagrant:~$ gzip -t /tmp/new/test.gz
vagrant@vagrant:~$ echo $?
0
```

16) Используя pvmove, переместите содержимое PV с RAID0 на RAID1

```bash
vagrant@vagrant:~$ sudo pvmove /dev/md1 /dev/md0
  /dev/md1: Moved: 16.00%
  /dev/md1: Moved: 100.00%
vagrant@vagrant:~$ lsblk
NAME                 MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
sda                    8:0    0   64G  0 disk
├─sda1                 8:1    0  512M  0 part  /boot/efi
├─sda2                 8:2    0    1K  0 part
└─sda5                 8:5    0 63.5G  0 part
  ├─vgvagrant-root   253:0    0 62.6G  0 lvm   /
  └─vgvagrant-swap_1 253:1    0  980M  0 lvm   [SWAP]
sdb                    8:16   0  2.5G  0 disk
├─sdb1                 8:17   0    2G  0 part
│ └─md0                9:0    0    2G  0 raid1
│   └─VG1-LV1        253:2    0  100M  0 lvm   /tmp/new
└─sdb2                 8:18   0  511M  0 part
  └─md1                9:1    0 1018M  0 raid0
sdc                    8:32   0  2.5G  0 disk
├─sdc1                 8:33   0    2G  0 part
│ └─md0                9:0    0    2G  0 raid1
│   └─VG1-LV1        253:2    0  100M  0 lvm   /tmp/new
└─sdc2                 8:34   0  511M  0 part
  └─md1                9:1    0 1018M  0 raid0
```

17) Сделайте --fail на устройство в вашем RAID1 md

```bash
vagrant@vagrant:~$ sudo mdadm /dev/md0 -f /dev/sdc1
mdadm: set /dev/sdc1 faulty in /dev/md0
```

18) Подтвердите выводом dmesg, что RAID1 работает в деградированном состоянии

```bash
[ 1818.731975] md/raid1:md0: not clean -- starting background reconstruction
[ 1818.731976] md/raid1:md0: active with 2 out of 2 mirrors
[ 1818.731989] md0: detected capacity change from 0 to 2144337920
[ 1818.732198] md: resync of RAID array md0
[ 1829.176795] md: md0: resync done.
[ 1929.120834] md1: detected capacity change from 0 to 1067450368
[ 1994.439334] md1: detected capacity change from 1067450368 to 0
[ 1994.439340] md: md1 stopped.
[ 2024.516437] md1: detected capacity change from 0 to 1067450368
[ 3392.262294] EXT4-fs (dm-2): mounted filesystem with ordered data mode. Opts: (null)
[ 3392.262302] ext4 filesystem being mounted at /tmp/new supports timestamps until 2038 (0x7fffffff)
[ 4495.206644] EXT4-fs (dm-2): mounted filesystem with ordered data mode. Opts: (null)
[ 4495.206650] ext4 filesystem being mounted at /tmp/new supports timestamps until 2038 (0x7fffffff)
[ 4627.418197] md/raid1:md0: Disk failure on sdc1, disabling device.
               md/raid1:md0: Operation continuing on 1 devices.
```

19) Протестируйте целостность файла, несмотря на "сбойный" диск он должен продолжать быть доступен

```bash
vagrant@vagrant:~$ gzip -t /tmp/new/test.gz
vagrant@vagrant:~$ echo $?
0
```
# Домашнее задание к занятию "3.6. Компьютерные сети, лекция 1"

1) Работа c HTTP через телнет
```bash
vagrant@vagrant:~$ telnet stackoverflow.com 80
Trying 151.101.193.69...
Connected to stackoverflow.com.
Escape character is '^]'.
GET /questions HTTP/1.0
HOST: stackoverflow.com

HTTP/1.1 301 Moved Permanently
cache-control: no-cache, no-store, must-revalidate
location: https://stackoverflow.com/questions
x-request-guid: be4dc91e-e90c-411f-9466-7950c628d2d4
feature-policy: microphone 'none'; speaker 'none'
content-security-policy: upgrade-insecure-requests; frame-ancestors 'self' https://stackexchange.com
Accept-Ranges: bytes
Date: Tue, 18 Jan 2022 15:07:59 GMT
Via: 1.1 varnish
Connection: close
X-Served-By: cache-fra19146-FRA
X-Cache: MISS
X-Cache-Hits: 0
X-Timer: S1642518479.498207,VS0,VE92
Vary: Fastly-SSL
X-DNS-Prefetch-Control: off
Set-Cookie: prov=cf86cdac-9087-8cd7-8721-53b6229df168; domain=.stackoverflow.com; expires=Fri, 01-Jan-2055 00:00:00 GMT; path=/; HttpOnly

Connection closed by foreign host. 
```

Получит ответ HTTP 301 Moved Permanently. Запрошенный ресурс был перещен (в данном случае в `https://stackoverflow.com/questions`).

2) Повторите задание 1 в браузере, используя консоль разработчика F12

Первый ответ сервера HTTP 307 Internal Redirect
```
Request URL: http://stackoverflow.com/
Request Method: GET
Status Code: 307 Internal Redirect
Referrer Policy: strict-origin-when-cross-origin
Location: https://stackoverflow.com/
Non-Authoritative-Reason: HSTS
```
Этот код означает, что ресурс был временно перемещён в URL, указанный в Location (`https://stackoverflow.com/`). При этом для перенаправленного запроса тело и метод запроса не будут изменены.

Страница была загружена за 1.89 сек
Самый долгий запрос `https://sb.scorecardresearch.com/cs/17440561/beacon.js`

3) Какой IP адрес у вас в интернете?

Из соображений безопасности и моей паранойи полный адрес говорить не буду (хоть он и не белый) ☺️

37.78.x.x

4) Какому провайдеру принадлежит ваш IP адрес? Какой автономной системе AS?

В запросе использовался мой реальный публичный IP адрес

Провайдер "Ростелеком", AS12389
```bash
vagrant@vagrant:~$ whois -h whois.radb.net 37.78.x.x
route:          37.78.0.0/16
descr:          PAO Rostelecom, Macroregional Branch South, Krasnodar, BRAS
origin:         AS12389
mnt-by:         STC-MNT
mnt-by:         ROSTELECOM-MNT
created:        2015-11-24T04:39:44Z
last-modified:  2015-11-24T04:39:44Z
source:         RIPE
remarks:        ****************************
remarks:        * THIS OBJECT IS MODIFIED
remarks:        * Please note that all data that is generally regarded as personal
remarks:        * data has been removed from this object.
remarks:        * To view the original object, please query the RIPE Database at:
remarks:        * http://www.ripe.net/whois
remarks:        ****************************
```

5) Через какие сети проходит пакет, отправленный с вашего компьютера на адрес 8.8.8.8? Через какие AS? Воспользуйтесь утилитой traceroute

В Virtual Box traceroute с UDP не работал корректно, пришлось использовать опцию `-I` для ICMP

```bash
vagrant@vagrant:~$ traceroute -AnI 8.8.8.8
traceroute to 8.8.8.8 (8.8.8.8), 30 hops max, 60 byte packets
 1  10.0.2.2 [*]  0.205 ms  0.169 ms  0.159 ms
 2  192.168.77.1 [*]  1.542 ms  1.535 ms  1.487 ms
 3  100.66.0.1 [*]  2.586 ms  2.693 ms  2.687 ms
 4  * * *
 5  * * *
 6  * * *
 7  * * *
 8  * * *
 9  74.125.253.109 [AS15169]  38.120 ms  38.151 ms  38.117 ms
10  216.239.48.85 [AS15169]  44.525 ms  44.675 ms  44.668 ms
11  * * *
12  * * *
13  * * *
14  * * *
15  * * *
16  * * *
17  * * *
18  * * *
19  * * *
20  8.8.8.8 [AS15169]  41.571 ms  41.710 ms  39.284 ms
```
6) Повторите задание 5 в утилите mtr. На каком участке наибольшая задержка - delay?

```bash
                                                 My traceroute  [v0.93]
vagrant (10.0.2.15)                                                                            2022-01-18T15:42:59+0000
Keys:  Help   Display mode   Restart statistics   Order of fields   quit
                                                                               Packets               Pings
 Host                                                                        Loss%   Snt   Last   Avg  Best  Wrst StDev
 1. AS???    10.0.2.2                                                         0.0%    20    0.3   0.4   0.1   2.4   0.5
 2. AS???    192.168.77.1                                                     0.0%    20    1.1   1.2   0.9   1.4   0.2
 3. AS???    100.66.0.1                                                       0.0%    20    2.1   2.1   1.7   2.5   0.1
 4. AS12389  178.34.130.56                                                   52.6%    20    2.1   2.1   2.0   2.3   0.1
 5. (waiting for reply)
 6. (waiting for reply)
 7. AS15169  108.170.250.146                                                 73.7%    20   21.9  19.4  18.6  21.9   1.4
 8. AS15169  142.251.71.194                                                  63.2%    20   41.4  41.8  41.3  43.3   0.7
 9. AS15169  74.125.253.109                                                   0.0%    20   33.4  36.3  33.0  61.6   7.2
10. AS15169  216.239.48.85                                                    0.0%    20   44.4  44.4  44.1  45.4   0.3
11. (waiting for reply)
12. (waiting for reply)
13. (waiting for reply)
14. (waiting for reply)
15. (waiting for reply)
16. (waiting for reply)
17. (waiting for reply)
18. (waiting for reply)
19. (waiting for reply)
20. AS15169  8.8.8.8                                                         37.5%    17   38.9  40.5  38.7  43.3   1.8
```
Наибольшая задержка на `AS15169  216.239.48.85`

7) Какие DNS сервера отвечают за доменное имя dns.google? Какие A записи? воспользуйтесь утилитой dig

```bash
vagrant@vagrant:~$ dig +short NS dns.google
ns1.zdns.google.
ns4.zdns.google.
ns2.zdns.google.
ns3.zdns.google.

vagrant@vagrant:~$ dig +short A dns.google
8.8.4.4
8.8.8.8
```
8) Проверьте PTR записи для IP адресов из задания 7. Какое доменное имя привязано к IP? воспользуйтесь утилитой dig

```bash
vagrant@vagrant:~$ dig +noall +answer -x 8.8.8.8
8.8.8.8.in-addr.arpa.   7158    IN      PTR     dns.google.
vagrant@vagrant:~$ dig +noall +answer -x 8.8.4.4
4.4.8.8.in-addr.arpa.   43835   IN      PTR     dns.google.
```