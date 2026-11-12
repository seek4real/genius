package token

import "testing"

// 表驱动测试：验证各 token 常量的值
// 运行方式：在项目根目录执行 go test ./token/ -v
func TestTokenConstants(t *testing.T) {
	tests := []struct {
		name     string
		token    TokenType
		expected string
	}{
		{"ILLEGAL 常量", ILLEGAL, "ILLEGAL"},
		{"EOF 常量", EOF, "EOF"},
		{"IDENT 常量", IDENT, "IDENT"},
		{"INT 常量", INT, "INT"},
		{"ASSIGN 常量", ASSIGN, "="},
		{"PLUS 常量", PLUS, "+"},
		{"COMMA 常量", COMMA, ","},
		{"SEMICOLON 常量", SEMICOLON, ";"},
		{"LRAREN 常量", LPAREN, "("},
		{"RPAREN 常量", RPAREN, ")"},
		{"LBRACE 常量", LBRACE, "{"},
		{"RBRACE 常量", RBRACE, "}"},
		{"FUNCTION 常量", FUNCTION, "FUNCTION"},
		{"LET 常量", LET, "LET"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.token != TokenType(tt.expected) {
				t.Errorf("常量值不对: got=%q, want=%q", tt.token, tt.expected)
			}
		})
	}
}
