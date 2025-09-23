# 2. Diseño y Arquitectura del Sistema

**Escenario:**
Una start up local de reparto de comida a domicilio
está experimentando un rápido crecimiento.
Necesita rediseñar su sistema backend para mejorar
la escalabilidad, la fiabilidad y la facilidad de
mantenimiento. El sistema actual es una aplicación
monolítica con un rendimiento limitado.
Diseñe una arquitectura de sistema distribuida que
aborde estos desafíos. Considere factores como el
diseño de la base de datos, el diseño de la API, las
colas de mensajes y el almacenamiento en caché.
Explique sus opciones de diseño y justifique su
arquitectura.

### Solución
![Component diagram](docs/diagram.svg)

### Explicación
El diseño se realizo pensando principalmente en el flujo de compra y entrega de ordenes.

###### Patrónes de arquitectura implementados
- DLQ (Cola de mensajes muertos)
- SAGA (Choreography-based)
- CQRS (se replica usando el WAL)
- Transactional Outbox

###### Decisiones de arquitectura
- Decidí no usar caché en el micro de Stock para consultar la disponibilidad porque al estar usando DynamoDB,
  si la información se indexa adecuadamente tiene un performance similar al uso de Redis (u otro Sistema Gestor de Base de Datos Clave-Valor)
- Opte por usar el patrón Transactional Outbox para garantizar la consistencia de información al confirmar una orden.
- Para el patrón de CQRS decidí usar DynamoDB para las lecturas y PostgreSQL para las escrituras.
- Si bien en el diagrama no se muestra explícitamente, al implementar el patrón CQRS yo replicaría la información aprovechando el WAL de PostgreSQL.