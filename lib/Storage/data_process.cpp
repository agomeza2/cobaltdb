#ifndef DATAPROCESSOR_H
#define DATAPROCESSOR_H
#include <string>
#include <vector>
#include "../Main_comp/relation.cpp"
#include <xlsxio_read.h>  // Include xlsxio for reading Excel files

class DataProcessor {
public:
    std::vector<Node> nodes;

    DataProcessor() {
    }

    // Method to get the vector of Nodes
    const std::vector<Node>& getNodes() {
        return nodes;
    }

    // Method to display all Nodes
    void showNodes() {
        for (auto& node : nodes) {
            node.show(); // Assuming Node has a show method
        }
    }

    // Method to process data from Excel using xlsxio
    void processDataExcel(std::string filePath) {
        std::cout << "Processing workbook: " << filePath << std::endl;

        xlsxioreader xlsxioread;
        if ((xlsxioread = xlsxioread_open(filePath.c_str())) == NULL) {
            std::cerr << "Error opening .xlsx file." << std::endl;
            return;
        }

        xlsxioreadersheet sheet;
        if ((sheet = xlsxioread_sheet_open(xlsxioread, NULL, XLSXIOREAD_SKIP_EMPTY_ROWS)) == NULL) {
            std::cerr << "Error opening sheet in .xlsx file." << std::endl;
            xlsxioread_close(xlsxioread);
            return;
        }

        std::vector<std::string> headers;
        char* value;
        bool firstRow = true;

        while (xlsxioread_sheet_next_row(sheet)) {
            std::vector<std::string> rowData;
            while ((value = xlsxioread_sheet_next_cell(sheet)) != NULL) {
                rowData.push_back(value);
                free(value);  // Free memory after reading the cell
            }

            if (firstRow) {
                headers = rowData;  // Capture the headers
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
                    node.add(headers[i], rowData[i]);  // Add each cell value to the node
                }

                nodes.push_back(node);  // Add the node to the list of nodes
            }
        }

        xlsxioread_sheet_close(sheet);
        xlsxioread_close(xlsxioread);

        std::cout << "Finished processing workbook." << std::endl;
    }
};

#endif // DATAPROCESSOR_H
