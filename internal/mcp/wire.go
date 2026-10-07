// SPDX-License-Identifier: Apache-2.0

package mcp

import "encoding/json"

// The JSON-RPC 2.0 envelope and the few result shapes of MCP 2025-06-18 this server
// produces, written out rather than taken from an SDK for the reason the package comment
// gives. They are here and not beside the tools so that the protocol is readable as one
// thing and the operations as another.

// request is one incoming message. The id is kept as raw JSON because the protocol allows a
// string or a number and a response has to carry back exactly what came in; decoding it into
// a Go type would turn 1 into 1 and 1.0 into the same thing, and the client matches on it.
//
// A request with no id is a notification, which is answered with silence.
type request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

// response is a result. Result is an interface rather than raw JSON so that each handler
// returns the shape it means and marshalling happens once, at the write.
type response struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result"`
}

type failure struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Error   rpcError        `json:"error"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// The JSON-RPC 2.0 codes this server uses, and only those. Nothing here invents a code: a
// tool that fails answers inside a result, so the protocol level errors are the three ways a
// request can fail to be a request.
const (
	codeParse     = -32700
	codeNoMethod  = -32601
	codeBadParams = -32602
)

func resultFor(id json.RawMessage, result any) response {
	return response{JSONRPC: "2.0", ID: id, Result: result}
}

// errorFor answers a request that could not be served. A null id is what the protocol asks
// for when the id could not be read, which is the parse error's case.
func errorFor(id json.RawMessage, code int, message string) failure {
	if id == nil {
		id = json.RawMessage("null")
	}
	return failure{JSONRPC: "2.0", ID: id, Error: rpcError{Code: code, Message: message}}
}

type initializeResult struct {
	ProtocolVersion string       `json:"protocolVersion"`
	Capabilities    capabilities `json:"capabilities"`
	ServerInfo      serverInfo   `json:"serverInfo"`
}

// capabilities is tools and nothing else, as a pointer so that an absent capability is an
// absent key rather than an empty object claiming the capability.
type capabilities struct {
	Tools *struct{} `json:"tools,omitempty"`
}

type serverInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type toolsResult struct {
	Tools []definition `json:"tools"`
}

// definition is what tools/list sends per operation, and the budget section 13 sets is a
// budget on these: name, one sentence, and a schema. Every one of them is in front of the
// model on every request of every phase whether it is called or not.
type definition struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema json.RawMessage `json:"inputSchema"`
}

type callResult struct {
	Content []content `json:"content"`
	IsError bool      `json:"isError,omitempty"`
}

type content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func textContent(s string) []content { return []content{{Type: "text", Text: s}} }
