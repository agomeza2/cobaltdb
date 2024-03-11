#include"lexer.cpp"
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
void parser_init(){
} 
};
