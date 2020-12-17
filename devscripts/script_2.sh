# run this under fabric folder

export PATH=$(pwd)/build/bin:$PATH 
export FABRIC_CFG_PATH=$(pwd)/sampleconfig 

FABRIC_LOGGING_SPEC=chaincode=debug CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:7052 peer node start --peer-chaincodedev=true

# if something wrong, check the port config in the yaml files under fabric/sampleconfig

# open another terminal and run script_3
