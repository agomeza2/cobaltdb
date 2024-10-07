cmake_minimum_required(VERSION 3.12)
project(MyExcelProject)

# Option to specify the vcpkg directory
set(VCPKG_ROOT "" CACHE PATH "Path to the vcpkg directory")
if(NOT VCPKG_ROOT)
    message(FATAL_ERROR "Please specify the VCPKG_ROOT path using -DVCPKG_ROOT=<path> during configuration.")
endif()

# Set the CMake toolchain file for vcpkg
set(CMAKE_TOOLCHAIN_FILE "${VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake" CACHE STRING "Toolchain file for vcpkg")

# Specify C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED TRUE)
set(CMAKE_CXX_EXTENSIONS FALSE)

# Set the import directory
set(IMPORT_DIR "${CMAKE_SOURCE_DIR}/import")

# Find xlsxio libraries
find_package(xlsxio REQUIRED)

# Print found directories for debugging
message(STATUS "Found xlsxio: ${xlsxio_DIR}")
message(STATUS "xlsxio_INCLUDE_DIRS: ${xlsxio_INCLUDE_DIRS}")
message(STATUS "xlsxio_LIBRARIES: ${xlsxio_LIBRARIES}")

# Find required libraries
find_package(ZLIB REQUIRED)
find_package(Expat REQUIRED)
# Print found directories for debugging
message(STATUS "Found Expat: ${Expat_DIR}")
message(STATUS "Expat_INCLUDE_DIRS: ${Expat_INCLUDE_DIRS}")
message(STATUS "Expat_LIBRARIES: ${Expat_LIBRARIES}")
find_package(miniz REQUIRED)

# Specify the source file and define the target
add_executable(cobalt src/main.cpp)  # Adjust the source file path as necessary

# Link required libraries after defining the target
target_link_libraries(cobalt PRIVATE
    ${xlsxio_LIBRARIES}  # Use the variable that holds the path to xlsxio libraries
    ZLIB::ZLIB
    ${Expat_LIBRARIES}
    miniz::miniz
)

# Set a definition to use in the source code
target_compile_definitions(cobalt PRIVATE IMPORT_DIR="${IMPORT_DIR}")
