<!-- Page Header -->
<header class="masthead" style="background-image: url('frontend/img/home-bg.jpg')">
    <div class="container">
        <div class="row">
            <div class="col-lg-8 col-md-10 mx-auto">
                <div class="site-heading">
                    <h1>Go Blog</h1>
                    <span class="subheading">Tiny Blog CMS writen in Google Go</span>
                </div>
            </div>
        </div>
    </div>
</header>

<!-- Main Content -->
<div class="container">
    <div class="row">
        <div class="col-lg-8 col-md-10 mx-auto">
            {{range .}}
            <div class="post-preview">
                <a href="/detail?id={{.post.Id}}">
                    <h2 class="post-title">
                        {{.post.Title}}
                    </h2>
                    <h3 class="post-subtitle">
                        {{.post.SubTitle}}
                    </h3>
                </a>
                <p class="post-meta">Posted by
                    {{.author.Username}} on {{ date .post.Date}}
                </p>
            </div>
            <hr>
            {{end}}
        </div>
    </div>
</div>