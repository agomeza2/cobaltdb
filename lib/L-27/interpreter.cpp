#include<lexer.cpp>
class interpreter{  

void lexer_init(string input){
	 std::ostringstream codeStream;
	 std::string line;
	     while (std::getline(std::cin, line)) {
	         codeStream << line << '\n';
	     }
	 
	     Lexer lexer(codeStream.str());
	 
	    
	     Token token;
	     do {
	         token = lexer.getNextToken();
	         std::cout << "Token Type: " << token.type << ", Value: " << token.value << std::endl;
	     } while (token.type != ERROR && token.type != EOF);
	 
    }
void parser_init(){
} 
};
