#include <iostream>
#include <vector>
#include <sstream>
#include <cctype>
#include <stdexcept>
#include <algorithm>
#include <string>
enum TokenType {
    CREATE, DATABASE, USE, CATEGORY, NODE, RELATION, AS, ASSIGN, SHOW, ALTER, REMOVE, ADD,
    COMMON, SEARCH, INT, FLOAT, DATE, STRING, BOOL, GROUP_BY, HOW_MANY, DESC, ASC, WHERE,
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
    Lexer(const std::string& input) : input(input), currentPosition(0) {}

    Token getNextToken() {
        skipWhitespace();

        if (currentPosition >= input.size()) {
            return {ERROR, "EOF"};
        }

        char currentChar = input[currentPosition];

        if (currentChar == '/' && currentPosition + 1 < input.size()) {
            if (input[currentPosition + 1] == '/') {
                processSingleLineComment();
                return getNextToken(); // Recursively get the next token after a comment
            } else if (input[currentPosition + 1] == '*') {
                processMultiLineComment();
                return getNextToken(); // Recursively get the next token after a comment
            }
        }

        if (isdigit(currentChar) || (currentChar == '-' && isdigit(input[currentPosition + 1]))) {
            return processNumber();
        } else if (isalpha(currentChar) || currentChar == '_') {
            return processIdentifier();
        } else if (currentChar == '\"') {
            return processStringLiteral();
        } else {
            return processSymbol();
        }
    }

private:
    const std::string& input;
    size_t currentPosition;

    void advance() {
        currentPosition++;
    }

    void skipWhitespace() {
        while (currentPosition < input.size() && isspace(input[currentPosition])) {
            advance();
        }
    }

    void processSingleLineComment() {
        while (currentPosition < input.size() && input[currentPosition] != '\n') {
            advance();
        }
    }

    void processMultiLineComment() {
        advance(); // Skip the opening /*
        while (currentPosition + 1 < input.size() &&
               (input[currentPosition] != '*' || input[currentPosition + 1] != '/')) {
            advance();
        }
        advance(); // Skip the closing */
        advance();
    }

    Token processNumber() {
        std::string number;
        while (currentPosition < input.size() &&
               (isdigit(input[currentPosition]) || input[currentPosition] == '.')) {
            number += input[currentPosition++];
        }

        // Check if it's a floating-point number
        if (number.find('.') != std::string::npos) {
            return {FLOAT_LITERAL, number};
        } else {
            return {INTEGER, number};
        }
    }

    Token processIdentifier() {
        std::string identifier;
        while (currentPosition < input.size() &&
               (isalnum(input[currentPosition]) || input[currentPosition] == '_')) {
            identifier += input[currentPosition++];
        }

        // Check if it's a keyword or boolean literal
        if (identifier == "TRUE" || identifier == "FALSE") {
            return {BOOL_LITERAL, identifier};
        } else {
            // Check if it's a keyword
            if (isKeyword(identifier)) {
                return {getKeywordType(identifier), identifier};
            } else {
                return {IDENTIFIER, identifier};
            }
        }
    }

    Token processStringLiteral() {
        std::string literal;
        currentPosition++; // Skip the opening quote
        while (currentPosition < input.size() && input[currentPosition] != '\"') {
            literal += input[currentPosition++];
        }
        currentPosition++; // Skip the closing quote
        return {STRING_LITERAL, literal};
    }

    Token processSymbol() {
        char currentChar = input[currentPosition++];
        switch (currentChar) {
            case ';': return {SEMICOLON, ";"};
            case ',': return {COMMA, ","};
            case '(': return {OPEN_PAREN, "("};
            case ')': return {CLOSE_PAREN, ")"};
            case '[': return {OPEN_BRACKET, "["};
            case ']': return {CLOSE_BRACKET, "]"};
            case ':': return {COLON, ":"};
            case '-':
                if (currentPosition < input.size() && isdigit(input[currentPosition])) {
                    return processNumber(); // Negative number
                } else {
                    return {ERROR, "Unexpected character: -"};
                }
            case '=':
                if (currentPosition < input.size() && input[currentPosition] == '>') {
                    advance(); // Skip '>'
                    return {ARROW, "=>"};
                } else {
                    return {ASSIGN, "="};
                }
            default: return {ERROR, "Unexpected character: " + std::string(1, currentChar)};
        }
    }

    bool isKeyword(const std::string& identifier) const {
        static const std::vector<std::string> keywords = {
            "CREATE", "DATABASE", "USE", "CATEGORY", "NODE", "RELATION", "AS", "SHOW",
            "ALTER", "REMOVE", "ADD", "COMMON", "SEARCH", "INT", "FLOAT", "DATE", "STRING",
            "BOOL", "GROUP_BY", "HOW_MANY", "DESC", "ASC", "WHERE"
        };
        return std::find(keywords.begin(), keywords.end(), identifier) != keywords.end();
    }

    TokenType getKeywordType(const std::string& identifier) const {
        // Implement a mapping from keyword string to TokenType
        // Add more cases as needed
        if (identifier == "CREATE") return CREATE;
        else if (identifier == "DATABASE") return DATABASE;
        else if (identifier == "USE") return USE;
        else if (identifier == "CATEGORY") return CATEGORY;
        else if (identifier == "NODE") return NODE;
        else if (identifier == "RELATION") return RELATION;
        else if (identifier == "AS") return AS;
        else if (identifier == "SHOW") return SHOW;
        else if (identifier == "ALTER") return ALTER;
        else if (identifier == "REMOVE") return REMOVE;
        else if (identifier == "ADD") return ADD;
        else if (identifier == "COMMON") return COMMON;
        else if (identifier == "SEARCH") return SEARCH;
        else if (identifier == "INT") return INT;
        else if (identifier == "FLOAT") return FLOAT;
        else if (identifier == "DATE") return DATE;
        else if (identifier == "STRING") return STRING;
        else if (identifier == "BOOL") return BOOL;
        else if (identifier == "GROUP_BY") return GROUP_BY;
        else if (identifier == "HOW_MANY") return HOW_MANY;
        else if (identifier == "DESC") return DESC;
        else if (identifier == "ASC") return ASC;
        else if (identifier == "WHERE") return WHERE;
        else return ERROR; // Default case, should not happen
    }
};
