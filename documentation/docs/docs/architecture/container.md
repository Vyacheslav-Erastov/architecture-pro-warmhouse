```puml
@startuml

!define RELATIVE_INCLUDE
!include C4_Container.puml

title Warm House Container Diagram

top to bottom direction

Person(user, User,  "Пользователь системы Теплый Дом")

System_Boundary(WarmHouseSystem, "Warm House System") {
    Container(spa, "SPA", "JavaScript, React", "Предоставляет функциолнальность системы Теплый Дом пользователю")
    ContainerDb(user_database, "User Database", "PostgreSQL Database", "Хранит информацию о пользователях")
    ContainerDb(sensor_database, "Sensor Database", "PostgreSQL Database", "Хранит информацию о датчиках")
    ContainerDb(device_database, "Device Database", "PostgreSQL Database", "Хранит информацию об устройствах")
    ContainerDb(telemetry_database, "Telemetry Database", "PostgreSQL Database", "Хранит информацию о телеметрии")
    Container(api_gateway, "API-Gateway", "Python, FastAPI", "Предоставляет функционал системы Умный Дом по API", $tags="apiContainer")
    ContainerQueue(events_queue, "System Events Queue", "RabbitMQ", "Очередь для системных событий")
    Container(sensor_service, "Sensor Service", "Go, Gin", "Предоставляет функционал управления датчиками")
    Container(temperature_api, "Temperature-API", "Python, FastAPI", "Предоставляет получения текущей температуры датчика")
    Container(telemetry_service, "Telemetry Service", "Python, FastAPI", "Предоставляет функционал сбора телеметрии")
    Container(device_service, "Device Service", "Go, Gin", "Предоставляет функционал управления устройствами")
    Container(user_service, "User Service", "Python, FastAPI", "Предоставляет функционал управления пользователями")
    ContainerQueue(telemetry_queue, "Telemetry Events Queue", "RabbitMQ", "Очередь для событий телеметрии")
    ContainerQueue(device_commands_queue, "Device Commands Queue", "RabbitMQ", "Очередь для команд управления устройствами")
}

System_Ext(SmartDeviceSystem, "Smart Device System", "Система, которая организует физическое подключение и интерфейс для управления и мониторинга умнымыми устройствами")

Rel(user, spa, "Ипсользует", "Веб-браузер/HTTPS")

Rel(telemetry_queue, telemetry_service, "Читает события телеметрии", "async, AMQP")
Rel(device_service, device_commands_queue, "Отправляет команды", "async, AMQP")
Rel(spa, api_gateway, "Использует", "sync, JSON/HTTPS")
Rel(api_gateway, spa, "Публикует события", "async, WebSocket")
Rel(api_gateway, events_queue, "Подписан на события", "async, AMQP")
Rel(device_service, events_queue, "Публикует события", "async, AMQP")
Rel(telemetry_service, events_queue, "Публикует события", "async, AMQP")
Rel(api_gateway, sensor_service, "Использует", "sync, JSON/HTTP")
Rel(api_gateway, telemetry_service, "Использует", "sync, JSON/HTTP")
Rel(telemetry_service, device_service, "Обновляет статус устройств", "sync, JSON/HTTP")
Rel(api_gateway, device_service, "Использует", "sync, JSON/HTTP")
Rel(api_gateway, user_service, "Использует", "sync, JSON/HTTP")
Rel(sensor_service, sensor_database, "Читает и пишет", "sync, pgx")
Rel(sensor_service, temperature_api, "Запрашивает текущую температуру датчика", "sync, JSON/HTTP")
Rel(telemetry_service, telemetry_database, "Читает и пишет", "sync, SQLModel")
Rel(device_service, device_database, "Читает и пишет", "sync, sqlx")
Rel(user_service, user_database, "Читает и пишет", "sync, SQLModel")

Rel(SmartDeviceSystem, telemetry_queue, "Публикует события телеметрии", "async, AMQP")
Rel(device_commands_queue, SmartDeviceSystem, "Читает команды устройствам", "async, AMQP")


@enduml
```
