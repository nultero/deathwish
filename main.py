#!/usr/bin/python3

from args import parse
from utils import unicodes
from sys import argv
from funcs import evalFuncs

# tracks dotfiles
# can put in zip to export
# can rebuild from zip
# configs to & from json

DEFAULT_CONFIG_PATH = '' # main dotfiles json $path
                         # i.e., something like
                         # '~/.config/$user.json' or
                         # '~/.dotfiles/*.json' or
                         # '~/.$user/novem.json'

def main():
    if len(argv) > 1:
        rgs = parse(argv[1:])
        rgs['conf'] = DEFAULT_CONFIG_PATH
        evalFuncs(rgs)

    else:
        print(f"  {unicodes('finger')}  no args provided")
        print("  (run '-h' or '--help' to see opts)")

if __name__ == '__main__':
    main()