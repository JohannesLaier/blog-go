# Tiny Blog in Google Go #

This project contains a tiny blog system written in golang. This project was developed as part of the go programming course at the DHBW Mosbach.

### How to set up? ###

* Install the Jetbains Gogland IDE: https://www.jetbrains.com/go/download/
* Install Google Go: https://golang.org/dl/. Ubuntu users could use: https://wiki.ubuntuusers.de/Go/
* Create your workspace folder in your home directory: mkdir -p ~/GoglandProjects
* Set up your GOENV Variables:
* For Linux: Add this files to your ~/.bashrc File:
```sh
        # golang
        export PATH=$PATH:$(go env GOPATH)/bin
        export GOPATH="$HOME/GoglandProjects"
        export GOROOT="/usr/share/go/
```
* For Windows: Add the following env variables 
```cmd
        GOPATH=%HOMEPATH%/GoglandProjects
        GOROOT=c:\Go
```
* Create the following folder structure into your workspace directory:
```cmd
        GoglandProjects
            bin
            src
            pkg
```
* Checkout the project to your workspace src folder (GoglandProjects/src)
* Open your Gogland-IDE and set your GOPATH to Your GoglandProjects Folder

### Screenshots ###

![Blog detail page](docs/imgs/blog_page_detail.png?raw=true "Blog detail page")
![Admin panel login](docs/imgs/blog_admin_login.png?raw=true "Admin panel login")
![Admin panel overview](docs/imgs/blog_admin_overview.png?raw=true "Admin panel overview")
![Admin panel edit post](docs/imgs/blog_admin_edit.png?raw=true "Admin panel edit post")

### Structure ###

The application has been developed to create a framework that can be reused to develop other web applications based on Google Go. For this purpose, a modular design and reusable components is essential. To achieve this goal the structure of the source code is divided into three main parts as follows:

| Folder                                                                                     |                                                                                                                                                                                                                                                                                                                                                                                              Description                                                                                                                                                                                                                                                                                                                                                                                              |
|--------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
|   app/     * controller/           * models/              * shared/                        |                                                                                                                                                                                                                                                             This folder contains the models and controllers of a specific application.  In this case, these are the controllers and models that affect the blog.  The shared / folder contains the wrapper classes, such as the session wrappers, which are used by different controllers.                                                                                                                                                                                                                                                            |
| core/     config/     controller/     db/     http/       server/       session/     util/ |                                                                                                                                                                  This folder contains basic modules that can be used to develop Go applications. There are classes for starting a server, classes that  allow applications to be developed according to the  model-view-controller principle, session management and the  implementation of the database.  Furthermore, this package includes a class for easier handling of the template engine.  Also in this folder is the config class which is responsible for writing classes.                                                                                                                                                                  |
| res/     ssl/     db/     session/     tpl/      ssl/     db/     session/     tpl/     | The resources folder contains all resources that have not been written to Google Go.  For example, the ssl / folder contains the SSL certificates for the web server.   The DB folder contains JSON files in which the databases used are stored.  By creating symbolic links, this can easily be stored in another area and backups can be created using rsync or other tools.   In the session folder, all HTTP sessions are stored in a JSON file.  Due to the file principle of Unix, this can also be easily stored in a folder or operated in a Docker container.   The tpl / folder contains the HTML templates (the views) of the blog application.   In  the www / folder are all static resources of the application such as  CSS and Javascript code, images, icons, fonts, and much more. |

The application's entry point is located in the app.go, which is located directly in the main directory.
By sharing the app / folder and sharing in the res / tpl / and res / www / folder, this design pattern makes it easy to create a new web application.
architecture

The blog application is based on different design paradigms. The model-view-controller principle, the singleton pattern and polymorphism are used. The application consists of the modules described below.

##### Model-View-Controller #####
To implement this application, the Model-View-Controller design pattern has been used. So the class controller exists, which can contain several handlers. The AddHandler () method allows multiple handlers to be passed to a controller, which can later be started by the server class. The View class can load a view, pass data to it, and embed that view in a layout and send it to the client. The models are passed to the view via the SetModel () method and defined in the folder app / models /.

##### Database #####
To simplify data management, a simple NoSQL database was implemented. This is essentially divided into three classes. The DB class (db.go) represents a database that groups multiple DBCollections. These are comparable to individual tables of a SQL database. In a DBCollection (db_collection.go) different kinds of records, in this case DBCollectionEntries, can be stored. Therefore, it can only store objects that own the methods of the DBCollectionEntrie interface and thus implement the interface. That is, the GetID () and SetID () methods must be implemented.
The DB class is implemented here according to the singleton pattern and has a global map in each of which an instance of all databases currently stored exist. These are identified by their name. A getter can be used to load the instance of the desired database from the map.
Since it is not possible to execute SQL queries on the collections, different methods from Functional Programming have been implemented on a DB Collection. This allows the same functions that are possible with SQL to be implemented. Among other things, filtering and aggregation methods have been implemented that can be used to process the data.
Using the Load () and Store () methods, databases can be stored in or read from a JSON file.

##### Session #####
To manage the HTTP sessions, the Session and SessionStore classes have been implemented. The SessionStore class is mainly used to collect all sessions. To create a new session, the SessionStore class can instantiate a new session and generate the cookie that is sent to the client. This class also checks whether a session is still valid or whether it has already expired. In a session, all information can be stored using the Get () or Put () methods using a key.

##### Server #####
The server class is used to start a web server. For this purpose, a server is given a configuration in the form of an instance of the config class and can then be started with the Run () method. The AddController () method allows different controllers to be passed to the server, which then become active when the server is started. The AddFileServer () methods can also be used to pass paths to the server in which to search for static files.

##### Configuration #####
The config class is used to read in the command line parameters and to save them. Getter methods can be used to read this information from the Config object. This class is used in the DB classes to determine the storage path for the DB, in the SessionStorage for the storage path of the sessions, in the server class to find the SSL certificates, and in the View class to find the templates and to be able to read.

### Unit Tests ###

Use the following command to install unit test library.

```sh
user@user:~/go/src/de/blog/projekt/01$ go get github.com/stretchr/testify/assert
```

Execute all unit tests

```sh
johannes@user:~/go/src/de/blog/projekt/01$ go test ./...
?   	de/blog/projekt/01	[no test files]
ok  	de/blog/projekt/01/app/controller/admin/author	0.034s
ok  	de/blog/projekt/01/app/controller/admin/category	0.018s
ok  	de/blog/projekt/01/app/controller/admin/comment	0.010s
ok  	de/blog/projekt/01/app/controller/admin/login	0.015s
ok  	de/blog/projekt/01/app/controller/admin/logout	0.023s
ok  	de/blog/projekt/01/app/controller/admin/main	0.037s
ok  	de/blog/projekt/01/app/controller/admin/post	0.014s
ok  	de/blog/projekt/01/app/controller/admin/profile	0.057s
ok  	de/blog/projekt/01/app/controller/blog/about	0.034s
ok  	de/blog/projekt/01/app/controller/blog/category	0.040s
ok  	de/blog/projekt/01/app/controller/blog/contact	0.004s
ok  	de/blog/projekt/01/app/controller/blog/detail	0.043s
ok  	de/blog/projekt/01/app/controller/blog/index	0.012s
ok  	de/blog/projekt/01/app/models/author	0.004s
ok  	de/blog/projekt/01/app/models/comment	0.014s
ok  	de/blog/projekt/01/app/models/keyword	0.019s
ok  	de/blog/projekt/01/app/models/post	0.011s
ok  	de/blog/projekt/01/app/shared	0.007s
ok  	de/blog/projekt/01/core/config	0.011s
ok  	de/blog/projekt/01/core/controller	0.014s
ok  	de/blog/projekt/01/core/db	0.025s
ok  	de/blog/projekt/01/core/http/server	1.028s
ok  	de/blog/projekt/01/core/http/session	0.004s
ok  	de/blog/projekt/01/core/util	0.007s
ok  	de/blog/projekt/01/core/view	0.003s
johannes@user:~/go/src/de/blog/projekt/01$ 
```

### Contribution guidelines ###

* Apply the TDD paradigm
* Writing tests
* Write clean code
* Refactor your code
* Pay attention on code quality
* Code coverage is allways 100%

### Useful Links ###
* [Gogland Docs](https://www.jetbrains.com/help/go/getting-started-with-gogland.html) - Documentation of the Gogland IDE
* [Go Documentation - Set up](https://golang.org/doc/code.html) - Official documentation of golang - Set up the envirement
* [Go Documentation](https://golang.org/doc/) - Official documentation of golang
* [Go standard library](https://golang.org/pkg/) - The go standard library api docs

### Books ###
* [Free German Gitbook](https://astaxie.gitbooks.io/build-web-application-with-golang/content/de/) - Introduction · Build web application with Golang
