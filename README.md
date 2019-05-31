# pipestack
Task / Project management application written in Go using the Buffalo tool kit

## Summary

Pipestack is a Task manager written specifically for software engineering teams. It utilises the computational concepts of Stack, Pipe and Heap.
Tasks are SMART :-

 * SPECIFIC and detailed in their requirement
 * MEASURED Track effort as Days, Hours, Minutes or Story points ( configurable )
 * ASSIGNED only to an individual
 * REALISTIC in defining the goal / outcome objective
 * TIME bound, with clear Start and End dates 
 
 Tasks are organised on an engineers workboard, which both the engineer and the manager have visibility of.
 
 More detailed information about how to use Pipestack, and its "DTR" methodology can be found in the applications contextual help system.
 
 ## Technology
 
 Pipestack is written in Go, because "Like why would you not ?" by default it anticipates access to a PostgreSQL database (MySQL, SQLite, CockroachDB are also supported).
 
 It uses the excellent Buffalo RAD ( Rapid Application Development) Toolkit see
 http://gobuffalo.io
 
 ## Getting Started
 
 Pipestack builds into a standalone binary which can be run on Linux, MAC OSX and Windows ( See note on Windows) using the buffalo build tool
 
 To get a development version running you will need to have the Go language installed for your system. For the smart kids using linux that should be a simple case of installing it from your distributions software package manager
  * Redhat based distros use the "yum" tool
  * Debian based distros use the "apt" tool
 
 On Kubuntu Linux (*no bias there at all*), and Ubuntu Linux you can use the following command
 
 ```sudo apt install golang```
 
 Next you'll need Buffalo installed, see https://gobuffalo.io/en/docs/installation
 
 Although I am not using SQLite, for this project you might wish to. Certainly I do use SQLite for some of my other applications.
 I would recommend installing the Custom Buffalo with SQLite3 support. Docs on how to do that can be found at
 
 https://gobuffalo.io/en/docs/installation#with-SQLite3
 
 ## Hacking on Pipestack
 
 I am really excited about working on this project, but the thing I am looking forward to the most is working with others on developing the application.
 
 Start first by reading the SPECIFICATION.md file in this repo. That sets out the overall vision for the application, along with a Roadmap of the things
 that are being targetted in each release.
 
 Next clone this repo, and get the application running in your local development environment. Then pick out one of the Issues from the [Issue Tracker](https://github.com/ricktimmis/pipestack/issues), work on a fix / feature implementation and send me a PR.
 Simples !
 
 ** A Note on Windows
 
 So if you're still using Windows it really is time to Stop! Come on man, the world has moved on all the designee types are using MAC and the Techies have all moved up to Linux. Let me give you a leg up. Get yourself over to [Kubuntu.org](https://kubuntu.org/getkubuntu/) download yourself a copy of the latest release and install it dual boot. You'll never believe how it will transform the performance of your PC. 
