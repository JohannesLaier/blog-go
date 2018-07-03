<div class="card">
    <div class="card-body">
        <h5 class="card-title">All Authors</h5>
        <p class="card-text">

            <div class="row">
                <div class="center-block">
                    <a href="author-detail" class="btn btn-primary" title="New Author">
                        <i class="fa fa-plus" aria-hidden="true"></i>
                    </a>
                </div>
            </div>

            {{ if  .success_delete }}
                <div class="alert alert-primary margin10" role="alert">
                    Successfully deleted
                </div>
            {{end}}

            {{ if not .authors }}
                <div class="center-text">There are no more authors available expexting yourself.</div>
            {{ end }}

            <ul class="list-group padding20">
            {{ range .authors }}
                <li class="list-group-item">
                {{ .Username }}
                    <span class="pull-right">
                        <a class="btn btn-xs btn-default" href="?action=delete&id={{ .Id }}" title="Delete Author">
                            <i class="fa fa-times"></i>
                        </a>
                    </span>
                </li>
            {{ end }}
            </ul>
        </p>
    </div>
</div>