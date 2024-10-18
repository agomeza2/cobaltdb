#ifndef DATAPROCESSOR_H
#define DATAPROCESSOR_H
#include <string>
#include <vector>
#include "../Main_comp/relation.cpp"
#include <xlnt/xlnt.hpp>
class DataProcessor {
public:
    std::vector<Node> nodes;
    std::vector<Relation> relations;
    DataProcessor() {
    }

    // Method to get the vector of Nodes
    const std::vector<Node>& getNodes() {
        return nodes;
    }

    // Method to display all Nodes
    void showNodes() {
        for (auto node : nodes) {
            node.show(); // Assuming Node has a show method
        }
    }
    void saveNodes(const std::string& db_path){
        for (int i=0; i<nodes.size(); i++){
            nodes[i].writeToJsonFile(db_path,nodes[i].name);
        }
    } 
    void saveRelations(const std::string& db_path){
        for (int i=0; i<relations.size(); i++){
            relations[i].writeToJsonFile(db_path,relations[i].name);
        }
    } 
    void processDataToNodeExcel(std::string filePath) {
        std::cout<<"workbook"<<std::endl;
        std::cout<<filePath<<std::endl;
        std::cout<<"Hello mom"; 
        xlnt::workbook wb;
        try {  
        wb.load(filePath);
        } catch (const std::exception& e) {
        std::cerr << "Error loading file: " << e.what() << std::endl;
        throw; 
        } catch (...) {
        std::cerr << "Unknown error occurred while loading the file." << std::endl;
        throw; 
        }


        auto ws = wb.active_sheet();
        std::vector<std::string> headers;

        bool firstRow = true;

        for (auto row : ws.rows(false)) {
            std::vector<std::string> rowData;
            for (auto cell : row) {
                rowData.push_back(cell.to_string());
            }

            if (firstRow) {
                headers = rowData;
                firstRow = false;
            } else {
                if (rowData.size() < headers.size()) {
                    std::cerr << "Row does not have enough columns: " << rowData.size() << std::endl;
                    continue;
                }

                std::string name = rowData[0];
                std::string category = headers[0];
                Node node(category, name);

                for (size_t i = 1; i < rowData.size(); ++i) {
                    node.add(headers[i], rowData[i]);
                }

                nodes.push_back(node);
            }
        }
    }
   void processDataToRelationExcel() {
    int count = 0;  // Counter for unique relation names

    // Compare each node with every other node
    for (size_t i = 0; i < nodes.size(); ++i) {
        for (size_t j = i + 1; j < nodes.size(); ++j) {
            const Node& node1 = nodes[i];
            const Node& node2 = nodes[j];

            // Find common key-value pairs between the two nodes
            for (const auto& [key, value1] : node1.properties) {
                auto it = node2.properties.find(key);
                if (it != node2.properties.end()) {
                    const std::any& value2 = it->second;

                    // Compare the values using std::any_cast
                    try {
                        if (std::any_cast<std::string>(value1) == std::any_cast<std::string>(value2)) {
                            // Create a relation for the matching key-value pair
                            Relation relation(
                                node1, node2, 
                                "RelatedBy_" + key,  // Category for this relation
                                "CommonAttribute_" + std::to_string(count), // Unique relation name
                                key, value1          // Store the attribute as a property
                            );
                            relations.push_back(relation);  // Add the relation to the list
                            ++count;  // Increment the counter for the next relation
                        }
                    } catch (const std::bad_any_cast& e) {
                        // Handle the case where the values are not strings
                        std::cerr << "Type mismatch for key: " << key << "\n";
                    }
                }
            }
        }
    }
}

     
};

#endif // DATAPROCESSOR_H