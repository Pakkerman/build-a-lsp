package analysis

import (
	"build-a-lsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	// map<filenames, content>
	Document map[string]string
}

func NewState() State {
	return State{Document: map[string]string{}}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "Common sense",
				Message:  "c'mon man, have some common sense",
			})
		}
	}

	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS C*de") {
			idx := strings.Index(line, "VS C*de")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS C*de")),
				Severity: 2,
				Source:   "Common sense",
				Message:  "Still not ideal, please change to some superior editor",
			})
		}
	}

	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Neovim")),
				Severity: 4,
				Source:   "Common sense",
				Message:  "Noice",
			})
		}
	}

	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Document[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {

	s.Document[uri] = text

	return getDiagnosticsForFile(text)
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

func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Document[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS Code with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{

					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Censor VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}

	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
	return response
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor and blazingly fast",
			Documentation: "Yayayayayayayaaaaaaaaas",
		}}

	// Ask your static analysis tools to figure out good completions

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{

			Line:      line,
			Character: end,
		},
	}
}
