# dcard-take-home-test

## 設計說明
此實作以 redis 作為 rate limit database，搭配 TTL 設定，優點在於
- 效能佳：in-memory database 存取效率快
- 可擴充：未來 deploy 到多台機器上，不管 auto scale 幾台都可以有 rate limit 效果（假設在 redis 效能足夠應付的狀況下）
- 設計簡單：透過 TTL 及 INCR 設計邏輯簡單，不需自行實作複雜計算邏輯

因為 redis 單台 single thread，所以尚未考量 race condition 的問題，若未來需要 redis cluster，就必須使用 distributed locks 來確保資源競爭下服務行為不會受影響。

## 環境設定

- Golang 1.13
- Redis 5.0.4

### Redis
此實作及測試皆需要本地端安裝 redis server，請利用啟動以下指令啟動：
```shell
docker-compose -f docker-compose.yml up -d redis
```

測試後手動關閉：
```shell
docker-compose stop redis
```

### 啟動 API
```shell
go run main.go
```
啟動後開啟 http://127.0.0.1:8080/rate 便可開始實測

## 執行測試
在 [main_test.go](main_test.go) 中，測試一分鐘60次內，60次以上預期錯誤，以及一分鐘過後預期成功，請透過以下指令執行：

```shell
go test -v
```

