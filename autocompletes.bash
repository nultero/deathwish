_novem() {
    COMPREPLY=()
    local cur=${COMP_WORDS[COMP_CWORD]}

    local STARTOPTS=( "puts" "list" "remove" "unpack" "zipto" )
    local SECONDOPTS=( "alias" "type" )
    local OPTS
    case $3 in 
    "list")
        OPTS="here"
    ;;
    "alias" | "type" | "source" | "search")
        OPTS=()
    ;;
    *)
        OPTS=${STARTOPTS[*]}
    ;;
    esac
    COMPREPLY=( `compgen -W "${OPTS[*]}" $cur` )
}

complete -F _novem novem