#include "../lib/L-27/interpreter.cpp"
#include "../lib/Main_comp/relation.cpp"
#include "../lib/Storage/storage.cpp"
int main(){
    interpreter interperter;
    Storage storage;
    Node Joe("people","Joe", "age", 42, "salary", 346.87, "Greeting", "Hello,Joe");
    Node Mia("people","Mia", "age", 22, "salary", 398.87, "Greeting", "Hello,Mia");
    Relation Teach(Joe,Mia,"Teach","Tutoria","time",12,"classroom",202);
    storage.create_folder();
    std::cout<<"creating foler \n";
    storage.create_db("test"); 
    std::cout<<"base de datos test"; 
    while(1){
        std::cout<<"User=>";
        std::ostringstream codeStream;
	    std::string line;
	    std::getline(std::cin, line); 
	    codeStream << line << '\n'; 
        interperter.lexer_init(codeStream.str());
        Joe.show();
        Joe.alter("age",35);
        Joe.remove("salary");
        Joe.add("service",true);
        Joe.show();
        Teach.show();
        Joe.show(); 

    } 
} 
