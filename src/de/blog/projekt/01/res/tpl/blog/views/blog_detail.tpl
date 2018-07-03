<!-- Page Header -->
<header class="masthead" style="background-image: url('frontend/img/post-bg.jpg')">
    <div class="container">
        <div class="row">
            <div class="col-lg-8 col-md-10 mx-auto">
                <div class="post-heading">
                    <h1>
                        {{.post.Title}}
                    </h1>
                    <h2 class="subheading">
                        {{.post.SubTitle}}
                    </h2>
                    <span class="meta">
                        Posted by {{.author.Username}} on {{ date .post.Date}}
                    </span>
                </div>
            </div>
        </div>
    </div>
</header>

<!-- Post Content -->
<article>
    <div class="container">
        <div class="row">
            <div class="col-lg-8 col-md-10 mx-auto">

                {{.post.Content}}

                <hr />

                <!-- Categories -->
                <h5 class="card-header">Categories</h5>
                <div class="card-body">
                    <div class="row">
                        <div class="col-lg-6">
                            <ul class="list-unstyled mb-0">
                                {{ range .keywords}}
                                <li>
                                    <a href="/category?id={{.Id}}">{{.Name}}</a>
                                </li>
                                {{ end }}
                            </ul>
                        </div>
                    </div>
                </div>

                <!-- Comment section -->

                <!-- Comments Form -->
                <div class="card my-4">
                    <h5 class="card-header">Leave a Comment:</h5>
                    <div class="card-body">
                        <form action="/detail?action=comment&id={{.post.Id}}" method="POST">
                            <div class="form-group">
                                <input type="text" class="form-control" name="username" placeholder="Nickname" value="{{ .nickname }}"></input>
                            </div>
                            <div class="form-group">
                                <textarea class="form-control" rows="3" name="text" placeholder="Your comment"></textarea>
                            </div>

                            <button type="submit" class="btn btn-primary">Send</button>

                            {{ if  .success }}
                                <div class="alert alert-primary" role="alert">
                                    Successfully saved. It will be displayed after approval by the author
                                </div>
                            {{end}}
                        </form>
                    </div>
                </div>

                <!-- Single Comment -->

                {{ range .comments}}
                    <div class="media mb-4">
                        <div class="media-body">
                            <h5 class="mt-0">{{.Username}}</h5>
                            <p>{{.Text}}</p>
                            <i>Posted at {{ date .Date}}</i>
                        </div>
                    </div>
                {{ end }}

            </div>
        </div>
    </div>
</article>