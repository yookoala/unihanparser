Unihan Parser
==============

This is a command line parser tool that helps studying the [Unihan database](http://www.unicode.org/reports/tr38/). The tool will build an sqlite database file of all the Unihan data. Researchers and developers may take advantage of a relational database to analyse and output data needed.

***

Dependencies
------------

- This tool is developed, and has only been tested working, on Linux

- Unihan database version 6 or up<br/>
  (This tool has built and checked against Unihan Database 6.1 and 6.3)

- Requires [Go](http://golang.org) version 1.1 or up

- Requires [mattn's go-sqlite library](http://github.com/mattn/go-sqlite3)<br />
  (with internet connection, the build script will install it for you)


Build the Tool and Database
---------------------------

  1. Unihan database is not included in the tool. Please download the database (version 6+) from [their official website](http://www.unicode.org/reports/tr38/) (usually at [this URL](http://www.unicode.org/Public/UCD/latest/ucd/Unihan.zip). Then extract and save all the ``*.txt`` files into ``data/Unihan``. 

  2. This software comes with a build script (build.sh). To build the commandline tool and the database, just run the build script in the same folder as this file is placed:

    <code>./build.sh</code>

  And things will be handled for you. A sqlite database file "unihan.db" will be generated in the "data" folder.


***

Legalities
----------

This software is offered under the terms of the [GNU Lesser General Public License, either version 3 or any later version](http://www.gnu.org/licenses/lgpl.html).

You don't need to sign a copyright assignment or any other kind of silly and tedious legal document, so just send patches and/or pull requests!
