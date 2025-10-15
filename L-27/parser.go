package l27

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Comando representa un comando parseado con sus parámetros
type Command struct {
	Name       string
	Args       []string
	Properties []map[string]interface{} // para insert/modify JSON-like
	Filters    []map[string]interface{} // para where / for
	RawQuery   []string                 // para find con varios filtros
}

// Parser estructura principal
type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expect(t TokenType) error {
	if p.curToken.Type != t {
		return fmt.Errorf("expected %v, got %v", t, p.curToken.Type)
	}
	return nil
}

func (p *Parser) ParseCommand() (*Command, error) {
	if p.curToken.Type != IDENT {
		return nil, fmt.Errorf("expected command name, got %s", p.curToken.Value)
	}
	cmd := &Command{Name: p.curToken.Value}
	p.nextToken()

	switch cmd.Name {
	case "ls":
		if err := p.parseLs(cmd); err != nil {
			return nil, err
		}
	case "cd":
		if err := p.parseCd(cmd); err != nil {
			return nil, err
		}
	case "touch":
		if err := p.parseTouch(cmd); err != nil {
			return nil, err
		}

	case "modify":
		if err := p.parseModify(cmd); err != nil {
			return nil, err
		}
	case "rm":
		if err := p.parseRm(cmd); err != nil {
			return nil, err
		}
	case "grep":
		if err := p.parseGrep(cmd); err != nil {
			return nil, err
		}
	case "import":
		if err := p.parseImport(cmd); err != nil {
			return nil, err
		}
	case "export":
		if err := p.parseExport(cmd); err != nil {
			return nil, err
		}
	case "common":
		if err := p.parseCommon(cmd); err != nil {
			return nil, err
		}
	case "cat":
		if err := p.parseCat(cmd); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown command %s", cmd.Name)
	}
	return cmd, nil
}

// Parse funciones individuales por comando

func (p *Parser) parseLs(cmd *Command) error {
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()
	return nil
}

func (p *Parser) parseCd(cmd *Command) error {
	// Agrega el comando "cd"
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()

	// Solo un argumento permitido: el nombre de la base de datos
	if p.curToken.Type == IDENT {
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.nextToken()
	} else {
		return fmt.Errorf("cd requiere un nombre de base de datos")
	}

	return nil
}

func (p *Parser) parseTouch(cmd *Command) error {
	if p.curToken.Value != "touch" {
		return fmt.Errorf("se esperaba 'touch', se encontró %s", p.curToken.Value)
	}
	cmd.Args = append(cmd.Args, p.curToken.Value) // opcional: guardar "touch"
	p.nextToken()

	var tipo string
	switch p.curToken.Value {
	case "node", "relation":
		tipo = p.curToken.Value
		cmd.Args = append(cmd.Args, tipo) // opcional: guardar tipo en Args
		p.nextToken()
	default:
		return fmt.Errorf("tipo desconocido %s, se esperaba node o relation", p.curToken.Value)
	}
	if p.curToken.Type != IDENT {
		return fmt.Errorf("se esperaba nombre de %s, se encontró %s", tipo, p.curToken.Value)
	}
	return nil
}

func (p *Parser) parseModify(cmd *Command) error {
	// modify {age:30} for {id:1} in document

	// propiedades a modificar
	props, err := p.parseProps()
	if err != nil {
		return err
	}
	cmd.Properties = props

	// filtro after "for"
	if p.curToken.Type != IDENT || p.curToken.Value != "for" {
		return errors.New("expected 'for' after modify properties")
	}
	p.nextToken()

	filters, err := p.parseProps()
	if err != nil {
		return err
	}
	cmd.Filters = filters

	// espera in document name
	if p.curToken.Type != IDENT || p.curToken.Value != "in" {
		return errors.New("expected 'in' after for clause")
	}
	p.nextToken()

	if p.curToken.Type != IDENT || p.curToken.Value != "document" {
		return errors.New("expected 'document' after 'in'")
	}
	cmd.Args = append(cmd.Args, "document")
	p.nextToken()

	if p.curToken.Type != IDENT {
		return errors.New("expected document name after 'document'")
	}
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()
	return nil
}

func (p *Parser) parseRm(cmd *Command) error {
	if p.curToken.Type != IDENT {
		return errors.New("expected db/node/relation after rm")
	}
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()

	return nil
}

func (p *Parser) parseGrep(cmd *Command) error {

	for p.curToken.Type == STRING {
		cmd.RawQuery = append(cmd.RawQuery, p.curToken.Value)
		p.nextToken()
	}

	for p.curToken.Type == IDENT {
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.nextToken()
	}
	return nil
}

func (p *Parser) parseImport(cmd *Command) error {
	// import filename_path
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()

	// Solo un argumento permitido: el nombre de la base de datos
	if p.curToken.Type == IDENT {
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.nextToken()
	} else {
		return fmt.Errorf("cd requiere un nombre de base de datos")
	}

	return nil
}

func (p *Parser) parseExport(cmd *Command) error {
	// export db/collection/document filename_path
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()

	// Solo un argumento permitido: el nombre de la base de datos
	if p.curToken.Type == IDENT {
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.nextToken()
	} else {
		return fmt.Errorf("cd requiere un nombre de base de datos")
	}

	return nil
}
func (p *Parser) parseCommon(cmd *Command) error {
	// Guardamos el nombre del comando
	cmd.Args = append(cmd.Args, p.curToken.Value) // "common"
	p.nextToken()

	// ID_source
	if p.curToken.Type != NUMBER {
		return fmt.Errorf("se esperaba ID_source numérico, se encontró %s", p.curToken.Value)
	}
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()

	// ID_target
	if p.curToken.Type != NUMBER {
		return fmt.Errorf("se esperaba ID_target numérico, se encontró %s", p.curToken.Value)
	}
	cmd.Args = append(cmd.Args, p.curToken.Value)
	p.nextToken()

	return nil
}

func (p *Parser) parseCat(cmd *Command) error {
	// Guardamos el nombre del comando
	cmd.Args = append(cmd.Args, p.curToken.Value) // "cat"
	p.nextToken()

	if p.curToken.Type != IDENT {
		return fmt.Errorf("se esperaba Node, Relation, Nodes o Relations, se encontró %s", p.curToken.Value)
	}

	switch p.curToken.Value {
	case "node", "relation":
		// Variante con ID: cat Node ID_node
		objType := p.curToken.Value
		cmd.Args = append(cmd.Args, objType)
		p.nextToken()

		if p.curToken.Type != NUMBER {
			return fmt.Errorf("se esperaba ID numérico para %s, se encontró %s", objType, p.curToken.Value)
		}
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.nextToken()

	case "nodes", "relations":
		// Variante sin ID: cat Nodes o cat Relations
		cmd.Args = append(cmd.Args, p.curToken.Value)
		p.nextToken()

	default:
		return fmt.Errorf("tipo desconocido %s, se esperaba Node, Relation, Nodes o Relations", p.curToken.Value)
	}

	return nil
}

// parseProps parsea estructuras tipo JSON simples {key:value,...} o listas [{...},{...}]
func (p *Parser) parseProps() ([]map[string]interface{}, error) {
	if p.curToken.Type == LBRACKET {
		// lista de objetos
		p.nextToken()
		var objs []map[string]interface{}
		for {
			if p.curToken.Type == RBRACKET {
				p.nextToken()
				break
			}
			if p.curToken.Type != LBRACE {
				return nil, errors.New("expected '{' in array of objects")
			}
			obj, err := p.parseSingleProp()
			if err != nil {
				return nil, err
			}
			objs = append(objs, obj)
			if p.curToken.Type == COMMA {
				p.nextToken()
				continue
			} else if p.curToken.Type == RBRACKET {
				p.nextToken()
				break
			} else {
				return nil, errors.New("expected ',' or ']' in array")
			}
		}
		return objs, nil
	} else if p.curToken.Type == LBRACE {
		// objeto único
		obj, err := p.parseSingleProp()
		if err != nil {
			return nil, err
		}
		return []map[string]interface{}{obj}, nil
	}
	return nil, errors.New("expected '{' or '[' to start properties")
}

// parseSingleProp parsea un objeto simple {key:value,...}
func (p *Parser) parseSingleProp() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if p.curToken.Type != LBRACE {
		return nil, errors.New("expected '{' at start of object")
	}
	p.nextToken()
	for p.curToken.Type != RBRACE && p.curToken.Type != EOF {
		if p.curToken.Type != IDENT && p.curToken.Type != STRING {
			return nil, errors.New("expected key in object")
		}
		key := p.curToken.Value
		p.nextToken()

		if p.curToken.Type != COLON {
			return nil, errors.New("expected ':' after key in object")
		}
		p.nextToken()

		val, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		m[key] = val

		if p.curToken.Type == COMMA {
			p.nextToken()
		} else if p.curToken.Type == RBRACE {
			break
		} else {
			return nil, errors.New("expected ',' or '}' in object")
		}
	}
	if p.curToken.Type != RBRACE {
		return nil, errors.New("expected '}' at end of object")
	}
	p.nextToken()
	return m, nil
}

func (p *Parser) parseValue() (interface{}, error) {
	switch p.curToken.Type {
	case STRING:
		val := p.curToken.Value
		p.nextToken()
		return val, nil
	case NUMBER:
		val := p.curToken.Value
		p.nextToken()
		if strings.Contains(val, ".") {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
			return f, nil
		} else {
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			return i, nil
		}
	case IDENT:
		// booleanos true/false o null
		switch p.curToken.Value {
		case "true":
			p.nextToken()
			return true, nil
		case "false":
			p.nextToken()
			return false, nil
		case "null":
			p.nextToken()
			return nil, nil
		default:
			// lo tratamos como string
			val := p.curToken.Value
			p.nextToken()
			return val, nil
		}
	default:
		return nil, fmt.Errorf("unexpected value token %v", p.curToken)
	}
}
