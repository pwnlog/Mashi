# Mashi

Generate username conventions that are commonly used by organizations.

# Help Menu

```
  __  __           _     _ 
 |  \/  |         | |   (_)
 | \  / | __ _ ___| |__  _ 
 | |\/| |/ _| / __| '_ \| |
 | |  | | (_| \__ \ | | | |
 |_|  |_|\__,_|___/_| |_|_|
          
        Author: pwnlog

Usage: mashi names.txt
Description: Reads names from a text file and generates possible usernames based on those names.
Option:
        -h, --help  Show this help message and exit.
```

# Run

Run without installing:

```sh
go run mashi.go names.txt
```

# Install

Build the binary:

```
go build mashi
```

Add the binary to a directory that's in `$PATH` environment variable:

```
sudo cp mashi /usr/bin/mashi
```

Run mashi:

```
mashi
```

# Usage

Generate usernames with `mashi` using a text file:

```
mashi names.txt
```

# Example

Example output:

```
❯ go run mashi.go names.txt

[+] Minimalized
evangelineburton
burtonevangeline
evangeline.burton
burton.evangeline
burtone
eburton
bevangeline
e.burton
b.evangeline

[+] Capitalized
EvangelineBurton
BurtonEvangeline
Evangeline.Burton
Burton.Evangeline
BurtonE
EBurton
BEvangeline
E.Burton
B.Evangeline

[+] Single Detected
evangeline
Evangeline

[+] Single Detected
burton
Burton

[+] Minimalized
aoiamawashi
amawashiaoi
aoi.amawashi
amawashi.aoi
amawashia
aamawashi
aaoi
a.amawashi
a.aoi

[+] Capitalized
AoiAmawashi
AmawashiAoi
Aoi.Amawashi
Amawashi.Aoi
AmawashiA
AAmawashi
AAoi
A.Amawashi
A.Aoi

[+] Single Detected
aoi
Aoi

[+] Single Detected
amawashi
Amawashi
```