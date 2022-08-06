#!/usr/bin/env python3

from os import mkdir, getlogin
from os.path import  realpath, expanduser as getHome
from typing import List
from sys import argv as args; args = args[1:]

class uvars_t():
    def __init__(self) -> None:
        self.UNAME = getlogin()
        self.HOME = getHome("~")
        self.NOV_DIR = self.HOME + "/.novem"
        self.NOV_FILE = self.NOV_DIR + "/novemManifest"
UVARS = uvars_t()

def notEnoughArgs():
    print("> not enough args; closing out")
    exit(0)

class fmtr():
    def yellow(s: str): return f"\x1b[33m{s}\x1b[0m"
    def red   (s: str): return f"\x1b[31m{s}\x1b[0m"
    def blue  (s: str): return f"\x1b[34m{s}\x1b[0m"


class Manifest():
    def getManifest(self) -> List[str]:
        '''   Creates the manifest if it doesn't exist -- reads if it does  '''
        nov_lines = []
        try:
            mkdir(UVARS.NOV_DIR)
        except Exception:
            pass

        try:
            with open(UVARS.NOV_FILE, 'r') as nf:
                for ln in nf.readlines():
                    nov_lines.append(ln)
        except FileNotFoundError:
            print(fmtr.yellow("novem index not found"), f"in {UVARS.NOV_DIR}, creating anyway")
            with open(UVARS.NOV_FILE, "w+") as nf:
                nf.write("")
        
        return nov_lines

    def __init__(self) -> None:
        self.Lines = self.getManifest()

    def WriteOut(self):
        cull = set() # dedupe in case I get the same lines
        for ln in self.Lines:
            cull.add(ln) # TODO add collision logic

        lines = [ln for ln in cull]
        lines.sort()

        with open(UVARS.NOV_FILE, "w") as f:
            f.write("\n".join(lines))


def ls(args: List[str], man: Manifest):
    if len(args) == 0:
        for ln in man.Lines:
            print(ln)

def put(args: List[str], man: Manifest):
    if len(args) == 0: 
        notEnoughArgs()

    for arg in args:
        canon: str = realpath(arg)
        canon = canon.replace(UVARS.HOME, "~")
        man.Lines.append(canon)
    man.WriteOut()
    print(fmtr.blue("written:"), "\n".join(args))


def main():
    dispatch = {
        "drop": 5,

        "ls"  : ls,
        "list": ls,

        "put" : put,
        "puts": put,
    }

    if len(args) < 1: 
        # TODO should zero args just stat the current dir?
        notEnoughArgs()


    man = Manifest()

    func = args[0]
    try:
        dispatch[func](args[1:], man)
    except Exception as e:
        print(e)
        print(fmtr.red("unrecognized arg:"), func)


if __name__ == "__main__":
    main()