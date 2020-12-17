# lifecycle things
source config.sh

export PATH=$(pwd)/build/bin:$PATH 
export FABRIC_CFG_PATH=$(pwd)/sampleconfig 


peer lifecycle chaincode approveformyorg  -o 127.0.0.1:7050 --channelID $channelID --name $name --version $version --sequence 1 --init-required --signature-policy "OR ('SampleOrg.member')" --package-id $name:$version
peer lifecycle chaincode checkcommitreadiness -o 127.0.0.1:7050 --channelID $channelID --name $name --version $version --sequence 1 --init-required --signature-policy "OR ('SampleOrg.member')"
peer lifecycle chaincode commit -o 127.0.0.1:7050 --channelID $channelID --name $name --version $version --sequence 1 --init-required --signature-policy "OR ('SampleOrg.member')" --peerAddresses 127.0.0.1:7051

# alias init='source config.sh; CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C $channelID -n $name -c '\'{\"Args\":\[\"InitLedger\"\]}\'' --isInit'

# then we may use the CLI cmds to invoke and query chaincodes
# note that the first call must be initialization
# CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C $channelID -n $name -c '{"Args":["InitLedger"]}' --isInit
# CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C $channelID -n $name -c '{"Args":["invoke","a","b","10"]}
# CORE_PEER_ADDRESS=127.0.0.1:7051 peer chaincode invoke -o 127.0.0.1:7050 -C $channelID -n $name -c '{"Args":["query","a"]}'