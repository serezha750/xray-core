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

	// 传输层：TCP + WebSocket + TLS（TLS 内含 ECH / uTLS fingerprint 支持）
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	_ "github.com/xtls/xray-core/transport/internet/websocket"
	_ "github.com/xtls/xray-core/transport/internet/tls" // 支持 echConfigList / fingerprint(chrome等)
	_ "github.com/xtls/xray-core/transport/internet/udp"

	// Headers
	_ "github.com/xtls/xray-core/transport/internet/headers/http"
	_ "github.com/xtls/xray-core/transport/internet/headers/noop"

	// 配置
	_ "github.com/xtls/xray-core/main/json"
	_ "github.com/xtls/xray-core/main/confloader/external"

	// 基本命令（含 tls ech 生成工具）
	_ "github.com/xtls/xray-core/main/commands/all"
)
