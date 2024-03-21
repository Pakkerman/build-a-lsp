package analysis

import (
	"build-a-lsp/lsp"
	"fmt"
)

type State struct {
	// map<filenames, content>
	Document map[string]string
}

func NewState() State {
	return State{Document: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// in a fully featured LSP this would look up the type in our type analysis code ...

	document := s.Document[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
		},
	}
}
