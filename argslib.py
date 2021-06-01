from utils import unicodes
# possibleArgs = ["buildfrom", "unpack"]

def parseAll():
    import argparse
    p = argparse.ArgumentParser()

    ### opts
    #p.add_argument("-c", "--config", help="specify a config to use for import or export", action='store_true')
    # do I need to specify a config? that'll be a necessary thing anyway, right?
    p.add_argument("-d", "--diff", help="echoes changes of a config's dotfile vs. specified file (only datestamps if not verbose)")
    # -v --verbose for tracking more stuff
    #  and a flag for special instructs

    #### MAINARGS >>>>>>>>>>>>>>>
    rgs = p.add_subparsers()
    #actions
    rgs.add_parser("add", help="add a dotfile (or similar) to a config for bundling")
    rgs.add_parser("edit", help="edit a dotfile (or similar) in a config for re-bundling")
    rgs.add_parser("remove", help="remove a dotfile (or similar) from a config for re-bundling")

    # acts. pt II -- bundles
    rgs.add_parser("buildfrom", help="fetches the dotfiles in config if they exist, and bundles into a zip")
    rgs.add_parser("unpack", help=f"if passed a dotfile pack zipped by {unicodes('9')} novem, puts files onto $PATHs on a live *nix system")

    ## info
    rgs.add_parser("list", help="echoes details of a dotfile from config")
    ##########  <<<<<<<<

    fn = p.parse_args()

    if fn.config:
        ...