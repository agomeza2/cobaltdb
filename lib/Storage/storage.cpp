#include <iostream>
#include <sys/stat.h>

class Storage {
public:
    void create_folder() {
        const char* path = "../../db/as";
        mode_t mode = 0777; // Permissions for the new directory
        
        // Change permissions of the parent directory if needed
        if (chmod("../../db", mode) != 0) {
            std::cerr << "Failed to change permissions of parent directory." << std::endl;
            return;
        }

        // Attempt to create the directory
        if (mkdir(path, mode) == 0) {
            std::cout << "Folder created successfully." << std::endl;
        } else {
            std::cerr << "Failed to create folder." << std::endl;
        }
    }
};
