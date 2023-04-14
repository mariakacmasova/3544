package config

import (
	"fmt"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

var (
	AppVersion             = "v4.0.0"
	AppPort                = "3000"
	AppDebug               = true
	AppOs                  = fmt.Sprintf("Chrome")
	AppPlatform            = waProto.DeviceProps_PlatformType(1)
	AppBasicAuthCredential string

	PathQrCode    = "statics/qrcode"
	PathSendItems = "statics/senditems"
	PathStorages  = "storages"

	DBName = "whatsapp.db"

	WhatsappLogLevel            = "ERROR"
	WhatsappAutoReplyMessage    string
	WhatsappAutoReplyWebhook    = "https://events.hookdeck.com/e/src_iuL6DUtZwqu4?x-hookdeck-allow-methods=get"
	WhatsappSettingMaxFileSize  int64 = 50000000  // 50MB
	WhatsappSettingMaxVideoSize int64 = 100000000 // 100MB
)
