const GIS = require('../')

describe('gis.protobuf.Polygon2D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point2D(); point1.setX(1); point1.setY(2)
    let point2 = new GIS.Point2D(); point2.setX(3); point2.setY(4)

    let polygon = new GIS.Polygon2D(); polygon.setPointList([point1, point2])

    expect(polygon.toGeoJSON()).toEqual({
      type: 'Polygon',
      coordinates: [[[1, 2], [3, 4]]]
    })
  })
})

describe('gis.protobuf.Polygon3D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point3D(); point1.setX(1); point1.setY(2); point1.setZ(3)
    let point2 = new GIS.Point3D(); point2.setX(4); point2.setY(5); point2.setZ(6)

    let polygon = new GIS.Polygon3D(); polygon.setPointList([point1, point2])

    expect(polygon.toGeoJSON()).toEqual({
      type: 'Polygon',
      coordinates: [[[1, 2, 3], [4, 5, 6]]]
    })
  })
})

describe('gis.protobuf.MultiPolygon2D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point2D(); point1.setX(1); point1.setY(2)
    let point2 = new GIS.Point2D(); point2.setX(3); point2.setY(4)
    let polygon1 = new GIS.Polygon2D(); polygon1.setPointList([point1, point2])

    let point3 = new GIS.Point2D(); point3.setX(5); point3.setY(6)
    let point4 = new GIS.Point2D(); point4.setX(7); point4.setY(8)
    let polygon2 = new GIS.Polygon2D(); polygon2.setPointList([point3, point4])

    let multiPolygon = new GIS.MultiPolygon2D(); multiPolygon.setPolygonList([polygon1, polygon2])

    expect(multiPolygon.toGeoJSON()).toEqual({
      type: 'MultiPolygon',
      coordinates: [[[[1, 2], [3, 4]]], [[[5, 6], [7, 8]]]]
    })
  })
})

describe('gis.protobuf.MultiPolygon3D has getGeoJSON function', function() {
  it('which returns a GeoJSON object', function() {
    let point1 = new GIS.Point3D(); point1.setX(1); point1.setY(2); point1.setZ(3)
    let point2 = new GIS.Point3D(); point2.setX(4); point2.setY(5); point2.setZ(6)
    let polygon1 = new GIS.Polygon3D(); polygon1.setPointList([point1, point2])

    let point3 = new GIS.Point3D(); point3.setX(7); point3.setY(8); point3.setZ(9)
    let point4 = new GIS.Point3D(); point4.setX(10); point4.setY(11); point4.setZ(12)
    let polygon2 = new GIS.Polygon3D(); polygon2.setPointList([point3, point4])

    let multiPolygon = new GIS.MultiPolygon3D(); multiPolygon.setPolygonList([polygon1, polygon2])

    expect(multiPolygon.toGeoJSON()).toEqual({
      type: 'MultiPolygon',
      coordinates: [[[[1, 2, 3], [4, 5, 6]]], [[[7, 8, 9], [10, 11, 12]]]]
    })
  })
})
