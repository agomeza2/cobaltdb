#include "../lib/L-27/interpreter.cpp"
#include "../lib/Storage/data_process.cpp"
#include "../lib/Storage/storage.cpp"
#include "../lib/Main_comp/users.cpp"
int main(){
    interpreter interperter;
    Storage storage;
    DataProcess data_process;
    Node Joe("people","Joe", "age", 42, "salary", 346.87, "Greeting", "Hello,Joe");
    Node Mia("people","Mia", "age", 22, "salary", 398.87, "Greeting", "Hello,Mia");
    Node Jose("people","Jose", "age",19, "salary", 98.87, "Greeting", "Hello,Jose");
    Node Makoto("people","Makoto", "age", 77, "salary", 88.87, "Greeting", "Hello,Makoto");
    Node Chito("animal","Chito","age",17,"spicies","feline","animal type","cat");
    Relation Teach(Joe,Mia,"Teach","Tutoria","time",12,"classroom",202);
    Relation Teach2(Joe,Mia,"Teach","Tutoria2","time",13,"classroom",205);
    Relation MakoChito(Makoto,Chito,"Owner","only owner","time",17,"place","yokohama");
    storage.create_user("Alex");
    std::cout<<"creating foler \n";
    storage.create_db("test","Alex"); 
    storage.create_db("DB","Alex");
    std::cout<<"base de datos test \n"; 
    std::string db_path = "../db/Alex/test";
    Mia.writeToJsonFile(db_path,Mia.name);
    Joe.writeToJsonFile(db_path,Joe.name);
    Jose.writeToJsonFile(db_path,Jose.name); 
    Makoto.writeToJsonFile(db_path,Makoto.name);
    Teach.writeToJsonFile(db_path,Teach.name);
    Teach2.writeToJsonFile(db_path,Teach2.name);
    Chito.writeToJsonFile(db_path,Chito.name);
    MakoChito.writeToJsonFile(db_path,MakoChito.name);
    data_process.readExcelToNodes("Project-Management-Sample-Data..xlsx",db_path);

    std::unordered_map<std::string, User*> users;

    users["admin"] = new AdminUser("admin", "admin123");
    users["user1"] = new StandardUser("user1", "user123");
    users["user2"] = new StandardUser("user2", "pass123");

    std::string username;
    std::string password;

    std::cout << "Enter username: ";
    std::cin >> username;
    std::cout << "Enter password: ";
    std::cin >> password;

    auto it = users.find(username);
    if (it != users.end() && it->second->authenticate(password)) {
        std::cout << "Authentication successful." << std::endl;
        it->second->displayInfo();
    } else {
        std::cout << "Incorrect username or password." << std::endl;
    }

    for (auto& pair : users) {
        delete pair.second;
    }
    while(1){
        std::cout<<"User=>";
        std::ostringstream codeStream;
	    std::string line;
	    std::getline(std::cin, line); 
	    codeStream << line << '\n'; 
        interperter.parser_init(codeStream.str()); 

    } 
} 
