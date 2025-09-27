```puml
@startuml
title Warm House Context Diagram

top to bottom direction

!include C4_Context.puml

Person(user, "Пользователь", "Пользователь системы Теплый Дом")
System(WarmHouseSystem, "Система Тёплый дом", "Система, которая организует удаленное управление отоплением в доме")

Rel(user, WarmHouseSystem, "Использует систему")

@enduml
```
