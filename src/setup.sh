#!/bin/bash
 
############################## PRE PROCESSING ################################
#check and process arguments
REQUIRED_NUMBER_OF_ARGUMENTS=1
if [ $# -lt $REQUIRED_NUMBER_OF_ARGUMENTS ]
then
    echo "Usage: $0 <path_to_config_file>"
    exit 1
fi

CONFIG_FILE=$1
 
echo "Config file is $CONFIG_FILE"
echo ""
 
#get the configuration parameters
source $CONFIG_FILE

############################## SETUP ################################
if [ "$USAGE" == "SETUP" ]
then
	counter=1
	for node in ${VM_NODES//,/ }
	do
		echo "Setting up $node ..."
	    COMMAND=''

	    if [ ! -e $DIRECTORY ]; then
	    	COMMAND=$COMMAND" git clone https://gitlab.engr.illinois.edu/tkao4/CS425-MP1.git;"
	    fi
	    COMMAND=$COMMAND" wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz;"
	    COMMAND=$COMMAND" sudo tar -C /usr/local -xvzf go1.7.3.linux-amd64.tar.gz;"

	    let counter=counter+1 
	    ssh -t -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no $node "
	            $COMMAND"
	done

elif [ "$USAGE" == "RUN" ]
then
	counter=1
	for node in ${VM_NODES//,/ }
	do
		echo "Running server $node ..."
		COMMAND=''
		if [ ! -e $DIRECTORY/src/machine.$counter.log ]; then
			COMMAND=$COMMAND" wget \"https://courses.engr.illinois.edu/cs425/fa2017/CS425_MP1_Demo_Logs_FA17/vm$counter.log\" -o \"machine.$counter.log\";"
		fi
		COMMAND=$COMMAND" export PATH=$PATH:/usr/local/go/bin;"
	    COMMAND=$COMMAND" export GOPATH=\"$HOME/CS425-MP1\";"
		COMMAND=$COMMAND" cd CS425-MP1/src;"
		COMMAND=$COMMAND" fuser -k 8000/tcp;"
		COMMAND=$COMMAND" nohup go run server/main.go > /dev/null 2>&1 &"

		let counter=counter+1
		ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no $node "
	            $COMMAND"
	done

elif [ "$USAGE" == "STOP" ]
then
	for node in ${VM_NODES//,/ }
	do
		echo "Stopping server $node ..."
		COMMAND=''
		COMMAND=$COMMAND" export PATH=$PATH:/usr/local/go/bin;"
	    COMMAND=$COMMAND" export GOPATH=\"$HOME/CS425-MP1\";"
		COMMAND=$COMMAND" fuser -k 8000/tcp"

		ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no $node "
	            $COMMAND"
	done
fi


