# rx-cli tool

A command-line "helper" tool to create folders, notes, generate shells and other common cli patterns for things like nmap, gobuster to assist with hacking challenges e.g. hackthebox, tryhackme etc.

<!-- build badges go here -->
<!-- asciicinema goes here -->
<!-- toc -->

## rx why?

While doing HTB challenges I found myself copying and pasting things over and over again just to get started on the new box, or go digging through notes or sites for a one-liner I had used many times before. I wanted a simple tool to generate my notes document from a template, with all the standard nmap, gobuster and reverse shell info as my starting point.

<!-- tocstop -->

<!-- steps -->

## Usage & Installation

<!-- usage -->
**Install from Binary Release (Recommended)**

Download the pre-compiled binary for your platform from the [Releases page](https://github.com/ronamosa/rx-cli/releases).

```bash
# Download the latest release for your platform
# Example for Linux (amd64)
wget https://github.com/ronamosa/rx-cli/releases/latest/download/rx_linux_amd64.zip

# Extract the binary
unzip rx_linux_amd64.zip

# Make it executable
chmod +x rx_linux_amd64

# Optional: Move to a directory in your PATH for easier access
sudo mv rx_linux_amd64 /usr/local/bin/rx
```

**Build from Source**

```bash
# clone repo
git clone https://github.com/ronamosa/rx-cli.git

# build binary - 'rx'
go build

# (optional) install 'rx' binary to your system
go install
```

**Using Makefile**

```bash
# clone repo
git clone https://github.com/ronamosa/rx-cli.git

# build binary - 'rx'
make build

# install binary to your system
make install

# build for multiple platforms
make release
```
<!-- usagestop -->

## Commands

<!-- commands -->

```bash
# create markdown notes file
rx create notes <target-name> <target-ip>
```

Example: `rx create notes Test 127.0.0.1`

![rx-create-notes](docs/img/rx-create-notes.png)

```bash
# create reverse shells (php, python, bash, binary)
rx create shell <type> --LHOST <listener ip> --LPORT <listener port>
```

**Available shell types:**
- `php` - Creates a PHP reverse shell
- `python` or `py` - Creates a Python reverse shell
- `bash` - Creates a Bash reverse shell
- `bin` - Creates a C source file for compilation into a binary reverse shell

Example: `rx create shell php --LHOST 127.0.0.1 --LPORT 4444`

![rx-create-shell](docs/img/rx-create-shell-php.png)

## Logo Font Options

```bash

██████  ██   ██ ██   ██  █████   ██████ ██   ██ 
██   ██  ██ ██  ██   ██ ██   ██ ██      ██  ██  
██████    ███   ███████ ███████ ██      █████   
██   ██  ██ ██  ██   ██ ██   ██ ██      ██  ██  
██   ██ ██   ██ ██   ██ ██   ██  ██████ ██   ██ 
                                                

██████╗ ██╗  ██╗██╗  ██╗ █████╗  ██████╗██╗  ██╗
██╔══██╗╚██╗██╔╝██║  ██║██╔══██╗██╔════╝██║ ██╔╝
██████╔╝ ╚███╔╝ ███████║███████║██║     █████╔╝ 
██╔══██╗ ██╔██╗ ██╔══██║██╔══██║██║     ██╔═██╗ 
██║  ██║██╔╝ ██╗██║  ██║██║  ██║╚██████╗██║  ██╗
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
                                                

 ██▀███  ▒██   ██▒ ██░ ██  ▄▄▄       ▄████▄   ██ ▄█▀
▓██ ▒ ██▒▒▒ █ █ ▒░▓██░ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ 
▓██ ░▄█ ▒░░  █   ░▒██▀▀██░▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ 
▒██▀▀█▄   ░ █ █ ▒ ░▓█ ░██ ░██▄▄▄▄██ ▒▓▓▄ ▄██▒▓██ █▄ 
░██▓ ▒██▒▒██▒ ▒██▒░▓█▒░██▓ ▓█   ▓██▒▒ ▓███▀ ░▒██▒ █▄
░ ▒▓ ░▒▓░▒▒ ░ ░▓ ░ ▒ ░░▒░▒ ▒▒   ▓▒█░░ ░▒ ▒  ░▒ ▒▒ ▓▒
  ░▒ ░ ▒░░░   ░▒ ░ ▒ ░▒░ ░  ▒   ▒▒ ░  ░  ▒   ░ ░▒ ▒░
  ░░   ░  ░    ░   ░  ░░ ░  ░   ▒   ░        ░ ░░ ░ 
   ░      ░    ░   ░  ░  ░      ░  ░░ ░      ░  ░   
                                    ░               

▄▄▄  ▐▄• ▄  ▄ .▄ ▄▄▄·  ▄▄· ▄ •▄ 
▀▄ █· █▌█▌▪██▪▐█▐█ ▀█ ▐█ ▌▪█▌▄▌▪
▐▀▀▄  ·██· ██▀▐█▄█▀▀█ ██ ▄▄▐▀▀▄·
▐█•█▌▪▐█·█▌██▌▐▀▐█ ▪▐▌▐███▌▐█.█▌
.▀  ▀•▀▀ ▀▀▀▀▀ · ▀  ▀ ·▀▀▀ ·▀  ▀


                                                                  ..      
                            .uef^"                          < .z@8"`      
   .u    .      uL   ..   :d88E                              !@88E        
 .d88B :@8c   .@88b  @88R `888E             u           .    '888E   u    
="8888f8888r '"Y888k/"*P   888E .z8k     us888u.   .udR88N    888E u@8NL  
  4888>'88"     Y888L      888E~?888L .@88 "8888" <888'888k   888E`"88*"  
  4888> '        8888      888E  888E 9888  9888  9888 'Y"    888E .dN.   
  4888>          `888N     888E  888E 9888  9888  9888        888E~8888   
 .d888L .+    .u./"888&    888E  888E 9888  9888  9888        888E '888&  
 ^"8888*"    d888" Y888*"  888E  888E 9888  9888  ?8888u../   888E  9888. 
    "Y"      ` "Y   Y"    m888N= 888> "888*""888"  "8888P'  '"888*" 4888" 
                           `Y"   888   ^Y"   ^Y'     "P'       ""    ""   
                                J88"                                      
                                @%                                        
                              :"                                          
```

## Recent Updates

### v0.0.2
- Added support for all shell types (Python, Bash, Binary/C source)
- Added Makefile for consistent builds and multi-platform releases
- Improved error handling and output messages
- Made Python and Bash shells executable upon creation
- Improved CLI help text with better documentation and examples
- Added binary releases for easy installation without compiling
