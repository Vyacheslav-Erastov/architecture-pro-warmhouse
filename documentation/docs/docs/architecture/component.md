```puml
@startuml

!define RELATIVE_INCLUDE
!include C4_Component.puml

title Warm House Component Diagram

top to bottom direction

Person(user, User,  "Пользователь системы Теплый Дом")

Container(spa, "SPA", "JavaScript, React", "Предоставляет функциональность системы Теплый Дом пользователю")

ContainerQueue(events_queue, "System Events Queue", "RabbitMQ", "Очередь для системных событий")
ContainerQueue(telemetry_queue, "Telemetry Events Queue", "RabbitMQ", "Очередь для событий телеметрии")
ContainerQueue(device_commands_queue, "Device Commands Queue", "RabbitMQ", "Очередь для команд управления устройствами")

ContainerDb(user_database, "User Database", "PostgreSQL Database", "Хранит информацию о пользователях")
ContainerDb(sensor_database, "Sensor Database", "PostgreSQL Database", "Хранит информацию о датчиках")
ContainerDb(device_database, "Device Database", "PostgreSQL Database", "Хранит информацию об устройствах")
ContainerDb(telemetry_database, "Telemetry Database", "PostgreSQL Database", "Хранит информацию о телеметрии")

System_Ext(SmartDeviceSystem, "Smart Device System", "Система, которая организует физическое подключение и интерфейс для управления и мониторинга умными устройствами")

Rel(SmartDeviceSystem, telemetry_queue, "Публикует события телеметрии", "async, AMQP")
Rel(SmartDeviceSystem, device_commands_queue, "Читает команды устройствам", "async, AMQP")

Rel(user, spa, "Использует", "Веб-браузер/HTTPS")

Container_Boundary(api_gateway, "API-Gateway") {
    Component(user_api_gateway, "User API Gateway", "FastAPI Router", "Позволяет обратиться к сервису управления пользователями")
    Component(sensor_api_gateway, "Sensor API Gateway", "FastAPI Router", "Позволяет обратиться к сервису управления датчиками")
    Component(device_api_gateway, "Device API Gateway", "FastAPI Router", "Позволяет обратиться к сервису управления устройствами")
    Component(telemetry_api_gateway, "Telemetry API Gateway", "FastAPI Router", "Позволяет обратиться к сервису телеметрии")
}

Rel(spa, api_gateway, "Делает API-вызовы", JSON/HTTPS)
Rel(api_gateway, spa, "Публикует события", "async, WebSocket")
Rel(api_gateway, events_queue, "Подписан на события", "async, AMQP")

Container_Boundary(user_service, "Сервис управления пользователями") {
    Component(user_api, "User API Controller", "FastAPI Router", "Обрабатывает HTTP-запросы пользователей")
    Component(user_logic, "User Service", "Python", "Бизнес-логика управления пользователями")
    Component(user_repo, "User Repository", "SQLModel", "Доступ к данным пользователей в БД")
}

Rel(user_api_gateway, user_api, "Делает API-вызовы", JSON/HTTPS)
Rel(user_repo, user_database, "Читает и пишет", "sync, SQLModel")


Container_Boundary(device_service, "Device Service") {
    Component(device_api, "Device API Controller", "FastAPI Router", "Обрабатывает запросы управления устройствами")
    Component(device_logic, "Device Service", "Python", "Бизнес-логика управления устройствами")
    Component(device_repo, "Device Repository", "SQLModel", "Доступ к данным устройств в БД")
    Component(device_producer, "Device Commands Producer", "AMQP", "Публикует команды в очередь устройств")
}

Rel(device_api_gateway, device_api, "Делает API-вызовы", JSON/HTTPS)
Rel(device_repo, device_database, "Читает и пишет", "sync, sqlx")
Rel(device_producer, device_commands_queue, "Отправляет команды", "async, AMQP")

Container_Boundary(telemetry_service, "Telemetry Service") {
    Component(telemetry_api, "Telemetry API Controller", "FastAPI Router", "Обрабатывает запросы телеметрии")
    Component(telemetry_repo, "Telemetry Repository", "SQLModel", "Доступ к данным телеметрии")
    Component(telemetry_consumer, "Telemetry Consumer", "AMQP", "Обрабатывает события телеметрии из очереди")
    Component(telemetry_producer, "Telemetry Producer", "AMQP", "Публикует события телеметрии")
}

Rel(telemetry_api_gateway, telemetry_api, "Делает API-вызовы", JSON/HTTPS)
Rel(telemetry_repo, telemetry_database, "Читает и пишет", "sync, SQLModel")
Rel(telemetry_producer, events_queue, "Публикует события", "async, AMQP")
Rel(telemetry_consumer, telemetry_queue, "Читает события телеметрии", "async, AMQP")


Container_Boundary(sensor_service, "Sensor Service") {
    Component(sensor_api, "Sensor API Controller", "Gin Router", "Обрабатывает запросы управления датчиками")
    Component(sensor_logic, "Sensor Service", "Go", "Бизнес-логика управления датчиками")
    Component(sensor_repo, "Sensor Repository", "pgx", "Доступ к данным датчиков в БД")
}

Rel(sensor_api_gateway, sensor_api, "Делает API-вызовы", JSON/HTTPS)
Rel(sensor_repo, sensor_database, "Читает и пишет", "sync, pgx")


Container_Boundary(temperature_api, "Temperature-API"){
    Component(temperature_contoller, "Temperature API Controller", "FastAPI APIRouter", "Обрабатывает запросы температуры датчиков")
}

Rel(sensor_logic, temperature_contoller, "Запрашивает текущую температуру датчика", "sync, JSON/HTTPS")

@enduml
```
