#include<lexer.cpp>
class interpreter{ 
    std::string input;

void lexer_init(string input){ 
    Lexer lexer(input);

    // Tokenize and print the tokens
    Token token = lexer.getNextToken();
    while (token.type != END_OF_FILE) {
        std::cout << "Token: ";
        switch (token.type) {
            case LPAREN: std::cout << "LPAREN"; break;
            case RPAREN: std::cout << "RPAREN"; break;
            case LBRACKET: std::cout << "LBRACKET"; break;
            case RBRACKET: std::cout << "RBRACKET"; break;
            case STRING: std::cout << "STRING, Value: " << token.value; break;
            case FLOAT: std::cout << "FLOAT, Value: " << token.value; break;
            case DATE: std::cout << "DATE, Value: " << token.value; break;
            case COLON: std::cout << "COLON"; break;
            case SEMICOLON: std::cout << "SEMICOLON"; break;
            case BACKSLASH: std::cout << "BACKSLASH"; break;
            case ASTERISK: std::cout << "ASTERISK"; break;
            case ASTERISK_BACKSLASH: std::cout << "ASTERISK_BACKSLASH"; break;
            default: std::cout << "UNKNOWN";
        }
        std::cout << std::endl;

        token = lexer.getNextToken();
    }
}
};
