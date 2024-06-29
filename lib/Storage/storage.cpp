#include <iostream>
#include <filesystem>

#ifdef _WIN32
bool isWindows() {
    return true;
}
#else
bool isWindows() {
    return false;
}
#endif

class Storage {
public:
    void create_folder_WIN() {
        const wchar_t* path = L"../db/as";
        
        try {
            std::filesystem::create_directories(path);
            std::wcout << L"Folder created successfully: " << path << std::endl;
        } catch (const std::filesystem::filesystem_error& e) {
            std::wcerr << "Failed to create folder: " << path << std::endl;
            std::wcerr << "Error: " << e.what() << std::endl;
        }
    }
    void create_folder_LINUX() {
        const char* path = "../db/as";
        
        try {
            std::filesystem::create_directories(path);
            std::cout << "Folder created successfully: " << path << std::endl;
        } catch (const std::filesystem::filesystem_error& e) {
            std::cerr << "Failed to create folder: " << path << std::endl;
            std::cerr << "Error: " << e.what() << std::endl;
        }
    }
    void create_folder(){
        if (isWindows()) {
        this->create_folder_WIN();
    }   else {
        this->create_folder_LINUX();
    }
    }
};