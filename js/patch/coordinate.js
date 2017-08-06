
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "Coordinate2D" type object.
 */
proto.gis.protobuf.Coordinate2D.prototype.toGeoJSON = function() {
  return [this.getX(), this.getY()];
};

/**
 * Creates an GeoJSON repesentation of a "Coordinate3D" type object.
 */
proto.gis.protobuf.Coordinate3D.prototype.toGeoJSON = function() {
  return [this.getX(), this.getY(), this.getZ()];
};
