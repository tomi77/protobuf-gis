
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "LineString2D/LineString3D" type object.
 */
proto.gis.protobuf.LineString2D.prototype.toGeoJSON = proto.gis.protobuf.LineString3D.prototype.toGeoJSON = function() {
  return {
    type: 'LineString',
    coordinates: this.getPointList().map(point => point.toGeoJSON().coordinates)
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiLineString2D/MultiLineString3D" type object.
 */
proto.gis.protobuf.MultiLineString2D.prototype.toGeoJSON = proto.gis.protobuf.MultiLineString3D.prototype.toGeoJSON = function() {
  return {
    type: 'MultiLineString',
    coordinates: this.getLineStringList().map(lineString => lineString.toGeoJSON().coordinates)
  };
};
