<div class="card">
    <div class="card-body">
        <h5 class="card-title">New Author</h5>
        <p class="card-text">
            <div class="row-fluid">
                <form action="?action=add" method="post">
                    <div class="form-group">
                        <label for="username">Username</label>
                        <input type="text" class="form-control" id="username" name="username" placeholder="Username" />
                    </div>
                    <div class="form-group">
                        <label for="subtitle">Password</label>
                        <input type="password" class="form-control" id="password" name="password" placeholder="Password" />
                    </div>

                    <button type="submit" class="btn btn-primary" title="Save Author">
                        <i class="fa fa-save" aria-hidden="true"></i>
                    </button>

                    {{ if  .error_empty }}
                        <div class="alert alert-danger margin10" role="alert">
                            All fields must be specified
                        </div>
                    {{end}}

                    {{ if  .error_username}}
                        <div class="alert alert-danger margin10" role="alert">
                            An author with this name already exist.
                        </div>
                    {{end}}
                </form>
            </div>
        </p>
    </div>
</div>