<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Go-Blog</title>

    <link rel="shortcut icon" type="image/x-icon" href="../backend/page/img/favicon.ico">

    <!-- Bootstrap core CSS -->
    <link href="../assets/bootstrap/css/bootstrap.min.css" rel="stylesheet" />

    <!-- Font Awesome -->
    <link href="../assets/fonts/font-awesome/css/font-awesome.min.css" rel="stylesheet" />

    <!-- Custom styles for this template -->
    <link href="../backend/page/css/style.css" rel="stylesheet" />
  </head>

  <body>
    <header>
      <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
        <a class="navbar-brand" href="/admin/posts">
          <img id="logo" src="../backend/page/img/logo.png" alt="Logo"/>
          Blog Admin Panel
        </a>
        <button class="navbar-toggler d-lg-none" type="button" data-toggle="collapse" data-target="#navBar" aria-controls="navBar" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navBar">
          <ul class="navbar-nav mr-auto">
            <li class="nav-item">
              <a class="nav-link" href="/admin/posts">Posts</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/admin/comments">Comments</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/admin/categories">Categories</a>
            </li>
              <li class="nav-item">
                  <a class="nav-link" href="/admin/authors">Authors</a>
              </li>
            <li class="nav-item">
              <a class="nav-link" href="/admin/profile">Profile</a>
            </li>
          </ul>
            <ul class="navbar-nav mr-right" style="margin-right: 20px;">
              <li class="nav-item">
                  <a class="nav-link" href="/">Blog</a>
              </li>
            </ul>
            <span class="mt-2 mt-md-0">
                <a href="/admin/logout" class="btn btn-primary" title="Logout">
                    <i class="fa fa-sign-out" aria-hidden="true"></i>
                </a>
            </span>
        </div>
      </nav>
    </header>

    <div class="container padding50">

          {{.content}}

    </div>

    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src= "../backend/page/js/jquery-slim.min.js"></script>
    <script src="../backend/page/js/popper.min.js"></script>
    <script src="../assets/bootstrap/js/bootstrap.min.js"></script>
  </body>
</html>