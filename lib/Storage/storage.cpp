#include <filesystem> // Works for C++17 and later
#include <iostream>
#include <string>
#include <vector>

#ifdef _WIN32
#include <Windows.h>
#endif

#ifdef _WIN32
std::wstring stringToWchar(const std::string& str);
#endif

class Storage {
public:
    void create_folder(std::string db_name = "") {
        std::string path = "../db";
        if (!db_name.empty()) {
            path += "/" + db_name;
        }

#ifdef _WIN32
        try {
            std::filesystem::create_directories(stringToWchar(path));
            std::cout << "Folder created successfully: " << path << std::endl;
        } catch (const std::filesystem::filesystem_error& e) {
            std::cerr << "Failed to create folder: " << path << std::endl;
            std::cerr << "Error: " << e.what() << std::endl;
        }
#else
        try {
            std::filesystem::create_directories(path);
            std::cout << "Folder created successfully: " << path << std::endl;
        } catch (const std::filesystem::filesystem_error& e) {
            std::cerr << "Failed to create folder: " << path << std::endl;
            std::cerr << "Error: " << e.what() << std::endl;
        }
#endif
    }

    void create_db(std::string db_name) {
        create_folder(db_name);
    }
};

#ifdef _WIN32
std::wstring stringToWchar(const std::string& str) {
    int size_needed = MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, nullptr, 0);
    std::vector<wchar_t> buffer(size_needed);
    MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, buffer.data(), size_needed);
    return std::wstring(buffer.begin(), buffer.end() - 1); // Exclude null terminator
}
#endif