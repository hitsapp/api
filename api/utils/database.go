package utils

import (
	"hits/api/prisma/db"
)

// testing workflow x3

var globalDb *db.PrismaClient

func SetGlobalDb(client *db.PrismaClient) {
	globalDb = client
}

func GetPrisma() *db.PrismaClient {
	return globalDb
}
