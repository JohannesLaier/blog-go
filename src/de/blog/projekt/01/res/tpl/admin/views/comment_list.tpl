<h1>Comments</h1>

{{ if not .newComments }}
    {{ if not .comments }}
        <h5>There are no comments.</h5>
    {{ end }}
{{ end }}

{{ if .newComments}}
<div class="card">
    <div class="card-body">
        <h5 class="card-title">Comments to release</h5>
        <p class="card-text">
            {{range .newComments}}
                <div class="row">
                    <div class="col-xs-12 col-sm-11 col-md-11 col-lg-11">
                        <div class="list-group">
                            <div class="list-group-item flex-column align-items-start">
                                <div class="d-flex w-100 justify-content-between">
                                    <div>
                                        <h5 class="mb-1">{{.comment.Username}}</h5>
                                        <h6 class="font-normal"> zu {{.post.Title}}</h6>
                                    </div>
                                    <small>{{ date .comment.Date}}</small>
                                </div>
                                <p class="mb-1">{{.comment.Text}}</p>
                             </div>
                         </div>
                    </div>
                    <div class="col-xs-12 col-sm-1 col-md-1 col-lg-1">
                        <a href="?action=share&id={{ .comment.Id }}" class="btn btn-primary" title="Release Comment">
                            <i class="fa fa-check" aria-hidden="true"></i>
                        </a>
                        <a href="?action=delete&id={{ .comment.Id }}" class="btn btn-primary margin5" title="Delete Comment">
                            <i class="fa fa-times" aria-hidden="true"></i>
                        </a>
                    </div>
                </div>
            {{end}}
        </p>
    </div>
</div>
{{end}}


{{ if .comments}}
    <div class="card">
        <div class="card-body">
            <h5 class="card-title">Released Comments</h5>
        <p class="card-text">
            {{range .comments}}
                <div class="row">
                    <div class="col-xs-12 col-sm-11 col-md-11 col-lg-11">
                        <div class="list-group">
                            <div class="list-group-item flex-column align-items-start">
                                <div class="d-flex w-100 justify-content-between">
                                    <div>
                                        <h5 class="mb-1">{{.comment.Username}}</h5>
                                        <h6 class="font-normal"> zu {{.post.Title}}</h6>
                                    </div>
                                    <small>{{ date .comment.Date}}</small>
                                </div>
                                <p class="mb-1">{{.comment.Text}}</p>
                            </div>
                         </div>
                     </div>
                    <div class="col-xs-12 col-sm-1 col-md-1 col-lg-1">
                        <a href="?action=delete&id={{ .comment.Id }}" class="btn btn-primary" title="Delete Comment">
                            <i class="fa fa-times" aria-hidden="true"></i>
                        </a>
                    </div>
                </div>
            {{end}}
        </p>
    </div>
</div>
{{end}}
