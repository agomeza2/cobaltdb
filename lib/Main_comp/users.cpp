#include <iostream>
#include <unordered_map>
#include <string>

class User {
protected:
    std::string username;
    std::string password;

public:
    User(const std::string& username, const std::string& password)
        : username(username), password(password) {}

    virtual ~User() {}

    std::string getUsername() const {
        return username;
    }

    bool authenticate(const std::string& enteredPassword) const {
        return password == enteredPassword;
    }

    virtual void displayInfo() const = 0; 
};

class AdminUser : public User {
public:
    AdminUser(const std::string& username, const std::string& password)
        : User(username, password) {}

    void displayInfo() const override {
        std::cout << "Admin User: " << username << std::endl;
    }
};

class StandardUser : public User {
public:
    StandardUser(const std::string& username, const std::string& password)
        : User(username, password) {}

    void displayInfo() const override {
        std::cout << "Standard User: " << username << std::endl;
    }
};

int main() {
    std::unordered_map<std::string, User*> users;

    users["admin"] = new AdminUser("admin", "admin123");
    users["user1"] = new StandardUser("user1", "user123");
    users["user2"] = new StandardUser("user2", "pass123");

    std::string username;
    std::string password;

    std::cout << "Enter username: ";
    std::cin >> username;
    std::cout << "Enter password: ";
    std::cin >> password;

    auto it = users.find(username);
    if (it != users.end() && it->second->authenticate(password)) {
        std::cout << "Authentication successful." << std::endl;
        it->second->displayInfo();
    } else {
        std::cout << "Incorrect username or password." << std::endl;
    }

    for (auto& pair : users) {
        delete pair.second;
    }

    return 0;
}

