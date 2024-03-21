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

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// in a fully featured LSP this would look up definition

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
