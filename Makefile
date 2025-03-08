schemadiff: internal/domain
internal/domain: migrations
	atlas migrate diff --env gorm

migrate: schemadiff
	atlas migrate apply --url "mysql://user:password@localhost:3306/demo_db"