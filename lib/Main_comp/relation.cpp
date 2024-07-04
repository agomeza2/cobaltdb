#include"node.cpp"
class Relation{
public:
Node  origin; 
Node  dest;
std::string name;
std::string category;
int id;
std::unordered_map<std::string, std::any> properties;

template <typename... Args>
Relation(Node origin, Node dest, std::string category, std::string name, Args&&... args)
    : origin(std::move(origin)), dest(std::move(dest)), category(std::move(category)), name(std::move(name)) {
    initializeProperties(std::forward<Args>(args)...);
    id=get_next_id(category);
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
    std::cout << "(" << this->category << ":" << this->name << ")(";
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

    static int get_next_id(const std::string& category) {
        return categoryCounters[category]++;
    }
     json toJson() const {
        json j;
        j["name"] = name;
        j["category"] = category;
        j["Relation ID"]=id;
        j["Origin"]= origin.name;
        j["Origin ID"]=origin.id;
        j["destination"]= dest.name;
        j["destination ID"]=dest.id;
        json props;
        for (const auto& [key, value] : properties) {
            if (value.type() == typeid(std::string)) {
                props[key] = std::any_cast<std::string>(value);
            } else if (value.type() == typeid(int)) {
                props[key] = std::any_cast<int>(value);
            } else if (value.type() == typeid(double)) {
                props[key] = std::any_cast<double>(value);
            } else if (value.type() == typeid(bool)) {
                props[key] = std::any_cast<bool>(value);
            }
            // Add more types as needed
        }
        j["properties"] = props;
        return j;
    }

    // Function to write Node attributes to a JSON file
    void writeToJsonFile(const std::string& db_path,const std::string& filename) const {
        std::ofstream file(db_path+filename);
        if (file.is_open()) {
            json j = toJson();
            file << std::setw(4) << j << std::endl; // Pretty print with indentation
            std::cout << "Node attributes written to " << filename << std::endl;
        } else {
            std::cerr << "Failed to open file: " << filename << std::endl;
        }
    }

    private:
    // Static map to keep track of the next ID for each category
    static std::unordered_map<std::string, int> categoryCounters;
};  
// Initialize the static member
std::unordered_map<std::string, int> Relation::categoryCounters;