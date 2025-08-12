package l27

import (
	"github.com/alecthomas/participle/v2"
)

// AST / gramática para CobaltDB (extensible)

// Programa: secuencia de sentencias
type Program struct {
	Stmts []*Stmt `@@*`
}

type Stmt struct {
	// Comandos soportados
	Create     *CreateStmt         `  @@`
	CreateUser *CreateUserStmt     `| @@`
	CreateDB   *CreateDatabaseStmt `| @@`
	RemoveDB   *RemoveDatabaseStmt `| @@`
	Use        *UseStmt            `| @@`
	Show       *ShowStmt           `| @@`
	List       *ListStmt           `| @@`
	Alter      *AlterStmt          `| @@`
	RemoveAttr *RemoveAttrStmt     `| @@`
	AddAttr    *AddAttrStmt        `| @@`
	HowMany    *HowManyStmt        `| @@`
	Common     *CommonStmt         `| @@`
	Search     *SearchStmt         `| @@`
	Import     *ImportStmt         `| @@`
	Export     *ExportStmt         `| @@`
}

// CREATE ...
type CreateStmt struct {
	CreateToken string        `"CREATE"`
	What        string        `@("DATABASE" | "NODE" | "NODES" | "RELATION" | "RELATIONS")`
	DBName      *string       `  @Ident?`                   // CREATE DATABASE name
	Obj         *JSONObject   `  @@?`                       // CREATE NODE { ... }
	ArrayObjs   []*JSONObject `( "[" @@ ( "," @@ )* "]" )?` // CREATE NODES [{...},...]
}

type CreateUserStmt struct {
	CreateToken string `"CREATE"`
	UserToken   string `"USER"`
	Name        string `@Ident`
	Role        string `@("ADMIN" | "STANDARD")`
}

type CreateDatabaseStmt struct {
	CreateToken string `"CREATE"`
	DBToken     string `"DATABASE"`
	Name        string `@Ident`
}

type RemoveDatabaseStmt struct {
	RemoveToken string `"REMOVE"`
	DBToken     string `"DATABASE"`
	Name        string `@Ident`
}

type UseStmt struct {
	UseToken string `"USE" | "SELECT"`
	Name     string `@Ident`
}

type ShowStmt struct {
	ShowToken string  `  "SHOW"`
	Type      string  `  @("NODE" | "RELATION" | "DATABASES" | "NODES" | "RELATIONS")`
	Filter    *Filter `@@?`
}

type ListStmt struct {
	ListToken string `  "LIST"`
	What      string `  @("DATABASES" | "NODES" | "RELATIONS")`
}

// ALTER (filter) : { ... }
type AlterStmt struct {
	AlterToken string      `"ALTER"`
	Object     *TargetSpec `@@` // e.g. NODE ID:10 or NODE name:Martyn OR (people:Martyn)
	Colon      *string     `( ":" )?`
	SetBody    *JSONObject `( @@ )?`
}

type TargetSpec struct {
	// allow either: NODE ID:10  OR  (people:Martyn)
	What   *string `  @("NODE" | "RELATION")?`
	Filter *Filter `  @@`
}

// REMOVE <attr> IN NODE name:Martyn
type RemoveAttrStmt struct {
	RemoveToken string  `  "REMOVE"`
	Key         string  `  @Ident`
	_           string  `  "IN"`
	What        string  `  @("NODE" | "RELATION")`
	Filter      *Filter ` @@`
}

// ADD age IN NODE name:Martyn
type AddAttrStmt struct {
	AddToken string  `  "ADD"`
	Key      string  `  @Ident`
	_        string  `  "IN"`
	What     string  `  @("NODE" | "RELATION")`
	Filter   *Filter ` @@`
}

type HowManyStmt struct {
	How   string ` "HOW"`
	Many  string ` "MANY"`
	Ident string ` @Ident`
}

// COMMON name:Martyn name:Lili
type CommonStmt struct {
	Common  string    `"COMMON"`
	Filters []*Filter `@@+`
}

// SEARCH people WHERE age = 20 GROUP BY name ORDER age DESC
type SearchStmt struct {
	SearchToken string     `"SEARCH"`
	Type        string     `@Ident`
	Where       *WhereStmt `@@?`
	GroupBy     *GroupBy   `@@?`
	OrderBy     *OrderBy   `@@?`
}

type WhereStmt struct {
	WhereToken string      `"WHERE"`
	Expr       *Expression `@@`
}

type GroupBy struct {
	GroupToken string `"GROUP"`
	ByToken    string `"BY"`
	Field      string `@Ident`
}

type OrderBy struct {
	OrderToken string  `@"ORDER"?`
	Field      string  `@Ident`
	Direction  *string `( @("ASC" | "DESC") )?`
}

type Expression struct {
	Left  string `@Ident`
	Op    string `@("=" | "!=" | ">" | "<" | ">=" | "<=")`
	Right *Lit   `@@`
}

// IMPORT "path"
type ImportStmt struct {
	Import string  `"IMPORT"`
	Path   *string `@String`
}

// EXPORT DATABASE name [TO "path"]
type ExportStmt struct {
	Export string  `"EXPORT"`
	DBTok  string  `"DATABASE"`
	Name   string  ` @Ident`
	ToTok  *string `( "TO" @String )?`
}

// FILTER e.g. ID:10 or name:Martyn or (people:Martyn)
type Filter struct {
	// Allow either parenthesized (key:value) or key:value
	LParen *string `  "("?`
	Key    string  `  @Ident ":"`
	Value  *Lit    `  @@`
	RParen *string `  ")?"`
}

type JSONObject struct {
	LBrace string    `"{"`
	Pairs  []*KVPair `(@@ ("," @@)*)?`
	RBrace string    `"}"`
}

type KVPair struct {
	Key   string `@Ident ":"`
	Value *Lit   `@@`
}

type Lit struct {
	Number *int    `  @Number`
	Str    *string `| @String`
	Ident  *string `| @Ident`
}

// Constructor del parser
func NewParser() (*participle.Parser[Program], error) {
	return participle.Build[Program](
		participle.Lexer(LexerDef),
		participle.Unquote("String"),
		participle.CaseInsensitive("Keyword"),
	)
}
