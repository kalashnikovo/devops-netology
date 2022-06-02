# Домашнее задание к занятию "7.2. Облачные провайдеры и синтаксис Terraform."

ДЗ выполнял для YC

Конфигурационные файлы Terraform \
[main.tf](./src/main.tf) \
[versions.tf](./src/versions.tf)

Создали сервисный аккаунт

```console
vagrant@ubuntu-20:~/terraform$ yc iam service-account list --folder-id b1g238j6jo2h41vp196i
+----------------------+-----------+
|          ID          |   NAME    |
+----------------------+-----------+
| aje1pqs1g29edfql05u9 | terraform |
+----------------------+-----------+
```

Далее получаем токен сервисного аккаунта

```console
vagrant@ubuntu-20:~/terraform$ export YC_TOKEN=`yc iam create-token`
```

Инициализируем провайдеры

```console
vagrant@ubuntu-20:~/terraform$ terraform init

Initializing the backend...

Initializing provider plugins...
- Reusing previous version of yandex-cloud/yandex from the dependency lock file
- Using previously-installed yandex-cloud/yandex v0.75.0

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

Выполняем применением описанной конфигурации

```console
vagrant@ubuntu-20:~/terraform$ terraform apply -auto-approve
Apply complete! Resources: 3 added, 0 changed, 0 destroyed.

Outputs:

external_ip_address_vm_1 = "51.250.76.247"
internal_ip_address_vm_1 = "192.168.10.16"
```

Удаляем ресурсы

```console
vagrant@ubuntu-20:~/terraform$ terraform destroy -auto-approve
Destroy complete! Resources: 3 destroyed.
```

---
