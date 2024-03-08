#include <iostream>
#include <cctype>
#include <vector>

// Token types
enum TokenType {
    LPAREN,   // (
    RPAREN,   // )
    LBRACKET, // [
    RBRACKET, // ]
    STRING,
    FLOAT,
    DATE,
    COLON,    // :
    SEMICOLON, // ;
    BACKSLASH, // \\
    ASTERISK,  // *
    ASTERISK_BACKSLASH, // *\

    END_OF_FILE
};

// Token structure
struct Token {
    TokenType type;
    std::string value; // For STRING token
};

class Lexer {
public:
    Lexer(const std::string& input) : input(input), position(0) {}

    // Get the next token from the input
    Token getNextToken() {
        while (position < input.length()) {
            char currentChar = input[position];

            if (currentChar == '(') {
                position++;
                return {LPAREN, ""};
            } else if (currentChar == ')') {
                position++;
                return {RPAREN, ""};
            } else if (currentChar == '[') {
                position++;
                return {LBRACKET, ""};
            } else if (currentChar == ']') {
                position++;
                return {RBRACKET, ""};
            } else if (currentChar == ':') {
                position++;
                return {COLON, ""};
            } else if (currentChar == ';') {
                position++;
                return {SEMICOLON, ""};
            } else if (currentChar == '\\') {
                position++;
                if (position < input.length() && input[position] == '*') {
                    position++;
                    return {ASTERISK_BACKSLASH, ""};
                } else {
                    return {BACKSLASH, ""};
                }
            } else if (currentChar == '*') {
                position++;
                return {ASTERISK, ""};
            } else if (currentChar == '"') {
                return parseString();
            } else if (isalpha(currentChar)) {
                return parseWord();
            } else if (isdigit(currentChar) || currentChar == '.') {
                return parseFloatOrDate();
            } else if (isspace(currentChar)) {
                // Skip whitespace
                position++;
            } else {
                // Handle unknown characters
                std::cerr << "Error: Unexpected character '" << currentChar << "'" << std::endl;
                exit(1);
            }
        }

        // End of file reached
        return {END_OF_FILE, ""};
    }

private:
    // Parse a string enclosed in double quotes
    Token parseString() {
        position++; // Skip the opening double quote
        size_t startPos = position;

        while (position < input.length() && input[position] != '"') {
            position++;
        }

        if (position == input.length()) {
            std::cerr << "Error: Unterminated string literal" << std::endl;
            exit(1);
        }

        std::string value = input.substr(startPos, position - startPos);
        position++; // Skip the closing double quote
        return {STRING, value};
    }

    // Parse a word (e.g., keyword)
    Token parseWord() {
        size_t startPos = position;

        while (position < input.length() && (isalnum(input[position]) || input[position] == '_')) {
            position++;
        }

        std::string value = input.substr(startPos, position - startPos);

        // Check if the word represents a keyword
        if (value == "float") {
            return {FLOAT, value};
        } else if (value == "date") {
            return {DATE, value};
        } else {
            return {STRING, value};
        }
    }

    // Parse a float or date
    Token parseFloatOrDate() {
        size_t startPos = position;

        while (position < input.length() && (isdigit(input[position]) || input[position] == '.')) {
            position++;
        }

        std::string value = input.substr(startPos, position - startPos);

        // Check if the value represents a float or date
        if (value.find('.') != std::string::npos) {
            return {FLOAT, value};
        } else {
            return {DATE, value};
        }
    }

    std::string input;
    size_t position;
};

