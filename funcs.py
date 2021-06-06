
def evalFuncs(args: dict):

    if args['diff']:
        from utils import diffArgs
        diffArgs(
            paths=args['paths'],
            verbosity=args['verb'],
            all=args['all']
            )

    else:
        functions = {
                'puts': fn.puts,
                'list': fn.ls,
                'remove': fn.remove,
                'buildfrom': fn.buildfrom,
                'unpack': fn.unpack
                }   

        _p = args['paths'][0] if args['funcs'][0] != 'list' else args['conf']

        functions[args['funcs'][0]](
                                pth=_p,
                                verbosity=args['verb'],
                                CONF=args['conf']
                                )


class fn():

    def puts(pth: str, verbosity: int, CONF: str):
    
        import json
        from datetime import datetime as dt
        from pathlib import Path
        from os.path import getmtime

        with open(CONF, 'r') as f:
            out = json.load(f)

        _p = Path(pth).absolute().resolve()
        timestamp = f"{dt.fromtimestamp(getmtime(_p))}"

        line = ""
        if verbosity > 1:
            line += f" ADDED new stamp for : '{_p}' \n\t timestamp : {timestamp}"
        elif verbosity == 1:
            line += f" ADDED : '{_p}'"
        elif verbosity < 1:
            line += f" ADDED : '{pth}'"

        out[str(_p)] = timestamp
        with open(CONF, 'w') as f:
            json.dump(out, f)
        print(line)

    ###################

    def ls(pth: str, verbosity: int, CONF: str):

        import json

        with open(CONF, 'r') as f:
            cflist = json.load(f)

        if verbosity < 1:
            for cf in cflist:
                print(cf.split('/')[-1])

        elif verbosity == 1:
            for cf, stamp in cflist.items():
                print(f"{cf.split('/')[-1]}  |  {stamp.split(' ')[0]}")
        
        elif verbosity == 2:
            for cf, stamp in cflist.items():
                print(f"{cf.split('/')[-1]}  |  {stamp}")

        elif verbosity == 3:
            for cf, stamp in cflist.items():
                print(f"{cf}  |  {stamp}")

    ###################


    def remove():
        ...

    def buildfrom():
        ...

    def unpack():
        ...