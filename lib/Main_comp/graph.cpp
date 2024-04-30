#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <fstream>
#include <iomanip> // Necesario para setw
#include <nlohmann/json.hpp>

using json = nlohmann::json;

class graph {
public:
    void addNode(const std::string& node) {
        if (nodes_.find(node) == nodes_.end()) {
            nodes_.insert(node);
            std::cout << "Nodo agregado: " << node << std::endl;
        } else {
            std::cerr << "El nodo ya existe: " << node << std::endl;
        }
    }

    void addRelation(const std::string& source, const std::string& target) {
        addNode(source);
        addNode(target);
        relaciones_[source].insert(target);
        std::cout << "Relación agregada: " << source << " -> " << target << std::endl;
    }

    void deleteNode(const std::string& node) {
        if (nodes_.erase(node) > 0) {
            relaciones_.erase(node);
            for (auto& pair : relaciones_) {
                pair.second.erase(node);
            }
            std::cout << "Nodo eliminado: " << node << std::endl;
        } else {
            std::cerr << "Nodo no encontrado: " << node << std::endl;
        }
    }

    void deleteRelation(const std::string& source, const std::string& target) {
        auto it = relaciones_.find(source);
        if (it != relaciones_.end()) {
            if (it->second.erase(target) > 0) {
                std::cout << "Relación eliminada: " << source << " -> " << target << std::endl;
            } else {
                std::cerr << "Relación no encontrada: " << source << " -> " << target << std::endl;
            }
        } else {
            std::cerr << "Nodo de origen no encontrado: " << source << std::endl;
        }
    }

    json toJSON() const {
        json grafoJson;
        for (const auto& node : nodes_) {
            grafoJson["nodos"].push_back(node);
        }
        for (const auto& pair : relaciones_) {
            for (const auto& target : pair.second) {
                grafoJson["relaciones"].push_back({pair.first, target});
            }
        }
        return grafoJson;
    }

    bool saveJSONToFile(const std::string& filename) const {
        json grafoJson = toJSON();
        std::ofstream outputFile(filename);
        if (outputFile.is_open()) {
            outputFile << std::setw(4) << grafoJson << std::endl;
            outputFile.close();
            std::cout << "JSON guardado en el archivo exitosamente.\n";
            return true;
        } else {
            std::cerr << "Error al abrir el archivo para escribir.\n";
            return false;
        }
    }

private:
    std::unordered_set<std::string> nodes_;
    std::unordered_map<std::string, std::unordered_set<std::string>> relaciones_;
};

int main() {
    graph grafo;

    grafo.addNode("A");
    grafo.addNode("B");
    grafo.addRelation("A", "B");

    grafo.saveJSONToFile("grafo.json");

    return 0;
}
