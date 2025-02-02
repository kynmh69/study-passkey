package utils

import (
	"github.com/kynmh69/study-passkey/logger"
	"github.com/kynmh69/study-passkey/prisma/db"
)

var Client *db.PrismaClient

func init() {
	Client = db.NewClient()
	if err := Client.Connect(); err != nil {
		logger.Logger.Panic(err.Error())
	}
}
