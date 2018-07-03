<div class="card">
    <div class="card-body">
        <h5 class="card-title">Info</h5>
        <p class="card-text">
            <div class="row-fluid">
                <form action="?action=savePost&id={{ .post.Id }}" method="post">
                    <div class="form-group">
                        <label for="title">Blog-Title</label>
                        <input type="text" class="form-control" id="title" name="title" placeholder="Blog Title" value="{{.post.Title}}" />
                    </div>
                    <div class="form-group">
                        <label for="subtitle">Blog-Subtitle</label>
                        <input type="text" class="form-control" id="subtitle" name="subtitle" placeholder="Blog Subtitle" value="{{.post.SubTitle}}" />
                    </div>
                    <div class="form-group">
                        <label for="content">Content</label>
                        <textarea class="form-control" id="content" name="content" rows="10">
                        {{.post.Content}}
                        </textarea>
                    </div>

                    <div class="form-group">
                        <div>Keywords</div>
                        {{ range .keywords }}

                            <div class="form-check form-check-inline">
                                <input class="form-check-input" type="checkbox" id="{{ .keyword.Id }}" name="keywords[]" value="{{ .keyword.Id }}" {{ if .checked}} checked {{ end }}>
                                <label class="form-check-label" for="{{ .keyword.Id }}">{{ .keyword.Name }}</label>
                            </div>

                        {{ end }}
                    </div>

                    <button type="submit" class="btn btn-primary" title="Save Post">
                        <i class="fa fa-save" aria-hidden="true"></i>
                    </button>

                    <a href="?action=deletePost&id={{ .post.Id }}" class="btn btn-primary" title="Delete Post">
                        <i class="fa fa-times" aria-hidden="true"></i>
                    </a>
                </form>
            </div>
        </p>
    </div>
</div>

{{ if .newComments}}
<div class="card">
    <div class="card-body">
        <h5 class="card-title">Comments to release</h5>
        <p class="card-text">
            {{range .newComments}}
                <div class="row">
                    <div class="col-xs-12 col-sm-11 col-md-11 col-lg-11">
                        <div class="list-group">
                            <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                                <div class="d-flex w-100 justify-content-between">
                                    <h5 class="mb-1">{{.Username}}</h5>
                                    <small>{{ date .Date}}</small>
                                </div>
                                <p class="mb-1">{{.Text}}</p>
                             </a>
                         </div>
                    </div>
                    <div class="col-xs-12 col-sm-1 col-md-1 col-lg-1">
                        <a href="?id={{ $.post.Id }}&action=shareComment&comment_id={{ .Id }}" class="btn btn-primary" title="Release Comment">
                            <i class="fa fa-check" aria-hidden="true"></i>
                        </a>
                        <a href="?id={{ $.post.Id }}&action=deleteComment&comment_id={{ .Id }}" class="btn btn-primary" title="Delete Comment">
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
                            <a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
                                <div class="d-flex w-100 justify-content-between">
                                    <h5 class="mb-1">{{.Username}}</h5>
                                    <small>{{ date .Date}}</small>
                                </div>
                                <p class="mb-1">{{.Text}}</p>
                                </a>
                         </div>
                     </div>
                    <div class="col-xs-12 col-sm-1 col-md-1 col-lg-1">
                        <a href="?id={{ $.post.Id }}&action=deleteComment&comment_id={{ .Id }}" class="btn btn-primary" title="Delete Comment">
                            <i class="fa fa-times" aria-hidden="true"></i>
                        </a>
                    </div>
                </div>
            {{end}}
        </p>
    </div>
</div>
{{end}}

<script src="../backend/page/lib/tinymce/js/tinymce/tinymce.min.js"></script>
<script src="../backend/page/lib/tinymce/js/tinymce/jquery.tinymce.min.js"></script>
<script>
    tinymce.init({
        mode : "textareas",
        height: 500,
        plugins: 'print preview fullpage searchreplace autolink directionality visualblocks visualchars fullscreen image link media template codesample table charmap hr pagebreak nonbreaking anchor toc insertdatetime advlist lists textcolor wordcount imagetools contextmenu colorpicker textpattern help',
        toolbar1: 'formatselect | bold italic strikethrough forecolor backcolor | link | alignleft aligncenter alignright alignjustify  | numlist bullist outdent indent  | removeformat',
    });
</script>