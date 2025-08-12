package l27

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var LexerDef = lexer.MustSimple([]lexer.SimpleRule{
	// Palabras clave
	{"Keyword", `(?i)\b(SHOW|COMMON|ALTER|REMOVE|DATABASE|DATABASES|RELATION|IN|HOW|MANY|NODE|SELECT|CREATE|NODES|RELATIONS|IMPORT|EXPORT|LIST|ADD|ADMIN|STANDARD)\b`},

	// Identificadores y literales
	{"Ident", `[a-zA-Z_][a-zA-Z0-9_]*`},
	{"Number", `[0-9]+`},
	{"String", `"(\\"|[^"])*"`},

	// Símbolos y puntuación
	{"Punct", `[\{\}\:\,\(\)]`},

	// Espacios en blanco — IMPORTANTÍSIMO: nil para ignorarlos
	{"Whitespace", `[ \t\n\r]+`},
})
