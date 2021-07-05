from os import path
from utils import unicodes

def parse(args: list) -> dict:

    from pathlib import Path
    from typing import Counter
    from os import getcwd

    validFuncs = set([
                        'puts',
                        'list',
                        'remove',
                        'unpack',
                        'zipto'
                    ])

    def _throwInvalid(err: int):
        errlogs = {
            1: f" {unicodes('9')}  error: unrecognized argument\n\t'xcxcxcx' is not valid $PATH, command, or flag",
            2: f" {unicodes('9')}  error: too many function arguments passed",
            3: f" {unicodes('9')}  error: --ALL flag modifies --DIFF to only use one $PATH - use a config .json",
            4: f" {unicodes('9')}  error: only $PATHs passed, and no --DIFF flag present to use multipath",
            5: f" {unicodes('9')}  error: --ALL flag only modifies --DIFF, are you sure you needed this flag?",
            6: f" {unicodes('9')}  error: no $PATH provided for func",
            7: f" {unicodes('9')}  error: --DIFF does not accept a function to call"
        }
        print(errlogs[err])
        exit(0)



    d = {
        "verb": 0,
        "diff": False,
        "help": False,
        "func": "",
    }


    while args:

        cur = args.pop(0)

        if '-' in cur:
            flags = Counter(cur.replace('-', ''))
            d["verb"] += flags['v'] + flags["verbose"]
            if 'h' in flags.keys() or "help" in flags.keys():
                d["help"] = True

    # for r in args:
    #     if (  r[0] == '-' 
    #           or r == 'h' 
    #           or r == 'help'):
    #         flagCounts = Counter(r.replace('-', ''))
    #         verbosity = flagCounts['v'] + flagCounts['verbose']

    #         # help flags here, exit with verbose level
    #         if flagCounts['h'] + flagCounts['help'] != 0:
    #             from helps import helper
    #             helper()

    #         if flagCounts['d'] + flagCounts['diff'] != 0:
    #             diff_flag = True
    #         if flagCounts['a'] + flagCounts['all'] != 0:
    #             all_flag = True

    #     else:
            
    #         def validPath(p: str) -> bool:
    #             here = Path(getcwd() + '/' + p)
    #             if here.exists() and not here.is_dir():
    #                 return True
    #             else:
    #                 return False

    #         if r not in validFuncs and not validPath(r):
    #             _throwInvalid(1)

    #         elif len(funcs) == 1 and not validPath(r):
    #             _throwInvalid(2)

    #         elif validPath(r) and len(paths) == 1 and all_flag:
    #             _throwInvalid(3)
            
        
    #         else:
    #             if validPath(r):
    #                 paths.append(r)
    #             else: 
    #                 funcs.append(r)
    
    # ## all these checks gotta be cleaned up
    # if len(funcs) == 0 and not diff_flag:
    #     _throwInvalid(4)

    # if len(funcs) == 1 and len(paths) == 1 and all_flag:
    #     _throwInvalid(5)

    # if len(funcs) == 1 and len(paths) == 0:
    #     if funcs[0] != 'list':
    #         _throwInvalid(6)

    # if len(funcs) == 1 and diff_flag:
    #     _throwInvalid(7)





    return d
