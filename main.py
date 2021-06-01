from argslib import parseAll
from utils import unicodes
from sys import argv

# tracks dotfiles
# can put in zip to export
# can rebuild from zip
# configs to & from json

def main():
    if len(argv) > 1:
        parseAll()
        print(f" {unicodes('9')} >")

    else:
        print(f"  {unicodes('finger')}  no args provided")
        print("  (run '-h' or '--help' to see opts)")

if __name__ == '__main__':
    main()