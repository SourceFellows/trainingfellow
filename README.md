# Beispiel Anwendung "Training Fellow"

Die Beispielanwendung "Training Fellow" wird im Buch "Microservices mit Go" des Rheinwerk Verlags beschrieben. Sie soll Schulungsanbieter bei der Schulungsbuchung, -vorbereitung und -durchführung unterstützen und vor Allem Ansätze für die Umsetzungen von Microservices in Go aufzeigen.

https://www.rheinwerk-verlag.de/microservices-mit-go-konzepte-werkzeuge-best-practices/

Der aktuelle Stand befinden sich immer unter: https://github.com/SourceFellows/trainingfellow

## Die Anwendung

Eine ausführliche Beschreibung der Anwendung befindet sich im Buch. Hier nur kurz die Domain Story der Anwendung:

![Domain Story der Anwendung](Training-Fellow.png)

## Start der Anwendung

Die komplette Anwendung kann über Docker-Compose mit folgendem Kommando gestartet werden:

```
docker-compose up
```

Daraufhin werden die Docker Container für die folgenden Services gestartet:

* Registrierungs-Service

* Vorbereitungs-Service

* NATS-Server

* MongoDB

* Mongo-Express