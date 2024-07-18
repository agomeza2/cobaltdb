#ifndef LEXER_H
#define LEXER_H

#include <iostream>
#include <vector>
#include <string>
#include <cctype>
#include <cstddef>  // For size_t
#include <new>

enum TokenType {
    CREATE, DATABASE, USE, CATEGORY, NODE, RELATION, AS, ASSIGN, SHOW, ALTER, REMOVE, ADD,
    COMMON, SEARCH, DBINT, DBFLOAT, DBDATE, DBSTRING, DBBOOL, GROUP_BY, HOW_MANY, DESC, ASC, WHERE,
    SEMICOLON, DOUBLE_SLASH_COMMENT, SLASH_STAR_COMMENT, COMMA, ARROW, OPEN_PAREN, CLOSE_PAREN,
    OPEN_BRACKET, CLOSE_BRACKET, COLON, IDENTIFIER, INTEGER, FLOAT_LITERAL, DATE_LITERAL,
    STRING_LITERAL, BOOL_LITERAL, NIL, 
    ERROR
};

struct Token {
    TokenType type;
    std::string value;
};

class Lexer {
public:
    Lexer(const std::string& input);

    Token getNextToken();

private:
    const std::string& input;
    size_t currentPosition;

    void advance();
    void skipWhitespace();
    void processSingleLineComment();
    void processMultiLineComment();
    Token processNumber();
    Token processIdentifier();
    Token processStringLiteral();
    Token processSymbol();
    bool isKeyword(const std::string& identifier) const;
    TokenType getKeywordType(const std::string& identifier) const;
};

#endif // LEXER_H