import unittest

from gis.protobuf.point_pb2 import Point2D, Point3D, MultiPoint2D, MultiPoint3D

class Point2DTestCase(unittest.TestCase):
    def test_togeojson(self):
        point = Point2D(x=1.0, y=2.0)

        self.assertEqual(point.toGeoJSON(), {
            'type': 'Point',
            'coordinates': [1.0, 2.0],
        })


class Point3DTestCase(unittest.TestCase):
    def test_togeojson(self):
        point = Point3D(x=1.0, y=2.0, z=3.0)

        self.assertEqual(point.toGeoJSON(), {
            'type': 'Point',
            'coordinates': [1.0, 2.0, 3.0],
        })


class MultiPoint2DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        multiPoint = MultiPoint2D(point=[Point2D(x=1.0, y=2.0),
                                         Point2D(x=3.0, y=4.0)])

        self.assertEqual(multiPoint.toGeoJSON(), {
            'type': 'MultiPoint',
            'coordinates': [[1.0, 2.0], [3.0, 4.0]]
        })


class MultiPoint3DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        multiPoint = MultiPoint3D(point=[Point3D(x=1.0, y=2.0, z=3.0),
                                         Point3D(x=4.0, y=5.0, z=6.0)])

        self.assertEqual(multiPoint.toGeoJSON(), {
            'type': 'MultiPoint',
            'coordinates': [[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]]
        })
