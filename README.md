# Mikrotik RIF decoder (Golang)

This is a tool that decodes the Mikrotik RIF files (the files that you  provide to their support team).

Mikrotik have a tool on their website to inspect RIF files and for comparing the sections side by side.
It's a very good tool, and I recommend using it if you can.

However, sometimes you can't upload your support file to their site. Maybe you want to vet what's in the
file before you submit it.

## Downloads

Download for Linux and Windows can be found on the [Releases tab](https://github.com/farseeker/go-mikrotik-rif/releases)

- [v0.1 for Linux](https://github.com/farseeker/go-mikrotik-rif/releases/download/v0.1/mikrotik-rif-linux)
- [v0.1 for Windows](https://github.com/farseeker/go-mikrotik-rif/releases/download/v0.1/mikrotik-rif-windows.exe)

## Acknowledgements

This code is a Go port of an [old Perl script](https://pastebin.com/pa30DNfw) originally written by `paul@unsup.sbrk.co.uk`.
The credit for all of the decoding logic goes to that script. I have also included that script in this repository
in case the Pastebin version ever goes offline.

## Why not an already existing tool?

The following tools I already found, and my rebuttals for why I created this tool:

- Mikrotik Website - Excellent. Highly recommend. You should use this unless you can't upload your RIF files
without reading what's in them first, or if you don't have internet access, or don't want to create a Mikrotik
account

- http://k3dt.eu/supout-reader/ - Works. But you're uploading your RIF file to an unknown 3rd party. Also you can't
view the entire RIF output at once (e.g. to copy/paste into a diff tool or something like that)

- [This old Perl script from the forum](https://pastebin.com/pa30DNfw) - it's been removed from the forum. This is the only
place I can find it. It's not cross-platform (*nix only), replies on Perl (I couldn't get it to work due to missing dependencies),
and is not portable (requires Perl, required dependencies, doesn't work on Windows). This script is what this work is based on.
