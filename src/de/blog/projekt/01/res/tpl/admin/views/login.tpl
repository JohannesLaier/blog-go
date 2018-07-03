<div class="login-box animated fadeInUp">
        <div class="box-header">
            <h2>Log In</h2>
        </div>
        {{ if  .error }}
            <div class="error">
                Username or password is incorrect
            </div>
        {{end}}
        <form action="" method="post">
               <label for="username">Username</label>
                <br/>
                <input type="text" id="username" name="username">
                <br/>
                <label for="password">Password</label>
                <br/>
                <input type="password" id="password" name="password">
                <br/>
                <button type="submit">Sign In</button>
                <br/>
        </form>
</div>