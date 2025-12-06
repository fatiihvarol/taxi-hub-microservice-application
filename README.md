# TaxiHub Microservice Application

TaxiHub is a real-time taxi management platform built using a modern microservice architecture. The system includes Authentication, Driver Service, Real-time Location Service (Redis GEO), WebSocket live tracking, API Gateway, Elasticsearch logging, and Kibana dashboard monitoring.

---

## Prerequisites
- Go (Golang)
- Docker & Docker Compose
- MongoDB
- Redis
- Elasticsearch
- Kibana

---

## Project Setup

### 1. Clone the Repository
```bash
git clone https://github.com/fatiihvarol/taxi-hub-microservice-application
cd taxi-hub-microservice-application
```

### 2. Create all `.env` file for each services
Create an `.env` file in the project:

```
JWT_SECRET=jwtsecret
REFRESH_SECRET=refreshsecret
API_KEY=12345
ELASTIC_URL=http://elasticsearch:9200
ELASTIC_INDEX=gateway-logs
```

### 3. Start All Services
```bash
docker compose up --build
```

---

## Kibana Dashboard
- URL: **http://localhost:5601**
- Index: **gateway-logs**

---

## General Information
TaxiHub is designed as a fully decoupled microservice ecosystem, where each service handles a dedicated responsibility such as authentication, driver management, location tracking, and logging. All traffic flows through an API Gateway, which also handles rate-limiting, validation, and logging into Elasticsearch.

---

## Service Specifications

---

### 🔐 Auth Service
- User Registration  
- User Login (JWT + Refresh Token)  
- Refresh Token endpoint  
- Token validation (`/auth/validate`)  
- Password hashing using BCrypt  
- JWT signing using HS256  

#### Example validate response:
```json
{
  "valid": true,
  "userId": "123",
  "role": "customer"
}
```

---

### 🚖 Driver Service
- CRUD operations for drivers  
- MongoDB as the data source  
- Routed through API Gateway  

---

### 📍 Location Service (Redis GEO + WebSocket)
- Stores driver locations using Redis GEO  
- Finds nearest drivers using GEORADIUS  
- Real-time WebSocket tracking  
- Built-in payload validation  

#### WebSocket endpoint:
```
ws://localhost:8081/ws
```

#### Example WebSocket message:
```json
{
  "driverId": "abc123",
  "lat": 40.992,
  "lon": 29.124
}
```

---

### 🛡️ API Gateway
- IP-based Rate Limiting (100 requests per minute)  
- JWT validation  
- API Key middleware  
- Full request/response logging to Elasticsearch  
- Reverse Proxy routing to microservices  

#### Log example:
```json
{
  "method": "POST",
  "path": "/auth/login",
  "status": 200,
  "duration_ms": 12
}
```

---

## Architecture Overview

```
                           +--------------------------+
                           |        API Gateway       |
 Clients ----------------->|  - JWT Validation        |
                           |  - Logging (Elastic)     |
                           |  - Rate Limiter          |
                           |  - API Key Validation    |
                           +--------------------------+
                                 ^        ^        ^
                                 |        |        |
        +------------------------+        |        +------------------------+
        |                                 |                                 |
+----------------------+      +----------------------+      +----------------------+
|     Auth Service     |      |    Driver Service    |      |   Location Service   |
|  - JWT, Refresh      |      |  - MongoDB           |      |  - Redis GEO         |
|  - User Auth         |      |  - Driver CRUD       |      |  - WebSocket         |
+----------------------+      +----------------------+      |  - Nearby Endpoint   |
                                                           +----------------------+

                   +------------------------------------------+
                   |        Elasticsearch + Kibana            |
                   |     - Centralized Logging                |
                   +------------------------------------------+
```


---

## Redis GEO Usage

### Add a driver position
```
GEOADD driver-locations <lon> <lat> <driverId>
```

### Find nearby drivers
```
GEORADIUS driver-locations <lon> <lat> 5 km WITHDIST
```



## License
MIT License
