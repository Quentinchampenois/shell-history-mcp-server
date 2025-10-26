# MCP - Shell History Server

__⚠️ Disclaimer: This project is only for discover and demo purpose, there is no active development planned. Furthermore, it is only for personal use because it exposes entire shell history and represents a very high confidentiality and security risk.__

__💡I recommend to clear history before using it like: "history -c"__

`shell-history-mcp-server` bridges your shell history and AI assistants using the Model Context Protocol (MCP). It reads shell commands from ZSH, then lets AI agents query them through structured function calls

Functions:
* `shell_history`: Returns shell history commands 

## Examples

Prompt:
> On which project did I work today ?

Output:
- You should see a tool call to the Shell History before output is generated

## Getting started

1. Clone the project 
```bash
$ git clone https://github.com/quentinchampenois/shell-history-mcp-server.git 
$ cd shell-history-mcp-server
```

2. Build the server
```bash
$ go mod tidy
$ go build -o shell-history-mcp-server
```
💡 Binary `shell-history-mcp-server` must be present in `$PATH` to be callable

🚀 MCP Server is ready to use ! But how ?  
If you try to run it directly `$ ./shell-history-mcp-server` process will run but nothing happens, and it's the expected behaviour.

## Use @modelcontextprotocol/inspector to debug

Package `@modelcontextprotocol/inspector` starts a WebUI to test directly your MCP server, test output based on provided input.

### Requirements

* Node

### How to

Start inspector as following

```bash
$ npx @modelcontextprotocol/inspector shell-history-mcp-server
```

Access WebUI at http://localhost:6274/?MCP_PROXY_AUTH_TOKEN=<GENERATED_TOKEN>

## Configure client using Cherry studio and Scaleway AI provider

In this example we will use [Cherry Studio](https://github.com/CherryHQ/cherry-studio) as client and [model qwen3-235b-a22b-instruct-2507 from Scaleway provider](https://www.scaleway.com/en/docs/generative-apis/quickstart/). 

### Requirements
* Configure account on Scaleway and generate API token (~5 minutes)
* Cherry studio installed

### Configure Qwen3-235b from Scaleway on Cherry Studio

Open Cherry Studio and  configure the Scaleway provider: 

![01_CONFIGURE_PROVIDER.png](./docs/01_CONFIGURE_PROVIDER.png)
![02_CONFIGURE_PROVIDER.png](./docs/02_CONFIGURE_PROVIDER.png)


Configure the MCP server just installed:


![03_CONFIGURE_MCP.png](./docs/03_CONFIGURE_MCP.png)
![04_CONFIGURE_MCP.png](./docs/04_CONFIGURE_MCP.png)
![05_CONFIGURE_MCP.png](./docs/05_CONFIGURE_MCP.png)

__Start a new chat with MCP__

![06_CHAT_WITH_MCP.png](./docs/06_CHAT_WITH_MCP.png)

__🚀 Once Cherry studio let's give a try with a basic prompt, example: "Which scripts I am using ?"__

## Extra

* Source Cherry studio: https://github.com/CherryHQ/cherry-studio
* Source Scaleway AI (Generative API): https://www.scaleway.com/en/docs/generative-apis/quickstart/
* Go MCP SDK: https://github.com/modelcontextprotocol/go-sdk/tree/main
* MCP Inspector: https://modelcontextprotocol.io/docs/tools/inspector