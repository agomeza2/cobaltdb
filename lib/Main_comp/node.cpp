#include <iostream>
#include <any>
#include <vector>
#include <typeinfo>

class Node {
public:
    std::string name;
    std::string label;
    std::vector<std::any> properties;

    template <typename... Args>
    Node(std::string name, std::string label, Args&&... args)
        : name(std::move(name)), label(std::move(label)), properties{std::forward<Args>(args)...} {}


void printElement(const std::any& element) {
    if (element.type() == typeid(int)) {
        int temp = std::any_cast<int>(element);
        std::cout << temp;
    } else if (element.type() == typeid(float)) {
        float temp = std::any_cast<float>(element);
        std::cout << temp;
    }else if (element.type() == typeid(double)){
        float temp = std::any_cast<double>(element);
        std::cout << temp;
	} 
	else if (element.type() == typeid(std::string)) {
        std::cout << std::any_cast<std::string>(element);
    } else if (element.type() == typeid(bool)) {
        bool temp = std::any_cast<bool>(element);
        std::cout << temp;
    } else {
        std::cerr << "Unsupported type" << std::endl;
    }
}
    void show() {
        std::cout << "(" << this->label << ":" << this->name << ")(";
        int i=0;
		for (const auto& item : this->properties) {
            if(i == this->properties.size()-1){
			printElement(item);	
			}else{
			printElement(item);
            std::cout << ",";
			}
			i++;
        }
        std::cout << ")\n";
    }

    void alter(std::any arg){
    }

	void delete(std::any arg){

	}

    void remove(std::any arg){
    }

    void add(std::any item){
    }
};
