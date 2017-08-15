
/** This code is included by protobuf-gis package **/

/**
 * Creates an GeoJSON repesentation of a "Polygon2D/Polygon3D" type object.
 */
proto.gis.protobuf.Polygon2D.prototype.toGeoJSON = proto.gis.protobuf.Polygon3D.prototype.toGeoJSON = function() {
  return {
    type: 'Polygon',
    coordinates: [this.getPointList().map(point => point.toGeoJSON().coordinates)]
  };
};

/**
 * Creates an GeoJSON repesentation of a "MultiPolygon2D/MultiPolygon3D" type object.
 */
proto.gis.protobuf.MultiPolygon2D.prototype.toGeoJSON = proto.gis.protobuf.MultiPolygon3D.prototype.toGeoJSON = function() {
  return {
    type: 'MultiPolygon',
    coordinates: this.getPolygonList().map(polygon => polygon.toGeoJSON().coordinates)
  };
};
