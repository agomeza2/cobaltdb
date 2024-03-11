#include <iostream>
#include <vector>
#include <sstream>

enum TokenType {
    SEARCH, WHERE, CATEGORY, SHOW, GRAPH, FUNC, CREATE, INT, FLOAT, STRING, BOOL,
    DATE, COMMON, WHICH, FOR, WHILE, LOOP, IF, ELSE, ELIF, ALTER, REMOVE, DATABASE,
    DATABASES, USE, RELATION, NIL, AND, NOT, EQ, OR, TRUE, FALSE, GROUP_BY, HOW_MANY,
    RETURN, ADD, OPEN_PAREN, CLOSE_PAREN, BACKSLASH, OPEN_BRACKET, CLOSE_BRACKET,
    COMMA, IDENTIFIER, NUMBER, STRING_LITERAL, COMMENT, ERROR
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

        // Check for comments
        if (currentChar == '/') {
            if (currentPosition + 1 < input.size() && input[currentPosition + 1] == '/') {
                processSingleLineComment();
                return getNextToken(); // Recursively get the next token after a comment
            } else if (currentPosition + 1 < input.size() && input[currentPosition + 1] == '*') {
                processMultiLineComment();
                return getNextToken(); // Recursively get the next token after a comment
            }
        }

        // Check for symbols
        switch (currentChar) {
            case '(': advance(); return {OPEN_PAREN, "("};
            case ')': advance(); return {CLOSE_PAREN, ")"};
            case '\\': advance(); return {BACKSLASH, "\\"};
            case '[': advance(); return {OPEN_BRACKET, "["};
            case ']': advance(); return {CLOSE_BRACKET, "]"};
            case ',': advance(); return {COMMA, ","};
            default: break;
        }

        if (isalpha(currentChar) || currentChar == '_') {
            return processIdentifier();
        } else if (isdigit(currentChar)) {
            return processNumber();
        } else if (currentChar == '\"') {
            return processStringLiteral();
        } else {
            return processOperator();
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

    Token processIdentifier() {
        std::string identifier;
        while (currentPosition < input.size() &&
               (isalnum(input[currentPosition]) || input[currentPosition] == '_')) {
            identifier += input[currentPosition++];
        }

        // Check if the identifier is a keyword
        if (identifier == "SEARCH") return {SEARCH, identifier};
        else if (identifier == "WHERE") return {WHERE, identifier};
        else if (identifier == "CATEGORY") return {CATEGORY, identifier};
        else if (identifier == "SHOW") return {SHOW, identifier};
        else if (identifier == "GRAPH") return {GRAPH, identifier};
        else if (identifier == "FUNC") return {FUNC, identifier};
        else if (identifier == "CREATE") return {CREATE, identifier};
        else if (identifier == "INT") return {INT, identifier};
        else if (identifier == "FLOAT") return {FLOAT, identifier};
        else if (identifier == "STRING") return {STRING, identifier};
        else if (identifier == "BOOL") return {BOOL, identifier};
        else if (identifier == "DATE") return {DATE, identifier};
        else if (identifier == "COMMON") return {COMMON, identifier};
        else if (identifier == "WHICH") return {WHICH, identifier};
        else if (identifier == "FOR") return {FOR, identifier};
        else if (identifier == "WHILE") return {WHILE, identifier};
        else if (identifier == "LOOP") return {LOOP, identifier};
        else if (identifier == "IF") return {IF, identifier};
        else if (identifier == "ELSE") return {ELSE, identifier};
        else if (identifier == "ELIF") return {ELIF, identifier};
        else if (identifier == "ALTER") return {ALTER, identifier};
        else if (identifier == "ADD") return {ADD, identifier};
        else if (identifier == "REMOVE") return {REMOVE, identifier};
        else if (identifier == "DATABASE") return {DATABASE, identifier};
        else if (identifier == "DATABASES") return {DATABAES,identifier};
}
 Token processNumber() {
        std::string number;
        while (currentPosition < input.size() && (isdigit(input[currentPosition]) || input[currentPosition] == '.')) {
            number += input[currentPosition++];
        }

        // Check if it's a floating-point number
        if (number.find('.') != std::string::npos) {
            return {FLOAT, number};
        } else {
            return {INT, number};
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

    Token processDate() {
        std::string date;
        while (currentPosition < input.size() && (isdigit(input[currentPosition]) || input[currentPosition] == '/')) {
            date += input[currentPosition++];
        }

        return {DATE, date};
    }
    };
