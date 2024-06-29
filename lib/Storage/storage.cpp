#include <filesystem> // Works for C++17 and later
#include <string>
#include <stringapiset.h>
#include <vector> 

// Function declaration for stringToWchar
wchar_t* stringToWchar(const std::string& str);

#ifdef _WIN32
bool isWindows() {
    return true; // Assuming this is Windows platform
}
#else
bool isWindows() {
    return false; // Assuming this is not Windows platform (Linux or other)
}
#endif

class Storage {
public:
    void create_folder_WIN(std::string db_name = "") {
        std::string path = "../db";
        if (!db_name.empty()) {
            path = path + "/" + db_name;
        }
        try {
            std::filesystem::create_directories(stringToWchar(path));
            std::cout <<"Folder created successfully: " << path << std::endl;
        } catch (const std::filesystem::filesystem_error& e) {
            std::cerr << "Failed to create folder: " << path << std::endl;
            std::cerr << "Error: " << e.what() << std::endl;
        }
    }

    void create_folder_LINUX(std::string db_name = "") {
        std::string path = "../db";
        if (!db_name.empty()) {
            path = path + "/" + db_name;
        }
        try {
            std::filesystem::create_directories(path);
            std::cout << "Folder created successfully: " << path << std::endl;
        } catch (const std::filesystem::filesystem_error& e) {
            std::cerr << "Failed to create folder: " << path << std::endl;
            std::cerr << "Error: " << e.what() << std::endl;
        }
    }

    void create_folder() {
        if (isWindows()) {
            this->create_folder_WIN();
        } else {
            this->create_folder_LINUX();
        }
    }

    void create_db(std::string db_name) {
        if (isWindows()) {
            this->create_folder_WIN(db_name);
        } else {
            this->create_folder_LINUX(db_name);
        }
    }
};

// Implementation of stringToWchar function
wchar_t* stringToWchar(const std::string& str) {
    
    int size_needed = MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, NULL, 0);

    std::vector<wchar_t> buffer(size_needed);

    MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, &buffer[0], size_needed);

    wchar_t* wideString = new wchar_t[size_needed];
    wcscpy_s(wideString, size_needed, buffer.data());

    return wideString;
}
