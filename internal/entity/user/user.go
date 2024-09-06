package entityuser

type (
	LoginReq struct {
		Email    string `json:"email" redis:"email"`
		Password string `json:"password" redis:"password"`
	}

	Token struct {
		AccessToken  string `json:"access_token" redis:"access_token"`
		RefreshToken string `json:"refresh_token" redis:"refresh_token"`
	}

	LoginRes struct {
		Token Token `json:"token" redis:"token"`
	}

	StatusUser struct {
		Successfully bool `json:"successfully" redis:"successfully"`
	}

	User struct {
		Id           string `json:"id" redis:"id"`
		UserName     string `json:"user_name" redis:"user_name"`
		Email        string `json:"email" redis:"email"`
		PasswordHash string `json:"password_hash" redis:"password_hash"`
	}
	UserRegisterReq struct {
		UserName     string `json:"user_name" redis:"user_name"`
		Email        string `json:"email" redis:"email"`
		PasswordHash string `json:"password_hash" redis:"password_hash"`
		SecretKey    string `json:"secret_key" redis:"secret_key"`
	}

	CreateUserReq struct {
		UserName        string `json:"user_name" redis:"user_name"`
		Email           string `json:"email" redis:"email"`
		Password        string `json:"password" redis:"password"`
		ConfirmPassword string `json:"confirm_password" redis:"confirm_password"`
	}
	StatusMessage struct {
		Message string `json:"message" redis:"message"`
	}
	VerifyUserReq struct {
		Email      string `json:"email" redis:"email"`
		SecretCode string `json:"secret_code" redis:"secret_code"`
	}
)
