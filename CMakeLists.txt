cmake_minimum_required(VERSION 3.12)
project(MyExcelProject)

# Set the path to vcpkg
set(VCPKG_ROOT "C:/Users/Alex/vcpkg")
set(CMAKE_TOOLCHAIN_FILE "${VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake")

# Find the xlnt package
find_package(Xlnt CONFIG REQUIRED)

# Specify C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED TRUE)
set(CMAKE_CXX_EXTENSIONS FALSE)

# Set the import directory
set(IMPORT_DIR "${CMAKE_SOURCE_DIR}/import")

# Specify the source file
add_executable(cobalt src/main.cpp)

# Link xlnt to your target
target_link_libraries(cobalt PRIVATE xlnt::xlnt)

# Set a definition to use in the source code
target_compile_definitions(cobalt PRIVATE IMPORT_DIR="${IMPORT_DIR}")

