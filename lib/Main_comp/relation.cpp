#include"node.cpp"
class Relation{
public:
Node  origin; 
Node  dest;
std::string name;
std::string label;
std::unordered_map<std::string, std::any> properties;

template <typename... Args>
Relation(Node origin, Node dest, std::string label, std::string name, Args&&... args)
    : origin(std::move(origin)), dest(std::move(dest)), label(std::move(label)), name(std::move(name)) {
    initializeProperties(std::forward<Args>(args)...);
    }

    void initializeProperties() {}

    template <typename T, typename... Rest>
    void initializeProperties(const std::string& key, T&& value, Rest&&... rest) {
        properties[key] = std::forward<T>(value);
        initializeProperties(std::forward<Rest>(rest)...);
    }

void printElement(const std::any& element) {
    if (element.type() == typeid(std::string)) {
        std::cout << std::any_cast<std::string>(element);
    }else if (element.type() == typeid(const char*)){
        std::cout<<"\""<<std::any_cast<const char*>(element)<<"\""; 
    } else if (element.type() == typeid(int)) {
        std::cout << std::any_cast<int>(element);
    } else if (element.type() == typeid(double)) {
        std::cout << std::any_cast<double>(element);
    } else if (element.type() == typeid(bool)) {
        std::cout << std::any_cast<bool>(element);
    } else {
        std::cerr << "Unsupported type";
    }
}

void show(){
    this->origin.show();
    std::cout<<"=>";
    std::cout << "(" << this->label << ":" << this->name << ")(";
        size_t i = 0;
        for (const auto& [key, value] : this->properties) {
            std::cout << key << ":";
            printElement(value);
            if (++i < this->properties.size()) {
                std::cout << ",";
            }
        }
        std::cout << ")\n";
    std::cout<<"=>";
    this->dest.show();
}
void alter(const std::string& key, std::any value) {
        properties[key] = value;
    }
    void alter_name(std::string name){
        this->name=name;
    } 
    void remove(const std::string& key) {
        properties.erase(key);
    }

    void add(const std::string& key, std::any value) {
        properties[key] = value;
    }
};  
