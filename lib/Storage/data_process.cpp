#include "../Main_comp/relation.cpp"
#include "xlnt/xlnt.hpp"
/*
class DataProcess {
private:
    Node node_array[1000000];  // Array to hold up to 1 million nodes
    int node_count;            // Count of nodes created

public:
    DataProcess() : node_count(0) {}

    // Function to read Excel file and create nodes
    void readExcelToNodes(const std::string& filePath, const std::string& db_path) {
        xlnt::workbook wb;
    wb.load("path_to_your_excel_file.xlsx");  // Replace with the actual path

    // Select the active worksheet
    xlnt::worksheet ws = wb.active_sheet();

    // Step 1: Get headers from the first row (row 1)
    std::vector<std::string> headers;
    for (auto cell : ws[1]) {  // First row is the header row
        headers.push_back(cell.to_string());
    }

    // Display the headers
    std::cout << "Headers: ";
    for (const auto &header : headers) {
        std::cout << header << " ";
    }
    std::cout << std::endl;

    // Step 2: Iterate over the rows and extract the data
    for (auto row : ws.rows(false)) {  // Iterate over rows, skipping empty cells
        std::vector<std::string> row_values;
        for (auto cell : row) {
            row_values.push_back(cell.to_string());
        }

        // Display the row data
        std::cout << "Row: ";
        for (const auto &value : row_values) {
            std::cout << value << " ";
        }
        std::cout << std::endl;
    }
    }
};*/