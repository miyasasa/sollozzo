
**_Overview_**

Sollozzo is command line tool developed with [GO(GoLang)](https://golang.org/). Sollozzo provide you can keep version info of your projects. Also you can make base _CRUD_ operations on these projects.
 
**_Getting Started_**

We apply all operation on sample project that it get the name "ibe".

**_Commands_**

Sollozzo offers two options to create a new project with add a command line parameter:

 sollozzo add ibe : This command creates a new project called ibe. It's version is 1.0.0 by default
 sollozzo add ibe 1.2.3 : In the same way this creates ibe project. In addition 1.2.3 is ibe's version. The important thing here is that version info should match with %d.%d.%d pattern. 
 
 sollozzo current ibe: Gets current version of ibe.
 
 sollozzo list       : Lists your all project with version info.
 
 sollozzo remove ibe : Removes ibe project without warning.
 
 Another feature of sollozzo is that we can update our project version with "release" command and flags as --major, --minor, --build. Optionally -M ,-m and -b can be used instead of above flags in order.
 
 sollozzo release ibe --major : 1 Increase major number in version of ibe.
 sollozzo release ibe --minor : 1 Increase minor number in version of ibe.
 sollozzo release ibe --build : 1 Increase build number in version of ibe.
 
 You can apply many flag on same command line. Also increse build number if you do not add any flag.
