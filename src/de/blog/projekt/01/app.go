package main

import (
	"de/blog/projekt/01/core/config"
	"de/blog/projekt/01/core/http/server"
	"de/blog/projekt/01/app/controller/blog/index"
	"de/blog/projekt/01/app/controller/blog/detail"
	"de/blog/projekt/01/app/controller/blog/about"
	"de/blog/projekt/01/app/controller/blog/contact"
	"de/blog/projekt/01/app/controller/blog/category"
	"de/blog/projekt/01/app/controller/admin/main"
	"de/blog/projekt/01/app/controller/admin/login"
	"de/blog/projekt/01/app/controller/admin/logout"
	"de/blog/projekt/01/app/controller/admin/post"
	"de/blog/projekt/01/app/controller/admin/comment"
	"de/blog/projekt/01/app/controller/admin/profile"
	"de/blog/projekt/01/app/controller/admin/category"
	"de/blog/projekt/01/app/controller/admin/author"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/db"
	"path"
)

func setupDB() {
	blog_db := db.Get("blog")

	blog_keywords := db.NewDBCollection("keyword")
	blog_keyword := keyword.NewKeyword("Cloud")
	blog_keywords.Add(blog_keyword)
	blog_keywords.Add(keyword.NewKeyword("IT"))
	blog_keywords.Add(keyword.NewKeyword("Security"))

	blog_authors := db.NewDBCollection("author")
	blog_author := author.NewAuthor("admin", "123456")
	blog_author2 := author.NewAuthor("Peter Miller", "secret")
	blog_authors.Add(blog_author)
	blog_authors.Add(blog_author2)

	blog_posts := db.NewDBCollection("post")
	blog_post_content := "<p>Lorem ipsum dolor sit amet, consetetur <strong>sadipscing</strong> elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.</p><p><img style=\"display: block; margin-left: auto; margin-right: auto;\" src=\"https://camo.githubusercontent.com/98ed65187a84ecf897273d9fa18118ce45845057/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67\" alt=\"\" width=\"132\" height=\"180\" />Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat.&nbsp; &nbsp;</p> <ul style=\"list-style-type: disc;\"><li>Eleifend</li><li>Molestie</li><li>Duis</li></ul><p>Ut wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi.</p> <ol><li>Ullamcorper suscipit</li><li>Vel illum dolore</li><li>Feugiat nulla facilisis</li></ol><p>Nam liber tempor cum soluta nobis eleifend option congue nihil imperdiet doming id quod mazim placerat facer</p>"
	blog_post := post.NewPost("Lorem ipsum dolor sit amet", "Consetetur sadipscing elitr", blog_post_content, blog_author.GetID())
	blog_post.AddKeyword(blog_keyword.GetID())
	blog_posts.Add(blog_post)

	blog_comment := comment.NewComment("Alex West", "At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet.", blog_post.GetID())
	blog_comment.SetShare(true)

	blog_comments := db.NewDBCollection("comment")
	blog_comments.Add(blog_comment)

	blog_db.AddCollection(blog_posts)
	blog_db.AddCollection(blog_authors)
	blog_db.AddCollection(blog_keywords)
	blog_db.AddCollection(blog_comments)

	blog_db.Store()
}

func main() {
	setupDB()

	config := config.GetConfig()

	session_store := session.GetSessionStore()
	session_store.Load()

	server := server.NewServer(config)

	server.AddFileServer("/frontend/", path.Join(config.GetResourceFolder(),"www","frontend"))
	server.AddFileServer("/backend/", path.Join(config.GetResourceFolder(), "www", "backend"))
	server.AddFileServer("/assets/", path.Join(config.GetResourceFolder(), "www", "assets"))

	server.AddController(blog_index.NewIndexController())
	server.AddController(blog_detail.NewDetailController())
	server.AddController(blog_about.NewAboutController())
	server.AddController(blog_contact.NewContactController())
	server.AddController(blog_category.NewCategoryController())

	server.AddController(admin_main.NewAdminMainController())
	server.AddController(admin_login.NewLoginController())
	server.AddController(admin_logout.NewLogoutController())
	server.AddController(admin_post.NewPostController())
	server.AddController(admin_comment.NewCommentController())
	server.AddController(admin_categorie.NewCategoryController())
	server.AddController(admin_author.NewAuthorController())
	server.AddController(admin_profile.NewProfileController())

	server.Run()
}
