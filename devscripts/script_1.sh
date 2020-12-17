# Clone the Fabric repository from GitHub, run this under fabric folder
make orderer peer configtxgen
export PATH=$(pwd)/build/bin:$PATH
export FABRIC_CFG_PATH=$(pwd)/sampleconfig

# initialize couchdb 
# docker run -p 5984:5984 couchdb:2

configtxgen -profile SampleDevModeSolo -channelID syschannel -outputBlock genesisblock -configPath $FABRIC_CFG_PATH -outputBlock $(pwd)/sampleconfig/genesisblock

ORDERER_GENERAL_GENESISPROFILE=SampleDevModeSolo orderer
# if something wrong, check the path /var/hyperledger... in the yaml files under fabric/sampleconfig.

# the orderer will be runing, open another terminal and run script_2