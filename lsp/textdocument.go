package lsp

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    string `json:"version"`
	Text       string `json:"text"`
}
