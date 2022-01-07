
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
