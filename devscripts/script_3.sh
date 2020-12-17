# run this under fabric folder
source ./config.sh

export PATH=$(pwd)/build/bin:$PATH 
export FABRIC_CFG_PATH=$(pwd)/sampleconfig 

# create channel called $channelID


configtxgen -channelID $channelID -outputCreateChannelTx $txID -profile SampleSingleMSPChannel -configPath $FABRIC_CFG_PATH

peer channel create -o 127.0.0.1:7050 -c $channelID -f $txID

peer channel join -b $blockID

# build the chaincode project with cmd
# go build - o name_of_chaincode path_to_chaincode
# change the path of the go excutable, watch out the name and version number


CORE_CHAINCODE_LOGLEVEL=debug CORE_PEER_TLS_ENABLED=false CORE_CHAINCODE_ID_NAME=$name:$version $path_to_excutables -peer.address 127.0.0.1:7052
