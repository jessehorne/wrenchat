wrenchat
===

<img src="/assets/logo.png" width="400">

A secure, asymmetrically (RSA) encrypted chat service written in Go.

___

*This project is a WIP and is not  ready for production!*

more coming soon...

___

# Documentation

## Join Server

To join the server, simply connect to the API using the IP/Port combination using a TCP connection.

## Create Room

Request
```json
{
  "type": "CMD",
  "cmd": "create room",
  "name": "whatever",
  "password": "unencrypted password!!!"
}
```

OK Response
```json
{
  "type": "OK",
  "msg": "SERVER_UUID"
}
```

Error Response
```json
{
  "type": "ERROR",
  "reason": "something ain't right"
}
```

## Join Room

Request
```json
{
  "type": "JOIN",
  "uuid": "SERVER_UUID",
  "password": "SERVER_PASSWORD"
}
```

OK Response
```json
{
  "type": "OK"
}
```

Error Response
```json
{
  "type": "ERROR",
  "reason": "something ain't right again"
}
```

## Send Message

Request
```json
{
  "type": "MSG",
  "room": "SERVER_UUID",
  "msg": "hello world!"
}
```

OK Response
```json
{
  "type": "OK"
}
```

Error Response
```json
{
  "type": "ERROR",
  "reason": "something ain't right yet again!!!"
}
```