<div class="card">
    <div class="card-body">
        <h5 class="card-title">Profile Settings</h5>
        <p class="card-text">
            <div class="row-fluid">
                <form action="?action=save" method="post">
                    <div class="form-group">
                        <label for="username">Username</label>
                        <input type="text" class="form-control" id="username" name="username" placeholder="Username" value="{{.user.Username}}" />
                    </div>
                    <div class="form-group">
                        <label for="password">Current password</label>
                        <input type="password" class="form-control" id="password" name="password" placeholder="Password" />
                    </div>
                    <div class="form-group">
                        <label for="newPassword">New password</label>
                        <input type="password" class="form-control" id="newPassword" name="newPassword" placeholder="New password" />
                    </div>
                    <div class="form-group">
                        <input type="password" class="form-control" id="newPassword" name="newPassword2" placeholder="Retype new password" />
                    </div>

                    {{ if  .success }}
                        <div class="alert alert-primary" role="alert">
                            Successfully saved
                        </div>
                    {{end}}

                    {{ if  .error_empty }}
                        <div class="alert alert-danger" role="alert">
                            All fields must be specified
                        </div>
                    {{end}}

                    {{ if  .error_password_invalid }}
                        <div class="alert alert-danger" role="alert">
                            Password is incorrect
                        </div>
                    {{end}}

                    {{ if  .error_passwords_doesnt_match }}
                        <div class="alert alert-danger" role="alert">
                            Passwords does not match
                        </div>
                    {{end}}

                    <button type="submit" class="btn btn-primary" title="Save Profile Settings">
                        <i class="fa fa-save" aria-hidden="true"></i>
                    </button>
                </form>
            </div>
        </p>
    </div>
</div>