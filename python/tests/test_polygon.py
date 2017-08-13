import unittest

from gis.protobuf.polygon_pb2 import Polygon2D, Polygon3D, MultiPolygon2D, MultiPolygon3D
from gis.protobuf.point_pb2 import Point2D, Point3D

class Polygon2DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        polygon = Polygon2D(point=[Point2D(x=1.0, y=2.0),
                                   Point2D(x=3.0, y=4.0)])

        self.assertEqual(polygon.toGeoJSON(), {
            'type': 'Polygon',
            'coordinates': [[[1.0, 2.0], [3.0, 4.0]]]
        })


class Polygon3DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        polygon = Polygon3D(point=[Point3D(x=1.0, y=2.0, z=3.0),
                                   Point3D(x=4.0, y=5.0, z=6.0)])

        self.assertEqual(polygon.toGeoJSON(), {
            'type': 'Polygon',
            'coordinates': [[[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]]
        })


class MultiPolygon2DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        multiPolygon = MultiPolygon2D(polygon=[Polygon2D(point=[Point2D(x=1.0, y=2.0),
                                                                Point2D(x=3.0, y=4.0)]),
                                               Polygon2D(point=[Point2D(x=5.0, y=6.0),
                                                                Point2D(x=7.0, y=8.0)])])

        self.assertEqual(multiPolygon.toGeoJSON(), {
            'type': 'MultiPolygon',
            'coordinates': [[[[1.0, 2.0], [3.0, 4.0]]], [[[5.0, 6.0], [7.0, 8.0]]]]
        })


class MultiPolygon3DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        multiPolygon = MultiPolygon3D(polygon=[Polygon3D(point=[Point3D(x=1.0, y=2.0, z=3.0),
                                                                Point3D(x=4.0, y=5.0, z=6.0)]),
                                               Polygon3D(point=[Point3D(x=7.0, y=8.0, z=9.0),
                                                                Point3D(x=10.0, y=11.0, z=12.0)])])

        self.assertEqual(multiPolygon.toGeoJSON(), {
            'type': 'MultiPolygon',
            'coordinates': [[[[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]], [[[7.0, 8.0, 9.0], [10.0, 11.0, 12.0]]]]
        })
