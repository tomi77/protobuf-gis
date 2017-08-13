import sys
from codecs import open

from setuptools import setup, find_packages


setup(
    name='protobuf-gis',
    version='1.0.3',
    author='Tomasz Jakub Rup',
    author_email='tomasz.rup@gmail.com',
    url='https://github.com/tomi77/protobuf-gis',
    description='GIS ProtoBuf module',
    long_description='' if sys.version_info[:2] == (3, 4) else open('README.rst').read(),
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Topic :: Scientific/Engineering :: GIS',
        'Topic :: Software Development',
        'Topic :: Software Development :: Libraries',
        'Topic :: Software Development :: Libraries :: Python Modules',
        'Topic :: Utilities',
    ],
    license='MIT',
    packages=find_packages(exclude=('tests',)),
    test_suite='tests'
)
