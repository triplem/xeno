// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// The plugin's declaration and the binary's surface are two halves of one thing, and this is
// the seam they can part at: mcp.json names a command and an argument, and nothing else
// compares them against what the dispatch answers. The vendored set had the same kind of
// seam and #199 is what it cost there.
func TestThePluginStartsTheSubcommandThisBinaryHas(t *testing.T) {
	b, err := os.ReadFile("../../.xeno/plugin/mcp.json")
	if err != nil {
		t.Fatalf("the plugin declares no server: %v", err)
	}
	var declared struct {
		Servers map[string]struct {
			Command string   `json:"command"`
			Args    []string `json:"args"`
		} `json:"mcpServers"`
	}
	if err := json.Unmarshal(b, &declared); err != nil {
		t.Fatalf("mcp.json is not JSON a client can read: %v", err)
	}
	s, ok := declared.Servers["xeno"]
	if !ok || len(declared.Servers) != 1 {
		t.Fatalf("mcp.json declares %d servers and this plugin has one", len(declared.Servers))
	}
	// Through the entry point and not the binary directly, which is section 7's arrangement:
	// the script normalises the environment, and the runner never learns which harness it is.
	if !strings.HasSuffix(s.Command, "bin/xeno-env.sh") {
		t.Errorf("the server is started as %q rather than through the plugin's entry point", s.Command)
	}
	if len(s.Args) != 1 || s.Args[0] != "mcp" {
		t.Fatalf("the declared arguments are %v", s.Args)
	}
	// And the subcommand is answered. Standard input in a test is empty, which is a client
	// that closed the pipe without sending anything: the server stops, writes nothing and
	// exits 0, rather than printing the usage as an unknown command would.
	var out, errw bytes.Buffer
	if code := run(s.Args, &out, &errw); code != 0 {
		t.Errorf("%v exited %d: %s", s.Args, code, errw.String())
	}
	if out.Len() != 0 {
		t.Errorf("the server wrote %q to a transport that was already closed", out.String())
	}
}
