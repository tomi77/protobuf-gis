
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "Point2D" type object.
 */
proto.gis.protobuf.Point2D.prototype.toGeoJSON = function() {
  return {
    type: 'Point',
    coordinates: [this.getX(), this.getY()]
  };
};

/**
 * Creates an GeoJSON repesentation of a "Point3D" type object.
 */
proto.gis.protobuf.Point3D.prototype.toGeoJSON = function() {
  return {
    type: 'Point',
    coordinates: [this.getX(), this.getY(), this.getZ()]
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiPoint2D/MultiPoint3D" type object.
 */
proto.gis.protobuf.MultiPoint2D.prototype.toGeoJSON = proto.gis.protobuf.MultiPoint3D.prototype.toGeoJSON = function() {
  return {
    type: 'MultiPoint',
    coordinates: this.getPointList().map(point => point.toGeoJSON().coordinates)
  };
};
