# protobuf-gis

## Installation

### proto files

`.proto` files must be installed in a `/usr/incude/` folder.

    git clone https://github.com/tomi77/protobuf-gis.git
    cd protobuf-gis
    make

### NodeJS

    npm install protobuf-gis --save

or

    yarn add protobuf-gis

## Additional functions

All messages has function `toGeoJSON` that returns a GeoJSON representation of a message.
