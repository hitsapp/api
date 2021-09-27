package utils

import (
	"hits/api/prisma/db"
)

// last workflow test maybe?

var globalDb *db.PrismaClient

func SetGlobalDb(client *db.PrismaClient) {
	globalDb = client
}

func GetPrisma() *db.PrismaClient {
	return globalDb
}
