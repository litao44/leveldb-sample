cmake_minimum_required(VERSION 3.10)
project(level)

set(CMAKE_CXX_STANDARD 11)

add_executable(level main.cpp)

target_link_libraries(level PUBLIC
        leveldb
        pthread
        snappy
)

