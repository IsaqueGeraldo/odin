package odin

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/IsaqueGeraldo/agni"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conn *gorm.DB

type Environment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Bootstrap() {
	var err error

	conn, err = gorm.Open(sqlite.Open("odin.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		agni.Println("[odin]: "+err.Error(), agni.RedText)
	}

	conn.AutoMigrate(&Environment{})
}

func Getenv(key string) (string, error) {
	if conn == nil {
		return "", errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	var env Environment

	if err := conn.Where("key = ?", key).First(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("key not found")
		}
		return "", err
	}

	return env.Value, nil
}

func Find(key string) ([]Environment, error) {
	if conn == nil {
		return nil, errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	var env []Environment

	if err := conn.Where("key LIKE ?", key).Find(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no records found matching the key")
		}
		return nil, err
	}

	return env, nil
}

func Setenv(key string, value string) error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	env := Environment{Key: key, Value: value}

	if err := conn.Where("key = ?", key).Assign(env).FirstOrCreate(&env).Error; err != nil {
		return err
	}

	return nil
}

func Unsetenv(key string) error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	key = sanitizeKey(key)
	env := Environment{Key: key}

	if err := conn.Where("key = ?", key).Delete(&env).Error; err != nil {
		return err
	}

	return nil
}

func Environ() ([]Environment, error) {
	if conn == nil {
		return nil, errors.New("the database connection is not initialized")
	}

	var env []Environment

	if err := conn.Find(&env).Error; err != nil {
		return nil, err
	}

	return env, nil
}

func Clearenv() error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	return conn.Exec("DELETE FROM environments").Error
}

func RenameKey(oldKey, newKey string) error {
	if conn == nil {
		return errors.New("the database connection is not initialized")
	}

	oldKey = sanitizeKey(oldKey)
	newKey = sanitizeKey(newKey)

	var env Environment
	if err := conn.Where("key = ?", oldKey).First(&env).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("key not found")
		}
		return err
	}

	var existingEnv Environment
	if err := conn.Where("key = ?", newKey).First(&existingEnv).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if existingEnv.ID != 0 {
		return errors.New("new key already exists")
	}

	env.Key = newKey
	if err := conn.Save(&env).Error; err != nil {
		return err
	}

	return nil
}

func sanitizeKey(key string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9_]+")
	cleaned := regex.ReplaceAllString(key, "_")

	cleaned = strings.TrimSuffix(cleaned, "_")

	return strings.ToUpper(cleaned)
}
