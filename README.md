
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

