#!/bin/bash

#MFI

echo "Building STC Compiler"

cd compiler
go build

echo "Compile go successful"

read -p "Enter path for STC compiler path (default \"~/stc\"): " path

path=${path:-"~/stc"}  # Default to '~/stc' if no input is given

# Expand the '~' symbol to the full home directory path
path=$(eval echo $path)

# Check if the path exists
if [ -e "$path" ]; then
    echo "The path '$path' exists."
    
    # Check if it's a directory
    if [ -d "$path" ]; then
        echo "It's a directory."
    elif [ -f "$path" ]; then
        echo "It's a file."
    else
        echo "It's something else (like a symbolic link)."
    fi
else
    echo "The path '$path' does not exist. Creating the directory..."
    mkdir -p "$path"  
    echo "Directory '$path' created."
    mv rpn $path

    echo "Create config file"


    read -p "Enter path for raw libs (default \"lib/src\") in $path directory: " lib_root
    lib_root=${lib_root:-"lib/src"}
    mkdir -p "$path/$lib_root"    
    echo "Directory '$path/$lib_root' created."

    read -p "Enter path for bin libs (default \"lib/bin\") in $path directory: " lib_path
    lib_path=${lib_path:-"lib/bin"}
    mkdir -p "$path/$lib_path"  
    echo "Directory '$path/$lib_path' created."

    read -p "Enter path for clang (default \"/usr/lib/llvm14/bin/clang\"): " clang_path
    clang_path=${clang_path:-"/usr/lib/llvm14/bin/clang"}

    read -p "Enter path for linker (default \"/usr/lib/llvm14/bin/llvm-link\"): " linker_path
    linker_path=${linker_path:-"/usr/lib/llvm14/bin/llvm-link"}

    read -p "Enter path for llc (default \"/usr/lib/llvm14/bin/llc\"): " llc_path
    llc_path=${llc_path:-"/usr/lib/llvm14/bin/llc"}

    read -p "Enter path for llc-disassembly (default \"/usr/lib/llvm14/bin/llvm-dis\"): " llc_dis_path
    llc_dis_path=${llc_dis_path:-"/usr/lib/llvm14/bin/llvm-dis"}

    
    touch stconfig.toml
    echo "[Config]" >> stconfig.toml
    echo "root_path = \"$path\"" >> stconfig.toml
    echo "clang_path = \"$clang_path\"" >> stconfig.toml
    echo "linker_path = \"$linker_path\"" >> stconfig.toml
    echo "llc_path =  \"$llc_path\"" >> stconfig.toml
    echo "llc_dis_path =  \"$llc_dis_path\"" >> stconfig.toml
  
    echo "[Libs]" >> stconfig.toml
    echo "lib_root = \"$lib_root\"" >> stconfig.toml
    echo "lib_path = \"$lib_path\"" >> stconfig.toml
    echo "lib_raw = ["  >> stconfig.toml
    echo "    \"arithmetic\", \"logic\", \"io\", \"array\" " >> stconfig.toml
    echo "]" >> stconfig.toml

    echo "Config file created"
    echo "Moving config file"
    sudo mkdir /etc/stc/
    sudo mv stconfig.toml /etc/stc/


    cd ..
    cp lib/src/* "$path/$lib_root"
    echo "Libs copied"

    read -p "Would you like to create a symlink (stc)? (y/n): " answer

    if [[ "$answer" =~ ^[Yy]$ ]]; then
        # Check if the symlink already exists
        if [ -e "/usr/local/bin/stc" ]; then
            echo "The symlink 'stc' already exists."
        else
            # Create the symlink
            sudo ln -s "$path/rpn" "/usr/local/bin/stc"
            echo "Symlink '/usr/local/bin/stc' created."
        fi
    else
        echo "No symlink created."
    fi

        read -p "Would you like to build libs (y/n): " answer

    if [[ "$answer" =~ ^[Yy]$ ]]; then
        # Check if the symlink already exists
        stc libs
    fi


fi