#!/usr/bin/env bash

# Copyright IBM Corp. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

if [[ -z "${FPC_PATH}" ]]; then
  echo "Error: FPC_PATH not set"
  exit 1
fi

. "${FPC_PATH}"/integration/test-network/scripts/utils.sh

channelID="mychannel"

outDir="."
cryptoConfigDir=${outDir}/crypto-config
channelArtifactsDir=${outDir}/channel-artifacts
packageDir=${outDir}/packages

CHAINCODE_SERVER_PORT=9999

CRYPPTOGEN_CMD=$GOPATH/bin/cryptogen
CONFIGTXGEN_CMD=$GOPATH/bin/configtxgen

echo "Clean existing deployment artifacts"
rm -rf ${cryptoConfigDir}
rm -rf ${channelArtifactsDir}
rm -rf ${packageDir}

echo "Generate crypto material for orgs"
$CRYPPTOGEN_CMD generate --output=${cryptoConfigDir} --config=./crypto-config.yaml

echo "Generate genesis block"
$CONFIGTXGEN_CMD -profile DemoGenesis -channelID testchainid -outputBlock ${channelArtifactsDir}/genesis.block
$CONFIGTXGEN_CMD -profile DemoChannel -outputCreateChannelTx ${channelArtifactsDir}/channel.tx -channelID ${channelID}



echo "Generate client connection profile"
function yaml_connection {
    sed -e "s/{{ORG}}/$1/g" \
        connection-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}
for org in $(shopt -s globstar; find ${cryptoConfigDir}/**/peerOrganizations/ -type d -depth 1 -execdir echo {} ';');
do
  echo "$(yaml_connection $org)" > ${cryptoConfigDir}/peerOrganizations/${org}/connection.yaml
done

echo "Package ercc and fpccc"
CC_TYPE="external"
ERCC_ID="ercc"
ERCC_VER="1.0"
FPCCC_ID="fpccc"
FPCCC_PATH=${FPC_PATH}/examples/helloworld
FPC_MRENCLAVE="$(cat "${FPCCC_PATH}"/_build/lib/mrenclave)"

for peer in $(shopt -s globstar; find ${cryptoConfigDir}/**/peers/ -type d -depth 1 -execdir echo {} ';');
do
    # ercc
    endpoint="${ERCC_ID}-${peer}:${CHAINCODE_SERVER_PORT}"
    packageName="${ERCC_ID}-${peer}.tgz"
    packageChaincode "${packageDir}" "${packageName}" "${ERCC_ID}" "${ERCC_VER}" "${CC_TYPE}" "${endpoint}" "${peer}"

    # fpc cc
    endpoint="${FPCCC_ID}-${peer}:${CHAINCODE_SERVER_PORT}"
    packageName="${FPCCC_ID}-${peer}.tgz"
    packageChaincode "${packageDir}" "${packageName}" "${FPCCC_ID}" "${FPC_MRENCLAVE}" "${CC_TYPE}" "${endpoint}" "${peer}"
done

echo "Store mrenclave for fpccc"
echo "FPC_MRENCLAVE=${FPC_MRENCLAVE}" >> ${packageDir}/chaincode-config.properties
