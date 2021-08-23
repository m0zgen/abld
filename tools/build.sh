#!/bin/bash
# Cretaed by Yevgeniy Gonvharov, https://sys-adm.in

# Envs
# ---------------------------------------------------\
PATH=$PATH:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin
SCRIPT_PATH=$(cd `dirname "${BASH_SOURCE[0]}"` && pwd)

DEST="/opt/blocky/"

# Build
cd $SCRIPT_PATH; cd ..

# Functions
# Help information
usage() {

    echo -e "" "\nParameters:\n"
    echo -e "-b - Build BLD"
    echo -e "-d \"srv1 srv2 srv3 \" - Deploy BLD to targets. Can't works without -u\n"
    echo -e "-u - Remote user name"
    exit 1

}

buildBLD() {

    if [[ ! -d $SCRIPT_PATH/builds ]]; then
        mkdir $SCRIPT_PATH/builds
    fi

    env GOOS=linux GOARCH=amd64 go build -o $SCRIPT_PATH/builds
}

deployBLD() {

    buildBLD

    echo "Process deployment to server: $1 .."

    local cmdStop="sudo systemctl stop blocky"
    local cmdStart="sudo systemctl start blocky"

    ssh -ttt $1 "sudo systemctl stop blocky"
    scp $SCRIPT_PATH/builds/blocky $1:$DEST 
    ssh -ttt $1 "sudo systemctl start blocky"

}

if [[ -z "$1" ]]; then
    usage;
fi

# Checks arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -b|--build) BUILD=1; ;;
        -u|--user) USER=1; USERNAME=$2;;
        -d|--deploy) DEPLOY=1; TARGETS=$2; ;;
        -h|--help) usage ;; 
    esac
    shift
done

if [[ "$BUILD" -eq "1" ]]; then
    buildBLD; echo "Binary saved to: $SCRIPT_PATH/builds"; echo "Done!"
fi

if [[ "$DEPLOY" -eq "1" ]]; then
    if [[ -z "$TARGETS" ]]; then
        usage
    else
        if [[ -z "$USER" ]]; then
            usage
        else
            for srv in $TARGETS; do
                deployBLD $srv $USERNAME
            done
        fi
    fi
fi

# echo $DEPLOY $TARGETS


