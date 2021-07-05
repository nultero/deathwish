class Helper:
    def __init__(self, rg: str, args: dict):
        self.args = args
        helpSet = {
            "puts": self.Puts,
            "list": self.Ls,
            "remove": self.Remove,
            "unpack": self.Unpack,
            "zipto": self.Zipto,
            "*": self.All,
        }
        helpSet[rg]()

    def Puts(self):
        print("\n*  novem PUTS [ PATH ]")
        print("\t        : logs the given PATH(s) as a dotfile")
        print("\t        : (this will also overwrite an existing dot if novem already logged it)")
        print("\n")

    def Ls(self):
        print("\n*  novem LIST [ {string}, here ]")
        print("\t none     : lists all dotfiles, details depend on verbosity")
        print("\t here     : tries to see if any logged dots are in current dir")
        print("\t{string}  : searches for {str} in dotfiles")

    def Remove(self):
        print("\n*  novem REMOVE [ PATH ]")
        print("\t        : removes PATH from logs")
        print("\t        : (after confirmation)")

    def Unpack(self):
        print("\n*  novem UNPACK")
        print("\t(no args) : unpacks all aliases to DEFAULT_SHELL_SOURCE_PATH")
        print(f"\t>>> SOURCE_PATH : {self.args['conf']}")

    def Zipto(self):
        print("\n*  novem UPDATE [ alias, type ]")
        print("\t'alias' : prompts for & edits alias content")
        print("\t'type'  : prompts for & alters type name in DEFAULT_CONFIG_DIR_PATH")
        print(f"\t>>> SOURCE_PATH : {self.args['zip']}")

    def All(self):
        self.Puts()
        self.Ls()
        self.Remove()
        self.Unpack()
        self.Zipto()
