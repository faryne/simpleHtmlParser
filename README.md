## SimpleHTMLParser

### 編譯方法
```shell script
go build -o simpleHTMLParser main.go 
```

### 使用方法
```shell script
./simpleHTMLParser -uri https://www.example.com -file ./crwaler.json
```

### JSON 設定檔結構
欄位名稱 | 欄位說明
---- | ----
encoding | 目標網頁的編碼
selectors | 要爬取的網頁元素，是一個 Selector 集合

Selector 欄位名稱 | Selector 欄位說明
---- | ----
identifier | string，輸出時所使用的名稱
selector | string，要爬取目標的 css selector
repeated | boolean，輸出時是否使用陣列輸出
property | enum，必須是 html / text / attr 中之一
target | string，當 property 為 attr 時使用。代表要爬指定的元素的屬性
type | enum，必須是 string / integer / boolean，輸出值時要轉出何種型態

