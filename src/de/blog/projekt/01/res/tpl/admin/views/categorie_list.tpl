<div class="card">
    <div class="card-body">
        <h5 class="card-title">All Categories</h5>
        <p class="card-text">
            <form class="input-group input-group-btn" action="?action=add" method="post">
				<input name="category" type="text" class="form-control input-lg" placeholder="Category" />
				<button type="submit" class="btn btn-lg orange button" title="New Category">
					<i class="fa fa-plus"></i>
				</button>
			</form>

			{{ if  .success_save }}
                <div class="alert alert-primary margin10" role="alert">
                    Successfully saved
                </div>
            {{end}}

            {{ if  .success_delete }}
                <div class="alert alert-primary margin10" role="alert">
                    Successfully deleted
                </div>
            {{end}}

            {{ if not .keywords }}
                <div class="center-text">There are no more categories available.</div>
            {{ end }}

            <ul class="list-group padding20">
                {{ range .keywords }}
                <li class="list-group-item">
                    {{ .Name }}
                    <span class="pull-right">
                        <a class="btn btn-xs btn-default" href="?action=delete&id={{ .Id }}" title="Delete Category">
                            <i class="fa fa-times"></i>
                        </a>
                    </span>
                </li>
                {{ end }}
            </ul>
        </p>
    </div>
</div>