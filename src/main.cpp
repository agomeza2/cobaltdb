#include "../lib/L-27/interpreter.cpp"
#include "../lib/Main_comp/node.cpp"

int main(){
    interpreter interperter;
    Node joe("Joe","people",2,35,40);
    while(1){
        std::cout<<"User=>";
        std::ostringstream codeStream;
	    std::string line;
	    std::getline(std::cin, line); 
	    codeStream << line << '\n'; 
        interperter.lexer_init(codeStream.str());
        joe.show();

    } 
} 