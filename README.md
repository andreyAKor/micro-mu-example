# micro-mu-example
Example of use a micro.mu framework for develop microservices on go-lang

---

Пример-каркас использования фреймворка micro.mu для разработки микросервисов на go-lang
Состоит из двух сервисов:
* geo.manager - управляющая сервис, ведущий учет данных в БД (MySQL) и распределяющий задания между исполнительными сервисами geo.worker
* geo.worker - исполнительный сервис (воркер), принимающий задачи от geo.manager и выполняющий его задачи

Сервисы общаются между собой используя системы очередей (queue) такие как RabbitMQ, NATS и прочее что поддерживает фреймворк micro.mu

Суть примера, показать как можно создать распределенное микросервисное приложение на micro.mu с центральной-головной частью (сервисом) и сетью сервисов обслуживающих его задачи.

Данный пример идеально подходит по парсингу/обработке вычислений чего либо на удаленно-распределенных узлах.

Для сборки приложений и внешних зависимостей (vendors) используется пакетный менеджер GB - https://andreykor.com/2018/01/24/%d1%83%d1%81%d1%82%d0%b0%d0%bd%d0%be%d0%b2%d0%ba%d0%b0-%d0%bf%d0%b0%d0%ba%d0%b5%d1%82%d0%bd%d0%be%d0%b3%d0%be-%d0%bc%d0%b5%d0%bd%d0%b5%d0%b4%d0%b6%d0%b5%d1%80%d0%b0-%d0%b4%d0%bb%d1%8f-go/