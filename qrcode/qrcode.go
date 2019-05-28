package qrcodeGeneratorService

import(
   "fmt"
   "net/http"
   // qrcode "github.com/skip2/go-qrcode"
   // "encoding/json"
   // log "github.com/sirupsen/logrus"
)

type QRCodeGenerator struct {
   Path  string
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

