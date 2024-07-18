#include"lexer.cpp"
#include <iostream>
#include <vector>
#include <stdexcept>
#include <string>

class Parser {
public:
    Parser(Lexer& lexer) : lexer(lexer), currentToken(lexer.getNextToken()) {}

    void parse() {
        while (currentToken.type != ERROR) {
            parseStatement();
        }
    }

private:
    Lexer& lexer;
    Token currentToken;

    void advance() {
        currentToken = lexer.getNextToken();
    }

    void expect(TokenType expectedType) {
        if (currentToken.type == expectedType) {
            advance();
        } else {
            throw std::runtime_error("Syntax error: Unexpected token " + currentToken.value);
        }
    }

    void parseStatement() {
        switch (currentToken.type) {
            case CREATE:
                parseCreateStatement();
                break;
            case USE:
                parseUseStatement();
                break;
            case SHOW:
                parseShowStatement();
                break;
            case ALTER:
                parseAlterStatement();
                break;
            case REMOVE:
                parseRemoveStatement();
                break;
            case ADD:
                parseAddStatement();
                break;
            default:
                throw std::runtime_error("Syntax error: Unexpected token " + currentToken.value);
        }
    }

    void parseCreateStatement() {
        expect(CREATE);
        if (currentToken.type == DATABASE) {
            parseCreateDatabase();
        } else if (currentToken.type == CATEGORY) {
            parseCreateCategory();
        } else {
            throw std::runtime_error("Syntax error: Expected DATABASE or CATEGORY after CREATE");
        }
        expect(SEMICOLON);
    }

    void parseCreateDatabase() {
        expect(DATABASE);
        expect(IDENTIFIER);
    }

    void parseCreateCategory() {
        expect(CATEGORY);
        expect(IDENTIFIER);

    }

    void parseUseStatement() {
        expect(USE);
        expect(DATABASE);
        expect(IDENTIFIER);

        expect(SEMICOLON);
    }

    void parseShowStatement() {
        expect(SHOW);

        expect(SEMICOLON);
    }

    void parseAlterStatement() {
        expect(ALTER);

        expect(SEMICOLON);
    }

    void parseRemoveStatement() {
        expect(REMOVE);

        expect(SEMICOLON);
    }

    void parseAddStatement() {
        expect(ADD);

        expect(SEMICOLON);
    }


};
