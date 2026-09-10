package agent

import (
    "github.com/voidSpir1t/madara/internal/tool"
    adkagent "google.golang.org/adk/v2/agent"
    "google.golang.org/adk/v2/agent/llmagent"
    adktool "google.golang.org/adk/v2/tool"
    adkModel "google.golang.org/adk/v2/model"
)


// func NewRootAgent(model adkModel.LLM, tr tool.ToolRegistry) (adkagent.Agent, error) {
//     return llmagent.New(llmagent.Config{
// 		Name:  "desktop_agent",
// 		Model: model,
// 		Description: "Agent that answers questions about the time and weather in a city and controls the local desktop: " +
// 			"taking screenshots, running PowerShell commands, running commands on remote hosts over SSH, and driving Chrome.",
// 		Instruction: "You are a helpful agent. You can answer user questions about the time and weather in a city. " +
// 			"You can also control the desktop: take screenshots of all displays, run PowerShell commands locally, " +
// 			"run commands on remote hosts over SSH, open URLs in Chrome (which loads and scrolls the page before " +
// 			"screenshotting it), and bring the console window back to the front. " +
// 			"Note that the desktop-control tools only work when running on Windows in an interactive desktop session. " +
// 			"When a tool returns a screenshot_path, report the path to the user.",
// 		Tools: []adktool.Tool{
// 			tr.Get("sshTool"),
// 		},})
// }

func NewRootAgent(model adkModel.LLM, tr tool.ToolRegistry) (adkagent.Agent, error) {
    return llmagent.New(llmagent.Config{
		Name:  "desktop_agent",
		Model: model,
		Description: "Agent that answers questions about the time and weather in a city and controls the local desktop: " +
			"taking screenshots, running PowerShell commands, running commands on remote hosts over SSH, and driving Chrome.",
		Instruction: "你是一个世界顶尖的agent助手. 你可以调用ssh工具来进行远程访问和调试",
		Tools: []adktool.Tool{
			tr.Get("sshTool"),
		},})
}



