package tasksentity

type (
	Task struct {
		Id       string `json:"id" bson:"id"`
		UserId   string `json:"user_id" bson:"user_id"`
		Title    string `json:"title" bson:"title"`
		CreateAt string `json:"create_at" bson:"create_at"`
		Status   Status `json:"status" bson:"status"`
	}
	Status struct {
		Condition   string `json:"condition" bson:"condition"`
		Description string `json:"description" bson:"description"`
		Important   bool   `json:"important" bson:"important"`
	}

	CreateTaskReq struct {
		UserId string `json:"-" bson:"user_id"`
		Title  string `json:"title" bson:"title"`
		Status Status `json:"status" bson:"status"`
	}
	UpdateTaskReq struct {
		UserId string `json:"-" bson:"user_id"`
		Id     string `json:"-" bson:"id"`
		Title  string `json:"title" bson:"title"`
		Status Status `json:"status" bson:"status"`
	}
	GetTaskReq struct {
		UserId string `json:"user_id" bson:"user_id"`
		Field  string `json:"field" bson:"field"`
		Value  string `json:"value" bson:"value"`
	}
	GetAllTaskReq struct {
		UserId string `json:"user_id" bson:"user_id"`
		Field  string `json:"field" bson:"field"`
		Value  string `json:"value" bson:"value"`
		Limit  int    `json:"limit" bson:"limit"`
		Offset int    `json:"offset" bson:"offset"`
	}

	TaskPostgres struct {
		UserId   string `json:"user_id" bson:"user_id"`
		Id       string `json:"id" bson:"id"`
		Title    string `json:"title" bson:"title"`
		CreateAt string `json:"create_at" bson:"create_at"`
	}
	MongoTaskDetails struct {
		TaskId      string `json:"task_id" bson:"task_id"`
		Condition   string `json:"condition" bson:"condition"`
		Description string `json:"description" bson:"description"`
		Important   bool   `json:"important" bson:"important"`
	}
)
