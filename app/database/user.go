package database

func Up() {
	_, err := DB.Query(`CREATE TABLE IF NOT EXISTS users (  
		user_id VARCHAR(36) NOT NULL,
		email VARCHAR(255) NOT NULL,
		address VARCHAR(1000),
		password VARCHAR(1000),
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		PRIMARY KEY (user_id) ,
		INDEX EMAIL_UNIQUE (email)
	  );`)
	if err != nil {
		panic(err)
	}
}
