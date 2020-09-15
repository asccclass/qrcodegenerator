package main

import(
   "os"
   "github.com/gorilla/mux"
   "github.com/asccclass/serverstatus"
   "github.com/asccclass/staticfileserver"
   "github.com/asccclass/staticfileserver/libs/qrcode"
)

func NewRouter(srv *SherryServer.ShryServer, documentRoot string)(*mux.Router)  {
   router := mux.NewRouter()

   // QRCode
   var QRCode  *qrcodeGeneratorService.QRCodeGenerator
   if os.Getenv("QRCodePath") != "" {
      var err error
      QRCode, err = qrcodeGeneratorService.NewQRCodeGenerator(srv, os.Getenv("QRCodePath"))
      if err == nil {
         QRCode.AddRouter(router) 
      }
   }

   //logger
   router.Use(SherryServer.ZapLogger(srv.Logger))

   // health check
   systemName := os.Getenv("SystemName")
   m := serverstatus.NewServerStatus(systemName)
   router.HandleFunc("/healthz", m.Healthz).Methods("GET")

   // Geo Location
   srv.GeoLocation.AddGeoLocationRouter(router)

   // Static File server
   staticfileserver := SherryServer.StaticFileServer{documentRoot, "index.html"}
   router.PathPrefix("/").Handler(staticfileserver)

   return router
}
