<div class="card">
    <div class="card-body">
        <div class="row">
            <div class="center-block">
                <a href="post-detail" class="btn btn-primary" title="New Post">
                    <i class="fa fa-plus" aria-hidden="true"></i>
                </a>
            </div>
        </div>

        <h5 class="card-title">All your posts</h5>
        <p class="card-text">

            {{ if not . }}
                <div class="center-text">There are no more posts available.</div>
            {{ end }}

            <div class="row">
                {{range .}}
                <div class="col-xs-12 col-sm-11">
                    <div class="list-group">
                      <a href="post-detail?id={{.Id}}" class="list-group-item list-group-item-action flex-column align-items-start">
                        <div class="d-flex w-100 justify-content-between">
                          <h5 class="mb-1">{{.Title}}</h5>
                          <small>{{ date .Date}}</small>
                        </div>
                        <p class="mb-1">{{.SubTitle}}</p>
                      </a>
                    </div>
                 </div>
                 <div class="col-xs-12 col-sm-1">
                    <a href="?action=delete&id={{ .Id }}" class="btn btn-primary" title="Delete Post">
                        <i class="fa fa-times" aria-hidden="true"></i>
                    </a>
                 </div>
                {{end}}
             </div>
        </p>
    </div>
</div>