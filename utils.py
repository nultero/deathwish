
def unicodes(match: str) -> str:
    tmp = ""
    cases = {
        '9': "\u277e >",
        'finger': "\u261E",
    }
    if match in cases.keys():
        tmp += cases[match]

    return tmp

def diffArgs(paths: list, verbosity: int, all: bool) -> None:
    # verbosity 0: has / has not been changed
    # verbosity 1: numlines changed
    # verbosity 2: numlines + datestamps
    # verbosity 3: list out lines changed + datestamps
    from pathlib import Path
    for p in paths:
        p = Path(p)
        with open(f"{p}", 'r') as f:
            dt = f.read()
            print(dt)
    print(f"vvv:{verbosity} ::: all:{all}")