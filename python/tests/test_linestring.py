import unittest

from gis.protobuf.linestring_pb2 import LineString2D, LineString3D, MultiLineString2D, MultiLineString3D
from gis.protobuf.point_pb2 import Point2D, Point3D

class LineString2DTestCase(unittest.TestCase):
    def test_togeojson(self):
        linestring = LineString2D(point=[Point2D(x=1.0, y=2.0),
                                         Point2D(x=3.0, y=4.0)])

        self.assertEqual(linestring.toGeoJSON(), {
            'type': 'LineString',
            'coordinates': [[1.0, 2.0], [3.0, 4.0]],
        })


class LineString3DTestCase(unittest.TestCase):
    def test_togeojson(self):
        linestring = LineString3D(point=[Point3D(x=1.0, y=2.0, z=3.0),
                                         Point3D(x=4.0, y=5.0, z=6.0)])

        self.assertEqual(linestring.toGeoJSON(), {
            'type': 'LineString',
            'coordinates': [[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]],
        })


class MultiLineString2DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        multiLineString = MultiLineString2D(line_string=[LineString2D(point=[Point2D(x=1.0, y=2.0),
                                                                             Point2D(x=3.0, y=4.0)]),
                                                         LineString2D(point=[Point2D(x=5.0, y=6.0),
                                                                             Point2D(x=7.0, y=8.0)])])

        self.assertEqual(multiLineString.toGeoJSON(), {
            'type': 'MultiLineString',
            'coordinates': [[[1.0, 2.0], [3.0, 4.0]], [[5.0, 6.0], [7.0, 8.0]]]
        })


class MultiLineString3DTestCase(unittest.TestCase):
    def test_toGeoJSON(self):
        multiLineString = MultiLineString3D(line_string=[LineString3D(point=[Point3D(x=1.0, y=2.0, z=3.0),
                                                                             Point3D(x=4.0, y=5.0, z=6.0)]),
                                                         LineString3D(point=[Point3D(x=7.0, y=8.0, z=9.0),
                                                                             Point3D(x=10.0, y=11.0, z=12.0)])])

        self.assertEqual(multiLineString.toGeoJSON(), {
            'type': 'MultiLineString',
            'coordinates': [[[1.0, 2.0, 3.0], [4.0, 5.0, 6.0]], [[7.0, 8.0, 9.0], [10.0, 11.0, 12.0]]]
        })
