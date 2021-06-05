
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
                'add': fn.add,
                'edit': fn.edit,
                'list': fn.ls,
                'remove': fn.remove,
                'buildfrom': fn.buildfrom,
                'unpack': fn.unpack
                }   

        functions[args['funcs'][0]](
                                pth=args['paths'][0],
                                verbosity=args['verb'],
                                CONF=args['conf']
                                )


class fn():

    def add(pth: str, verbosity: int, CONF: str):
        import json
        from datetime import datetime as dt
        from pathlib import Path
        from os import stat
        with open(CONF, 'r') as f:
            out = json.load(f)
        _path = Path(pth).absolute().resolve()
        out[str(_path)] = f"{stat(_path).st_mtime}"
        with open(CONF, 'w') as f:
            json.dump(out, f)

    def edit():
        ...

    def ls():
        ...

    def remove():
        ...

    def buildfrom():
        ...

    def unpack():
        ...