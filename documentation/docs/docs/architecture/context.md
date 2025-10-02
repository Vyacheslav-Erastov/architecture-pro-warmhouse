```puml
@startuml
title Warm House Context Diagram

top to bottom direction

!define RELATIVE_INCLUDE
!include C4_Context.puml

Person(user, "User", "Пользователь системы Теплый Дом")
System(WarmHouseSystem, "Warm House System", "Система, которая организует удаленное управление отоплением в доме")
System_Ext(SmartDeviceSystem, "Smart Device System", "Система, которая организует физическое подключение и интерфейс для управления и мониторинга умными устройствами")

Rel(user, WarmHouseSystem, "Использует систему")
Rel(WarmHouseSystem, SmartDeviceSystem, "Использует систему")

@enduml
```
