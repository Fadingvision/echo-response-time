# responsetime Middleware

responsetime Middleware add HTTP response Header which shows how many time spent in this request.

### Usage

```go
e := echo.New()
e.Use(responsetime.ResponseTime())
```

### Custom Configuration

```js
e := echo.New()
e.Use(responsetime.ResponseTimeWithConfig(responsetime.ResponseTimeConfig{
  Suffix: true,
  Digits: 2,
}))
```

### Default Configuration

```go
DefaultResponseTimeConfig = ResponseTimeConfig{
  Digits:     3,
  HeaderName: "X-Response-Time",
  Suffix:     false,
}
```

