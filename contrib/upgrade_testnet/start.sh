#!/usr/bin/env bash

NODES=$1
GENESIS_VERSION=$2
GENESIS_URL=$3
UPGRADE_NAME=$4

BUILDDIR=$(pwd)/build
CONTRIBFOLDER=$(pwd)/contrib
TESTNETDIR=$CONTRIBFOLDER/upgrade_testnet
STATICLIBDIR=$CONTRIBFOLDER/lib

# Remove the build folder
echo "===> Removing build folder"
rm -r -f $BUILDDIR

# Create the 4 nodes folders with the correct denom
echo "===> Creating $NODES nodes localnet"
make setup-localnet COIN_DENOM="udaric" NODES=$NODES > /dev/null > /dev/null

# Run the Python script to setup the genesis
echo "===> Setting up the genesis file"
docker run --rm --user $UID:$GID \
  -v $TESTNETDIR:/usr/src/app \
  -v $BUILDDIR:/desmos:Z \
  desmoslabs/desmos-python python setup_genesis.py /desmos $NODES $GENESIS_URL > /dev/null

# Build the new Desmos-Cosmovisor image
echo "===> Building the new Desmos-Cosmovisor image"
make -C $CONTRIBFOLDER/images desmos-cosmovisor DESMOS_VERSION=$GENESIS_VERSION > /dev/null

# Set the correct Desmos image version inside the docker compose file
echo "===> Setting up the Docker compose file"
sed -i "s|image: \".*\"|image: \"desmoslabs/desmos-cosmovisor:$GENESIS_VERSION\"|g" $TESTNETDIR/docker-compose.yml


# Download and check of static libraries

echo "===> Create static libraries dir"
if [[ -d "$STATICLIBDIR" ]]
  echo "$STATICLIBDIR already exists"
then
  mkdir $STATICLIBDIR
fi
echo "===> Download static libraries"
if [[ ! -f $STATICLIBDIR/libwasmvm_muslc.a ]]
then
  wget https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta5/libwasmvm_muslc.a -P $STATICLIBDIR/
fi

echo "===> Checking static libraries"
sha256sum $STATICLIBDIR/libwasmvm_muslc.a | grep d16a2cab22c75dbe8af32265b9346c6266070bdcf9ed5aa9b7b39a7e32e25fe0

# Build the current code using Alpine to make sure it's later compatible with the devnet
echo "===> Building Desmos"
docker run --rm --user $ID:$GID -v $(pwd):/desmos desmoslabs/desmos-build BUILD_TAGS=muslc make build-linux > /dev/null

# Copy the Desmos binary into the proper folders
UPGRADE_FOLDER="$BUILDDIR/node0/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin"
if [ ! -d "$UPGRADE_FOLDER" ]; then
  echo "===> Setting up upgrade binary"

  for ((i = 0; i < $NODES; i++)); do
    echo "====> Node $i"
    mkdir -p "$BUILDDIR/node$i/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin"
    cp "$BUILDDIR/desmos" "$BUILDDIR/node$i/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin/desmos"
  done
fi

# Start the devnet
echo "===> Starting the devnet"
docker-compose -f $TESTNETDIR/docker-compose.yml up -d
