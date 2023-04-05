# 2023-Backend-Intern-Homework

## project layout
```
jkf-star  
├─ cmd/             # 本專案的主要應用程式
│   ├─ main.go/      
├─ db/              # 資料庫DDL
│   ├─ migrations/     # ddl
│   └─ samples/        # mock data
├─ docker/
│   ├─ postgresql/     # postgres servies
│   ├─ workspace/      # workspace series
├─ internal/        # 私有應用程式和函示庫的程式碼
│   ├─ interceptor/    # grpc 攔截器
│   ├─ Error/          # grpc error code 
├─ model/           # gorm model
├─ pb/              # grpc pb
├─ pkg/             # 支援工具
│   ├─ cache/       # cache 目前暫時不引入
│   ├─ protocol/    # protocol 協議相關 目前暫時不用 
│   ├─ repote/      # repote 對應的server 目前單純print
├─ proto/           # proto file
├─ service/         # api service
├─ repository/      # api repository
├─ test/            # 測試應用程式和測試資料
├─ transport/       # GRPC API
├─ utils/           # GRPC API  
├─ go.mod           # config相關
├─ go.sum
├─ Dockerfile       # for api golang  
├─ docker-compose   
├─ makefile
├─ .env
├─ .env.development
├─ app.go          # 主程式執行點
└─ README.md        
```

## docker start
```shell
docker-compose up postgres -d
```

## test
```shell
go test ./test -v -failfast
```
## storage 的選擇

