# Random Password Generator
This is a simple command-line application written in Go that generates a random password and copies it to the clipboard.


## How it works
When you run the application, it generates a random password of a specified length. The length is determined by a command-line argument. **If no argument is provided or the argument is not a positive integer, the length defaults to ```12```.**

The generated password is then copied to the clipboard using the ```atotto/clipboard``` package.

If the password is successfully copied to the clipboard a success message will be shown. If the password cannot be copied to the clipboard, the application prints an error message and exits.

## Dependencies
This application uses the following external Go packages:

```github.com/atotto/clipboard``` for clipboard access.

## Distribution
There are a few excecutable files in the ```dist``` directory. These files are compiled versions of the application for different operating systems. The files are named according to the operating system they are compiled for. 


1. MacOS
    - `rpg_darwin_amd64`: This is the version for MacOS systems running on AMD64 (x86_64) architecture. This is suitable for most modern Macs.

2. Linux
    - `rpg_linux_amd64`: This is the version for Linux systems running on AMD64 (x86_64) architecture. This is suitable for most modern Linux systems.
    - `rpg_linux_386`: This is the version for Linux systems running on 386 (x86) architecture. This is suitable for older Linux systems.

3. Windows
    - `rpg_windows_amd64`: This is the version for Windows systems running on AMD64 (x86_64) architecture. This is suitable for most modern Windows systems.
    - `rpg_windows_386`: This is the version for Windows systems running on 386 (x86) architecture. This is suitable for older Windows systems.

If you want to run the application on a different operating system, you will need to compile it yourself. There is a bash script named ```build.sh``` in the ```scripts``` directory that will compile the application for your operating system. All you have to is edit the array of platforms in the script and run it.
The platforms can be seen by running the command ```go tool dist list```.

## Download
You can download the compiled executables from the ```dist``` directory in this repository. 

### MacOS
```bash
curl -O https://raw.githubusercontent.com/edsonjaramillo/rpg/main/dist/rpg_darwin_amd64
```

### Linux AMD64
```bash
curl -O https://raw.githubusercontent.com/edsonjaramillo/rpg/main/dist/rpg_linux_amd64
```

### Linux 386
```bash
curl -O https://raw.githubusercontent.com/edsonjaramillo/rpg/main/dist/rpg_linux_386
```

### Windows AMD64
```bash
curl -O https://raw.githubusercontent.com/edsonjaramillo/rpg/main/dist/rpg_windows_amd64
```

### Windows 386
```bash
curl -O https://raw.githubusercontent.com/edsonjaramillo/rpg/main/dist/rpg_windows_386 
```

## Usage
Be sure to add the executable to your PATH so you can run it from anywhere in the terminal. You will need to rename it to rpg for shorthand because the name is quite long due to the platform and version information.


From CLI:
```bash
rpg [length]
```

From Go:
```bash
go run cmd/rpg/main.go [length]
```
