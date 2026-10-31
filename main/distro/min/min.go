package min

import (
	// Mandatory core
	_ "github.com/xtls/xray-core/app/dispatcher"
	_ "github.com/xtls/xray-core/app/proxyman/inbound"
	_ "github.com/xtls/xray-core/app/proxyman/outbound"

	// 最小必要功能
	_ "github.com/xtls/xray-core/app/log"
	_ "github.com/xtls/xray-core/app/router"

	// Fix dependency cycle
	_ "github.com/xtls/xray-core/transport/internet/tagged/taggedimpl"

	// 只保留 VLESS
	_ "github.com/xtls/xray-core/proxy/vless/inbound"
	_ "github.com/xtls/xray-core/proxy/vless/outbound"
	_ "github.com/xtls/xray-core/proxy/freedom"
	_ "github.com/xtls/xray-core/proxy/blackhole"

	// 只保留 VLESS + WS 需要的传输
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	_ "github.com/xtls/xray-core/transport/internet/websocket"
	_ "github.com/xtls/xray-core/transport/internet/tls"
	_ "github.com/xtls/xray-core/transport/internet/udp"

	// Headers
	_ "github.com/xtls/xray-core/transport/internet/headers/http"
	_ "github.com/xtls/xray-core/transport/internet/headers/noop"

	// 配置
	_ "github.com/xtls/xray-core/main/json"
	_ "github.com/xtls/xray-core/main/confloader/external"

	// 基本命令
	_ "github.com/xtls/xray-core/main/commands/all"
)
