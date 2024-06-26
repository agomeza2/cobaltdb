#include <iostream>
#include <any>
#include <unordered_map>
#include <typeinfo>
#include "json.hpp"
#include "json_fwd.hpp"
#include <fstream>
#include <iomanip>

using json = nlohmann::json;
class Node {
public:
    std::string name;
    std::string label;
    std::unordered_map<std::string, std::any> properties;

    template <typename... Args>
    Node(std::string label, std::string name, Args&&... args)
        :label(std::move(label)),name(std::move(name))  {
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


    void show() {
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

     json toJson() const {
        json j;
        j["name"] = name;
        j["label"] = label;
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
    void writeToJsonFile(const std::string& filename) const {
        std::ofstream file(filename);
        if (file.is_open()) {
            json j = toJson();
            file << std::setw(4) << j << std::endl; // Pretty print with indentation
            std::cout << "Node attributes written to " << filename << std::endl;
        } else {
            std::cerr << "Failed to open file: " << filename << std::endl;
        }
    }
};