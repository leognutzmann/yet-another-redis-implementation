# Redis Clone

Uma implementação simples de um servidor Redis em Go, com suporte ao protocolo [RESP (Redis Serialization Protocol)](https://redis.io/docs/latest/develop/reference/protocol-spec/).

## O que foi implementado

- **Servidor TCP** na porta 6379
- **Protocolo RESP** para comunicação (reader e writer)
- **Comandos suportados:**
  - `PING`
  - `SET key value`
  - `GET key`
  - `HSET hash field value`
  - `HGET hash field`
  - `HGETALL hash`

## Como rodar

```bash
go run .
```

Use o `redis-cli` ou qualquer outro client que implemente o protocolo RESP para conectar:

```bash
redis-cli
127.0.0.1:6379> PING
PONG
127.0.0.1:6379> SET name Leonardo
OK
127.0.0.1:6379> GET name
"Leonardo"
127.0.0.1:6379> HSET user:1 name Leonardo age 30
(integer) 2
127.0.0.1:6379> HGET user:1 name
"Leonardo"
127.0.0.1:6379> HGETALL user:1
1) "name"
2) "Leonardo"
3) "age"
4) "30"
```
