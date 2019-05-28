package qrcodeGeneratorService

import(
   "fmt"
   "time"
   "bytes"
   "image/color"
   "strconv"
   "io/ioutil"
   "net/http"
   qrcode "github.com/skip2/go-qrcode"
   log "github.com/sirupsen/logrus"
   // "encoding/json"
)

type QRCodeGenerator struct {
   Path  string
}

func (qrg *QRCodeGenerator) WritePNGFile(url, filename string)(error) {
   if err := qrcode.WriteColorFile(url, qrcode.Medium, 256, color.Black, color.White, filename); err != nil {
      return err
   }
   return nil
}

func (qrg *QRCodeGenerator) createPNGImage(qrcodeurl string) ([]byte, error) {
   if err := qrg.WritePNGFile(qrcodeurl, qrg.Path + "test.png"); err != nil {
      return nil, err
   }
   png, err := ioutil.ReadFile(qrg.Path + "test.png")
   if err != nil {
      log.Printf("error in myHandler - error: %v", err)
      return nil, err 
   }
/*
   var png []byte
   png, err := qrcode.Encode(qrcodeurl, qrcode.Medium, 256)
   if err != nil {
      log.Printf("error in myHandler - error: %v", err)
      return nil, err
   }
*/
   return png, nil
}

// Write file to client for download
func (qrg *QRCodeGenerator) QRCodeImage(w http.ResponseWriter, r *http.Request) {
   png, err := qrg.createPNGImage("https://www.justdrink.com.tw/")
   if err != nil {
      fmt.Fprintf(w, "Something error:%v", err)
      
   } else {
      w.Header().Set("Content-type", "image/png")
      w.Write(png)
   }
}

// Write file to client for download
func (qrg *QRCodeGenerator) DownloadQRCode(w http.ResponseWriter, r *http.Request) {
   png, err := qrg.createPNGImage("https://www.justdrink.com.tw/")
   if err != nil {
      fmt.Fprintf(w, "Something error:%v", err)
   } else {
      mime := http.DetectContentType(png)
      fileSize := len(string(png))
      w.Header().Set("Content-type", mime)
      w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
      w.Header().Set("Expires", "0")
      w.Header().Set("Content-Transfer-Encoding", "binary")
      w.Header().Set("Content-Length", strconv.Itoa(fileSize))
      w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
      http.ServeContent(w, r, "qrcode.png", time.Now(), bytes.NewReader(png))
   }
}

// Initial
func NewQRCodeGenerator(path string) (*QRCodeGenerator, error) {
   if len(path) <= 0 {
      return nil, fmt.Errorf("Must have path.")
   }
   return &QRCodeGenerator {
      Path: path,
   }, nil
}


// check health
func (qrg *QRCodeGenerator) Healthz(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
}

