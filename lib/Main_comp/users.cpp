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

