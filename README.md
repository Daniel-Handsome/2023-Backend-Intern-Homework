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
step 1
將.env的DB_HOST改成postgres

step 2
```shell
docker-compose up postgres api -d
```

## Precautions
golang version 請用1.18以上
1.17有相容性問題

## test
```shell
go test ./test -v -failfast
```
## storage 的選擇
選擇postgres的原因有幾點
   1 在transaction方面postgres處理比較好，面對高併發的處理上也比較好
   2 postgres是使用head, 所以在 Range Query 的時候select多的欄位的話postgres比較有優勢，因為不用從non-clustered index 取得 clustered index 再去查詢
   3 對於當前的的個人推薦因為需要大量刪除跟新增node對應的文章，所以如果用fk去用在效能及修改上會比較麻煩，而用postgres的array會比較方便動態增減
   4 postgres的分區有分範圍分區跟子分區，比起mysql更加方便，可以依照想要的欄位去分區
綜合以上幾點對於目前的功能開發來說，postgres會是較優的選擇。

可以想想除了⼀筆筆刪之外怎麼清除更有效率
   1 目前是想到用分區去做，可以依照時間去分區，再依照時間一個一個刪除
   2 直接刪除table在create建立
   3 TRUNCATE TABLE 清空全部
