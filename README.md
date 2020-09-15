## QRCode Generator API Server 程式
主要用來自動產生QRCODE圖檔，並提供下載

### Usage

* Download QRCode image

```
POST http://10.109.10.224/qrcodegenerator/generateQRCodeAndDownload

params:  your data
size: 160/256 or other size
```

* Show QRCode image

```
GET http://10.109.10.224/qrcodegenerator/generateQRCode/160/www.justdrink.com.tw
```

#### 使用函式庫
```
go get -u github.com/skip2/go-qrcode/...
```
