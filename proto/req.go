package proto

type (
	ReqCommon struct {
		PageNum  int    `json:"pagenum"`
		PageSize int    `json:"pagesize"`
	}
	ReqFindUser struct {
		PageNum  int    `json:"pagenum"`
		PageSize int    `json:"pagesize"`
		IdOrName string `json:"idorname"`
	}
	ReqAddUser struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Role     int    `json:"role"`
	}
	ReqEditUser struct {
		Id int `json:"id"`
		UserName string `json:"username"`
		Email    string `json:"email"`
		Role     int    `json:"role"`
	}
	ReqFindArticle struct {
		PageNum  int    `json:"pagenum"`
		PageSize int    `json:"pagesize"`
		Title string `json:"title"`
	}
	ReqCateArticle struct {
		PageNum  int    `json:"pagenum"`
		PageSize int    `json:"pagesize"`
		Id int `json:"id"`
	}
)
