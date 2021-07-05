#!/usr/bin/python3

from args import parse
from utils import unicodes
from sys import argv
from funcs import evalFuncs

DEFAULT_CONFIG_PATH = '+++' # main dotfiles json $path
                                        # i.e., something like
                                        # '~/.config/$user.json' or
                                        # '~/.dotfiles/*.json' or
                                        # '~/.$user/novem.json'

DEFAULT_ZIP_PATH = '&&&' # main zip target used in 'zipto' argument

def isConfigEmpty(args: dict) -> bool:
    return not (len(args["conf"]) * len(args["zip"]))

def throwInvalidConst(conf: str):
    print(f"  {unicodes('9')} error: the '{conf}' setting in `novem.py` is empty")
    quit()


def main():

    if len(argv) > 1:
        args = {}
        args["conf"] = DEFAULT_CONFIG_PATH
        args["zip"] = DEFAULT_ZIP_PATH

        if not isConfigEmpty(args):
            args.update(parse(argv[1:]))
            print("\n")
            print(args)

        else:
            ## INITIALIZE SCRIPT
            print("nfrgioj")
        # if not isEmpty(DEFAULT_CONFIG_PATH):
        # else:
        #     throwInvalidConst("DEFAULT_CONFIG_PATH")


        # if not isEmpty(DEFAULT_ZIP_PATH):
        # else:
        #     throwInvalidConst("DEFAULT_ZIP_PATH")
        

        # evalFuncs(rgs)


    else:
        print(f"  {unicodes('finger')}  no args provided")
        print("  (run '-h' or '--help' to see opts)")

if __name__ == '__main__':
    main()