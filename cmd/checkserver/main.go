package main

import (
	"fmt"
	"os"

	"cyberstrike-ai/internal/config"
)

func main() {
	cfg, err := config.Load(os.Args[1])
	if err != nil {
		fmt.Println("LOAD_ERROR:", err)
		os.Exit(1)
	}
	fmt.Println("server.host =", cfg.Server.Host)
	fmt.Println("server.port =", cfg.Server.Port)
	fmt.Println("server.tls_enabled =", cfg.Server.TLSEnabled)
	c2 := "nil(=enabled!)"
	if cfg.C2.Enabled != nil {
		c2 = fmt.Sprintf("%v", *cfg.C2.Enabled)
	}
	fmt.Println("c2.enabled =", c2)
	fmt.Println("mcp.enabled =", cfg.MCP.Enabled)
	fmt.Println("audit.enabled =", cfg.Audit.Enabled)
	fmt.Println("hitl.tool_whitelist =", cfg.Hitl.ToolWhitelist)
	fmt.Println("hitl.default_reviewer =", cfg.Hitl.DefaultReviewer)
	fmt.Printf("tool_guard = %+v\n", cfg.ToolGuard)
	fmt.Println("external_mcp.servers =", len(cfg.ExternalMCP.Servers))
	fmt.Println("ai.default_channel =", cfg.AI.DefaultChannel)
	for id, ch := range cfg.AI.Channels {
		masked := "***set***"
		if len(ch.APIKey) < 8 {
			masked = "***EMPTY***"
		}
		fmt.Printf("  channel %s: provider=%s base_url=%s model=%s key=%s\n", id, ch.Provider, ch.BaseURL, ch.Model, masked)
	}
	fmt.Println("security.tools_dir =", cfg.Security.ToolsDir, "inline tools =", len(cfg.Security.Tools))
}
