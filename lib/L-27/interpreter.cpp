#include"parser.cpp"
class interpreter{  
public:
void lexer_init(std::string input){
	 
	     Lexer lexer(input);
		 std::cout<<"lexer lexing";
	 
	    
	     Token token;
	     do {
	         token = lexer.getNextToken();
	         std::cout << "Token Type: " << token.type << ", Value: " << token.value << std::endl;
	     } while (token.type != ERROR && token.type != EOF);
	 
    }
void parser_init(std::string input ){
	Lexer lexer(input);
    Parser parser(lexer);
      try {
        parser.parse();
        std::cout << "Parsing completed successfully." << std::endl;
    } catch (const std::exception& e) {
        std::cerr << e.what() << std::endl;
    }
} 
};
