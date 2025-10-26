package functions_shell_history

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"shell-history-mcp-server/internal/shell"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Command   string `json:"command,omitempty" jsonschema:"Specify the command to search"`
	Contains  string `json:"contains,omitempty" jsonschema:"Looks for any command containing this substring"` // any command with this substring
	AfterTime int64  `json:"after_time,omitempty" jsonschema:"Browse history after this time"`                // Unix timestamp
	Limit     int    `json:"limit,omitempty" jsonschema:"Limit the number of results to return"`
}

type Output struct {
	Commands string `json:"commands" jsonschema:"the shell commands found in history"`
}

func GetShellHistory(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {

	home, _ := os.UserHomeDir()
	file, err := os.Open(home + "/.zsh_history")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cmds []shell.ZshHistory
	for scanner.Scan() {
		line := scanner.Text()
		zshHistory := shell.NewZshHistory(line)

		// If zsh row does not match regexp, zshHistory will be nil
		if zshHistory == nil {
			continue
		}

		if input.Command != "" {
			if len(zshHistory.Command) == 0 {
				continue
			}
			// Match the first word of the command (e.g., "docker", "uv")
			firstWord := strings.Fields(zshHistory.Command)[0]
			if firstWord != input.Command {
				continue
			}
		}

		// Filter by contain if defined
		if input.Contains != "" {
			if !strings.Contains(zshHistory.Command, input.Contains) {
				continue
			}
		}

		cmds = append(cmds, *zshHistory)

		if input.Limit > 0 && len(cmds) > input.Limit {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	commands, err := json.Marshal(cmds)
	if err != nil {
		log.Fatal(err)
	}
	return nil, Output{Commands: string(commands)}, nil
}
