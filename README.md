## SimpleHTMLParser

### 編譯方法
```shell script
go build -o simpleHTMLParser main.go 
```

### 使用方法
```shell script
./simpleHTMLParser -uri https://www.example.com -file ./crwaler.json [-parse_type html|csv]
```

### JSON 設定檔結構
欄位名稱 | 欄位說明
---- | ----
encoding | 目標網頁的編碼
selectors | 要爬取的網頁元素，是一個 ```Selector``` 集合

Selector 欄位名稱 | 欄位類型 | Selector 欄位說明
---- | ---- | ----
identifier | string | 輸出時所使用的名稱
selector | string | 要爬取目標的 css selector
repeated | boolean | 輸出時是否使用陣列輸出
children | ```[]Selector``` | 查詢該元素下的特定子元素適用
output | ```Output``` | 輸出 

Output 欄位名稱 | 類型 | Output 欄位說明
---- | ---- | ----
property | enum | 必須是 html / text / attr 中之一
target | string | 當 property 為 attr 時使用。代表要爬指定的元素的屬性
type | enum | 必須是 string / integer / boolean，輸出值時要轉出何種型態

### 使用範例

#### 爬 BookWalker 輕小說列表頁
```shell script
go run main.go -uri "https://www.bookwalker.com.tw/more/fiction/1/3" -file ./samples/crawler2.json
```

#### 爬 DMM 女優頁
```shell script
go run main.go -uri "http://actress.dmm.co.jp/-/detail/=/actress_id=17/" -file ./samples/crawler.json  
```

#### 爬 Getchu 排名頁
```shell script
go run main.go -uri "http://www.getchu.com/rank/?genre=pc_soft" -file ./samples/crawler3.json
```

#### 爬證交所上市公司列表
```shell script
go run main.go -uri "https://isin.twse.com.tw/isin/C_public.jsp?strMode=2" -file ./samples/test1.json
```

#### 爬取 csv 內容（教育部-振興三倍券通訊交易適用業者名單）
```shell script
go run main.go  -parse_type csv -uri "https://transform.cloud.sa.gov.tw/DataSets/DataSetResource.ashx?rId=A09010000E-000195-001" -file ./samples/csvtest1.json  
```
### @TODO
* 加上爬取網頁時可用的選項，例如 User-Agent 等 header 變更