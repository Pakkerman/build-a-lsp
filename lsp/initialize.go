package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... and among other truch load of stuff
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilites `json:"capabilities"`
	ServerInfo   ServerInfo        `json:"serverInfo"`
}

type ServerCapabilites struct {
	TextDocumentSync int `json:"textDocumentSync"`

	HoverProvider bool `json:"hoverProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilites{
				TextDocumentSync: 1,
				HoverProvider:    true,
			},
			ServerInfo: ServerInfo{
				Name:    "build-a-lsp",
				Version: "0.0.0.0.0.0-alaph1.final"},
		},
	}
}
