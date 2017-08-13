from .linestring_pb2 import LineString2D, LineString3D, MultiLineString2D, MultiLineString3D
from .point_pb2 import Point2D, Point3D, MultiPoint2D, MultiPoint3D
from .polygon_pb2 import Polygon2D, Polygon3D, MultiPolygon2D, MultiPolygon3D


def _2d(point):
    return [point.x, point.y]


def _3d(point):
    return [point.x, point.y, point.z]


def pointToGeoJSON(dimension):
    def inner(self):
        return {
            'type': 'Point',
            'coordinates': dimension(self)
        }
    return inner


def multiPointToGeoJSON(type, dimension):
    def inner(self):
        return {
            'type': type,
            'coordinates': [dimension(point) for point in self.point]
        }
    return inner


def multiLineStringToGeoJSON(type, dimension):
    def inner(self):
        return {
            'type': type,
            'coordinates': [[dimension(point) for point in lineString.point] for lineString in self.line_string]
        }
    return inner


def polygonToGeoJSON(dimension):
    def inner(self):
        return {
            'type': 'Polygon',
            'coordinates': [[dimension(point) for point in self.point]]
        }
    return inner


def multiPolygonToGeoJSON(dimension):
    def inner(self):
        return {
            'type': 'MultiPolygon',
            'coordinates': [[[dimension(point) for point in polygon.point]] for polygon in self.polygon]
        }
    return inner


setattr(Point2D, 'toGeoJSON', pointToGeoJSON(_2d))
setattr(Point3D, 'toGeoJSON', pointToGeoJSON(_3d))
setattr(MultiPoint2D, 'toGeoJSON', multiPointToGeoJSON('MultiPoint', _2d))
setattr(MultiPoint3D, 'toGeoJSON', multiPointToGeoJSON('MultiPoint', _3d))
setattr(LineString2D, 'toGeoJSON', multiPointToGeoJSON('LineString', _2d))
setattr(LineString3D, 'toGeoJSON', multiPointToGeoJSON('LineString', _3d))
setattr(MultiLineString2D, 'toGeoJSON', multiLineStringToGeoJSON('MultiLineString', _2d))
setattr(MultiLineString3D, 'toGeoJSON', multiLineStringToGeoJSON('MultiLineString', _3d))
setattr(Polygon2D, 'toGeoJSON', polygonToGeoJSON(_2d))
setattr(Polygon3D, 'toGeoJSON', polygonToGeoJSON(_3d))
setattr(MultiPolygon2D, 'toGeoJSON', multiPolygonToGeoJSON(_2d))
setattr(MultiPolygon3D, 'toGeoJSON', multiPolygonToGeoJSON(_3d))
