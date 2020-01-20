#!/bin/bash

# get go deps

cd $FPC_PATH && make godeps

# build ecc_enclave
cd $FPC_PATH/ecc_enclave && make clean && make

# build ecc
cd $FPC_PATH/ecc &&  make sym-links

# build auction

cd $FPC_PATH/examples/auction && make

cd $FPC_PATH/demo/client/backend/mock
export LD_LIBRARY_PATH=${LD_LIBRARY_PATH:+"${LD_LIBRARY_PATH}:"}${FPC_PATH}/ecc_enclave/_build/lib
ln -s $FPC_PATH/examples/acution /_build/ enclave